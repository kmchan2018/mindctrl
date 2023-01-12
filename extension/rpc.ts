

import * as Errors from './errors';
import * as Util from './util';


//////////////////////////////////////////////////////////////////////////
//
// Re-export error types, factories and type guards.
//

export type {
	DispatchError,
	ValidationError,
	InternalError,
	ExecutionError,
} from './errors';

export {
	createDispatchError, isDispatchError,
	createValidationError, isValidationError,
	createInternalError, isInternalError,
	createExecutionError, isExecutionError
} from './errors';


//////////////////////////////////////////////////////////////////////////
//
// Definitions of generic RPC input, output and results.
//

export interface Input {
	[ key: string ]: any;
}

export interface Result {
	success: true;
	[ key: string ]: any;
}

export type Output =
	Result |
	Errors.DispatchError |
	Errors.ValidationError |
	Errors.InternalError |
	Errors.ExecutionError


//////////////////////////////////////////////////////////////////////////
//
// Registry of RPC methods.
//

const registry = <Record<string,(input: Input) => Promise<Output>>>{
	// empty
};


//////////////////////////////////////////////////////////////////////////
//
// Event channels.
//

const onRequestChannel = Util.createEventChannel<[string,Input]>();
const onResponseChannel = Util.createEventChannel<[string,Input,Output]>();


//////////////////////////////////////////////////////////////////////////
//
// Register RPC method. The function accepts three parameters. The first
// one is the name of the method; the second one is an function that
// validates if the input is acceptable; and the last one is an function
// that actually executes the method.
//

export function register<RpcInput extends Input, RpcOutput extends Output>(method: string, validate: (input: Input) => input is RpcInput, execute: (input: RpcInput) => Promise<RpcOutput>) {
	registry[method] = async function(input: Input): Promise<Output> {
		try {
			if (validate(input)) {
				return await execute(input);
			} else {
				return Errors.createValidationError(`invalid input for method ${method}`);
			}
		} catch (err) {
			return Errors.createInternalError(`unexpected exception thrown when calling method ${method}`);
		}
	};
}


//////////////////////////////////////////////////////////////////////////
//
// Dispatch the incoming request to the appropriate method for execution.
//

export async function dispatch(method: string, input: Input): Promise<Output> {
	const unknownMethod = async function(input: Input): Promise<Output> {
		return Errors.createDispatchError(`unknown method ${method}`);
	};

	onRequestChannel.emit(method, input);
	const callable = registry[method] || unknownMethod;
	const output = await callable(input);
	onResponseChannel.emit(method, input, output);
	return output;
}


//////////////////////////////////////////////////////////////////////////
//
// Listenable events.
//

export const onRequest = onRequestChannel.observer;
export const onResponse = onResponseChannel.observer;


