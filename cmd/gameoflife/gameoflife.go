// Package gameoflife is a "Game of Life" application.
package main

import (
	"context"

	"github.com/nsf/termbox-go"
	"github.com/pierrre/cellauto"
	"github.com/pierrre/cellauto/wireworld"
	"github.com/pierrre/go-libs/goroutine"
)

//nolint:gocyclo // TODO: Fix this cyclomatic complexity.
func main() {
	ctx := context.Background()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	evQueue := make(chan termbox.Event)
	goroutine.Start(ctx, func(context.Context) {
		for {
			evQueue <- termbox.PollEvent()
		}
	})
	width, height := termbox.Size()

	game := &cellauto.Game{
		Rule: wireworld.Rule,
		Grid: cellauto.NewGrid(cellauto.Point{X: width, Y: height}),
	}
	for y := range game.Grid.Size.Y {
		for x := range game.Grid.Size.X {
			if (x+y)%2 == 0 {
				game.Grid.Set(cellauto.Point{X: x, Y: y}, 1)
			}
		}
	}

	for {
		select {
		case ev := <-evQueue:
			if ev.Type == termbox.EventKey {
				return
			}
		default:
		}

		for y := range game.Grid.Size.Y {
			for x := range game.Grid.Size.X {
				var bg termbox.Attribute
				if game.Grid.Get(cellauto.Point{X: x, Y: y}) > 0 {
					bg = termbox.ColorRed
				} else {
					bg = termbox.ColorDefault
				}
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, bg)
			}
		}
		err = termbox.Flush()
		if err != nil {
			panic(err)
		}

		game.Step(ctx)
	}
}
