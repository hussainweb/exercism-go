package wordcount

import (
	"strings"
	"unicode"
)

// Frequency represents the word frequency
type Frequency map[string]int

// WordCount returns word count in a sentence
func WordCount(sentence string) Frequency {
	frequency := make(Frequency)
	words := getWords(strings.ToLower(sentence))
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func getWords(sentence string) []string {
	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return r != '\'' && !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Remove single quotes around words
	for i, word := range words {
		words[i] = strings.Trim(word, "'")
	}

	return words
}
