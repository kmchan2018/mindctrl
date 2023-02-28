package codec

// Options for creating a new codec.
//
// The 'Username' and 'Password' fields contains the credentials
// used to authenticate with the intermediate MQTT broker. They
// can be empty if authentication is not needed. By default, they
// are both empty meaning no authentication is needed.
//
// The 'WebsocketMessageSize' field contains the maximum size
// of a single websocket message. The minimum size is 32KB, and
// the default 32MB.
//
// The 'Capacity' field contains the number of MQTT messages to
// be buffered for processing. The minimum size and the default
// sizes are both 100.
//
type Options struct {
	Username           string // username for the intermediate MQTT broker
	Password           string // password for the intermediate MQTT broker
	WebsocketFrameSize int64  // maximum size of websocket packet
	MqttMessageBuffer  int    // number of MQTT messages buffered by the codec
	MqttKeepAlive      int    // keepalive duration of MQTT connection
}

func (options *Options) getPahoUsernameFlag() bool {
	if options == nil {
		return false
	} else if options.Username == "" {
		return false
	} else {
		return true
	}
}

func (options *Options) getPahoPasswordFlag() bool {
	if options == nil {
		return false
	} else if options.Password == "" {
		return false
	} else {
		return true
	}
}

func (options *Options) getPahoUsername() string {
	if options == nil {
		return ""
	} else {
		return options.Username
	}
}

func (options *Options) getPahoPassword() []byte {
	if options == nil {
		return []byte{}
	} else if options.Password == "" {
		return []byte{}
	} else {
		return []byte(options.Password)
	}
}

func (options *Options) getWebsocketFrameSize() int64 {
	if options == nil {
		return 16 * 1024 * 1024
	} else if options.WebsocketFrameSize < 32768 {
		return 16 * 1024 * 1024
	} else {
		return options.WebsocketFrameSize
	}
}

func (options *Options) getMqttMessageBuffer() int {
	if options == nil {
		return 100
	} else if options.MqttMessageBuffer < 100 {
		return 100
	} else {
		return options.MqttMessageBuffer
	}
}

func (options *Options) getMqttKeepAlive() uint16 {
	if options == nil {
		return 1800
	} else if options.MqttKeepAlive < 0 {
		return 0
	} else if options.MqttKeepAlive == 0 {
		return 1800
	} else if options.MqttKeepAlive <= 300 {
		return 300
	} else if options.MqttKeepAlive >= 65535 {
		return 65535
	} else {
		return uint16(options.MqttKeepAlive)
	}
}
