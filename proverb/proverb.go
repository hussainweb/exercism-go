package proverb

import "fmt"

// Proverb returns standard proverb
func Proverb(rhyme []string) []string {
	var proverb []string

	if len(rhyme) == 0 {
		return proverb
	}

	lineTemplate := "For want of a %s the %s was lost."
	lastLineTemplate := "And all for the want of a %s."
	for i := range rhyme {
		if i+1 < len(rhyme) {
			proverb = append(proverb, fmt.Sprintf(lineTemplate, rhyme[i], rhyme[i+1]))
		}
	}
	proverb = append(proverb, fmt.Sprintf(lastLineTemplate, rhyme[0]))

	return proverb
}
