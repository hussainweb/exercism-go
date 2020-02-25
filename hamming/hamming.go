package hamming

import (
	"fmt"
)

// Distance returns the Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Input strings are not of same length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
