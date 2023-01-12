package protocol

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Generate an unique MQTT client ID to identify the client with the
// intermediate MQTT broker.
//
func GenerateMqttClientId(prefix string) string {
	builder := strings.Builder{}
	now := time.Now().UnixMicro()
	pid := os.Getpid()
	randbytes := make([]byte, 8)

	builder.WriteString("mindctrl_")

	if prefix != "" {
		builder.WriteString(prefix)
		builder.WriteString("_")
	}

	builder.WriteString(strconv.FormatInt(int64(pid), 10))
	builder.WriteString("_")
	builder.WriteString(strconv.FormatInt(now, 10))

	if _, err := crand.Read(randbytes); err == nil {
		builder.WriteByte('_')
		builder.WriteString(strconv.FormatUint(binary.LittleEndian.Uint64(randbytes), 10))
	} else {
		mrand.Seed(now)
		builder.WriteByte('_')
		builder.WriteString(strconv.FormatUint(mrand.Uint64(), 10))
	}

	return builder.String()
}
