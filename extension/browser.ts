

import * as WebExtension from 'webextension-polyfill';


//////////////////////////////////////////////////////////////////////////
//
// Data types for browser name and identity.
//

export type BrowserName = 
	'chrome' |        // Google Chrome and Chromium
	'edge' |          // Microsoft Edge
	'edge-legacy' |   // Microsoft Edge (Legacy version)
	'firefox' |       // Mozilla Firefox
	'opera' |         // Opera
	'safari' |        // Apple Safari
	'vivaldi' |       // Vivaldi
	'unknown'

export interface BrowserInfo {
	name: BrowserName;
	version?: string;
}


//////////////////////////////////////////////////////////////////////////
//
// Data used for browser detection.
//

interface Matcher {
	pattern: RegExp;
	name: BrowserName;
	priority: number;
}

const hintMatchers: Array<Matcher> = [
	{ pattern: /Chromium/i,   name: 'chrome',    priority: 2 }, // many chromium based browser pretend to be chrome, so lower priority
	{ pattern: /Chrome/i,     name: 'chrome',    priority: 2 }, // many chromium based browser pretend to be chrome, so lower priority
	{ pattern: /Edge/i,       name: 'edge',      priority: 4 }, // https://learn.microsoft.com/en-us/microsoft-edge/web-platform/user-agent-guidance
];

const useragentMatchers: Array<Matcher> = [
	{ pattern: /Chrome\/(\d+(?:\.\d+)*)/i,     name: 'chrome',        priority: 2 }, // many chromium based browser pretend to be chrome, so lower priority
	{ pattern: /Chromium\/(\d+(?:\.\d+)*)/i,   name: 'chrome',        priority: 2 },
	{ pattern: /Edg\/(\d+)/i,                  name: 'edge',          priority: 4 }, // https://learn.microsoft.com/en-us/microsoft-edge/web-platform/user-agent-guidance
	{ pattern: /Edge\/(\d+)/i,                 name: 'edge-legacy',   priority: 4 }, // https://learn.microsoft.com/en-us/microsoft-edge/web-platform/user-agent-guidance
	{ pattern: /Firefox\/(\d+(?:\.\d+)*)/i,    name: 'firefox',       priority: 4 },
	{ pattern: /OPR\/(\d+(?:\.\d+)*)/i,        name: 'opera',         priority: 4 },
	{ pattern: /Safari\/(\d+(?:\.\d+)*)/i,     name: 'safari',        priority: 1 }, // even chrome pretend to be Safari, so lowest priority
	{ pattern: /Vivaldi\/(\d+(?:\.\d+)*)/i,    name: 'vivaldi',       priority: 4 },
];


//////////////////////////////////////////////////////////////////////////
//
// Cache for the browser detection result.
//

let cache: BrowserInfo | undefined = undefined;


//////////////////////////////////////////////////////////////////////////
//
// Identifies the browser running the extension.
//

export async function detectBrowser(): Promise<BrowserInfo> {
	if (cache) {
		return cache;
	}

	const info = await (async function(): Promise<BrowserInfo> {
		try {
			// @ts-ignore
			if (typeof browser !== 'undefined' && browser.runtime && browser.runtime.getBrowserInfo) {
				const raw = await WebExtension.runtime.getBrowserInfo();
				const rawName = raw.name.toLowerCase();
				const rawVersion = raw.version;
            
				if (rawName.indexOf('firefox') >= 0) {
					return { name: 'firefox', version: rawVersion };
				}
  	  }
    
			let name = 'unknown' as BrowserName;
			let version = undefined;
			let priority = 0;
    
	    if (typeof navigator !== 'undefined') {
				if (typeof navigator.userAgentData !== 'undefined') {
					for (const brand of navigator.userAgentData.brands) {
						for (const matcher of hintMatchers) {
							if (matcher.pattern.test(brand.brand) && matcher.priority > priority) {
								name = matcher.name;
								version = brand.version;
								priority = matcher.priority;
							}
						}
					}
				}

				if (typeof navigator.userAgent !== 'undefined') {
					for (const matcher of useragentMatchers) {
						const match = matcher.pattern.exec(navigator.userAgent);

						if (match && matcher.priority > priority) {
							name = matcher.name;
							version = match[1];
							priority = matcher.priority;
						}
					}
				}
			}

			if (version) {
				return { name, version };
			} else {
				return { name };
			}
		} catch (err) {
			return { name: 'unknown' };
		}
	})();

	console.log("Result of browser detection: ", info);

	cache = info;
	return info;
}


//////////////////////////////////////////////////////////////////////////
//
// Returns if the current browser is Google Chrome.
//

export async function isChrome(): Promise<boolean> {
	const info = await detectBrowser();
	return (info.name === 'chrome');
}


//////////////////////////////////////////////////////////////////////////
//
// Returns if the current browser is Microsoft Edge.
//

export async function isEdge(): Promise<boolean> {
	const info = await detectBrowser();
	return (info.name === 'edge');
}


//////////////////////////////////////////////////////////////////////////
//
// Returns if the current browser is Microsoft Edge Legacy.
//

export async function isEdgeLegacy(): Promise<boolean> {
	const info = await detectBrowser();
	return (info.name === 'edge-legacy');
}


//////////////////////////////////////////////////////////////////////////
//
// Returns if the current browser is Mozilla Firefox.
//

export async function isFirefox(): Promise<boolean> {
	const info = await detectBrowser();
	return (info.name === 'firefox');
}


//////////////////////////////////////////////////////////////////////////
//
// Returns if the current browser is Opera.
//

export async function isOpera(): Promise<boolean> {
	const info = await detectBrowser();
	return (info.name === 'opera');
}


//////////////////////////////////////////////////////////////////////////
//
// Returns if the current browser is Apple Safari.
//

export async function isSafari(): Promise<boolean> {
	const info = await detectBrowser();
	return (info.name === 'safari');
}


