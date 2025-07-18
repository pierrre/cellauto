// Package cellauto provide a cellular automaton implementation.
package cellauto

import (
	"context"
	"fmt"
)

// Point is a point on a [Grid].
type Point struct {
	X, Y int
}

// Add adds a [Point] to another.
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

// NewGrid creates a new [Grid].
func NewGrid(size Point) *Grid {
	return &Grid{
		Size:    size,
		Squares: make([]uint8, size.X*size.Y),
	}
}

func (g *Grid) index(p Point) int {
	return p.Y*g.Size.X + p.X
}

// Get returns the value of a [Point].
func (g *Grid) Get(p Point) uint8 {
	return g.Squares[g.index(p)]
}

// Set sets the value of a [Point].
func (g *Grid) Set(p Point, v uint8) {
	g.Squares[g.index(p)] = v
}

// Contains returns true if the Grid contains the [Point], and false otherwise.
func (g *Grid) Contains(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < g.Size.X && p.Y < g.Size.Y
}

// Neighbors returns the neighbors value of a [Point].
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

// Rule represents a rule applied to a [Point] on a [Grid].
type Rule func(p Point, g *Grid) uint8

// Game is a cellular automaton game.
type Game struct {
	Rule    Rule
	Grid    *Grid
	tmpGrid *Grid
}

// Step runs the [Game] for 1 step.
func (g *Game) Step(ctx context.Context) {
	if g.tmpGrid == nil {
		g.tmpGrid = NewGrid(g.Grid.Size)
	}
	g.step(Point{0, 0}, g.Grid.Size)
	g.Grid, g.tmpGrid = g.tmpGrid, g.Grid
}

func (g *Game) step(minPoint, maxPoint Point) {
	for y := minPoint.Y; y < maxPoint.Y; y++ {
		for x := minPoint.X; x < maxPoint.X; x++ {
			p := Point{x, y}
			v := g.Rule(p, g.Grid)
			g.tmpGrid.Set(p, v)
		}
	}
}
