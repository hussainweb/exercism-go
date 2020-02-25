package romannumerals

import (
	"errors"
)

var romanPowers map[int][]string = map[int][]string{
	0: []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	1: []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	2: []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	3: []string{"", "M", "MM", "MMM"},
}

// ToRomanNumeral converts number to Roman representation
func ToRomanNumeral(number int) (roman string, err error) {
	if number <= 0 || number > 3000 {
		err = errors.New("We only support numbers between 1 to 3000")
		return
	}

	iter := 0
	for number > 0 {
		roman = romanPowers[iter][number%10] + roman
		number /= 10
		iter++
	}

	return
}
