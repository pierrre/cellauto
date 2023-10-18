package gameoflife

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/cellauto"
)

func TestRuleGameOfLife(t *testing.T) {
	g := cellauto.NewGrid(cellauto.Point{X: 3, Y: 3})
	for _, p := range []cellauto.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
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
		assert.Equal(t, res, tc.expected)
	}
}
