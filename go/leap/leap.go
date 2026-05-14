// Package leap for checking if a leap year.
package leap

// IsLeapYear checks if given year is leap, returns boolean.
func IsLeapYear(year int) bool {

	return year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// if year%100 == 0 && year%400 == 0 {
	// 	return true
	// }

	// if year%4 == 0 && year%100 != 0 {
	// 	return true
	// }

	// return false
}
