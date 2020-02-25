package raindrops

import "strconv"

// Convert converts an integer into a raindrop string
func Convert(n int) string {
	raindrop := ""
	if n%3 == 0 {
		raindrop += "Pling"
	}
	if n%5 == 0 {
		raindrop += "Plang"
	}
	if n%7 == 0 {
		raindrop += "Plong"
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(n)
	}
	return raindrop
}
