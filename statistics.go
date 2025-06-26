package mathplus

import (
	"math"
	"sort"
)

// Average returns the arithmetic mean of a set of float64 numbers.
// If no values are provided, it returns 0.
func Average(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}

// Median returns the median of a set of float64 numbers.
// If no values are provided, it returns 0.
func Median(nums ...float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sorted := append([]float64(nil), nums...)
	sort.Float64s(sorted)

	mid := n / 2
	if n%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

// StdDev returns the sample standard deviation of a set of numbers.
// If fewer than 2 values are provided, it returns 0.
func StdDev(nums ...float64) float64 {
	n := len(nums)
	if n < 2 {
		return 0
	}

	mean := Average(nums...)
	var sumSqDiff float64
	for _, v := range nums {
		diff := v - mean
		sumSqDiff += diff * diff
	}

	variance := sumSqDiff / float64(n-1)
	return math.Sqrt(variance)
}

// Mode returns the most frequent value(s) in the input data.
// If no mode exists (i.e., all values appear only once), it returns an empty slice.
func Mode(nums ...float64) []float64 {
	if len(nums) == 0 {
		return nil
	}

	freq := make(map[float64]int)
	maxFreq := 0

	for _, n := range nums {
		freq[n]++
		if freq[n] > maxFreq {
			maxFreq = freq[n]
		}
	}

	var modes []float64
	for num, count := range freq {
		if count == maxFreq && maxFreq > 1 {
			modes = append(modes, num)
		}
	}
	return modes
}
