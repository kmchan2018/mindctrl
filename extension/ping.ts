

import * as Rpc from './rpc';


//////////////////////////////////////////////////////////////////////////
//
// Register all RPC methods.
//

export function registerAllMethods() {
	registerPingMethod();
}


//////////////////////////////////////////////////////////////////////////
//
// Register ping RPC method.
//
// The method performs nothing. It is provided as a simple method to test
// connectivity between the clients and the extension.
//

interface PingInput {
	// empty
}

interface PingResult {
	success: true;
}

export function registerPingMethod() {
	Rpc.register(
		'ping',

		function (input: Rpc.Input): input is PingInput {
			return true;
		},

		async function (input: PingInput): Promise<PingResult> {
			return { success: true };
		}
	);
}


