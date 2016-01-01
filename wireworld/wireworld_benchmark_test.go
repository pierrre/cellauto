package wireworld

import (
	"testing"

	"github.com/pierrre/cellauto"
)

func BenchmarkRuleEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rule(cellauto.Point{X: 0, Y: 0}, testGrid)
	}
}

func BenchmarkRuleHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rule(cellauto.Point{X: 2, Y: 1}, testGrid)
	}
}

func BenchmarkRuleTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rule(cellauto.Point{X: 3, Y: 1}, testGrid)
	}
}

func BenchmarkRuleConductorChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rule(cellauto.Point{X: 1, Y: 1}, testGrid)
	}
}

func BenchmarkRuleConductorNoChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rule(cellauto.Point{X: 4, Y: 1}, testGrid)
	}
}
