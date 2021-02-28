package geometry

type Circle struct {
	p Point
	r float64
}

func NewCircle(x float64, y float64, r float64) *Circle {
	return &Circle{Point{x, y}, r}
}

func CircumscribedCircle(p Point, q Point, r Point) *Circle {
	a := (q.Sub(p)).Mul(2)
	b := (r.Sub(p)).Mul(2)
	c := Point{p.Dot(p) - q.Dot(q), p.Dot(p) - r.Dot(r)}
	tmp := Point{a.y*c.y - b.y*c.x, b.x*c.x - a.x*c.y}
	center := tmp.Div(a.Cross(b))
	return &Circle{center, p.Dist(center)}
}
