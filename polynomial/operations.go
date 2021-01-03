package polynomial

import (
	"math"
)

// Sum obtains a Polynomial representing the sum of the caller and other Polynomial
func (p Polynomial) Sum(other Polynomial) Polynomial {
	m := make(map[uint]float64)

	grade := max(p.grade, other.grade)

	for i := uint(0); i <= grade; i++ {
		sum := p.coefficients[i] + other.coefficients[i]

		if checkNotZero(p.coefficients[i]) {
			m[i] = sum
		}
	}

	return Polynomial{m, getGradeFromMap(m)}
}

// Subtract obtains a Polynomial substracting other from the caller
func (p Polynomial) Subtract(other Polynomial) Polynomial {
	m := make(map[uint]float64)

	grade := max(p.grade, other.grade)

	for i := uint(0); i <= grade; i++ {
		sub := p.coefficients[i] - other.coefficients[i]

		if checkNotZero(sub) {
			m[i] = sub
		}
	}

	return Polynomial{m, getGradeFromMap(m)}
}

// Multiply obtains a Polynomial as a result of multiplying the caller and other Polynomial
func (p Polynomial) Multiply(other Polynomial) Polynomial {
	m := make(map[uint]float64)

	grade := p.grade + other.grade

	for i := uint(0); i <= p.grade; i++ {
		for j := uint(0); j <= other.grade; j++ {
			a := p.coefficients[i]
			b := other.coefficients[j]

			if checkNotZero(a * b) {
				m[i+j] += a * b
			}
		}
	}

	return Polynomial{m, grade}
}

// Derive retrieves the derivative Polyinomial as a result of deriving the caller Polynomial
func (p Polynomial) Derive() Polynomial {
	m := make(map[uint]float64)

	for i := uint(0); i <= p.grade; i++ {
		if checkNotZero(p.coefficients[i+1]) {
			m[i] = p.coefficients[i+1] * (float64(i + 1))
		}
	}

	return Polynomial{m, p.grade - 1}
}

// Integrate retrieves the Polynomial result of the integration of the caller Polynomial
// The argument will be the constant of integration
func (p Polynomial) Integrate(constant float64) Polynomial {
	m := make(map[uint]float64)

	m[0] = constant

	for i := uint(0); i <= p.grade+1; i++ {
		if checkNotZero(p.coefficients[i]) {
			m[i+1] = p.coefficients[i] / float64(i+1)
		}
	}

	return Polynomial{m, p.grade + 1}
}

// Evaluate retrieves the value of the caller Polynomial expression in a given point x
func (p Polynomial) Evaluate(x float64) float64 {
	var sum float64

	for i := uint(0); i <= p.grade; i++ {
		if checkNotZero(p.coefficients[i]) {
			sum += p.coefficients[i] * math.Pow(x, float64(i))
		}
	}

	return sum
}

func max(a uint, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func checkNotZero(a float64) bool {
	return (a*a - 0.) > precission
}
