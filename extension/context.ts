

import * as WebExtension from 'webextension-polyfill';


//////////////////////////////////////////////////////////////////////////
//
// Identity of the window and tab the console is running in.
//

let consoleWindow: number | undefined | 'unknown' = 'unknown';
let consoleTab: number | undefined | 'unknown' = 'unknown';


//////////////////////////////////////////////////////////////////////////
//
// Get the window the caller is running in. The function will cache the
// result to avoid unnecessary calls to the web browser.
//

export async function getConsoleWindowId(): Promise<number|undefined> {
	if (consoleWindow !== 'unknown') {
		return consoleWindow;
	} else {
		const result = await WebExtension.windows.getCurrent();

		if (result && result.id) {
			console.log(`context.ts: console window = ${result.id}`);
			consoleWindow = result.id;
			return result.id;
		} else {
			consoleWindow = undefined;
			return undefined;
		}
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Get the window the caller is running in. The function will cache the
// result to avoid unnecessary calls to the web browser.
//

export async function getConsoleTabId(): Promise<number|undefined> {
	if (consoleTab !== 'unknown') {
		return consoleTab;
	} else {
		const result = await WebExtension.tabs.getCurrent();

		if (result && result.id) {
			console.log(`context.ts: console tab = ${result.id}`);
			consoleTab = result.id;
			return result.id;
		} else {
			consoleTab = undefined;
			return undefined;
		}
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Check if the given window id identifies the window the caller is
// running in.
//

export async function isConsoleWindow(windowId: number): Promise<boolean> {
	return (windowId === await getConsoleWindowId());
}


//////////////////////////////////////////////////////////////////////////
//
// Check if the given tab id identifies the tab the caller is running
// in.
//

export async function isConsoleTab(tabId: number): Promise<boolean> {
	return (tabId === await getConsoleTabId());
}


//////////////////////////////////////////////////////////////////////////
//
// Check if the given window id identifies the window the caller is
// running in.
//

export async function ensureNotConsoleWindow(windowId: number): Promise<number> {
	if (windowId !== await getConsoleWindowId()) {
		return windowId;
	} else {
		throw new Error('Cannot operate on console window');
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Check if the given tab id identifies the tab the caller is running
// in.
//

export async function ensureNotConsoleTab(tabId: number): Promise<number> {
	if (tabId !== await getConsoleTabId()) {
		return tabId;
	} else {
		throw new Error('Cannot operate on console tab');
	}
}


//////////////////////////////////////////////////////////////////////////
//
// List of normal windows ordered by the time they receive focus. The
// most recently focused window will be at the start.
//

const windowStack: Array<number> = [];


//////////////////////////////////////////////////////////////////////////
//
// Return the last normal window that has received focus.
//

export async function getLastFocusedWindow(): Promise<number> {
	if (windowStack.length == 0) {
		WebExtension.windows.onCreated.addListener(function(win) {
			if (win.id !== undefined && win.type === 'normal') {
				if (win.focused) {
					windowStack.unshift(win.id);
				} else {
					windowStack.push(win.id);
				}
			}
			console.log("on_created", win, windowStack);
		});

		WebExtension.windows.onRemoved.addListener(function(id) {
			for (let i = 0; i < windowStack.length; i++) {
				if (windowStack[i] === id) {
					windowStack.splice(i, 1);
					console.log("on_removed", id, windowStack);
					return;
				}
			}
		});

		WebExtension.windows.onFocusChanged.addListener(function(id) {
			for (let i = 1; i < windowStack.length; i++) {
				if (windowStack[i] === id) {
					windowStack.splice(i, 1);
					windowStack.unshift(id);
					console.log("on_focus_changed", id, windowStack);
					return;
				}
			}
		});

		for (const win of await WebExtension.windows.getAll({ populate: false })) {
			if (win.id !== undefined && win.type === 'normal') {
				if (win.focused) {
					windowStack.unshift(win.id);
				} else {
					windowStack.push(win.id);
				}
			}
		}
	}

	return windowStack[0];
}


