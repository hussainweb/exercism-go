package lasagna

func PreparationTime(slices []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}
	return len(slices) * averageTime
}

func Quantities(slices []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, slice := range slices {
		switch slice {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[(len(friendsList) - 1)]
}

func ScaleRecipe(portions []float64, numPortions int) []float64 {
	newPortions := make([]float64, len(portions))
	for i := range newPortions {
		newPortions[i] = portions[i] * float64(numPortions) / 2
	}
	return newPortions
}
