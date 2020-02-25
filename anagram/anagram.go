package anagram

import (
	"strings"
)

// Detect returns a list of valid anagrams
func Detect(word string, candidates []string) (anagrams []string) {
	for _, cand := range candidates {
		if isAnagram(word, cand) {
			anagrams = append(anagrams, cand)
		}
	}
	return
}

func isAnagram(word string, anagram string) bool {
	word = strings.ToLower(word)
	anagram = strings.ToLower(anagram)
	if word == anagram {
		return false
	}

	wordLetters := []rune(word)
	for _, letter := range anagram {
		if i, ok := findElement(wordLetters, letter); ok {
			wordLetters = remElement(wordLetters, i)
		} else {
			return false
		}
	}

	// If we have no letters left, we have an anagram
	if len(wordLetters) == 0 {
		return true
	}
	return false
}

func findElement(slice []rune, letter rune) (int, bool) {
	for i, c := range slice {
		if c == letter {
			return i, true
		}
	}

	return -1, false
}

func remElement(slice []rune, index int) []rune {
	slice[index] = slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice
}
