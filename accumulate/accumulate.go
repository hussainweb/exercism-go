package accumulate

// Accumulate transforms a given collection using a converter
func Accumulate(collection []string, converter func(string) string) []string {
	converted := []string{}
	for _, i := range collection {
		converted = append(converted, converter(i))
	}
	return converted
}
