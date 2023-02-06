// Package cellautotest provides internal tools for tests.
package cellautotest

import (
	"github.com/pierrre/assert/ext/davecghspew"
	"github.com/pierrre/assert/ext/pierrrecompare"
	"github.com/pierrre/assert/ext/pierrreerrors"
)

// Configure configures tools used in tests.
func Configure() {
	pierrrecompare.Configure()
	davecghspew.ConfigureDefault()
	pierrreerrors.Configure()
}
