package mathplus

import (
	"errors"
	"math"
)

// Add returns the sum of all provided float64 numbers.
func Add(nums ...float64) float64 {
	var sum float64
	for _, val := range nums {
		sum += val
	}
	return sum
}

// Sub returns the cumulative subtraction starting from 0.
func Sub(nums ...float64) float64 {
	var sub float64
	for _, val := range nums {
		sub -= val
	}
	return sub
}

// Mul returns the product of all provided float64 numbers.
// If any number is 0, the result will be 0.
func Mul(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	if containsZero(nums) {
		return 0
	}
	product := nums[0]
	for i := 1; i < len(nums); i++ {
		product *= nums[i]
	}
	return product
}

// Div returns the result of sequential division of all inputs from left to right.
// Returns an error if dividing by zero or input is empty.
func Div(nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("cannot divide empty list")
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		result /= nums[i]
	}
	return result, nil
}

// Pow returns base raised to the given exponent.
func Pow(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Sqrt returns the square root of a number using Newton's method.
// If the number is negative, it returns NaN.
func Sqrt(num float64) float64 {
	if num < 0 {
		return math.NaN()
	}
	if num == 0 {
		return 0
	}
	x := num
	for i := 0; i < 10; i++ {
		x = (x + num/x) / 2
	}
	return x
}

// containsZero checks if any number in the slice is 0.
func containsZero(nums []float64) bool {
	for _, v := range nums {
		if v == 0 {
			return true
		}
	}
	return false
}
