

import * as WebExtension from 'webextension-polyfill';

import * as Context from './context';
import * as Rpc from './rpc';
import * as Util from './util';
import * as Validator from './validator';


//////////////////////////////////////////////////////////////////////////
//
// Register all RPC methods.
//

export function registerAllMethods() {
	registerQueryMethod();
}


//////////////////////////////////////////////////////////////////////////
//
// Register documents.query RPC method.
//
// The method executes a GraphQL query over the HTML document inside the
// given tab. The method is implemented by injecting the GraphQL engine
// into the document via content script, and then calling the injected
// GraphQL engine to execute the query and return the result.
//

interface QueryInput {
	tabId: number;
	query: string;
	operation?: string;
	variables?: Record<string,any>;
}

interface QueryResult {
	success: true;
	result: any;
}

export function registerQueryMethod() {
	function isQueryResult(input: Record<string,any>): input is QueryResult {
		if (Validator.validateType(input.success, Validator.isTrue) === false) {
			return false;
		} else {
			return true;
		}
	}

	Rpc.register(
		'documents.query',

		function (input: Rpc.Input): input is QueryInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else if (Validator.validateType(input.query, Validator.isString) === false) {
				return false;
			} else if (Validator.validateType(input.operation, Validator.isString, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.variables, Validator.isRecord, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: QueryInput): Promise<QueryResult|Rpc.ExecutionError|Rpc.InternalError> {
			try {
				// Note that args has to be JSON serializable. In Chrome, the value 'undefined'
				// is not JSON serializable and will cause the scripting.executeScript call to
				// fail (on the other hand, Firefox is more lenient in this regard). Therefore,
				// missing arguments in the args array have to be converted to null. 

				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const target = { tabId };

				const injections = await WebExtension.scripting.executeScript({
					target: { tabId },
					files: [ '/content_scripts/query.js' ],
				});
			
				if (injections.length == 0) {
					return Rpc.createExecutionError(`tab ${tabId} cannot be injected`);
				} else if (injections[0] === undefined) {
					return Rpc.createExecutionError(`tab ${tabId} cannot be injected`);
				} else if (injections[0].error) {
					console.error('[BUG] Query content script throws unexpected error: ', injections[0].error);
					return Rpc.createInternalError(`unexpected error thrown when injecting query content script to tab ${tabId}`);
				}

				const invocations = await WebExtension.scripting.executeScript({
					target: { tabId },
					args: [ input.query, input.operation || null, input.variables || null ],
					func: async function(query: string, operation?: string, variables?: Record<string,any>) {
						// @ts-ignore
						return await Query.Invoke(query, operation, variables);
					},
				});

				if (invocations.length == 0) {
					return Rpc.createExecutionError(`tab ${tabId} cannot be injected`);
				} else if (invocations[0] === undefined) {
					return Rpc.createExecutionError(`tab ${tabId} cannot be injected`);
				} else if (invocations[0].error) {
					console.error('[BUG] Invocation of query content script throws unexpected error: ', invocations[0].error);
					return Rpc.createInternalError(`unexpected error thrown when invoking query content script in tab ${tabId}`);
				} else {
					const result = invocations[0].result;

					if (typeof result === 'object') {
						if (isQueryResult(result)) {
							return result;
						} else if (Rpc.isExecutionError(result)) {
							return result;
						} else if (Rpc.isInternalError(result)) {
							return result;
						}
					}

					console.error('[BUG] Query content script returns malformed data: ', result);
					return Rpc.createInternalError(`malformed data returned after invoking query content script in tab ${tabId}`);
				}
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


