package darts

import "math"

// Score returns the score of a dart throw
func Score(x, y float64) int {
	r := math.Sqrt((x * x) + (y * y))
	switch {
	case r <= 1:
		return 10
	case r <= 5:
		return 5
	case r <= 10:
		return 1
	}
	return 0
}
