package protein

import (
	"errors"
)

// ErrStop indicates we have encountered a STOP codon
var ErrStop = errors.New("STOP")

// ErrInvalidBase indicates an invalid base
var ErrInvalidBase = errors.New("Invalid Base")

var codonProteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon converts a codon string into a list of proteins
func FromCodon(codon string) (string, error) {
	var protein string
	protein, ok := codonProteins[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// FromRNA converts a RNA string to a list of proteins
func FromRNA(rna string) (proteins []string, err error) {
	l := len(rna)
	for i := 0; i < l; i += 3 {
		codon := rna[i : i+3]
		protein, ok := FromCodon(codon)
		if ok == ErrStop {
			return
		}
		if ok == ErrInvalidBase {
			return proteins, ok
		}
		proteins = append(proteins, protein)
	}
	return
}
