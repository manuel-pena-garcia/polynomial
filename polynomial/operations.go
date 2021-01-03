package polynomial

// Sum obtains a Polynomial representing the sum of the caller and other Polynomial
func (p Polynomial) Sum(other Polynomial) Polynomial {
	m := make(map[uint]float64)

	grade := max(p.grade, other.grade)

	for i := 0; i <= int(grade); i++ {
		sum := p.coefficients[uint(i)] + other.coefficients[uint(i)]

		if (sum-0.)*(sum-0.) > precission {
			m[uint(i)] = sum
		}
	}

	return Polynomial{m, getGradeFromMap(m)}
}

// Subtract obtains a Polynomial substracting other from the caller
func (p Polynomial) Subtract(other Polynomial) Polynomial {
	m := make(map[uint]float64)

	grade := max(p.grade, other.grade)

	for i := 0; i <= int(grade); i++ {
		sub := p.coefficients[uint(i)] - other.coefficients[uint(i)]

		if (sub-0.)*(sub-0.) > precission {
			m[uint(i)] = sub
		}
	}

	return Polynomial{m, getGradeFromMap(m)}
}

// Multiply obtains a Polynomial as a result of multiplying the caller and other Polynomial
func (p Polynomial) Multiply(other Polynomial) Polynomial {
	m := make(map[uint]float64)

	grade := p.grade + other.grade

	for i := 0; i <= int(p.grade); i++ {
		for j := 0; j <= int(other.grade); j++ {
			a := p.coefficients[uint(i)]
			b := other.coefficients[uint(j)]

			if (a-0.)*(a-0.) > precission && (b-0.)*(b-0.) > precission {
				m[uint(i+j)] += a * b
			}
		}
	}

	return Polynomial{m, grade}
}

func max(a uint, b uint) uint {
	if a > b {
		return a
	}
	return b
}
