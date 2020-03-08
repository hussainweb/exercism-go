package clock

import "fmt"

// Clock represents a time object
type Clock struct {
	hour   int
	minute int
}

// New creates a new clock
func New(h int, m int) Clock {
	for m < 0 {
		m += 60
		h--
	}
	for h < 0 {
		h += 24
	}

	if m >= 60 {
		h += m / 60
		m %= 60
	}
	if h >= 24 {
		h %= 24
	}

	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract subtracts minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
