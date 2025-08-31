package utils

import (
	"regexp"
	"strings"
)

// phoneRe is a compiled regular expession used to validate phone numbers.
// Explanation of the regex pattern:
//
//	^         → start of the string
//	\+?       → optional '+' at the beginning (for international format)
//	[0-9]{8,15} → allows only digits, length must be between 8 and 15
//	$         → end of the string
var phoneRe = regexp.MustCompile(`^\+?[0-9]{8,15}$`)

// IsValidPhone checks whether the given string is a valid phone number
// based on the phoneRe regex pattern.
// - Trims spaces from input
// - Returns true if the phone number matches the regex, false otherwise
func IsValidPhone(s string) bool {
	return phoneRe.MatchString(strings.TrimSpace(s))
}
