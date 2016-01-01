package cellauto

import (
	"testing"
)

func BenchmarkGridNeighborsCenter(b *testing.B) {
	g := NewGrid(Point{3, 3})
	for i := 0; i < b.N; i++ {
		g.Neighbors(Point{1, 1})
	}
}

func BenchmarkGridNeighborsBorder(b *testing.B) {
	g := NewGrid(Point{3, 3})
	for i := 0; i < b.N; i++ {
		g.Neighbors(Point{0, 1})
	}
}

func BenchmarkGridNeighborsCorner(b *testing.B) {
	g := NewGrid(Point{3, 3})
	for i := 0; i < b.N; i++ {
		g.Neighbors(Point{0, 0})
	}
}

func BenchmarkGameStepSmall(b *testing.B) {
	benchmarkGameStep(b, NewGrid(Point{16, 16}))
}

func BenchmarkGameStepMedium(b *testing.B) {
	benchmarkGameStep(b, NewGrid(Point{128, 128}))
}

func BenchmarkGameStepLarge(b *testing.B) {
	benchmarkGameStep(b, NewGrid(Point{1024, 1024}))
}

func benchmarkGameStep(b *testing.B, g *Grid) {
	game := &Game{
		Rule: func(p Point, g *Grid) uint8 {
			return 0
		},
		Grid: g,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Step()
	}
}
