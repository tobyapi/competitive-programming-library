package geometry

type Polygon struct {
	v []Point
}

func NewPolygonEmpty(size int) *Polygon {
	vertexes := make([]Point, size)
	return &Polygon{vertexes}
}

func NewPolygon(points ...Point) *Polygon {
	return &Polygon{points}
}

func (p *Polygon) Prev(i int) Point {
	return p.v[(i+len(p.v)-1)%len(p.v)]
}

func (p *Polygon) Curr(i int) Point {
	return p.v[i%len(p.v)]
}

func (p *Polygon) Next(i int) Point {
	return p.v[(i+1)%len(p.v)]
}

func (p *Polygon) Size() {
	return len(p.v)
}

func (pol *Polygon) Contains(p Point) int {
	in := false
	for i := 0; i < pol.Size(); i++ {
		a := pol.Curr(i).Sub(p)
		b := pol.Next(i).Sub(p)
		if b.y < a.y {
			a, b = b, a
		}
		if a.y <= 0 && 0 < b.y && a.Cross(b) < 0 {
			in = !in
		}
		if a.Cross(b) == 0 && a.Dot(b) <= 0 {
			return 2
		}
	}
	if in {
		return 1
	}
	return 0
}
