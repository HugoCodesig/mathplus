package mathplus

// This file will contain math utils and will probably contain new functions that haven't been sorted yet.

// Abs returns the absolute value of a float64 number.
func Abs(num float64) float64 {
	if num < 0 {
		return -num
	} else if num == 0 {
		return 0
	}
	return num
}
