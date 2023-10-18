package wireworld

import (
	"testing"

	"github.com/pierrre/assert"
	"github.com/pierrre/cellauto"
)

var testGrid *cellauto.Grid

func init() {
	testGrid = cellauto.NewGrid(cellauto.Point{X: 6, Y: 3})
	testGrid.Set(cellauto.Point{X: 1, Y: 1}, StateConductor)
	testGrid.Set(cellauto.Point{X: 2, Y: 1}, StateHead)
	testGrid.Set(cellauto.Point{X: 3, Y: 1}, StateTail)
	testGrid.Set(cellauto.Point{X: 4, Y: 1}, StateConductor)
	testGrid.Set(cellauto.Point{X: 5, Y: 2}, 255)
}

func TestRuleWireworld(t *testing.T) {
	type TC struct {
		point    cellauto.Point
		expected uint8
	}
	for _, tc := range []TC{
		{
			point:    cellauto.Point{X: 0, Y: 0},
			expected: StateEmpty,
		},
		{
			point:    cellauto.Point{X: 1, Y: 1},
			expected: StateHead,
		},
		{
			point:    cellauto.Point{X: 2, Y: 1},
			expected: StateTail,
		},
		{
			point:    cellauto.Point{X: 3, Y: 1},
			expected: StateConductor,
		},
		{
			point:    cellauto.Point{X: 4, Y: 1},
			expected: StateConductor,
		},
		{
			point:    cellauto.Point{X: 5, Y: 2},
			expected: StateEmpty,
		},
	} {
		res := Rule(tc.point, testGrid)
		assert.Equal(t, res, tc.expected)
	}
}
