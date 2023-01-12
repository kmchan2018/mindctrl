

import * as WebExtension from 'webextension-polyfill';


//
// Get the extension manifest. Since manifest v3 has a slightly
// different API than manifest v2, the script needs to determine
// the manifest version and uses the appropriate API.
//

const manifest = WebExtension.runtime.getManifest();


//
// Get the URL of the frontend page.
//

const url = WebExtension.runtime.getURL('frontend.html');


//
// Bring an existing frontend dialog to the front or open a new one
// when the browser action is triggered.
//

(function() {
	async function onActionClicked() {
		const tabs = await WebExtension.tabs.query({ url });

		if (tabs.length > 0 && tabs[0].windowId !== undefined) {
			await WebExtension.windows.update(tabs[0].windowId, {
				//state: 'normal',
				drawAttention: true,
			});
		} else {
			await WebExtension.windows.create({
				type: 'panel',
				url: url,
				width: 480,
				height: 800,
			});
		}
	}

	if (manifest["manifest_version"] == 2) {
		WebExtension.browserAction.setTitle({ title: "Open Mindctrl Console" });
		WebExtension.browserAction.onClicked.addListener(onActionClicked);
	} else {
		WebExtension.action.setTitle({ title: "Open Mindctrl Console" });
		WebExtension.action.onClicked.addListener(onActionClicked);
	}
})();


//
// Erase browser history for the frontend page. Users do not expect
// these entries in their browser history, so we need to clean them
// up.
//

WebExtension.history.onVisited.addListener(function(item: WebExtension.History.HistoryItem) {
	if (url === item.url) {
		console.log(`Removing unwanted history record(s) for url ${url}`);
		WebExtension.history.deleteUrl({ url });
	}
});


