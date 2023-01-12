package codec

import (
	"context"
	"encoding/json"
	"github.com/eclipse/paho.golang/paho"
	"github.com/kmchan2018/mindctrl/client/protocol"
	"net/rpc"
	"nhooyr.io/websocket"
	"strconv"
	"time"
)

// Implementation of net/rpc client codec for communicating with
// the mindctrl browser extension ("server") via websocket
// connection to an intermediate MQTT broker.
//
type Codec struct {
	ctx         context.Context
	mqtt        *paho.Client
	name        string
	server      string
	channel     chan *paho.Publish
	response    protocol.ResponsePacket
	invalidated bool
}

// Create a new net/rpc client codec by connecting to the MQTT
// broker, subscribing to the relevant topics and watching for
// alive message from the server.
//
func NewCodec(url string, name string, server string, options *Options) (*Codec, error) {
	ctx := context.TODO()

	socketOptions := &websocket.DialOptions{
		Subprotocols: []string{"mqtt"},
	}

	if socket, _, err := websocket.Dial(ctx, url, socketOptions); err != nil {
		return nil, err
	} else {
		socket.SetReadLimit(options.getWebsocketFrameSize())

		connection := websocket.NetConn(ctx, socket, websocket.MessageBinary)
		channel := make(chan *paho.Publish, options.getMqttMessageBuffer())
		topic1 := protocol.GetClientTopic(name)
		topic2 := protocol.GetServerStatusTopic(server)

		mqtt := paho.NewClient(paho.ClientConfig{
			Conn: connection,
			Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
				channel <- m
			}),
		})

		connectPacket := &paho.Connect{
			KeepAlive:    0,
			ClientID:     protocol.GenerateMqttClientId(name),
			CleanStart:   true,
			UsernameFlag: options.getPahoUsernameFlag(),
			PasswordFlag: options.getPahoPasswordFlag(),
			Username:     options.getPahoUsername(),
			Password:     options.getPahoPassword(),
		}

		subscribePacket := &paho.Subscribe{
			Subscriptions: map[string]paho.SubscribeOptions{
				topic1: {},
				topic2: {},
			},
		}

		if ack, err := mqtt.Connect(ctx, connectPacket); err != nil {
			socket.Close(websocket.StatusNormalClosure, "connection error")
			return nil, err
		} else if ack.ReasonCode != 0 {
			socket.Close(websocket.StatusNormalClosure, "connection error")
			return nil, err
		}

		if _, err := mqtt.Subscribe(ctx, subscribePacket); err != nil {
			mqtt.Disconnect(&paho.Disconnect{ReasonCode: 0})
			return nil, err
		}

		// The server will post a retained message with text "alive"
		// to its status topic when it starts. It will also post a
		// retained message with text "dead" to its status topic when
		// it quits.
		//
		// Therefore, the server status can be checked by subscribing
		// to the status topic of the server and waiting for the alive
		// message. If one is received, the server is alive. Otherwise,
		// the server is dead.
		//
		// The timer exists to cap the waiting time to 5 seconds if
		// no status message is coming.

		codec := &Codec{ctx: ctx, mqtt: mqtt, name: name, server: server, channel: channel, invalidated: false}
		timer := time.NewTimer(5 * time.Second)

		for {
			select {
			case received := <-channel:
				if protocol.IsServerStatusTopic(received.Topic) {
					if protocol.IsAliveStatusMessage(received.Payload) {
						timer.Stop()
						return codec, nil
					} else if protocol.IsDeadStatusMessage(received.Payload) {
						timer.Stop()
						return nil, ErrServerDead
					}
				}

			case <-timer.C:
				return nil, ErrServerDead
			}
		}
	}
}

// Submit a RPC request to the server for execution by publishing a
// message to the server topic.
//
// Note that if the codec is invalidated by previous dead status
// message, the function will do nothing and return the error
// [ErrServerDead] directly.
//
func (codec *Codec) WriteRequest(request *rpc.Request, input interface{}) error {
	if codec.invalidated == false {
		packet := &protocol.RequestPacket{}
		packet.Type = "request"
		packet.Method = request.ServiceMethod
		packet.Id = strconv.FormatUint(request.Seq, 10)
		packet.Client = codec.name
		packet.Server = codec.server
		packet.Input = input

		if mPacket, err := json.Marshal(packet); err != nil {
			return err
		} else {
			publishPacket := &paho.Publish{
				Topic:   protocol.GetServerRpcTopic(codec.server),
				QoS:     0,
				Retain:  false,
				Payload: mPacket,
			}

			if _, err := codec.mqtt.Publish(codec.ctx, publishPacket); err != nil {
				return err
			} else {
				return nil
			}
		}
	} else {
		return ErrServerDead
	}
}

// Receive a RPC response from the server. The function will decode
// the RPC response and cache the "body" part for the subsequent
// [Codec.ReadResponseBody] calls.
//
// The function will also check for status messages. If any "dead"
// status message is received, the codec will be invalidated and
// the error [ErrServerDead] is returned.
//
func (codec *Codec) ReadResponseHeader(response *rpc.Response) error {
	if codec.invalidated == false {
		for received := range codec.channel {
			//
			// A will that contains the text "dead" will be posted to the
			// status topic when the server disconnects. After receiving
			// the will, no more message is expected from the server. In
			// this situation, the codec can only be closed. Any other
			// operations will fail with io.EOF as error.
			//

			if protocol.IsServerStatusTopic(received.Topic) {
				if protocol.IsDeadStatusMessage(received.Payload) {
					codec.invalidated = true
					return ErrServerDead
				} else {
					continue
				}
			}

			//
			// Apart from the special case above, the message should be
			// the response to an outstanding request.
			//

			if err := json.Unmarshal(received.Payload, &codec.response); err != nil {
				continue
			} else if id, err := strconv.ParseUint(codec.response.Id, 10, 64); err != nil {
				continue
			} else {
				if codec.response.Type != "response" {
					continue
				} else if codec.response.Client != codec.name {
					continue
				} else if codec.response.Server != codec.server {
					continue
				}

				response.ServiceMethod = codec.response.Method
				response.Seq = id
				return nil
			}
		}
	}

	return ErrServerDead
}

// Decode the result from the previous RPC response to the given
// output object.
//
// Note that if the codec is invalidated by previous dead status
// message, the function will do nothing and return the error
// [ErrServerDead] directly.
//
func (codec *Codec) ReadResponseBody(output interface{}) error {
	if codec.invalidated == false {
		if output == nil {
			return nil
		} else {
			if err := json.Unmarshal(codec.response.Output, output); err != nil {
				return err
			} else {
				return nil
			}
		}
	} else {
		return ErrServerDead
	}
}

// Close the connection to the intermediate MQTT broker and
// free any resources used by the codec.
//
func (codec *Codec) Close() error {
	err := codec.mqtt.Disconnect(&paho.Disconnect{ReasonCode: 0})
	close(codec.channel)
	return err
}
