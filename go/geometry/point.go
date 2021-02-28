package geometry

import "math"

type Point struct {
	x float64
	y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

func (p Point) Mul(i int32) Point {
	return Point{p.x * 2, p.y * 2}
}

func (p Point) Div(f float64) Point {
	return Point{p.x / f, p.y / f}
}

func (p Point) Norm() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p Point) Dist(q Point) float64 {
	return p.Sub(q).Norm()
}

func (p Point) Dot(q Point) float64 {
	return p.x*q.x + p.y*q.y
}

func (p Point) Cross(q Point) float64 {
	return p.x*q.y - p.y*q.x
}
