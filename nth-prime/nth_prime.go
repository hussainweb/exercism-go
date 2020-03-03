package prime

import (
	"math"
)

// Nth returns the n-th prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	foundPrimes := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			foundPrimes++
			if foundPrimes == n {
				// We have found our n-th prime
				return i, true
			}
		}
	}
}

func isPrime(p int) bool {
	// Doesn't affect the benchmark much in terms of percentage,
	// but this is a useful shortcut in other situations.
	if p > 7 && (p%2 == 0 || p%3 == 0 || p%5 == 0 || p%7 == 0) {
		return false
	}

	s := int(math.Sqrt(float64(p)))
	for i := 2; i <= s; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}
