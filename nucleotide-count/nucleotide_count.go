package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, err error) {
	nucleotides := string(d)
	h = make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	for _, n := range nucleotides {
		h[n]++
	}
	if len(nucleotides) != h['A']+h['C']+h['G']+h['T'] {
		err = errors.New("Invalid nucleotide")
	}
	return
}
