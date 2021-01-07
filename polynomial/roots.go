package polynomial

import "math"

// IsRoot checks if a float64 value is a root of the Polynomial
func (p Polynomial) IsRoot(x float64) bool {
	var sum float64

	for i := uint(0); i <= p.grade; i++ {
		sum += p.coefficients[i] * math.Pow(x, float64(i))
	}

	if checkNotZero(sum) {
		return false
	}

	return true
}

// Simplify returns a simplified version of the Polynomial, if possible
// If not, returns the same Polynomial, as it is the simplest version of itself
func Simplify(p Polynomial) Polynomial {
	if checkNotZero(p.coefficients[0]) {
		return p
	}

	for i := uint(0); i <= p.grade; i++ {
		p.coefficients[i] = p.coefficients[i+1]
	}

	return Simplify(p)
}
