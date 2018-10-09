// Leap package provides a function to handle leap years
package leap

// IsLeapYear gets an int and returns true or false if that 
// int is leap year or not
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
