package cellauto

import (
	"testing"

	"github.com/pierrre/assert"
)

func TestPointAdd(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := p1.Add(p2)
	assert.Equal(t, p3, Point{4, 6})
	p4 := p2.Add(p1)
	assert.Equal(t, p4, p3)
}

func TestPointString(t *testing.T) {
	p := Point{1, 2}
	s := p.String()
	assert.Equal(t, s, "(1,2)")
}

func TestNewGrid(t *testing.T) {
	g := NewGrid(Point{2, 3})
	assert.Equal(t, g.Size, Point{2, 3})
	assert.SliceLen(t, g.Squares, 6)
}

func TestGridGetSet(t *testing.T) {
	g := NewGrid(Point{2, 3})
	g.Set(Point{1, 1}, 6)
	v := g.Get(Point{1, 1})
	assert.Equal(t, v, 6)
}

func TestGridContains(t *testing.T) {
	g := NewGrid(Point{2, 3})
	type TC struct {
		point    Point
		expected bool
	}
	for _, tc := range []TC{
		{
			point:    Point{0, 0},
			expected: true,
		},
		{
			point:    Point{1, 1},
			expected: true,
		},
		{
			point:    Point{1, 2},
			expected: true,
		},
		{
			point:    Point{-1, -1},
			expected: false,
		},
		{
			point:    Point{2, 3},
			expected: false,
		},
		{
			point:    Point{10, 10},
			expected: false,
		},
	} {
		res := g.Contains(tc.point)
		assert.Equal(t, res, tc.expected)
	}
}

func TestGridNeighbors(t *testing.T) {
	g := NewGrid(Point{3, 3})
	for y := range g.Size.Y {
		for x := range g.Size.X {
			g.Set(Point{x, y}, 1)
		}
	}
	type TC struct {
		point    Point
		expected [8]uint8
	}
	for _, tc := range []TC{
		{
			point:    Point{1, 1},
			expected: [8]uint8{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			point:    Point{0, 0},
			expected: [8]uint8{0, 0, 0, 1, 1, 1, 0, 0},
		},
		{
			point:    Point{0, 1},
			expected: [8]uint8{0, 1, 1, 1, 1, 1, 0, 0},
		},
	} {
		res := g.Neighbors(tc.point)
		assert.Equal(t, res, tc.expected)
	}
}

func TestGameStep(t *testing.T) {
	ctx := t.Context()
	g := &Game{
		Rule: func(p Point, g *Grid) uint8 {
			return 1
		},
		Grid: NewGrid(Point{10, 10}),
	}
	g.Step(ctx)
	for y := range g.Grid.Size.Y {
		for x := range g.Grid.Size.X {
			p := Point{x, y}
			v := g.Grid.Get(p)
			assert.Equal(t, v, 1)
		}
	}
}

func TestParallel(t *testing.T) {
	ctx := t.Context()
	p := Point{10, 10}
	f := func(minPoint, maxPoint Point) {}
	parallelAuto(ctx, p, f)
	parallel(ctx, p, 1, f)
	parallel(ctx, p, 2, f)
}
