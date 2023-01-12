

import * as WebExtension from 'webextension-polyfill';

import * as Config from './config';
import * as Logger from './logger';
import * as Server from './server';
import * as Timestamp from './timestamp';
import * as Util from './util';

import * as Documents from './documents';
import * as Downloads from './downloads';
import * as Info from './info';
import * as Ping from './ping';
import * as Tabs from './tabs';
import * as Windows from './windows';


//////////////////////////////////////////////////////////////////////////
//
// Register RPC methods.
//

Documents.registerAllMethods();
Downloads.registerAllMethods();
Info.registerAllMethods();
Ping.registerAllMethods();
Tabs.registerAllMethods();
Windows.registerAllMethods();


//////////////////////////////////////////////////////////////////////////
//
// Advise the browser that this tab should not be discarded, since the
// service will stop whenever the page is unloaded.
//

try {
	WebExtension.tabs.update({ autoDiscardable: false });
} catch (err) {
	// this is only an advise; it should not really matter if the
	// call fails
}


//////////////////////////////////////////////////////////////////////////
//
// Enable the Ctrl-Alt-D hotkey to toggle whether requests and responses
// are logged to the developer tool console.
//

(async function() {
	await Util.waitDocumentLoaded();

	let debug = false;

	Server.onRequest.addListener(function(request) {
		if (debug) {
			console.log('Received request %s from client %s:', request.id, request.client, request.method, request);
		}
	});

	Server.onResponse.addListener(function(response, request) {
		if (debug) {
			console.log('Delivered response for request %s from client %s:', request.id, request.client, response);
		}
	});

	document.addEventListener('keyup', function(ev: KeyboardEvent) {
		if (ev.altKey !== true) return;
		if (ev.ctrlKey !== true) return;
		if (ev.code !== "KeyD") return;
		debug = (!debug);
		console.log('debug: ', debug);
	});
})();


//////////////////////////////////////////////////////////////////////////
//
// Handles the server and the console tab.
//

(async function() {
	await Util.waitDocumentLoaded();

	const rootElement = document.querySelector<HTMLDivElement>('#root')!;
	const actionElement = document.querySelector<HTMLButtonElement>('#back-action')!;
	const messagesElement = document.querySelector<HTMLTableSectionElement>('#messages')!;
	const templateElement = document.querySelector<HTMLTableRowElement>('#templates tr:first-of-type')!;
	const statusElement = document.querySelector<HTMLElement>('#footer')!;
	const startElement = document.querySelector<HTMLButtonElement>('#start-action')!;
	const stopElement = document.querySelector<HTMLButtonElement>('#stop-action')!;

	function updateStatus(status: string) {
		statusElement.innerText = status;
	}

	function updatePermittedActions(start: boolean, stop: boolean) {
		if (stop) rootElement.classList.add('stoppable');
		if (start) rootElement.classList.add('startable');
		if (!stop) rootElement.classList.remove('stoppable');
		if (!start) rootElement.classList.remove('startable');
	}

	Logger.onWrite.addListener(function(message: string, timestamp: Date) {
		if (messagesElement.firstElementChild && messagesElement.childElementCount >= 1000) {
			const rowElement = messagesElement.firstElementChild;
			const timeElement = rowElement.querySelector<HTMLTableCellElement>('td.time')!;
			const textElement = rowElement.querySelector<HTMLTableCellElement>('td.text')!;
			rowElement.remove();

			textElement.innerText = message;
			timeElement.innerText = Timestamp.formatShortTimestamp(timestamp);
			timeElement.setAttribute('title', Timestamp.formatFullTimestamp(timestamp));
			messagesElement.appendChild(rowElement);
		} else {
			const rowElement = templateElement.cloneNode(true) as HTMLTableRowElement;
			const timeElement = rowElement.querySelector<HTMLTableCellElement>('td.time')!;
			const textElement = rowElement.querySelector<HTMLTableCellElement>('td.text')!;

			textElement.innerText = message;
			timeElement.innerText = Timestamp.formatShortTimestamp(timestamp);
			timeElement.setAttribute('title', Timestamp.formatFullTimestamp(timestamp));
			messagesElement.appendChild(rowElement);
		}
	});

	Server.onStarting.addListener(function() {
		Logger.write('Mindctrl is starting');
		updateStatus('Starting');
		updatePermittedActions(false, false);
	});

	Server.onServing.addListener(function() {
		Logger.write('Mindctrl is running');
		updateStatus('Running');
		updatePermittedActions(false, true);
	});

	Server.onStopping.addListener(function() {
		Logger.write('Mindctrl is stopping');
		updateStatus('Stopping');
		updatePermittedActions(false, false);
	});

	Server.onStopped.addListener(function() {
		Logger.write('Mindctrl has stopped');
		updateStatus('Idle');
		updatePermittedActions(true, false);
	});

	Server.onUnconfigured.addListener(function() {
		Logger.write('Mindctrl has to be configured before starting');
		updateStatus('Idle');
		updatePermittedActions(true, false);
	});

	Server.onUnreachable.addListener(function() {
		Logger.write('Mindctrl cannot connect to the remote MQTT server');
		updateStatus('Idle');
		updatePermittedActions(true, false);
	});

	Server.onDisconnected.addListener(function() {
		Logger.write('Mindctrl is disconnected from the remote MQTT server');
		updateStatus('Idle');
		updatePermittedActions(true, false);
	});

	actionElement.addEventListener('click', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();
		rootElement.classList.replace('about', 'console');
		rootElement.classList.replace('options', 'console');
	});

	startElement.addEventListener('click', function(ev: Event) {
		Server.start();
	});

	stopElement.addEventListener('click', function(ev: Event) {
		Server.stop();
	});

	rootElement.classList.remove('development');
	rootElement.classList.add('production');

	Logger.write('Mindctrl is ready');
	updateStatus('Idle');
	updatePermittedActions(true, false);
})();


//////////////////////////////////////////////////////////////////////////
//
// Handles the preferences tab. It handles opening the preferences tab,
// updating the form with current preferences, validating form values
// and saving the preferences.
//

(async function() {
	await Util.waitDocumentLoaded();

	const rootElement = document.querySelector<HTMLDivElement>('#root')!;
	const actionElement = document.querySelector<HTMLButtonElement>('#options-action')!;
	const formElement = document.querySelector<HTMLFormElement>('#options')!;
	const urlElement = document.querySelector<HTMLInputElement>('#options input[name="url"]')!;
	const nameElement = document.querySelector<HTMLInputElement>('#options input[name="name"]')!;
	const usernameElement = document.querySelector<HTMLInputElement>('#options input[name="username"]')!;
	const passwordElement = document.querySelector<HTMLInputElement>('#options input[name="password"]')!;
	const reloadElement = document.querySelector<HTMLButtonElement>('#options button[name="reload"]')!;

	actionElement.addEventListener('click', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();

		Config.load().then(function(config) {
			if (config !== undefined) {
				urlElement.setAttribute('value', config.url);
				nameElement.setAttribute('value', config.name);
				usernameElement.setAttribute('value', config.username || '');
				passwordElement.setAttribute('value', config.password || '');
				formElement.reset();
				rootElement.classList.replace('console', 'options');
			} else {
				urlElement.setAttribute('value', '');
				nameElement.setAttribute('value', '');
				usernameElement.setAttribute('value', '');
				passwordElement.setAttribute('value', '');
				formElement.reset();
				rootElement.classList.replace('console', 'options');
			}
		});
	});

	urlElement.addEventListener('change', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();
		urlElement.setCustomValidity(Config.getUrlValidationError(urlElement.value));
		urlElement.reportValidity();
	});

	nameElement.addEventListener('change', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();
		nameElement.setCustomValidity(Config.getNameValidationError(nameElement.value));
		nameElement.reportValidity();
	});

	reloadElement.addEventListener('click', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();

		Config.load().then(function(config) {
			if (config !== undefined) {
				urlElement.setAttribute('value', config.url);
				nameElement.setAttribute('value', config.name);
				usernameElement.setAttribute('value', config.username || '');
				passwordElement.setAttribute('value', config.password || '');
				formElement.reset();
			} else {
				urlElement.setAttribute('value', '');
				nameElement.setAttribute('value', '');
				usernameElement.setAttribute('value', '');
				passwordElement.setAttribute('value', '');
				formElement.reset();
			}
		});
	});

	formElement.addEventListener('submit', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();

		const url = urlElement.value;
		const name = nameElement.value;
		const username = usernameElement.value;
		const password = passwordElement.value;

		Config.save({ url, name, username, password }).then(function() {
			urlElement.setAttribute('value', url);
			nameElement.setAttribute('value', name);
			usernameElement.setAttribute('value', username);
			passwordElement.setAttribute('value', password);
		});
	});

	formElement.addEventListener('reset', function(ev: Event) {
		urlElement.setCustomValidity(Config.getUrlValidationError(urlElement.value));
		urlElement.reportValidity();
		nameElement.setCustomValidity(Config.getNameValidationError(nameElement.value));
		nameElement.reportValidity();
	});
})();


//////////////////////////////////////////////////////////////////////////
//
// Handles the about tab.
//

(async function() {
	await Util.waitDocumentLoaded();

	const rootElement = document.querySelector<HTMLElement>('#root')!;
	const actionElement = document.querySelector<HTMLButtonElement>('#about-action')!;

	actionElement.addEventListener('click', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();
		rootElement.classList.replace('console', 'about');
	});
})();


