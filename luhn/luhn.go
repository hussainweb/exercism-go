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

	// We need to double every digit from the RIGHT, but scanning
	// a string from the right isn't efficient. Hence, we set up
	// a flag called 'double' which will indicate if the current
	// digit should be doubled and we toggle that flag every loop.
	// The initial value of the flag is dependent on if the string
	// has even length or odd.
	double := len(normalized)%2 == 0

	checksum := 0
	for _, l := range normalized {
		if !unicode.IsNumber(l) {
			return false
		}

		// Convert the number in the character to an integer.
		digit := int(l - '0')
		if double {
			digit *= 2
			if digit >= 10 {
				digit -= 9
			}
		}
		checksum += digit
		// Toggle the 'doubling' of the next digit.
		double = !double
	}

	return checksum%10 == 0
}
