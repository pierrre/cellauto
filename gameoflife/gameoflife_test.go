package gameoflife

import (
	"testing"

	"github.com/pierrre/cellauto"
)

func TestRuleGameOfLife(t *testing.T) {
	g := cellauto.NewGrid(cellauto.Point{X: 3, Y: 3})
	for _, p := range []cellauto.Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
	} {
		g.Set(p, 1)
	}
	type TC struct {
		point    cellauto.Point
		expected uint8
	}
	for _, tc := range []TC{
		{
			point:    cellauto.Point{X: 0, Y: 0},
			expected: 1,
		},
		{
			point:    cellauto.Point{X: 1, Y: 0},
			expected: 0,
		},
		{
			point:    cellauto.Point{X: 2, Y: 0},
			expected: 1,
		},
		{
			point:    cellauto.Point{X: 0, Y: 1},
			expected: 1,
		},
		{
			point:    cellauto.Point{X: 1, Y: 1},
			expected: 0,
		},
		{
			point:    cellauto.Point{X: 2, Y: 1},
			expected: 1,
		},
		{
			point:    cellauto.Point{X: 0, Y: 2},
			expected: 0,
		},
		{
			point:    cellauto.Point{X: 1, Y: 2},
			expected: 0,
		},
		{
			point:    cellauto.Point{X: 2, Y: 2},
			expected: 0,
		},
	} {
		res := Rule(tc.point, g)
		if res != tc.expected {
			t.Fatalf("unexpected result for %s: got %v, want %v", tc.point, res, tc.expected)
		}
	}
}
