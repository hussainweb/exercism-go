package atbash

import (
	"unicode"
)

// Atbash returns a string encoded with Atbash cipher
func Atbash(str string) string {
	encoded := ""
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			continue
		}

		if unicode.IsLetter(c) {
			c = unicode.ToLower(c)
			// Transpose the characters by adding or subtracting from the
			// mid-point of the alphabet set. The mid-point can be either
			// 'm' (109) or 'n' (110). We have to add/subtract twice the
			// distance from the mid-point and normalize by subtracting 1.
			if c <= 109 {
				c += 2*(110-c) - 1
			} else {
				c -= 2*(c-109) - 1
			}
		}

		// Write a blank space at every 5th position in
		// a 6 character slot.
		if len(encoded)%6 == 5 {
			encoded += " "
		}
		encoded += string(c)
	}

	return encoded
}
