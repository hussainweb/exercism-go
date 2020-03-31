package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	var strings []string
	for i := 0; i <= len(s)-n; i++ {
		strings = append(strings, s[i:i+n])
	}

	return strings
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}

	return s[0:n], true
}
