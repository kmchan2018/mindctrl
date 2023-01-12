package options

import (
	"regexp"
	"strconv"
)

var (
	idPattern *regexp.Regexp = regexp.MustCompile(`^(?:0|[1-9][0-9]*)$`)
)

// Return if the input string contains decimal digits only. Note
// that the function is not intended for validating valid number
// and will reject other characters that may appears in number such
// as negative sign and decimal point.
//
func IsId(input string) bool {
	return idPattern.MatchString(input)
}

// Return if the input string contains hexadecimal digits only. Note
// that the function is not intended for validating valid number
// and will reject other characters that may appears in number such
// as negative sign and decimal point.
//
func ParseId(input string) int {
	if result, err := strconv.Atoi(input); err == nil {
		return result
	} else {
		panic(err)
	}
}
