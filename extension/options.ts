

import * as Config from './config';
import * as Util from './util';


//////////////////////////////////////////////////////////////////////////
//
// Handles the options page. It handles opening the options page,
// updating the form with current preferences, validating form values
// and saving the preferences.
//

(async function() {
	await Util.waitDocumentLoaded();

	const formElement = document.querySelector<HTMLFormElement>('#options')!;
	const urlElement = document.querySelector<HTMLInputElement>('#options input[name="url"]')!;
	const nameElement = document.querySelector<HTMLInputElement>('#options input[name="name"]')!;
	const usernameElement = document.querySelector<HTMLInputElement>('#options input[name="username"]')!;
	const passwordElement = document.querySelector<HTMLInputElement>('#options input[name="password"]')!;
	const reloadElement = document.querySelector<HTMLButtonElement>('#options button[name="reload"]')!;

	urlElement.addEventListener('change', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();
		urlElement.setCustomValidity(Config.getUrlValidationError(urlElement.value));
		urlElement.reportValidity();
	});

	nameElement.addEventListener('change', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();
		nameElement.setCustomValidity(Config.getNameValidationError(nameElement.value));
		nameElement.reportValidity();
	});

	reloadElement.addEventListener('click', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();

		Config.load().then(function(config) {
			if (config !== undefined) {
				urlElement.setAttribute('value', config.url);
				nameElement.setAttribute('value', config.name);
				usernameElement.setAttribute('value', config.username || '');
				passwordElement.setAttribute('value', config.password || '');
				formElement.reset();
			} else {
				urlElement.setAttribute('value', '');
				nameElement.setAttribute('value', '');
				usernameElement.setAttribute('value', '');
				passwordElement.setAttribute('value', '');
				formElement.reset();
			}
		});
	});

	formElement.addEventListener('submit', function(ev: Event) {
		ev.preventDefault();
		ev.stopPropagation();

		const url = urlElement.value;
		const name = nameElement.value;
		const username = usernameElement.value;
		const password = passwordElement.value;

		Config.save({ url, name, username, password }).then(function() {
			urlElement.setAttribute('value', url);
			nameElement.setAttribute('value', name);
			usernameElement.setAttribute('value', username);
			passwordElement.setAttribute('value', password);
		});
	});

	formElement.addEventListener('reset', function(ev: Event) {
		urlElement.setCustomValidity(Config.getUrlValidationError(urlElement.value));
		urlElement.reportValidity();
		nameElement.setCustomValidity(Config.getNameValidationError(nameElement.value));
		nameElement.reportValidity();
	});

	Config.load().then(function(config) {
		if (config !== undefined) {
			urlElement.setAttribute('value', config.url);
			nameElement.setAttribute('value', config.name);
			usernameElement.setAttribute('value', config.username || '');
			passwordElement.setAttribute('value', config.password || '');
			formElement.reset();
		} else {
			urlElement.setAttribute('value', '');
			nameElement.setAttribute('value', '');
			usernameElement.setAttribute('value', '');
			passwordElement.setAttribute('value', '');
			formElement.reset();
		}
	});
})();



