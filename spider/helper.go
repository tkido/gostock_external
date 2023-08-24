package spider

import (
	"strings"
)

// Sanitizers
func pass(s string) string {
	return s
}
func rmComma(s string) string {
	if s == "---" {
		return "-"
	}
	return strings.Replace(s, ",", "", -1)
}
