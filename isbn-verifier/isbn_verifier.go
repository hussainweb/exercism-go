package isbn

import (
	"errors"
	"unicode"
)

// IsValidISBN checks if the isbn is valid
func IsValidISBN(isbn string) bool {
	checksum, err := getIsbnChecksum(isbn)
	if err != nil {
		return false
	}

	return checksum%11 == 0
}

func getIsbnChecksum(isbn string) (int, error) {
	digit := 0
	checksum := 0
	for _, l := range isbn {
		if l == '-' {
			continue
		}

		// ISBN may only contain numbers, except 'X' at the
		// 10th position
		if !unicode.IsNumber(l) && (l != 'X' || digit != 9) {
			return 0, errors.New("invalid ISBN")
		}

		d := 0
		switch {
		case l == 'X':
			d = 10
		default:
			d = int(l) - 48
		}

		checksum += (10 - digit) * d
		digit++
	}

	if digit != 10 {
		return 0, errors.New("invalid ISBN")
	}

	return checksum, nil
}
