package protocol

import (
	"encoding/json"
)

// Request packet defines the shape of the on-the-wire request data.
// The fields are mostly self explanatory, but some needs further
// explanation:
//
// 1. The Type field should always contain the string "request". It
// identifies the type of packet in log messages and network
// captures.
//
// 2. The Id field stores the identifier of the request. It equals to
// the sequence number assigned by the RPC client.
//
// 3. The Input field is called "params" on the wire to make the
// message JSON-RPC compliant.
//
type RequestPacket struct {
	Type   string      `json:"type"`   // type of the packet; always "request"
	Id     string      `json:"id"`     // unique id of the call
	Method string      `json:"method"` // method to be called
	Client string      `json:"client"` // client who makes the call
	Server string      `json:"server"` // server who executes the call
	Input  interface{} `json:"params"` // input of the call
}

// Response packet defines the shape of the on-the-wire response data.
// The fields are mostly self explanatory, but some needs further
// explanation:
//
// 1. The Type field should always contain the string "response". It
// identifies the type of packet in log messages and network
// captures.
//
// 2. The Id field stores the identifier of the request. It equals to
// the sequence number assigned by the RPC client.
//
// 3. The Output field is called "result" on the wire to make the
// message JSON-RPC compliant.
//
type ResponsePacket struct {
	Type   string          `json:"type"`   // type of the packet; always "response"
	Id     string          `json:"id"`     // unique id of the call
	Method string          `json:"method"` // method to be called
	Client string          `json:"client"` // client who makes the call
	Server string          `json:"server"` // server who executes the call
	Output json.RawMessage `json:"result"` // output of the call
}
