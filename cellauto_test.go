package cellauto

import (
	"testing"
)

func TestPointAdd(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := p1.Add(p2)
	if p3 != (Point{4, 6}) {
		t.Fatal("not equal")
	}
	p4 := p2.Add(p1)
	if p4 != p3 {
		t.Fatal("not equal")
	}
}

func TestPointString(t *testing.T) {
	p := Point{1, 2}
	s := p.String()
	if s != "(1,2)" {
		t.Fatal("not equal")
	}
}

func TestNewGrid(t *testing.T) {
	g := NewGrid(Point{2, 3})
	if g.Size != (Point{2, 3}) {
		t.Fatal("not equal")
	}
	if len(g.Squares) != 6 {
		t.Fatal("not equal")
	}
}

func TestGridGetSet(t *testing.T) {
	g := NewGrid(Point{2, 3})
	g.Set(Point{1, 1}, 6)
	v := g.Get(Point{1, 1})
	if v != 6 {
		t.Fatal("not equal")
	}
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
		if res != tc.expected {
			t.Fatalf("unexpected result for %s: got %t, want %t", tc.point, res, tc.expected)
		}
	}
}

func TestGridNeighbors(t *testing.T) {
	g := NewGrid(Point{3, 3})
	for y := 0; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
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
		if res != tc.expected {
			t.Fatalf("unexpected result for %s: got %v, want %v", tc.point, res, tc.expected)
		}
	}
}

func TestGameStep(t *testing.T) {
	g := &Game{
		Rule: func(p Point, g *Grid) uint8 {
			return 1
		},
		Grid: NewGrid(Point{10, 10}),
	}
	g.Step()
	for y := 0; y < g.Grid.Size.Y; y++ {
		for x := 0; x < g.Grid.Size.X; x++ {
			p := Point{x, y}
			v := g.Grid.Get(p)
			if v != 1 {
				t.Fatalf("unexpected value for %s: got %d, want %d", p, v, 1)
			}
		}
	}
}

func TestParallel(t *testing.T) {
	p := Point{10, 10}
	f := func(min, max Point) {}
	parallelAuto(p, f)
	parallel(p, 1, f)
	parallel(p, 2, f)
}
