

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
	registerMoveMethod();
	registerResizeMethod();
	registerMinimizeMethod();
	registerMaximizeMethod();
	registerFullscreenMethod();
	registerRestoreMethod();
	registerFocusMethod();
	registerUnfocusMethod();
	registerRemoveMethod();
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.find RPC method.
//
// The method searches the browser for windows that matches the given
// criteria and reports the result back to the caller.
//
// The method is a simple wrapper over the windows.getAll Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-getAll
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/getAll
//

interface FindInput {
	// empty
}

interface FindResult {
	success: true;
	result: Array<WebExtension.Windows.Window>;
}

export function registerFindMethod() {
	Rpc.register(
		'windows.find',

		function (input: Rpc.Input): input is FindInput {
			return true;
		},

		async function (input: FindInput): Promise<FindResult|Rpc.ExecutionError> {
			try {
				const result = await WebExtension.windows.getAll({ populate: true, windowTypes: [ 'normal' ] });
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.get RPC method.
//
// The method gets information on the window identified by the given
// window ID, and reports the result back to the caller.
//
// The method is a simple wrapper over the windows.get Web Extension API.
// Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-get
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/get
//

interface GetInput {
	windowId: number;
}

interface GetResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerGetMethod() {
	Rpc.register(
		'windows.get',

		function (input: Rpc.Input): input is GetInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: GetInput): Promise<GetResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.get(windowId, { populate: true });
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.get_current RPC method.
//
// The method gets information on the current window (the most recent
// window to be focused), and reports the result back to the caller.
//
// Note that the method is not a wrapper over the windows.getLastFocused
// Web Extension API. For more information, check out the following:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-getLastFocused
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/getLastFocused
// https://bugs.chromium.org/p/chromium/issues/detail?id=546696
//
// Instead, the extension listens to active change events and keep track
// of the window that received focus most recently. Details on the event
// can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#event-onFocusChanged
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/onFocusChanged
//

interface GetCurrentInput {
	// empty
}

interface GetCurrentResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerGetCurrentMethod() {
	Rpc.register(
		'windows.get_current',

		function (input: Rpc.Input): input is GetCurrentInput {
			return true;
		},

		async function (input: GetCurrentInput): Promise<GetCurrentResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.getLastFocusedWindow();
				const result = await WebExtension.windows.get(windowId, { populate: true });
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.create RPC method.
//
// The method opens a new window, and reports its details back to the
// caller.
//
// The url option specifies the url of the initial tab of the new
// window. If not given, the initial tab of the new window will show
// a blank page.
//
// The state option specifies the state of the new window. The allowed
// values include 'normal', 'minimized', 'maximized' and 'fullscreen'.
// If the option is not specified, the new window will be created with
// 'normal' state.
//
// The focus option specifies if the new window will receive focus
// automatically.
//
// The top and left options specifies the position of the new window
// on the screen. If they are not specified, the window will be placed
// at the browsers' discretion.
//
// The width and height options specifies the size of the new window.
// If they are not specified, the window will be sized at the browsers'\
// discretion.
//
// The method is a simple wrapper over the windowss.create Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-create
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/create
//

interface CreateInput {
	url?: string;
	state?: Exclude<WebExtension.Windows.WindowState,'docked'>;
	focused?: boolean;
	top?: number;
	left?: number;
	width?: number;
	height?: number;
}

interface CreateResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerCreateMethod() {
	Rpc.register(
		'windows.create',

		function (input: Rpc.Input): input is CreateInput {
			if (Validator.validateType(input.url, Validator.isUrl, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.focused, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.top, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.left, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.width, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.height, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: CreateInput): Promise<CreateResult|Rpc.ExecutionError> {
			try {
				const url = input.url || 'about:blank';
				const state = input.state || 'normal';
				const focused = (state === 'normal' ? input.focused || false : false);
				const top = (state === 'normal' ? input.top || undefined : undefined);
				const left = (state === 'normal' ? input.left || undefined : undefined);
				const width = (state === 'normal' ? input.width || undefined : undefined);
				const height = (state === 'normal' ? input.height || undefined : undefined);
				const options = { url, state, focused, top, left, width, height } as WebExtension.Windows.CreateCreateDataType;
				const result = await WebExtension.windows.create(options);
				Logger.write(`Window ${result.id} created`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.move RPC method.
//
// The method moves the window identified by the given window ID to the
// given screen position, and reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface MoveInput {
	windowId: number;
	top: number;
	left: number;
}

interface MoveResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerMoveMethod() {
	Rpc.register(
		'windows.move',

		function (input: Rpc.Input): input is MoveInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.top, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.left, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: MoveInput): Promise<MoveResult|Rpc.ExecutionError> {
			try {
				const top = input.top;
				const left = input.left;
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { top, left });
				Logger.write(`Window ${windowId} moved to (${left}, ${top})`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.resize RPC method.
//
// The method resizes the window identified by the given window ID to
// the given dimension, and reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface ResizeInput {
	windowId: number;
	width: number;
	height: number;
}

interface ResizeResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerResizeMethod() {
	Rpc.register(
		'windows.resize',

		function (input: Rpc.Input): input is ResizeInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.width, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.height, Validator.isNumber, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: ResizeInput): Promise<ResizeResult|Rpc.ExecutionError> {
			try {
				const width = input.width;
				const height = input.height;
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { width, height });
				Logger.write(`Window ${windowId} resized to ${width}x${height}`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.minimize RPC method.
//
// The method minimize the window identified by the given window ID, and
// reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface MinimizeInput {
	windowId: number;
}

interface MinimizeResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerMinimizeMethod() {
	Rpc.register(
		'windows.minimize',

		function (input: Rpc.Input): input is MinimizeInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: MinimizeInput): Promise<MinimizeResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { state: 'minimized' });
				Logger.write(`Window ${windowId} minimized`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.maximize RPC method.
//
// The method maximize the window identified by the given window ID, and
// reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface MaximizeInput {
	windowId: number;
}

interface MaximizeResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerMaximizeMethod() {
	Rpc.register(
		'windows.maximize',

		function (input: Rpc.Input): input is MaximizeInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: MaximizeInput): Promise<MaximizeResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { state: 'maximized' });
				Logger.write(`Window ${windowId} maximized`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.fullscreen RPC method.
//
// The method fullscreen the window identified by the given window ID,
// and reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface FullscreenInput {
	windowId: number;
}

interface FullscreenResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerFullscreenMethod() {
	Rpc.register(
		'windows.fullscreen',

		function (input: Rpc.Input): input is FullscreenInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: FullscreenInput): Promise<FullscreenResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { state: 'fullscreen' });
				Logger.write(`Window ${windowId} fullscreened`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.restore RPC method.
//
// The method restore the window identified by the given window ID, and
// reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface RestoreInput {
	windowId: number;
}

interface RestoreResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerRestoreMethod() {
	Rpc.register(
		'windows.restore',

		function (input: Rpc.Input): input is RestoreInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: RestoreInput): Promise<RestoreResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { state: 'normal' });
				Logger.write(`Window ${windowId} restored`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.focus RPC method.
//
// The method focus the window identified by the given window ID, and
// report the result back to the caller. Note that calling this method
// on a window also makes the window the current window, since the
// current window is the one that are focused most recently.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface FocusInput {
	windowId: number;
}

interface FocusResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerFocusMethod() {
	Rpc.register(
		'windows.focus',

		function (input: Rpc.Input): input is FocusInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: FocusInput): Promise<FocusResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { focused: true });
				Logger.write(`Window ${windowId} focused`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.unfocus RPC method.
//
// The method removes the focus on the window identified by the given
// window ID, and reports the result back to the caller.
//
// The method is a simple wrapper over the windows.update Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-update
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/update
//

interface UnfocusInput {
	windowId: number;
}

interface UnfocusResult {
	success: true;
	result: WebExtension.Windows.Window;
}

export function registerUnfocusMethod() {
	Rpc.register(
		'windows.unfocus',

		function (input: Rpc.Input): input is UnfocusInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: UnfocusInput): Promise<UnfocusResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				const result = await WebExtension.windows.update(windowId, { focused: false });
				Logger.write(`Window ${windowId} unfocused`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register windows.remove RPC method.
//
// The method closes the window identified by the given window ID, along
// with all the tabs inside.
//
// The method is a simple wrapper over the windows.remove Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/windows/#method-remove
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/windows/remove
//

interface RemoveInput {
	windowId: number;
}

interface RemoveResult {
	success: true;
}

export function registerRemoveMethod() {
	Rpc.register(
		'windows.remove',

		function (input: Rpc.Input): input is RemoveInput {
			if (Validator.validateType(input.windowId, Validator.isWindowId, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: RemoveInput): Promise<RemoveResult|Rpc.ExecutionError> {
			try {
				const windowId = await Context.ensureNotConsoleWindow(input.windowId);
				await WebExtension.windows.remove(windowId);
				Logger.write(`Window ${windowId} removed`);
				return { success: true };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


