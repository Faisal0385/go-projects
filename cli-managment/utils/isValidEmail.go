package utils

import (
	"regexp"
	"strings"
)

// emailRe is a compiled regular expession used to validate email addresses.
// Explanation of the regex pattern:
//
//	^                      → start of the string
//	[A-Za-z0-9._%+\-]+     → one or more allowed characters before '@'
//	                         (letters, digits, dot, underscore, percent, plus, minus)
//	@                      → mandatory '@' symbol
//	[A-Za-z0-9.\-]+        → domain name part (letters, digits, dot, hyphen)
//	\.                     → mandatory '.' before TLD
//	[A-Za-z]{2,}           → top-level domain must be at least 2 letters (e.g., com, org, net)
//	$                      → end of the string
var emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// IsValidEmail checks whether the given string is a valid email address.
// - Trims spaces from input
// - Returns true if the email matches the regex, false otherwise
func IsValidEmail(s string) bool {
	return emailRe.MatchString(strings.TrimSpace(s))
}
