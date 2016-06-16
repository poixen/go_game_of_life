package components

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) Neighbors() []*Point {
	return []*Point{
		NewPoint(p.x - 1, p.y - 1), NewPoint(p.x, p.y - 1), NewPoint(p.x + 1, p.y - 1),
		NewPoint(p.x - 1, p.y), 			    NewPoint(p.x + 1, p.y),
		NewPoint(p.x - 1, p.y + 1), NewPoint(p.x , p.y + 1), NewPoint(p.x + 1, p.y + 1),
	}
}