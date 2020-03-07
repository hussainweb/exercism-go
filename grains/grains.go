package grains

import (
	"errors"
)

// Square returns the number of grains on a specific square
// of a chess-board
func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, errors.New("Invalid size")
	}

	// Same as 2 ^ (square - 1)
	return 1 << (square - 1), nil
}

// Total returns the number of grains for a chess-board
func Total() uint64 {
	return 1<<64 - 1
}
