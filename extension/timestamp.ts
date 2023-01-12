

//////////////////////////////////////////////////////////////////////////
//
// Helper function to convert a timestamp to a brief string.
//

export function formatShortTimestamp(timestamp: Date): string {
	const hour   = timestamp.getHours().toString(10).padStart(2, '0');
	const minute = timestamp.getMinutes().toString(10).padStart(2, '0');
	return `${hour}:${minute}`;
}


//////////////////////////////////////////////////////////////////////////
//
// Helper function to convert a timestamp to a complete string.
//

export function formatFullTimestamp(timestamp: Date): string {
	const year   = timestamp.getFullYear().toString(10).padStart(4, '0');
	const month  = (timestamp.getMonth() + 1).toString(10).padStart(2, '0');
	const date   = timestamp.getDate().toString(10).padStart(2, '0');
	const hour   = timestamp.getHours().toString(10).padStart(2, '0');
	const minute = timestamp.getMinutes().toString(10).padStart(2, '0');
	const second = timestamp.getSeconds().toString(10).padStart(2, '0');
	return `${year}:${month}:${date} ${hour}:${minute}:${second}`;
}


