// Package protocol contains information required to communicate
// with the extension.
//
// First of all, the package provides some helper function to get
// and check the MQTT topics used by the extension and the clients.
//
// Then, the package also provides data types for the request and
// response packets, plus functions to determine the type of status
// packets.
//
// Moreover, the package also contains a list of method supported
// by the extension, and their corresponding input/output data.
//
// Lastly, the package also provides some helper functions for
// miscellaneous tasks like generating unique MQTT client ID.
//
package protocol
