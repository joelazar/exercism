package triangle

import "math"

type Kind int

const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

func triangleInequality(a, b, c float64) bool {
	if (a+b < c) || (a+c < b) || (b+c < a) {
		return false
	}
	return true
}

func checkTriangleData(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 ||
		!triangleInequality(a, b, c) ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	return true
}

// KindFromSides - calculate the triangle kind from side attributes
func KindFromSides(a, b, c float64) Kind {
	if !checkTriangleData(a, b, c) {
		return NaT
	}
	if a == b || b == c || a == c {
		if a == b && a == c {
			return Equ
		}
		return Iso
	}
	return Sca
}
