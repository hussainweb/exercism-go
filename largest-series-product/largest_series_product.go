package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct returns the largest product of
// a series of span digits
func LargestSeriesProduct(digits string, span int) (max int64, err error) {
	l := len(digits)
	if span > l {
		err = errors.New("Span is greater than length of the input string")
		return
	}
	if span < 0 {
		err = errors.New("Span must be greater than 0")
		return
	}

	for i := 0; i <= l-span; i++ {
		spannedDigits := digits[i : i+span]
		prod, err := getProduct(spannedDigits)
		if err != nil {
			return -1, err
		}

		if prod > max {
			max = prod
		}
	}

	return
}

func getProduct(digits string) (prod int64, err error) {
	if digits == "" {
		return 1, nil
	}

	prod = 1
	for _, d := range digits {
		if !unicode.IsNumber(d) {
			err = errors.New("not a number")
		}
		prod *= int64(d) - 48
	}
	return
}
