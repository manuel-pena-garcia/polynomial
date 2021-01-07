package polynomial

// Polynomial is a map where:
// Keys are the exponents. If not present, its coefficient is zero
// Values are the corresponding coefficients
type Polynomial struct {
	coefficients map[uint]float64
	grade        uint
}

const precission float64 = 0.0000000000001

// CreateFromMap allows to create a Polynomial from a map.
func CreateFromMap(m map[uint]float64) Polynomial {
	grade := getGradeFromMap(m)

	return Polynomial{m, grade}
}

// CreateFromSlice llows to create a Polynomial from a slice of coefficients
// The coefficientes must be ordered, starting from 0
// A coefficient which is zero will not be added to the Polynomial object
func CreateFromSlice(s []float64) Polynomial {
	ns := removeTrailingZeros(s)

	grade := uint(len(ns))

	m := make(map[uint]float64)

	for index, value := range ns {
		if checkNotZero(value) {
			m[uint(index)] = value
		}
	}

	return Polynomial{m, grade}
}

func removeTrailingZeros(s []float64) []float64 {
	l := len(s)

	if !checkNotZero(s[l-1] * s[l-1]) {
		return removeTrailingZeros(s[:l-2])
	}

	return s
}

func getGradeFromMap(m map[uint]float64) uint {
	var grade uint

	for key := range m {
		if key > grade {
			grade = key
		}
	}

	return grade
}
