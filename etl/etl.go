package etl

import "strings"

// Transform transforms input to return
func Transform(input map[int][]string) map[string]int {
	out := map[string]int{}

	for score, sequence := range input {
		for _, letter := range sequence {
			out[strings.ToLower(letter)] = score
		}
	}

	return out
}
