package isogram

import (
	"unicode"
)

// IsIsogram returns true if the word is an isogram
func IsIsogram(word string) bool {
	var characters = map[rune]bool{}

	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}

		b := unicode.ToUpper(letter)
		if characters[b] != false {
			// We have seen this letter before, early return
			return false
		}
		characters[b] = true
	}

	// If we reached here, we haven't seen any character twice
	// which means that the word is an isogram.
	return true
}
