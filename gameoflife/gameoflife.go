// Package gameoflife implements the "Game Of Life" variant of the cellular automaton.
package gameoflife

import (
	"github.com/pierrre/cellauto"
)

// Rule is the [cellauto.Rule] for the "Game Of Life" variant.
func Rule(p cellauto.Point, g *cellauto.Grid) uint8 {
	a := 0
	for _, v := range g.Neighbors(p) {
		if v > 0 {
			a++
		}
	}
	if a < 2 || a > 3 {
		return 0
	}
	if a == 3 {
		return 1
	}
	return g.Get(p)
}
