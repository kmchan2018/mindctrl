

//////////////////////////////////////////////////////////////////////////
//
// Event channel implements an event API that resembles the API provided
// by Web Extension.
//

type EventCallback<P extends any[]> = (...args: P) => void;

interface EventObserver<P extends any[]> {
	hasListener: (callback: EventCallback<P>) => boolean;
	addListener: (callback: EventCallback<P>) => void;
	removeListener: (callback: EventCallback<P>) => void;
}

interface EventChannel<P extends any[]> {
	emit: (...args: P) => void;
	observer: EventObserver<P>;
}

export function createEventChannel<P extends any[]>(): EventChannel<P> {
	const callbacks = [] as Array<EventCallback<P>>;

	return Object.freeze({
		emit: function(...args: P) {
			for (const callback of callbacks) {
				callback(...args);
			}
		},

		observer: Object.freeze({
			hasListener: function(callback: EventCallback<P>): boolean {
				return (callbacks.indexOf(callback) >= 0);
			},

			addListener: function(callback: EventCallback<P>): void {
				if (callbacks.indexOf(callback) === -1) {
					callbacks.push(callback);
				}
			},

			removeListener: function(callback: EventCallback<P>): void {
				const index = callbacks.indexOf(callback);

				if (index >= 0) {
					callbacks.splice(index, 1);
				}
			},
		}),
	});
}


/*
//////////////////////////////////////////////////////////////////////////
//
// Convert error into user-visible message.
//

export function toMessage(error: any): string {
	if (typeof error === 'object' && error instanceof Error) {
		return error.message;
	} else if (typeof error === 'string') {
		return error;
	} else if (typeof error === 'undefined') {
		return 'unidentified error';
	} else if (typeof error === 'object' && error === null) {
		return 'unidentified error';
	} else {
		return `unexpected error ${error.toString()}`;
	}
}
*/


//////////////////////////////////////////////////////////////////////////
//
// Wait until some time has elapsed. Note that the method is implemented
// using setTimeout and therefore may be throttled by the browser.
//

export function waitDuration(duration: number): Promise<void> {
	return new Promise<void>((resolve, reject) => {
		setTimeout(resolve, duration);
	});
}


//////////////////////////////////////////////////////////////////////////
//
// Wait until the document is ready.
//

export function waitDocumentLoaded(): Promise<void> {
	return new Promise<void>((resolve, reject) => {
		if (document.readyState !== 'complete') {
			document.addEventListener('readystatechange', function listen() {
				if (document.readyState === 'complete') {
					document.removeEventListener('readystatechange', listen);
					resolve();
				}
			});
		} else {
			resolve();
		}
	});
}


