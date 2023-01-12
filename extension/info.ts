

import * as WebExtension from 'webextension-polyfill';

import * as Browser from './browser';
import * as Rpc from './rpc';


//////////////////////////////////////////////////////////////////////////
//
// Register all RPC methods.
//

export function registerAllMethods() {
	registerGetBrowserMethod();
	registerGetPlatformMethod();
}


//////////////////////////////////////////////////////////////////////////
//
// Register info.get_browser RPC method.
//
// The method detects the name and version of the browser running the
// extension and reports the result back to the caller.
//
// The method is a simple wrapper over the runtime.getBrowserInfo Web
// Extension API on Firefox. Details on the API can be found in:
//
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/runtime/getBrowserInfo
//
// However, other browsers currently do not implement this API, and the
// extension will emulate the API by checking the user agent string or
// client hints.
//

interface GetBrowserInput {
	// empty
}

interface GetBrowserResult {
	success: true;
	result: Browser.BrowserInfo;
}

export function registerGetBrowserMethod() {
	Rpc.register<GetBrowserInput,GetBrowserResult>(
		'info.get_browser',

		function (input: Rpc.Input): input is GetBrowserInput {
			return true;
		},

		async function (input: GetBrowserInput): Promise<GetBrowserResult> {
			const result = await Browser.detectBrowser();
			return { success: true, result };
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Register info.get_platform RPC method.
//
// The method detects the operating system and processor architecture
// the host browser is running on, and reports the result back to the
// caller.
//
// The method is a simple wrapper over the runtime.getPlatformInfo Web
// Extension API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/runtime/#method-getPlatformInfo
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/runtime/getPlatformInfo
//

interface GetPlatformInput {
	// empty
}

interface GetPlatformResult {
	success: true;
	result: WebExtension.Runtime.PlatformInfo;
}

export function registerGetPlatformMethod() {
	Rpc.register(
		'info.get_platform',

		function (input: Rpc.Input): input is GetPlatformInput {
			return true;
		},

		async function (input: GetPlatformInput): Promise<GetPlatformResult> {
			const result = await WebExtension.runtime.getPlatformInfo();
			return { success: true, result };
		}
	);
}


