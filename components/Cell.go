package components


// A Cell is something that can be alive or dead.
// The contains some logic that lets it know if it
// should be alive or dead Next.
type Cell struct {
	alive bool
}

// Calling `Cell` will create a new struct, we then return
// a pointer to that struct, rather than a copy (which would
// happen if we just returned `Cell`)
func NewCell(alive bool) *Cell {
	return &Cell{alive}
}

// a live cell with 3 neighbors lives in the next step
// a dead cell with 3 neighbors lives in the next step
// a live cell with 2 or 3 neighbors lives in the next step
func (c *Cell) Next(count int) *Cell {
	return &Cell{count == 3 || (c.alive && count == 2)}
}

// returns the integer value of the Cell state
func (c *Cell) Value() int {
	if c.alive {
		return 1
	}
	return 0
}

// returns the state of the Cell
func (c *Cell) IsAlive() bool {
	return c.alive
}
