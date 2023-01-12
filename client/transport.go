package mindctrl

import (
	"github.com/kmchan2018/mindctrl/client/codec"
	"net/rpc"
)

// Options contains additional data for the transport.
//
type Options = codec.Options

// Callback is function that is invoked after an async call has
// finished.
//
type Callback = func(method string, args interface{}, reply interface{}, err error)

// Transport handles communication with the server. Note that the
// transport implementation is NOT thread-safe and therefore
// should not be used in multiple goroutines.
//
type Transport struct {
	client  *rpc.Client
	channel chan *rpc.Call
	pending map[*rpc.Call]Callback
}

// Create a new transport with the given options.
//
func NewTransport(url string, client string, server string, options *Options) (*Transport, error) {
	if c, err := codec.NewCodec(url, client, server, options); err != nil {
		return nil, err
	} else {
		return &Transport{
			client:  rpc.NewClientWithCodec(c),
			channel: make(chan *rpc.Call, 100),
			pending: make(map[*rpc.Call]Callback),
		}, nil
	}
}

// Call a remote method synchronously.
//
func (transport *Transport) call(method string, args interface{}, reply interface{}) error {
	return transport.client.Call(method, args, reply)
}

// Call a remote method asynchronously.
//
func (transport *Transport) start(method string, args interface{}, reply interface{}, callback Callback) {
	call := transport.client.Go(method, args, reply, transport.channel)
	transport.pending[call] = callback
}

// Watch for a single completed asynchronous call and invoke the
// corresponding callback function. Return true if there are still
// outstanding async calls to watch, and false otherwise.
//
func (transport *Transport) Dispatch() bool {
	for {
		if remainder := len(transport.pending); remainder > 0 {
			call := <-transport.channel

			if callback, found := transport.pending[call]; found == false {
				continue
			} else {
				callback(call.ServiceMethod, call.Args, call.Reply, call.Error)
				delete(transport.pending, call)
				return remainder > 1
			}
		} else {
			return false
		}
	}
}
