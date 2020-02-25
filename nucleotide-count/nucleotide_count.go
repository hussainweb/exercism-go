package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (h Histogram, err error) {
	nucleotides := string(d)
	h = make(Histogram)
	h['A'] = strings.Count(nucleotides, "A")
	h['C'] = strings.Count(nucleotides, "C")
	h['G'] = strings.Count(nucleotides, "G")
	h['T'] = strings.Count(nucleotides, "T")
	if len(nucleotides) != h['A']+h['C']+h['G']+h['T'] {
		err = errors.New("Invalid nucleotide")
	}
	return
}
