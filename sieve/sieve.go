package sieve

// Sieve finds prime numbers up to a limit.
func Sieve(limit int) []int {
	var primes []int
	candidates := make(map[int]bool)

	for i := 2; i <= limit; i++ {
		isPrime, ok := candidates[i]
		if ok && !isPrime {
			// If the key exists and is false, this number is already
			// marked as not-prime.
			continue
		}

		// Mark the number as prime and mark multiples as not prime.
		candidates[i] = true
		primes = append(primes, i)

		n := 2
		for n*i <= limit {
			candidates[n*i] = false
			n++
		}
	}

	return primes
}
