package allergies

var allergyList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies lists all the allergies for the score
func Allergies(score uint) []string {
	var result = make([]string, 0)
	bitmask := uint(1)
	for _, allergen := range allergyList {
		if score&bitmask != 0 {
			result = append(result, allergen)
		}
		bitmask <<= 1
	}
	return result
}

// AllergicTo checks if allergy is indicated.
func AllergicTo(score uint, allergen string) bool {
	for i, a := range allergyList {
		if allergen == a {
			return score&(1<<i) != 0
		}
	}
	return false
}
