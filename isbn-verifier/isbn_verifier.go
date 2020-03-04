package isbn

import (
	"errors"
	"strings"
	"unicode"
)

// IsValidISBN checks if the isbn is valid
func IsValidISBN(isbn string) bool {
	isbnString, err := normalizeISBN(isbn)
	if err != nil {
		return false
	}

	// Running the loop through the string twice has no significant
	// impact on the performance. On my machine, it was only
	// 3% slower which is trivial.
	checksum := 0
	for i, l := range isbnString {
		d := 0
		switch {
		case l == 'X':
			d = 10
		default:
			d = int(l) - 48
		}

		checksum += (10 - i) * d
	}

	return checksum%11 == 0
}

func normalizeISBN(isbn string) (string, error) {
	var sb strings.Builder
	digit := 0
	for _, l := range isbn {
		if l == '-' {
			continue
		}

		// ISBN may only contain numbers, except 'X' at the
		// 10th position
		if !unicode.IsNumber(l) && (l != 'X' || digit != 9) {
			return "", errors.New("invalid ISBN")
		}
		digit++
		sb.WriteRune(l)
	}

	if sb.Len() != 10 {
		return "", errors.New("invalid ISBN")
	}

	return sb.String(), nil
}
