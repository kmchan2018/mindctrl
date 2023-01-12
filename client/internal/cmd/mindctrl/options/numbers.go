package options

import (
	"regexp"
	"strconv"
)

var (
	numberPattern *regexp.Regexp = regexp.MustCompile(`^(?:0|[1-9][0-9]*)$`)
)

// Return if the input string represents a non-negative number. Mote
// that the function is not intended for validating valid number
// and will reject other characters that may appears in number such
// as negative sign and decimal point.
//
func IsNumber(input string) bool {
	return idPattern.MatchString(input)
}

// Return if the input string represents a non-negative number. Note
// that the function is not intended for validating valid number
// and will reject other characters that may appears in number such
// as negative sign and decimal point.
//
func ParseNumber(input string) int {
	if result, err := strconv.Atoi(input); err == nil {
		return result
	} else {
		panic(err)
	}
}
