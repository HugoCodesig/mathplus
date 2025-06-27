package mathplus

import (
	"errors"
	"math"
)

// Log returns the logarithm of num using the specified base.
// Accepted bases: 0 (natural log), 2 (log base 2), 10 (log base 10).
func Log(base int, num float64) (float64, error) {
	if num <= 0 {
		return 0, errors.New("logarithm undefined for non-positive values")
	}
	switch base {
	case 0:
		return math.Log(num), nil
	case 2:
		return math.Log2(num), nil
	case 10:
		return math.Log10(num), nil
	default:
		return 0, errors.New("unsupported log base; use 0, 2, or 10")
	}
}

// LogSpec returns special logarithmic results based on symbolic base names.
// Supported: "1p" for Log1p(x), "b" for Logb(x)
func LogSpec(base string, num float64) (float64, error) {
	switch base {
	case "1p":
		if num <= -1 {
			return 0, errors.New("log1p undefined for num <= -1")
		}
		return math.Log1p(num), nil
	case "b":
		if num <= 0 {
			return 0, errors.New("logb undefined for num <= 0")
		}
		return math.Logb(num), nil
	default:
		return 0, errors.New("unsupported symbolic log base; use '1p' or 'b'")
	}
}
