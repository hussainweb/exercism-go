package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encodes the string
func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	var sb strings.Builder

	// Initialise the variables with the first character
	// in the string
	var lastRune rune = rune(input[0])
	lastRuneCount := 1
	for _, l := range input[1:] {
		if l == lastRune {
			// Keep counting repeated characters
			lastRuneCount++
			continue
		}

		// We should write the number to the string only if
		// it is more than 1
		if lastRuneCount > 1 {
			sb.WriteString(strconv.Itoa(lastRuneCount))
		}
		sb.WriteRune(lastRune)
		lastRuneCount = 1
		lastRune = l
	}

	// Write the last rune to the string
	if lastRuneCount > 1 {
		sb.WriteString(strconv.Itoa(lastRuneCount))
	}
	sb.WriteRune(lastRune)
	return sb.String()
}

// RunLengthDecode encodes the string
func RunLengthDecode(input string) string {
	if input == "" {
		return ""
	}

	var countBuilder strings.Builder
	var decoded strings.Builder
	for _, l := range input {
		if unicode.IsNumber(l) {
			countBuilder.WriteRune(l)
			continue
		}

		// It's not a number, which means we should write the
		// character to the string and reset our count string
		count, _ := strconv.Atoi(countBuilder.String())
		countBuilder.Reset()

		// If there was no count, it was 1
		if count == 0 {
			count = 1
		}
		decoded.WriteString(strings.Repeat(string(l), count))
	}

	return decoded.String()
}
