// Package leap to check for leap years
package leap

// IsLeapYear return true if the specified year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
