package components

import (
	"sync"
	"fmt"
	"bytes"
)

// a board has a pointer to a map of string, *Cell and a lock mutex
type Board struct {
	mu sync.Mutex
	board map[Point]*Cell
}

// returns a pointer to a new board
func NewBoard() *Board {
	return &Board{sync.Mutex{}, map[Point]*Cell{}}
}

// returns the total cells alive
func (b *Board)TotalAlive() int {
	count := 0
	for _, p := range b.board {
		count+= p.Value()
	}
	return count
}

// sets the point on the board to alive
func (b *Board) SetAlive(p* Point) {
	b.mu.Lock()
	b.board[*p] = NewCell(true)
	b.mu.Unlock()
}


func (b *Board) Transfer(next *Board, cp <-chan Point, w *sync.WaitGroup) {
	for pv := range cp {
		p := &pv
		// get the next cell based on current cells neighbors
		c := b.GetCell(p).Next(b.AliveNeighbors(p))
		// if the cell is alive the next board cell should be
		if c.IsAlive() {
			next.SetAlive(p)
		}
	}
	// specify that the routine has finished
	w.Done()
}


func (b *Board) AliveNeighbors(p *Point) int {
	count := 0
	for _, pv := range p.Neighbors() {
		count+= b.GetCell(pv).Value()
	}
	return count
}


func (b *Board) GetCell(p *Point) *Cell {
	c:= b.board[*p]
	if c != nil {
		return c
	}
	return NewCell(false)
}


func (b *Board) Next() *Board {
	// return pointer to new board
	next := NewBoard()

	// wait for 4 go routines to finish
	w := &sync.WaitGroup{}
	w.Add(9)

	pointsChannel := pointsGenerator(b, w)

	// 8 tasks
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)
	go b.Transfer(next, pointsChannel, w)

	// blocks until count is at 0
	w.Wait()
	return next
}


func pointsGenerator(b *Board,  w *sync.WaitGroup) <-chan Point {
	pointsGenerator := make(chan Point)
	// start a go routine to pass each point in the board to the channel
	go func() {
		for p := range (b.board) {
			// first process the current point
			pointsGenerator <- p
			for _, n := range p.Neighbors() {
				// then update the neighbors based on the current point
				pointsGenerator <- *n
			}
		}
		close(pointsGenerator)
		w.Done()
	}()
	return pointsGenerator
}


// prints the board
func (b *Board) Print(w, h int) {
	var buffer bytes.Buffer
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if b.GetCell(NewPoint(x, y)).IsAlive() {
				buffer.WriteString("X")
			} else {
				buffer.WriteString(".")
			}
		}
		buffer.WriteString("\n")
	}
	fmt.Println(buffer.String())
}