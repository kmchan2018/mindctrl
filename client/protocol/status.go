package protocol

// Return if the given payload is an alive status message which
// contains the literal string "alive".
//
func IsAliveStatusMessage(payload []byte) bool {
	if string(payload) == "alive" {
		return true
	} else {
		return false
	}
}

// Return if the given payload is a dead status message which
// contains the literal string "dead".
//
func IsDeadStatusMessage(payload []byte) bool {
	if string(payload) == "dead" {
		return true
	} else {
		return false
	}
}
