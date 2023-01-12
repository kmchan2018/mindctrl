package protocol

import (
	"fmt"
	"strings"
)

// Get the MQTT topic where the client monitors for RPC responses.
// The client can monitor the topic to receive any response from
// the server.
//
func GetClientTopic(client string) string {
	return fmt.Sprintf("mindctrl/clients/%s", client)
}

// Get the MQTT topic where the server monitors for RPC requests.
// Clients can publish request packets to this topic for the
// server to execute.
//
func GetServerRpcTopic(server string) string {
	return fmt.Sprintf("mindctrl/servers/%s", server)
}

// Get the MQTT topic where the server publishes status updates.
// The client can monitor the topic to determine if the server
// is online.
//
func GetServerStatusTopic(server string) string {
	return fmt.Sprintf("mindctrl/statuses/%s", server)
}

// Return if the given MQTT topic is a server status topic. It
// checks if the topic starts with the string mindctrl/statuses/.
//
func IsServerStatusTopic(topic string) bool {
	return strings.HasPrefix(topic, "mindctrl/statuses/")
}
