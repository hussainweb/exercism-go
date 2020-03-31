package prime

// Factors returns the prime factors.
func Factors(n int64) []int64 {
	factors := make([]int64, 0)

	remainingProduct := n
	for i := int64(2); i <= n; i++ {
		for remainingProduct%i == 0 {
			factors = append(factors, i)
			remainingProduct /= i
		}
		// If we have reached 1 as our last factor, we're done.
		if remainingProduct == 1 {
			break
		}
	}

	return factors
}
