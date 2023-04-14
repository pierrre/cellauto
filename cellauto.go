// Package cellauto provide a cellular automaton implementation.
package cellauto

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/pierrre/go-libs/goroutine"
)

// Point is a point on a Grid.
type Point struct {
	X, Y int
}

// Add adds a Point to another.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// Grid is a grid of squares with multiple states.
//
// By convention, the X-axis goes from left to right, and the Y-axis goes from top to bottom.
type Grid struct {
	Size    Point
	Squares []uint8
}

// NewGrid creates a new Grid.
func NewGrid(size Point) *Grid {
	return &Grid{
		Size:    size,
		Squares: make([]uint8, size.X*size.Y),
	}
}

// SquareIndex return the internal index of a square.
func (g *Grid) index(p Point) int {
	return p.Y*g.Size.X + p.X
}

// Get returns the value of a square.
func (g *Grid) Get(p Point) uint8 {
	return g.Squares[g.index(p)]
}

// Set sets the value of a square.
func (g *Grid) Set(p Point, v uint8) {
	g.Squares[g.index(p)] = v
}

// Contains returns true if the Grid contains the Point, and false otherwise.
func (g *Grid) Contains(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < g.Size.X && p.Y < g.Size.Y
}

// Neighbors returns the neighbors value of a Point.
// Start from top-left, and go clockwise.
// Out of bounds neighbors are equal to 0.
func (g *Grid) Neighbors(p Point) [8]uint8 {
	if p.X > 0 && p.X < g.Size.X-1 && p.Y > 0 && p.Y < g.Size.Y-1 {
		return [8]uint8{
			g.Get(p.Add(Point{-1, -1})),
			g.Get(p.Add(Point{0, -1})),
			g.Get(p.Add(Point{1, -1})),
			g.Get(p.Add(Point{1, 0})),
			g.Get(p.Add(Point{1, 1})),
			g.Get(p.Add(Point{0, 1})),
			g.Get(p.Add(Point{-1, 1})),
			g.Get(p.Add(Point{-1, 0})),
		}
	}
	var res [8]uint8
	for i, q := range [8]Point{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	} {
		np := p.Add(q)
		if g.Contains(np) {
			res[i] = g.Get(np)
		}
	}
	return res
}

// Rule represents a rule applied to a Point on a Grid.
type Rule func(p Point, g *Grid) uint8

// Game is a cellular automaton game.
type Game struct {
	Rule    Rule
	Grid    *Grid
	tmpGrid *Grid
}

// Step runs the Game for 1 step.
func (g *Game) Step() {
	if g.tmpGrid == nil {
		g.tmpGrid = NewGrid(g.Grid.Size)
	}
	parallelAuto(g.Grid.Size, g.step)
	g.Grid, g.tmpGrid = g.tmpGrid, g.Grid
}

func (g *Game) step(min, max Point) {
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			p := Point{x, y}
			v := g.Rule(p, g.Grid)
			g.tmpGrid.Set(p, v)
		}
	}
}

func parallel(p Point, pr int, f func(min, max Point)) {
	if pr == 1 {
		f(Point{0, 0}, p)
		return
	}
	wg := new(sync.WaitGroup)
	for y := 0; y < pr; y++ {
		min := Point{0, p.Y * y / pr}
		max := Point{p.X, p.Y * (y + 1) / pr}
		if max.X > min.X && max.Y > min.Y {
			goroutine.WaitGroup(wg, func() {
				f(min, max)
			})
		}
	}
	wg.Wait()
}

func parallelAuto(p Point, f func(min, max Point)) {
	parallel(p, runtime.GOMAXPROCS(0), f)
}
