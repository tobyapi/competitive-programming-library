package geometry

import "math"

const EPS float64 = 1e-8

func Sign(a float64) int {
	if a < -EPS {
		return -1
	} else if EPS < a {
		return 1
	}
	return 0
}

func Sign2(a float64, b float64) int {
	if a < b-EPS {
		return -1
	} else if b+EPS < a {
		return 1
	}
	return 0
}

func Sqrt(a float64) float64 {
	return math.Sqrt(math.Max(a, 0.0))
}
