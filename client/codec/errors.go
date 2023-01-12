package codec

import (
	"errors"
)

// Error reported by [Codec] to indicate that the server (not the
// intermediate MQTT broker) is dead.
//
var ErrServerDead = errors.New("server dead")
