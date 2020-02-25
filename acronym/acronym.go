package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns acronyms
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = regexp.MustCompile("[^A-Za-z ]*").ReplaceAllString(s, "")

	words := strings.Split(s, " ")
	acronym := ""
	for _, word := range words {
		if word != "" {
			acronym += string(word[0])
		}
	}

	return strings.ToUpper(acronym)
}
