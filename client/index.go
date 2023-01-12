// Package mindctrl provides a simple library to control Mindctrl
// web extension ("server"). This is achieved by providing a
// fluent API to execute rmote calls to the server.
//
// Operation Types
//
// Each supported method is handled by its own operation type, like
// PingOperation for "ping" method, FindDownloadsOperation for
// "downloads.find" method, etc.
//
// Arguments
//
// The create function of operation types demands all required
// arguments of the method as its parameters.
//
// Then, all method arguments can be updated by setters. For
// required arguments, the setters demand the new value as its
// sole parameter. For optional arguments, the setter takes
// two parameters - the first one controls if the argument is
// specified or not, and the second one the actual value of the
// argumenmt. All setter returns the operation itself.
//
// Finally, all method arguments can be retrieved by the respective
// getters.
//
// Execution
//
// Each operation types provide several methods to execute the
// operation.
//
// The "Execute" function executes the call over the given transport
// synchronously. It will wait until the call is finished and return
// the result.
//
// The "Start" function executes the call over the given transport
// asynchronously, and invoke the provided callback after the call
// is finished. The end result can be retrieved by the "Result"
// function.
//
// The "StartChannel" function executes the call over the given
// transport asynchronously, and writes to the given channel after
// the call is finished. The end result can be retrieved by the
// "Result" function.
//
// For the asynchronous methods, progress checks and post-finish
// actions are handled by the [Transport.Dispatch] function.
//
package mindctrl
