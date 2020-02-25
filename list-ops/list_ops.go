package listops

// IntList list of ints
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append appends
func (a IntList) Append(b IntList) IntList {
	for _, item := range b {
		a = append(a, item)
	}
	return a
}

// Concat appends
func (a IntList) Concat(b []IntList) IntList {
	for _, list := range b {
		for _, item := range list {
			a = append(a, item)
		}
	}
	return a
}

// Filter filters
func (a IntList) Filter(filterFunc predFunc) IntList {
	filtered := IntList{}
	for _, item := range a {
		if filterFunc(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Length find length
func (a IntList) Length() int {
	c := 0
	for range a {
		c++
	}
	return c
}

// Map maps items
func (a IntList) Map(mapFunc unaryFunc) IntList {
	mapped := IntList{}
	for _, item := range a {
		mapped = append(mapped, mapFunc(item))
	}
	return mapped
}

// Foldr folds
func (a IntList) Foldr(folder binFunc, initial int) int {
	folded := initial
	for _, item := range a {
		folded = folder(folded, item)
	}
	return folded
}

// Foldl folds
func (a IntList) Foldl(folder binFunc, initial int) int {
	folded := initial
	for _, item := range a {
		folded = folder(folded, item)
	}
	return folded
}

// Reverse folds
func (a IntList) Reverse() IntList {
	reversed := IntList{}
	for i := len(a) - 1; i >= 0; i-- {
		reversed = append(reversed, a[i])
	}
	return reversed
}
