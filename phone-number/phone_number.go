package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// Number returns a NANP number after cleaning it up.
func Number(phone string) (string, error) {
	phone = regexp.MustCompile("[^0-9]*").ReplaceAllString(phone, "")

	if len(phone) > 11 || len(phone) < 10 {
		return "", errors.New("Invalid number")
	}
	if len(phone) == 11 {
		if phone[0] != '1' {
			return "", errors.New("Not a NANP number")
		}
		phone = phone[1:]
	}
	if phone[0] == '0' || phone[0] == '1' {
		return "", errors.New("Invalid area code")
	}
	if phone[3] == '0' || phone[3] == '1' {
		return "", errors.New("Invalid exchange code")
	}

	return phone, nil
}

// AreaCode returns the area code for a NANP number.
func AreaCode(phone string) (string, error) {
	phone, err := Number(phone)
	if err != nil {
		return "", err
	}
	return phone[:3], nil
}

// Format returns the formatted version of a NANP number.
func Format(phone string) (string, error) {
	phone, err := Number(phone)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phone[:3], phone[3:6], phone[6:]), nil
}
