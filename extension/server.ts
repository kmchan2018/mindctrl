

import * as Mqtt from 'mqtt';

import * as Config from './config';
import * as Rpc from './rpc';
import * as Util from './util';


//////////////////////////////////////////////////////////////////////////
//
// Server states contains information on whether the server is running,
// plus any information required to work with the server.
//

interface IdleState {
	type: 'idle';
}

interface StartingState {
	type: 'starting';
	config: Config.Config;
	client: Mqtt.Client;
}

interface ServingState {
	type: 'serving';
	config: Config.Config;
	client: Mqtt.Client;
	timer: ReturnType<typeof setInterval>;
}

interface StoppingState {
	type: 'stopping';
	config: Config.Config;
	client: Mqtt.Client;
}

type State =
	IdleState |
	StartingState |
	ServingState |
	StoppingState


//////////////////////////////////////////////////////////////////////////
//
// Requests are instruction from clients to execute some method plus
// some extra metadata. Responses are replies from server to report any
// result from the method, or any error that arises from execution of
// the method.
//

export interface Request {
	type: 'request';
	id: string;
	method: string;
	params: Rpc.Input;
	client: string;
	server: string;
}

export interface Response {
	type: 'response';
	id: string;
	method: string;
	result: Rpc.Output;
	client: string;
	server: string;
}


//////////////////////////////////////////////////////////////////////////
//
// Type guard for request types. Note that type guard is not needed for
// response because they are generated here and no validation is needed.
//

function isRequest(input: any): input is Request {
	if (typeof input !== 'object') return false;
	if (typeof input['type'] !== 'string') return false;
	if (typeof input['id'] !== 'string') return false;
	if (typeof input['method'] !== 'string') return false;
	if (typeof input['params'] !== 'object') return false;
	if (typeof input['client'] !== 'string') return false;
	if (typeof input['server'] !== 'string') return false;

	if (input['type'] !== 'request') return false;
	if (input['id'] === '') return false;
	if (input['method'] === '') return false;
	if (input['client'] === '') return false;
	if (input['server'] === '') return false;

	return true;
}


//////////////////////////////////////////////////////////////////////////
//
// Current state of the server.
//

let state = { type: 'idle' } as State;


//////////////////////////////////////////////////////////////////////////
//
// Event channels.
//

const onStartingChannel = Util.createEventChannel<[]>();
const onServingChannel = Util.createEventChannel<[]>();
const onStoppingChannel = Util.createEventChannel<[]>();
const onStoppedChannel = Util.createEventChannel<[]>();
const onRequestChannel = Util.createEventChannel<[Request]>();
const onResponseChannel = Util.createEventChannel<[Response,Request]>();
const onUnconfiguredChannel = Util.createEventChannel<[]>();
const onUnreachableChannel = Util.createEventChannel<[]>();
const onDisconnectedChannel = Util.createEventChannel<[]>();
const onGarbageChannel = Util.createEventChannel<[any]>();


//////////////////////////////////////////////////////////////////////////
//
// Setup the server and update the server state after connection to the
// server. It involves:
//
// 1. Subscribing to the server topic for incoming requests
// 2. Publishing an alive message to the status topic
// 3. Starting a timer for periodic alive message
// 4. Transiting the server state to serving
//

function whenConnect() {
	if (state.type === 'starting') {
		const client = state.client;
		const config = state.config;
		const name = state.config.name;

		client.subscribe(`mindctrl/servers/${name}`);
		client.publish(`mindctrl/statuses/${name}`, 'alive', { retain: true });

		const timer = setInterval(function() {
			client.publish(`mindctrl/statuses/${name}`, 'alive', { retain: true });
		}, 60000);

		state = { type: 'serving', config, client, timer };
		onServingChannel.emit();
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Clean up the server after being disconnected from the server. This
// function may be called in multiple situations.
//
// In the first case, this function is called during starting state. It
// means that the MQTT server cannot be connected. This function will
// report the failure and transit the server state back to idle.
//
// In the second case, this function is called during stopping state,
// after a call to stopServer. Since most of the cleanup is already done
// by stopServer, this function will report the completion of shutdown
// sequence and transit the server state back to idle.
//
// In the last case, the function is called during serving state, due
// to network or server problems. The function will stop the lingering
// timer, report the failure and move the server state back to idle.
//

function whenDisconnect() {
	if (state.type === 'starting') {
		state = { type: 'idle' };
		onUnreachableChannel.emit();
	} else if (state.type === 'stopping') {
		state = { type: 'idle' };
		onStoppedChannel.emit();
	} else if (state.type === 'serving') {
		clearInterval(state.timer);
		state = { type: 'idle' };
		onDisconnectedChannel.emit();
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Process incoming messages from the clients.
//

async function whenMessage(topic: string, payload: any) {
	if (state.type === 'serving') {
		const mqtt = state.client;

	  try {
			const message = payload.toString();
			const request = JSON.parse(message);

			if (isRequest(request)) {
				onRequestChannel.emit(request);

				const id = request.id;
				const method = request.method;
				const params = request.params;
				const client = request.client;
				const server = request.server;
				const result = await Rpc.dispatch(method, params);
				const response = { type: 'response', id, method, result, client, server } as Response;

				mqtt.publish(`mindctrl/clients/${client}`, JSON.stringify(response));
				onResponseChannel.emit(response, request);
			} else {
				onGarbageChannel.emit(payload);
			}
		} catch (err) {
			onGarbageChannel.emit(payload);
		}
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Start the server.
//

export async function start() {
	const config = await Config.load();

	if (config) {
		const client = Mqtt.connect(config.url, {
			username: (config.username != '' ? config.username : undefined),
			password: (config.password != '' ? config.password : undefined),
			keepalive: 0,
			connectTimeout: 10 * 1000,
			reconnectPeriod: 0,
			clean: true,
			will: {
				topic: `mindctrl/statuses/${config.name}`,
				payload: 'dead',
				retain: true,
				qos: 0,
			},
		});

		state = { type: 'starting', config, client };
		onStartingChannel.emit();

		client.once('connect', whenConnect);
		client.once('close', whenDisconnect);
		client.on('message', whenMessage);
	} else {
		onUnconfiguredChannel.emit();
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Stop the server.
//

export function stop() {
	if (state.type === 'serving') {
		const client = state.client;
		const config = state.config;
		const name = state.config.name;

		clearInterval(state.timer);
		state.client.publish(`mindctrl/statuses/${name}`, 'dead', { retain: true });
		state.client.end(false, undefined);

		state = { type: 'stopping', config, client };
		onStoppingChannel.emit();
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Listenable events.
//

export const onStarting = onStartingChannel.observer;
export const onServing = onServingChannel.observer;
export const onStopping = onStoppingChannel.observer;
export const onStopped = onStoppedChannel.observer;
export const onRequest = onRequestChannel.observer;
export const onResponse = onResponseChannel.observer;
export const onUnconfigured = onUnconfiguredChannel.observer;
export const onUnreachable = onUnreachableChannel.observer;
export const onDisconnected = onDisconnectedChannel.observer;
export const onGarbage = onGarbageChannel.observer;


