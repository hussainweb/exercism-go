package cipher

import (
	"strings"
	"unicode"
)

// Caesar implements the Caesar cipher.
type Caesar struct {
	shift        string
	shiftCounter int
}

// NewCaesar returns a new Caesar implementation.
func NewCaesar() *Caesar {
	return &Caesar{"d", 0}
}

// NewShift returns a new Caesar implementation.
func NewShift(shift int) *Caesar {
	if shift <= -26 || shift >= 26 || shift == 0 {
		return nil
	}
	if shift < 0 {
		shift += 26
	}
	return &Caesar{string(rune(shift + 'a')), 0}
}

// NewVigenere returns
func NewVigenere(shift string) *Caesar {
	seenOtherThanA := false
	for _, r := range shift {
		if !unicode.IsLower(r) {
			return nil
		}
		if r != 'a' {
			seenOtherThanA = true
		}
	}

	if !seenOtherThanA {
		return nil
	}

	return &Caesar{shift, 0}
}

// Encode encodes a string
func (c *Caesar) Encode(str string) string {
	var sb strings.Builder
	c.shiftCounter = 0
	for _, r := range str {
		if !unicode.IsLetter(r) {
			continue
		}

		r = unicode.ToLower(r)

		er := r + c.getNextShiftCharacter()
		if er > 'z' {
			er -= 26
		}
		sb.WriteRune(er)
	}

	return sb.String()
}

// Decode decodes a string
func (c *Caesar) Decode(str string) string {
	var sb strings.Builder
	c.shiftCounter = 0
	for _, r := range str {
		er := r - c.getNextShiftCharacter()
		if er < 'a' {
			er += 26
		}
		sb.WriteRune(er)
	}

	return sb.String()
}

func (c *Caesar) getNextShiftCharacter() rune {
	r := (rune(c.shift[c.shiftCounter]) - 'a')
	c.shiftCounter++
	if c.shiftCounter >= len(c.shift) {
		c.shiftCounter = 0
	}
	return r
}
