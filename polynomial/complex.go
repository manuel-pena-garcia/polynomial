package polynomial

import (
	"math"
)

// Real represents the real part of a Complex number
type Real float64

// Imaginary represents the imaginary part of a Complex number
type Imaginary float64

func (i Imaginary) multiply(other Imaginary) Imaginary {
	return i * other * Imaginary(-1.)
}

// Complex represents a Complex number composed of real and imaginary parts
type Complex struct {
	real      Real
	imaginary Imaginary
}

// NewComplex is a factory that can be used to create Complex numbers given real and imaginary parts
func NewComplex(r float64, i float64) Complex {
	c := Complex{}

	c.real = Real(r)
	c.imaginary = Imaginary(i)

	return c
}

// Sum adds another Complex number to the caller one
func (c *Complex) Sum(other *Complex) {
	c.real += other.real
	c.imaginary += other.imaginary
}

// Subtract subtracts another Complex number from the caller one
func (c *Complex) Subtract(other *Complex) {
	c.real -= other.real
	c.imaginary -= other.imaginary
}

// Multiply multiplies the caller Complex number by another one
func (c *Complex) Multiply(other *Complex) {
	r := c.real*other.real + Real(c.imaginary.multiply(other.imaginary))
	i := Imaginary(c.real)*other.imaginary + c.imaginary*Imaginary(other.real)

	c.real, c.imaginary = r, i
}

// MultiplyByScalar multiplies the caller Complex by a float64 number
func (c *Complex) MultiplyByScalar(f float64) {
	c.real *= Real(f)
	c.imaginary *= Imaginary(f)
}

// Pow elevates a Complex number to the given exponent
func (c Complex) Pow(exponent uint) Complex {
	for i := uint(0); i < exponent; i++ {
		c.Multiply(&c)
	}

	return c
}

// Module gets the module of a complex number
func (c *Complex) Module() float64 {
	return math.Sqrt(float64(c.real*c.real) + float64(c.imaginary.multiply(c.imaginary)))
}
