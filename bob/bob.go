package bob

import (
	"strings"
)

// Hey returns a response from Bob
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	if remark == "" {
		return "Fine. Be that way!"
	}

	if IsQuestion(remark) {
		if HasLetters(remark) && IsYell(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if HasLetters(remark) && IsYell(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

// IsYell checks if the string is a yell
func IsYell(remark string) bool {
	return remark == strings.ToUpper(remark)
}

// IsQuestion checks if the remark is a question
func IsQuestion(remark string) bool {
	return remark[len(remark)-1] == '?'
}

// HasLetters checks if the remark has any of the alphabets
func HasLetters(remark string) bool {
	return strings.ContainsAny(strings.ToUpper(remark), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
