

//////////////////////////////////////////////////////////////////////////
//
// Helper functions.
//

function isFound(selector: string): boolean {
	if (document.querySelector(selector)) {
		return true;
	} else {
		return false;
	}
}

function isMissing(selector: string): boolean {
	if (document.querySelector(selector)) {
		return false;
	} else {
		return true;
	}
}

function checkCondition(ready: boolean, required?: Array<string>, forbidden?: Array<string>): boolean {
	if (ready && document.readyState !== 'complete') {
		return false;
	} else if (required && required.every(isFound) === false) {
		return false;
	} else if (forbidden && forbidden.every(isMissing) === false) {
		return false;
	} else {
		return true;
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Public API.
//

export function Invoke(timeout: number, ready: boolean, required?: Array<string>, forbidden?: Array<string>): Promise<boolean> {
	if (checkCondition(ready, required, forbidden)) {
		return Promise.resolve(true);
	} else {
		let until = Date.now().valueOf() + timeout;
		let timer: number;

		return new Promise((resolve, reject) => {
			function handler() {
				const now = Date.now().valueOf();

				if (checkCondition(ready, required, forbidden) === true) {
					resolve(true);
					clearInterval(timer);
				} else if (now >= until) {
					resolve(false);
					clearInterval(timer);
				}
			}

			timer = window.setInterval(handler, 500);
		});
	}
}


