package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if the given string is valid per Luhn formula
func Valid(sequence string) bool {
	normalized := strings.ReplaceAll(strings.Trim(sequence, " "), " ", "")
	if len(normalized) <= 1 {
		return false
	}

	// We need to double every digit from the RIGHT, but we can
	// do it more efficiently by calculating the modulus to expect
	// when scanning from the left.
	// When the string has odd number of digits, it is every second
	// digit from the left, but if there are even number of digits,
	// it is every digit but starting at the first.
	// We can do this by just calculating the modulus of 2 we
	// should expect. For strings of even length, the modulus is 0
	// which gets us every second digit starting from the first.
	modulus := 1
	if len(normalized)%2 == 0 {
		modulus = 0
	}

	checksum := 0
	for i, l := range normalized {
		if !unicode.IsNumber(l) {
			return false
		}

		// Convert the number in the character to an integer.
		digit := int(l & 0xF)
		if i%2 == modulus {
			digit *= 2
			if digit >= 10 {
				digit -= 9
			}
		}
		checksum += digit
	}

	return checksum%10 == 0
}
