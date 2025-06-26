package mathplus

import (
	"fmt"
	"math"
)

// Max returns the maximum value in the input slice.
// It ignores NaN and -Inf, and returns +Inf immediately if found.
// Logs warnings for special values.
func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := math.Inf(-1)
	for i, v := range nums {
		if math.IsNaN(v) {
			fmt.Printf("NaN found in position %d of array %v\n", i, nums)
			continue
		}
		if math.IsInf(v, -1) {
			fmt.Printf("-Inf found in position %d of array %v\n", i, nums)
			continue
		}
		if math.IsInf(v, 1) {
			fmt.Printf("+Inf found in position %d of array %v\n", i, nums)
			return v
		}
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum value in the input slice.
// Returns -Inf immediately if found. NaN leads to NaN result.
func Min(nums ...float64) float64 {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	min := nums[0]
	for i, v := range nums {
		if math.IsNaN(v) {
			fmt.Printf("NaN found in position %d of array %v\n", i, nums)
			return math.NaN()
		}
		if math.IsInf(v, -1) {
			fmt.Printf("-Inf found in position %d of array %v\n", i, nums)
			return math.Inf(-1)
		}
		if math.IsInf(v, 1) {
			fmt.Printf("+Inf found in position %d of array %v\n", i, nums)
			continue
		}
		if v < min {
			min = v
		}
	}
	return min
}
