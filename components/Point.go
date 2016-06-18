package components

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) Neighbors(w, h int) []*Point {

	// default neighbors
	neighbors := []*Point{
		NewPoint(p.x - 1, p.y - 1), NewPoint(p.x, p.y - 1), NewPoint(p.x + 1, p.y - 1),
		NewPoint(p.x - 1, p.y), 			    NewPoint(p.x + 1, p.y),
		NewPoint(p.x - 1, p.y + 1), NewPoint(p.x , p.y + 1), NewPoint(p.x + 1, p.y + 1),
	}

	// collect all valid neighbors
	validNeighbors := []*Point{}
	for _, neighbor := range neighbors {
		if (neighbor.x >= 0) && (neighbor.x < w) &&
		(neighbor.y >= 0) && (neighbor.y < h) {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}

	return validNeighbors
}