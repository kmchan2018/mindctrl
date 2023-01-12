

import * as WebExtension from 'webextension-polyfill';

import * as Browser from './browser';
import * as Context from './context';
import * as Logger from './logger';
import * as Rpc from './rpc';
import * as Util from './util';
import * as Validator from './validator';


//////////////////////////////////////////////////////////////////////////
//
// Register all RPC methods.
//

export function registerAllMethods() {
	registerFindMethod();
	registerGetMethod();
	registerGetCurrentMethod();
	registerCreateMethod();
	registerLoadMethod();
	registerReloadMethod();
	registerActivateMethod();
	registerMuteMethod();
	registerUnmuteMethod();
	registerPinMethod();
	registerUnpinMethod();
	registerMoveMethod();
	registerDiscardMethod();
	registerRemoveMethod();
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.find RPC method.
//
// The method searches the browser for tabs that matches the given
// criteria and reports the matches back to the caller.
//
// The method is a simple wrapper over the tabs.query Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-query
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/query
//

interface FindInput {
	windowId?: number;
	url?: string;
	status?: WebExtension.Tabs.TabStatus;
	active?: boolean;
	audible?: boolean;
	discarded?: boolean;
	muted?: boolean;
	pinned?: boolean;
}

interface FindResult {
	success: true;
	result: Array<WebExtension.Tabs.Tab>;
}

export function registerFindMethod() {
	Rpc.register(
		'tabs.find',

		function (input: Rpc.Input): input is FindInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.url, Validator.isMatchPattern, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.status, Validator.isTabStatus, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.active, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.audible, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.discarded, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.muted, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.pinned, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: FindInput): Promise<FindResult|Rpc.ExecutionError> {
			try {
				const windowType = 'normal' as const;
				const url = input.url;
				const status = input.status;
				const active = input.active;
				const audible = input.audible;
				const discarded = input.discarded;
				const muted = input.muted;
				const pinned = input.pinned;

				if (input.windowId) {
					const windowId = await Context.ensureNotConsoleWindow(input.windowId);
					const result = await WebExtension.tabs.query({ windowType, windowId, url, status, active, audible, discarded, muted, pinned });
					return { success: true, result };
				} else {
					const result = await WebExtension.tabs.query({ windowType, url, status, active, audible, discarded, muted, pinned });
					return { success: true, result };
				}
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.get RPC method.
//
// The method retrieves information on the tab identified by the given
// tab ID, and report the information back to the caller.
//
// The method is a simple wrapper over the tabs.get Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-get
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/get
//

interface GetInput {
	tabId: number;
}

interface GetResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerGetMethod() {
	Rpc.register(
		'tabs.get',

		function (input: Rpc.Input): input is GetInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: GetInput): Promise<GetResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.get(tabId);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.get_current RPC method.
//
// The method retrieves information on the active tab of the current
// window, and reports the information back to the caller.
//
// The methodis a simple wrapper over the tabs.query Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-query
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/query
//

interface GetCurrentInput {
	// empty
}

interface GetCurrentResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerGetCurrentMethod() {
	Rpc.register(
		'tabs.get_current',

		function (input: Rpc.Input): input is GetCurrentInput {
			return true;
		},

		async function (input: GetCurrentInput): Promise<GetCurrentResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.getLastFocusedWindow();
				const results = await WebExtension.tabs.query({ windowId, active: true });

				if (results.length >= 1) {
					const result = results[0]
					return { success: true, result };
				} else {
					return Rpc.createExecutionError('current tab cannot be found');
				}
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.create RPC method.
//
// The method opens a new tab in the browser, and reports information on
// the new tab back to the caller.
//
// The windowId option determines where the new tab will be opened in.
// If the option is not specified, the new tab will be created in the
// current window.
//
// Next, the active option determines whether the new tab will be
// activate or not. If the option is not specified, the new tab will
// be inactive.
//
// Moreover, the url options determines the initial page the tab will
// load. If the options is not specified, a blank page (about:blank)
// will be loaded.
//
// Finally, the noWait option determines when the method will return.
// If the option is false or unspecified, the method will return after
// the tab is opened and fully loaded. If the option is true, the method
// will return early, not waiting the tab to be fully loaded.
//
// The method is a simple wrapper over the tabs.create Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-create
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/create
//

interface CreateInput {
	windowId?: number;
	url?: string;
	active?: boolean;
	noWait?: boolean;
}

interface CreateResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerCreateMethod() {
	Rpc.register(
		'tabs.create',

		function (input: Rpc.Input): input is CreateInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.url, Validator.isUrl, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.active, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.noWait, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: CreateInput): Promise<CreateResult|Rpc.ExecutionError> {
			try {
				const url = input.url || 'about:blank';
				const active = input.active || false;
				const noWait = input.noWait || false;

				const windowId = input.windowId || await Context.getLastFocusedWindow();
				const result = await WebExtension.tabs.create({ url, active, windowId });
				const tabId = result.id!;

				return await new Promise<CreateResult|Rpc.ExecutionError>((resolve, reject) => {
					function handleRetrieved(tab: WebExtension.Tabs.Tab) {
						if (tab.id === tabId && tab.status === 'complete') {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							Logger.write(`Tab ${tabId} created`);
							resolve({ success: true, result: tab });
						}
					}

					function handleUpdated(id: number, info: WebExtension.Tabs.OnUpdatedChangeInfoType, tab: WebExtension.Tabs.Tab) {
						if (id === tabId && tab.status === 'complete') {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							Logger.write(`Tab ${tabId} created`);
							resolve({ success: true, result: tab });
						}
					}

					function handleRemoved(id: number, info: WebExtension.Tabs.OnRemovedRemoveInfoType) {
						if (id === tabId) {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							resolve(Rpc.createExecutionError(`tab ${tabId} closed before fully loaded`));
						}
					}

					function handleFailed(error: any) {
						WebExtension.tabs.onUpdated.removeListener(handleUpdated);
						WebExtension.tabs.onRemoved.removeListener(handleRemoved);
						resolve(Rpc.createExecutionError(error));
					}

					if (noWait) {
						Logger.write(`Tab ${tabId} created`);
						resolve({ success: true, result });
					} else {
						WebExtension.tabs.onUpdated.addListener(handleUpdated);
						WebExtension.tabs.onRemoved.addListener(handleRemoved);
						WebExtension.tabs.get(tabId).then(handleRetrieved, handleFailed);
					}
				});
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.load RPC method.
//
// The method loads a new page in the tab identified by the given tab ID,
// and reports updated information on the loaded tab back to the caller.
//
// The replace option determines how the new page will be inserted into
// the history stack. If the option is true, the new page will replace
// the current page in the history stack. Otherwise, the new page will
// be pushed on top of the current page.
//
// Moreover, the noWait option determines when the method will return.
// If the option is false or unspecified, the method will return after
// the tab is fully loaded. If the option is true, the method will
// return early, not waiting the tab to be fully loaded.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface LoadInput {
	tabId: number;
	url: string;
	replace?: boolean;
	noWait?: boolean;
}

interface LoadResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerLoadMethod() {
	Rpc.register(
		'tabs.load',

		function (input: Rpc.Input): input is LoadInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else if (Validator.validateType(input.url, Validator.isUrl, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.replace, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.noWait, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: LoadInput): Promise<LoadResult|Rpc.ExecutionError> {
			try {
				const url = input.url;
				const updates = { url } as WebExtension.Tabs.UpdateUpdatePropertiesType;
				const noWait = input.noWait || false;

				// Only Firefox supports the loadReplace flag to control if the current
				// page should be replaced in the history stack. Other browser does not.
				// Therefore, we check for need to be selective here.

				if (await Browser.isFirefox()) {
					updates.loadReplace = input.replace;
				}

				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, updates);

				// Gives browser some time to begin the load process. Without
				// the wait, the browser may return even before the process
				// starts.

				await Util.waitDuration(1000);

				return await new Promise<LoadResult|Rpc.ExecutionError>((resolve, reject) => {
					function handleRetrieved(tab: WebExtension.Tabs.Tab) {
						if (tab.id === tabId && tab.status === 'complete') {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							Logger.write(`Tab ${tabId} navigated`);
							resolve({ success: true, result: tab });
						}
					}

					function handleUpdated(id: number, info: WebExtension.Tabs.OnUpdatedChangeInfoType, tab: WebExtension.Tabs.Tab) {
						if (id === tabId && tab.status === 'complete') {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							Logger.write(`Tab ${tabId} navigated`);
							resolve({ success: true, result: tab });
						}
					}

					function handleRemoved(id: number, info: WebExtension.Tabs.OnRemovedRemoveInfoType) {
						if (id === tabId) {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							resolve(Rpc.createExecutionError(`tab ${tabId} closed before fully loaded`));
						}
					}

					function handleFailed(error: any) {
						WebExtension.tabs.onUpdated.removeListener(handleUpdated);
						WebExtension.tabs.onRemoved.removeListener(handleRemoved);
						resolve(Rpc.createExecutionError(error));
					}

					if (noWait) {
						Logger.write(`Tab ${tabId} navigated`);
						resolve({ success: true, result });
					} else {
						WebExtension.tabs.onUpdated.addListener(handleUpdated);
						WebExtension.tabs.onRemoved.addListener(handleRemoved);
						WebExtension.tabs.get(tabId).then(handleRetrieved, handleFailed);
					}
				});
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.reload RPC method.
//
// The method reloads the page in the tab identified by the given tab ID.
//
// The bypassCache option controls whether browser cache is used for the
// reload. If the option is true, the browser cache will NOT be used and
// resources will be reloaded from the Internet. Otherwise, the browser
// cache is used.
//
// Moreover, the noWait option determines when the method will return.
// If the option is false or unspecified, the method will return after
// the tab is fully loaded. If the option is true, the method will
// return early, not waiting the tab to be fully loaded.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface ReloadInput {
	tabId: number;
	bypassCache?: boolean;
	noWait?: boolean;
}

interface ReloadResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerReloadMethod() {
	Rpc.register(
		'tabs.reload',

		function (input: Rpc.Input): input is ReloadInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else if (Validator.validateType(input.bypassCache, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.noWait, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: ReloadInput): Promise<ReloadResult|Rpc.ExecutionError> {
			try {
				const bypassCache = input.bypassCache || false;
				const noWait = input.noWait || false;
				const tabId = await Context.ensureNotConsoleTab(input.tabId);

				await WebExtension.tabs.reload(tabId, { bypassCache });

				// Gives browser some time to begin the reload process. Without
				// the wait, the browser may return even before the process
				// starts.

				await Util.waitDuration(1000);

				return await new Promise<ReloadResult|Rpc.ExecutionError>((resolve, reject) => {
					function handleRetrieved(tab: WebExtension.Tabs.Tab) {
						if (tab.id == tabId && (noWait || tab.status === 'complete')) {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							Logger.write(`Tab ${tabId} reloaded`);
							resolve({ success: true, result: tab });
						}
					}

					function handleUpdated(id: number, info: WebExtension.Tabs.OnUpdatedChangeInfoType, tab: WebExtension.Tabs.Tab) {
						if (id === tabId && tab.status === 'complete') {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							Logger.write(`Tab ${tabId} reloaded`);
							resolve({ success: true, result: tab });
						}
					}

					function handleRemoved(id: number, info: WebExtension.Tabs.OnRemovedRemoveInfoType) {
						if (id === tabId) {
							WebExtension.tabs.onUpdated.removeListener(handleUpdated);
							WebExtension.tabs.onRemoved.removeListener(handleRemoved);
							resolve(Rpc.createExecutionError(`tab ${tabId} closed before fully reloaded`));
						}
					}

					function handleFailed(error: any) {
						WebExtension.tabs.onUpdated.removeListener(handleUpdated);
						WebExtension.tabs.onRemoved.removeListener(handleRemoved);
						resolve(Rpc.createExecutionError(error));
					}

					if (noWait) {
						WebExtension.tabs.get(tabId).then(handleRetrieved, handleFailed);
					} else {
						WebExtension.tabs.onUpdated.addListener(handleUpdated);
						WebExtension.tabs.onRemoved.addListener(handleRemoved);
						WebExtension.tabs.get(tabId).then(handleRetrieved, handleFailed);
					}
				});
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.activate RPC method.
//
// The method activates the tab identified by the given tab ID, making it
// to be the current active tab of the parent window.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface ActivateInput {
	tabId: number;
}

interface ActivateResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerActivateMethod() {
	Rpc.register(
		'tabs.activate',

		function (input: Rpc.Input): input is ActivateInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: ActivateInput): Promise<ActivateResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, { active: true });
				Logger.write(`Tab ${tabId} activated`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.deactivate RPC method.
//
// The method deactivates the tab identified by the given tab ID. When
// the target tab is deactivated, the browser may activate another
// random tab in the same window.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface DeactivateInput {
	tabId: number;
}

interface DeactivateResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerDeactivateMethod() {
	Rpc.register(
		'tabs.deactivate',

		function (input: Rpc.Input): input is DeactivateInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: DeactivateInput): Promise<DeactivateResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, { active: false });
				Logger.write(`Tab ${tabId} deactivated`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.mute RPC method.
//
// The method mutes sound playing from the tab identified by the given
// tab ID. Audio from muted tabs are silenced by the browser.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface MuteInput {
	tabId: number;
}

interface MuteResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerMuteMethod() {
	Rpc.register(
		'tabs.mute',

		function (input: Rpc.Input): input is MuteInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: MuteInput): Promise<MuteResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, { muted: true });
				Logger.write(`Tab ${tabId} muted`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.unmute RPC method.
//
// The method unmutes the tab identified by the given tab ID and makes
// audio playbook from the tab audible.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface UnmuteInput {
	tabId: number;
}

interface UnmuteResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerUnmuteMethod() {
	Rpc.register(
		'tabs.unmute',

		function isUnmuteInput(input: Rpc.Input): input is UnmuteInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: UnmuteInput): Promise<UnmuteResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, { muted: false });
				Logger.write(`Tab ${tabId} unmuted`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.pin RPC method.
//
// The method pins the tab specified by the given tab ID.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface PinInput {
	tabId: number;
}

interface PinResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerPinMethod() {
	Rpc.register(
		'tabs.pin',

		function (input: Rpc.Input): input is PinInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: PinInput): Promise<PinResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, { pinned: true });
				Logger.write(`Tab ${tabId} pinned`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.unpin RPC method.
//
// The method unpins the tab specified by the given tab ID.
//
// The method is a simple wrapper over the tabs.update Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/update
//

interface UnpinInput {
	tabId: number;
}

interface UnpinResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerUnpinMethod() {
	Rpc.register(
		'tabs.unpin',

		function (input: Rpc.Input): input is UnpinInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: UnpinInput): Promise<UnpinResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				const result = await WebExtension.tabs.update(tabId, { pinned: false });
				Logger.write(`Tab ${tabId} unpinned`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.move RPC method.
//
// The method moves the tab identified by the given tab ID to a different
// position in the same or different window.
//
// The method is a simple wrapper over the tabs.move Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-move
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/move
//

interface MoveInput {
	tabId: number;
	index: number;
	windowId?: number;
}

interface MoveResult {
	success: true;
	result: WebExtension.Tabs.Tab;
}

export function registerMoveMethod() {
	Rpc.register(
		'tabs.move',

		function (input: Rpc.Input): input is MoveInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else if (Validator.validateType(input.index, Validator.isNumber) === false) {
				return false;
			} else if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: MoveInput): Promise<MoveResult|Rpc.ExecutionError> {
			try {
				const index = input.index;

				if (input.windowId) {
					const tabId = await Context.ensureNotConsoleTab(input.tabId);
					const windowId = await Context.ensureNotConsoleWindow(input.windowId);
					const result = await WebExtension.tabs.move(tabId, { index, windowId }) as WebExtension.Tabs.Tab;
					Logger.write(`Tab ${tabId} moved to position ${index} of window ${windowId}`);
					return { success: true, result };
				} else {
					const tabId = await Context.ensureNotConsoleTab(input.tabId);
					const result = await WebExtension.tabs.move(tabId, { index }) as WebExtension.Tabs.Tab;
					Logger.write(`Tab ${tabId} moved to position ${index}`);
					return { success: true, result };
				}
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.discard RPC method.
//
// The method discards the tab specified by the given tab ID. When a tab
// is discarded, the tab still exists in the browser window but its
// content is unloaded from memory. Returning to the tab requires a page
// reload.
//
// Calling the method on an already discarded tab is an no-op. On the
// other hand, the behavior of the method on an active tab depends on
// the browser and therefore implementation defined.
//
// Note that the browser may implement discard by replacing the current
// tab with a new, uninitialized tab, leading to change of tab ID.
//
// The method is a simple wrapper over the tabs.discard Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-discard
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/discard
//

interface DiscardInput {
	tabId: number;
}

interface DiscardResult {
	success: true;
}

export function registerDiscardMethod() {
	Rpc.register(
		'tabs.discard',

		function (input: Rpc.Input): input is DiscardInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: DiscardInput): Promise<DiscardResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				await WebExtension.tabs.discard(tabId);;
				Logger.write(`Tab ${tabId} discarded`);
				return { success: true };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register tabs.remove RPC method.
//
// The method closes the tab identified by the given tab ID. After a tab
// is closed, the tab will disappear from the browser window.
//
// The method is a simple wrapper over the tabs.remove Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/tabs/#method-remove
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/tabs/remove
//

interface RemoveInput {
	tabId: number;
}

interface RemoveResult {
	success: true;
}

export function registerRemoveMethod() {
	Rpc.register(
		'tabs.remove',

		function (input: Rpc.Input): input is RemoveInput {
			if (Validator.validateType(input.tabId, Validator.isTabId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: RemoveInput): Promise<RemoveResult|Rpc.ExecutionError> {
			try {
				const tabId = await Context.ensureNotConsoleTab(input.tabId);
				await WebExtension.tabs.remove(tabId);
				Logger.write(`Tab ${tabId} removed`);
				return { success: true };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


