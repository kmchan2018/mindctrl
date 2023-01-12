

import * as Util from './util';


//////////////////////////////////////////////////////////////////////////
//
// Event channels.
//

const onWriteChannel = Util.createEventChannel<[string,Date]>();


//////////////////////////////////////////////////////////////////////////
//
// Write a log message.
//

export function write(message: string, timestamp?: Date) {
	if (timestamp) {
		onWriteChannel.emit(message, timestamp);
	} else {
		timestamp = new Date();
		onWriteChannel.emit(message, timestamp);
	}
}


//////////////////////////////////////////////////////////////////////////
//
// Listenable events.
//

export const onWrite = onWriteChannel.observer;



