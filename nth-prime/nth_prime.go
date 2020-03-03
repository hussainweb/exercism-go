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
	s := int(math.Sqrt(float64(p)))
	for i := 2; i <= s; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}
