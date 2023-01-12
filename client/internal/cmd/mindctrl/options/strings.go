package options

import (
	"regexp"
)

var (
	DEC_PATTERN *regexp.Regexp = regexp.MustCompile(`^[0-9]+$`)
	HEX_PATTERN *regexp.Regexp = regexp.MustCompile(`^[0-9A-Fa-f]+$`)
)

// Return if the input string contains decimal digits only. Note
// that the function is not intended for validating valid number
// and will reject other characters that may appears in number such
// as negative sign and decimal point.
//
func IsDigit(input string) bool {
	return DEC_PATTERN.MatchString(input)
}

// Return if the input string contains hexadecimal digits only. Note
// that the function is not intended for validating valid number
// and will reject other characters that may appears in number such
// as negative sign and decimal point.
//
func IsXDigit(input string) bool {
	return HEX_PATTERN.MatchString(input)
}
