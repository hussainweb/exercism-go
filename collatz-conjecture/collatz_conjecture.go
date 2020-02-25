package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture returns the number of steps required for Collatz Conjecture
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("n must be a positive integer")
	}

	steps := 0
	for n != 1 {
		switch n % 2 {
		case 0:
			n /= 2
		case 1:
			n = (3 * n) + 1
		}
		steps++
	}

	return steps, nil
}
