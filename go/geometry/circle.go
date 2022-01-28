package geometry

type Circle struct {
	p Point
	r float64
}

func EqualsCircle(a, b Circle) bool {
	return EqualsPoint(a.p, b.p) && Sign2(a.r, b.r) == 0
}

func CircumscribedCircle(p Point, q Point, r Point) Circle {
	a := (q.Sub(p)).Mul(2)
	b := (r.Sub(p)).Mul(2)
	c := Point{Dot(p, p) - Dot(q, q), Dot(p, p) - Dot(r, r)}
	tmp := Point{a.y*c.y - b.y*c.x, b.x*c.x - a.x*c.y}
	center := tmp.Div(Cross(a, b))
	return Circle{center, Dist(p, center)}
}
