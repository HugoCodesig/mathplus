package mathplus

// Factorial returns the factorial of a non-negative integer n.
// For n <= 1, it returns 1. This implementation uses recursion.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// IsPrime checks whether an integer n is a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GCD returns the greatest common divisor of integers a and b using the Euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of integers a and b.
// If either input is 0, it returns 0.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	absA := int(Abs(float64(a)))
	absB := int(Abs(float64(b)))
	return (absA / GCD(absA, absB)) * absB
}
