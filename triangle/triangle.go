package triangle

import (
	"math"
)

// Kind indicates the type of triangle
type Kind int

const (
	_   = iota
	NaT // not a triangle
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	p := (a + b + c) / 2
	area2 := p * (p - a) * (p - b) * (p - c)
	if area2 < 0 || math.IsNaN(area2) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	var k Kind
	switch {
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
