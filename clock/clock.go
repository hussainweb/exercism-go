package clock

import "fmt"

// Clock represents a time object
type Clock struct {
	minutes int
}

// New creates a new clock
func New(h int, m int) Clock {
	mins := h*60 + m
	for mins < 0 {
		mins += 24 * 60
	}

	mins %= 60 * 24

	return Clock{mins}
}

func (c Clock) String() string {
	h := c.minutes / 60
	m := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

// Add adds minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract subtracts minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
