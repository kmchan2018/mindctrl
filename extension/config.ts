

import * as WebExtension from 'webextension-polyfill';


//////////////////////////////////////////////////////////////////////////
//
// Definition of config
//

export interface Config {
	url: string;
	name: string;
	username: string;
	password: string;
}


//////////////////////////////////////////////////////////////////////////
//
// Return validation error for the given url.
//

export function getUrlValidationError(url: string): string {
	try {
		if (url === '' || url.trim() === '') {
			return 'URL cannot be empty';
		} else {
			const parsed = new URL(url);
			const protocol = parsed.protocol;
			const hostname = parsed.hostname;
			const port = parsed.port;

			if (protocol !== 'ws:' && protocol !== 'wss:') {
				return 'Mindctrl can only connect to MQTT server via ws/wss protocols';
			} else if (hostname === '' || hostname.trim() === '') {
				return 'URL should contain host name for the MQTT server';
			} else if (port === '' || port.trim() === '') {
				return 'URL should contain port number for the MQTT server';
			} else {
				return '';
			}
		}
	} catch (err) {
		return 'URL seems to be invalid';
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Return validation error for the given name.
//

export function getNameValidationError(name: string): string {
	if (name === '' || name.trim() === '') {
		return 'Name cannot be empty';
	} else {
		return '';
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Load the config data from the extension storage area.
//

export async function load(): Promise<Config|undefined> {
	//
	// Note that the key to default value map should use null
	// instead of undefined; Chrome will ignore any key that
	// has undefined as its default value; Firefox does not
	// have this problem.
	//

	const data = await WebExtension.storage.local.get({
		version: null,
		url: null,
		name: null,
		username: "",
		password: "",
	});

	if (data.version === 5) {
		const url = data.url as string;
		const name = data.name as string;
		const username = data.username as string;
		const password = data.password as string;
		return { url, name, username, password };
	} else {
		return undefined;
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Save the given configuration data to the extension storage area.
//

export async function save(config: Config) {
	await WebExtension.storage.local.set({
		version: 5,
		url: config.url,
		name: config.name,
		username: config.username,
		password: config.password,
	});
}


