package pangram

import (
	"strings"
	"unicode"
)

// IsPangram checks if a string is a pangram
func IsPangram(sentence string) bool {
	seenLetters := make(map[rune]bool)
	lettersCounted := 0
	for _, letter := range strings.ToLower(sentence) {
		if unicode.IsLetter(letter) && seenLetters[letter] == false {
			seenLetters[letter] = true
			lettersCounted++
		}
	}
	return lettersCounted == 26
}
