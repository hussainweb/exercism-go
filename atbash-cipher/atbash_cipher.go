package atbash

import (
	"strings"
	"unicode"
)

// Atbash returns a string encoded with Atbash cipher
func Atbash(str string) string {
	var encoded strings.Builder
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			continue
		}

		if unicode.IsLetter(c) {
			// Transpose the letter to its opposite one.
			// e.g. 'a' becomes 'z', 'm' becomes 'n', and so on.
			c = 'z' - unicode.ToLower(c) + 'a'
		}

		// Write a blank space at every 5th position in
		// a 6 character slot.
		if encoded.Len()%6 == 5 {
			encoded.WriteRune(' ')
		}
		encoded.WriteRune(c)
	}

	return encoded.String()
}
