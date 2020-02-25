package diffsquares

import "math"

// SquareOfSum returns square of sum of first n integers
func SquareOfSum(n int) int {
	f := float64(n * (n + 1) / 2)
	return int(math.Pow(f, 2))
}

// SumOfSquares returns the sum of squares of first n integers
func SumOfSquares(n int) int {
	// Formula from https://brilliant.org/wiki/sum-of-n-n2-or-n3/#sum-of-the-squares-of-the-first-n-positive-integers
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference of squares for the first n integers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
