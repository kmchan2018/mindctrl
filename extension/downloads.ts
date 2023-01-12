

import * as WebExtension from 'webextension-polyfill';

import * as Browser from './browser';
import * as Logger from './logger';
import * as Pattern from './pattern';
import * as Rpc from './rpc';
import * as Util from './util';
import * as Validator from './validator';


//////////////////////////////////////////////////////////////////////////
//
// Enable all RPC methods.
//

export function registerAllMethods() {
	registerFindMethod();
	registerGetMethod();
	registerCreateMethod();
	registerPauseMethod();
	registerResumeMethod();
	registerCancelMethod();
	registerRemoveMethod();
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.find RPC method.
//
// The method searches the browser for downloads that matches the given
// criteria and reports the matches back to the caller.
//
// The url option expects a match pattern for the download url. If it is
// specified, only downloads whose url matches the pattern will be
// included. Otherwise, no downloads will be excluded.
//
// Moreover, the state option expects a download state. If it is given,
// only downloads in the given state are included. Otherwise, no download
// will be excluded.
//
// The method is a simple wrapper over the downloads.search Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-search
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/search
//

interface FindInput {
	url?: string;
	state?: WebExtension.Downloads.State;
}

interface FindResult {
	success: true;
	result: Array<WebExtension.Downloads.DownloadItem>;
}

export function registerFindMethod() {
	Rpc.register(
		'downloads.find',

		function (input: Rpc.Input): input is FindInput {
			if (Validator.validateType(input.url, Validator.isMatchPattern, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.state, Validator.isDownloadState, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: FindInput): Promise<FindResult|Rpc.ExecutionError> {
			try {
				const state = input.state;
				const urlRegex = (input.url ? Pattern.convertMatchPattern(input.url) : undefined);
				const result = await WebExtension.downloads.search({ urlRegex, state });
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.get RPC method.
//
// The method retrieve information on the download identified by the
// given download ID and reports the information back to the caller.
//
// The method is a simple wrapper over the downloads.search Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-search
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/search
//

interface GetInput {
	downloadId: number;
}

interface GetResult {
	success: true;
	result: WebExtension.Downloads.DownloadItem;
}

export function registerGetMethod() {
	Rpc.register(
		'downloads.get',

		function (input: Rpc.Input): input is GetInput {
			if (Validator.validateType(input.downloadId, Validator.isDownloadId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: GetInput): Promise<GetResult|Rpc.ExecutionError> {
			try {
				const downloadId = input.downloadId;
				const result = await getDownload(downloadId, true);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.create RPC method.
//
// The method initiates a new download that fetches the given URL to the
// given file under the downloads folder, and reports information on the
// created download back to the caller.
//
// The referrer option controls the referrer header used for the http
// download requests. If it is not specified, no referrer header will
// be included. Note that the option is not supported on Chrome based
// browsers.
//
// Moreover, the noWait option determines when the method will return.
// If the option is false or unspecified, the method will return until
// the download reaches a stable state. If the option is true, the
// method will return early, right after the download has started.
//
// The operation is a simple wrapper over the downloads.create Web
// Extension API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-download
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/download
//

interface CreateInput {
	url: string;
	filename: string;
	referrer?: string;
	noWait?: boolean;
}

interface CreateResult {
	success: true;
	result: WebExtension.Downloads.DownloadItem;
}

export function registerCreateMethod() {
	Rpc.register(
		'downloads.create',

		function (input: Rpc.Input): input is CreateInput {
			if (Validator.validateType(input.url, Validator.isUrl) === false) {
				return false;
			} else if (Validator.validateType(input.filename, Validator.isFilename) === false) {
				return false;
			} else if (Validator.validateType(input.referrer, Validator.isUrl, Validator.isUndefined) === false) {
				return false;
			} else if (Validator.validateType(input.noWait, Validator.isBoolean, Validator.isUndefined) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: CreateInput): Promise<CreateResult|Rpc.ExecutionError> {
			try {
				const url = input.url;
				const filename = input.filename;
				const headers = [] as Array<WebExtension.Downloads.DownloadOptionsTypeHeadersItemType>;
				const conflictAction = 'uniquify' as WebExtension.Downloads.FilenameConflictAction;
				const options = { url, filename, conflictAction, headers } as WebExtension.Downloads.DownloadOptionsType;
				const noWait = input.noWait || false;

				// There are two differences between Firefox and Chrome regarding the
				// behavior of the downloads.create method.
				//
				// First of all, Firefox allows referer to be set via headers while Chrome
				// does not. In fact, if the referer header is specified in the header
				// argument, Chrome will throw an error indicating that forbidden header
				// is given.
				// 
				// Next, the meaning of saveAs key is different between Firefox and Chrome.
				// For Firefox, the value of the key determines whether the save dialog is
				// shown or not. For Chrome, the existence of the key determine whether
				// the save dialog is shown or not.

				if (await Browser.isFirefox()) {
					options.saveAs = false;

					if (typeof input.referrer === 'string') {
						const name = 'Referer';
						const value = input.referrer;
						headers.push({ name, value });
					}
				}

				const downloadId = await WebExtension.downloads.download(options);
				const result = await getDownload(downloadId, noWait);
				Logger.write(`Download ${downloadId} created to save to file ${filename}`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.pause RPC method.
//
// The method pauses the download identified by the given download ID,
// and reports updated information on the paused download back to the
// caller.
//
// The method is a simple wrapper over the downloads.pause Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-pause
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/pause
//

interface PauseInput {
	downloadId: number;
}

interface PauseResult {
	success: true;
	result: WebExtension.Downloads.DownloadItem;
}

export function registerPauseMethod() {
	Rpc.register(
		'downloads.pause',

		function (input: Rpc.Input): input is PauseInput {
			if (Validator.validateType(input.downloadId, Validator.isDownloadId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: PauseInput): Promise<PauseResult|Rpc.ExecutionError> {
			try {
				const downloadId = input.downloadId;
				const result = await WebExtension.downloads.pause(downloadId).then(() => getDownload(downloadId, true));
				Logger.write(`Download ${downloadId} paused`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.resume RPC method.
//
// The method resumes the download identified by the given download ID,
// and reports updated information on the resumed download back to the
// caller.
//
// The method is a simple wrapper over the downloads.resume Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-resume
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/resume
//

interface ResumeInput {
	downloadId: number;
	noWait?: boolean;
}

interface ResumeResult {
	success: true;
	result: WebExtension.Downloads.DownloadItem;
}

export function registerResumeMethod() {
	Rpc.register(
		'downloads.resume',

		function (input: Rpc.Input): input is ResumeInput {
			if (Validator.validateType(input.downloadId, Validator.isDownloadId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: ResumeInput): Promise<ResumeResult|Rpc.ExecutionError> {
			try {
				const downloadId = input.downloadId;
				const noWait = input.noWait || false;
				const result = await WebExtension.downloads.resume(downloadId).then(() => getDownload(downloadId, noWait));
				Logger.write(`Download ${downloadId} resumed`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.cancel RPC method.
//
// The method cancels the download identified by the given download ID,
// and reports updated information on the cancelled download back to the
// caller.
//
// The method is a simple wrapper over the downloads.cancel Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-cancel
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/cancel
//

interface CancelInput {
	downloadId: number;
}

interface CancelResult {
	success: true;
	result: WebExtension.Downloads.DownloadItem;
}

export function registerCancelMethod() {
	Rpc.register(
		'downloads.cancel',

		function (input: Rpc.Input): input is CancelInput {
			if (Validator.validateType(input.downloadId, Validator.isDownloadId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: CancelInput): Promise<CancelResult|Rpc.ExecutionError> {
			try {
				const downloadId = input.downloadId;
				const result = await WebExtension.downloads.cancel(downloadId).then(() => getDownload(downloadId, true));
				Logger.write(`Download ${downloadId} cancelled`);
				return { success: true, result };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Enable downloads.remove RPC method.
//
// The method removes the download identified by the given download ID.
// Note that the method removes the download from browser history only;
// It will NOT remove the downloaded file from the disk.
//
// The method is a simple wrapper over the downloads.erase Web Extension
// API. Details on the API can be found in:
//
// https://developer.chrome.com/docs/extensions/reference/downloads/#method-erase
// https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/API/downloads/erase
//

interface RemoveInput {
	downloadId: number;
}

interface RemoveResult {
	success: true;
}

export function registerRemoveMethod() {
	Rpc.register(
		'downloads.remove',

		function (input: Rpc.Input): input is RemoveInput {
			if (Validator.validateType(input.downloadId, Validator.isDownloadId) === false) {
				return false;
			} else {
				return true;
			}
		},

		async function (input: RemoveInput): Promise<RemoveResult|Rpc.ExecutionError> {
			try {
				const downloadId = input.downloadId;
				await WebExtension.downloads.erase({ id: downloadId });
				Logger.write(`Download ${downloadId} removed`);
				return { success: true };
			} catch (error) {
				return Rpc.createExecutionError(error);
			}
		}
	);
}


//////////////////////////////////////////////////////////////////////////
//
// Helper functions to get a single download from its ID. The function
// can optionally wait until the download is stablized (aka the download
// is completed or interrupted) before returning.
//

function getDownload(downloadId: number, noWait: boolean): Promise<WebExtension.Downloads.DownloadItem> {
	if (noWait) {
		return WebExtension.downloads.search({ id: downloadId }).then(function(results) {
			if (results.length == 0) {
				throw new Error(`unknown download id ${downloadId}`);
			} else if (results[0].id !== downloadId) {
				throw new Error(`unknown download id ${downloadId}`);
			} else {
				return results[0];
			}
		});
	}

	return new Promise<WebExtension.Downloads.DownloadItem>((resolve, reject) => {
		function handleRetrieved(downloads: Array<WebExtension.Downloads.DownloadItem>) {
			if (downloads.length === 0) {
				reject(new Error(`unknown download id ${downloadId}`));
			} else if (downloads[0].id !== downloadId) {
				reject(new Error(`unknown download id ${downloadId}`));
			} else {
				if (downloads[0].state === 'interrupted') {
					resolve(downloads[0]);
				} else if (downloads[0].state === 'complete') {
					resolve(downloads[0]);
				} else if (noWait) {
					resolve(downloads[0]);
				}
			}
		}

		function handleChanged(delta: WebExtension.Downloads.OnChangedDownloadDeltaType) {
			if (delta.id === downloadId) {
				if (delta.state !== undefined) {
					if (delta.state.current === 'interrupted' || delta.state.current === 'complete') {
						WebExtension.downloads.search({ id: downloadId }).then(handleRetrieved, handleFailed);
					}
				}
			}
		}

		function handleErased(id: number) {
			if (id === downloadId) {
				WebExtension.downloads.onChanged.removeListener(handleChanged);
				WebExtension.downloads.onErased.removeListener(handleErased);
				reject(new Error(`unknown download id ${downloadId}`));
			}
		}

		function handleFailed(error: any) {
			WebExtension.downloads.onChanged.removeListener(handleChanged);
			WebExtension.downloads.onErased.removeListener(handleErased);
			reject(error);
		}

		if (noWait) {
			WebExtension.downloads.search({ id: downloadId }).then(handleRetrieved, handleFailed);
		} else {
			WebExtension.downloads.onChanged.addListener(handleChanged);
			WebExtension.downloads.onErased.addListener(handleErased);
			WebExtension.downloads.search({ id: downloadId }).then(handleRetrieved, handleFailed);
		}
	});
}


