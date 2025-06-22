// Package wireworld implements the "Wireworld" variant of the cellular automaton.
package wireworld

import (
	"context"

	"github.com/pierrre/cellauto"
)

// All available states.
const (
	StateEmpty = uint8(iota)
	StateHead
	StateTail
	StateConductor
)

// Rule is the [cellauto.Rule] for the "Wireworld" variant.
//
//nolint:gocyclo // Cyclomatic complexity is too high.
func Rule(p cellauto.Point, g *cellauto.Grid) uint8 {
	v := g.Get(p)
	//nolint:nestif // TODO fix nested if.
	if v == StateConductor {
		c := 0
		n := g.Neighbors(p)
		if n[0] == StateHead {
			c++
		}
		if n[1] == StateHead {
			c++
		}
		if n[2] == StateHead {
			c++
		}
		if n[3] == StateHead {
			c++
		}
		if n[4] == StateHead {
			c++
		}
		if n[5] == StateHead {
			c++
		}
		if n[6] == StateHead {
			c++
		}
		if n[7] == StateHead {
			c++
		}
		if c == 1 || c == 2 {
			return StateHead
		}
		return StateConductor
	}
	if v == StateEmpty {
		return StateEmpty
	}
	if v == StateHead {
		return StateTail
	}
	if v == StateTail {
		return StateConductor
	}
	return StateEmpty
}

// Game is specialized in Wireword.
type Game struct {
	Grid    *cellauto.Grid
	tmpGrid *cellauto.Grid
	points  []cellauto.Point
}

// Step runs 1 step.
func (g *Game) Step(ctx context.Context) {
	if g.tmpGrid == nil {
		g.tmpGrid = cellauto.NewGrid(g.Grid.Size)
	}
	if g.points == nil {
		g.initPoints()
	}
	g.step(g.points)
	g.Grid, g.tmpGrid = g.tmpGrid, g.Grid
}

func (g *Game) initPoints() {
	g.points = make([]cellauto.Point, 0)
	for y := range g.Grid.Size.Y {
		for x := range g.Grid.Size.X {
			p := cellauto.Point{X: x, Y: y}
			if g.Grid.Get(p) != StateEmpty {
				g.points = append(g.points, p)
			}
		}
	}
}

func (g *Game) step(ps []cellauto.Point) {
	for _, p := range ps {
		v := Rule(p, g.Grid)
		g.tmpGrid.Set(p, v)
	}
}
