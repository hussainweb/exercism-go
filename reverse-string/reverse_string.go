package reverse

import (
	"strings"
)

// Reverse reverses a string
func Reverse(s string) string {
	var sb strings.Builder
	b := []rune(s)
	for i := len(b) - 1; i >= 0; i-- {
		sb.WriteRune(b[i])
	}
	return sb.String()
}
