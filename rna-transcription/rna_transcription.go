package strand

import (
	"strings"
)

// ToRNA converts a DNA strand to RNA
func ToRNA(dna string) string {
	s := dna
	s = strings.ReplaceAll(s, "G", "X")
	s = strings.ReplaceAll(s, "C", "G")
	s = strings.ReplaceAll(s, "X", "C")
	s = strings.ReplaceAll(s, "A", "U")
	s = strings.ReplaceAll(s, "T", "A")
	return s
}
