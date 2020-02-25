package strain

// Ints integers
type Ints []int

// Lists array of integers
type Lists [][]int

// Strings array of strings
type Strings []string

// Keep is array_filter
func (c Ints) Keep(filter func(int) bool) Ints {
	if c == nil {
		return nil
	}

	r := Ints{}

	for _, i := range c {
		if filter(i) {
			r = append(r, i)
		}
	}
	return r
}

// Discard is array_filter
func (c Ints) Discard(filter func(int) bool) Ints {
	if c == nil {
		return nil
	}

	r := Ints{}

	for _, i := range c {
		if !filter(i) {
			r = append(r, i)
		}
	}
	return r
}

// Keep is array_filter
func (c Lists) Keep(filter func([]int) bool) Lists {
	if c == nil {
		return nil
	}

	r := Lists{}

	for _, i := range c {
		if filter(i) {
			r = append(r, i)
		}
	}
	return r
}

// Keep is array_filter
func (c Strings) Keep(filter func(string) bool) Strings {
	if c == nil {
		return nil
	}

	r := Strings{}

	for _, i := range c {
		if filter(i) {
			r = append(r, i)
		}
	}
	return r
}
