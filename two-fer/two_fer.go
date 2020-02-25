// Package twofer contains logic for two-fer strings.
package twofer

import "fmt"

// ShareWith returns a standard two-fer string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
