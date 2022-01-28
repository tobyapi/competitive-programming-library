package geometry

import "math"

type Point struct {
	x float64
	y float64
}

func add(a, b float64) float64 {
	if math.Abs(a+b) < EPS*(math.Abs(a)+math.Abs(b)) {
		return 0
	}
	return a + b
}

func (p Point) Add(q Point) Point {
	return Point{add(p.x, q.x), add(p.y, q.y)}
}

func (p Point) Sub(q Point) Point {
	return Point{add(p.x, -q.x), add(p.y, -q.y)}
}

func (p Point) Mul(d float64) Point {
	return Point{p.x * d, p.y * d}
}

func (p Point) Div(f float64) Point {
	return Point{p.x / f, p.y / f}
}

func (p Point) Norm() float64 {
	return math.Sqrt(add(p.x*p.x, p.y*p.y))
}

func EqualsPoint(a, b Point) bool {
	return Sign(Dist(a, b)) == 0
}

func Dist(p, q Point) float64 {
	return p.Sub(q).Norm()
}

func Dot(p, q Point) float64 {
	return add(p.x*q.x, p.y*q.y)
}

func Cross(p, q Point) float64 {
	return add(p.x*q.y, -p.y*q.x)
}
