// Package wireworld is a WireWorld application.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/pierrre/cellauto"
	"github.com/pierrre/cellauto/wireworld"
)

func main() {
	parseFlags()
	g := newGame()
	run(g)
}

var (
	flagSteps = 100
	argFile   = "primes.wi"
)

func parseFlags() {
	flag.IntVar(&flagSteps, "steps", flagSteps, "Steps")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [FILE]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	f := flag.Arg(0)
	if f != "" {
		argFile = f
	}
}

func newGame() *wireworld.Game {
	return &wireworld.Game{
		Grid: loadGrid(),
	}
}

//nolint:gocyclo // TODO: Fix this cyclomatic complexity.
func loadGrid() *cellauto.Grid {
	b, err := os.ReadFile(argFile)
	if err != nil {
		log.Panic(err)
	}
	s := bufio.NewScanner(bytes.NewReader(b))
	if !s.Scan() {
		log.Panic("unexpected end")
	}
	var width int
	var height int
	_, err = fmt.Sscanf(s.Text(), "%d %d", &width, &height)
	if err != nil {
		log.Panic(err)
	}
	g := cellauto.NewGrid(cellauto.Point{X: width, Y: height})
	y := 0
	for s.Scan() {
		for x, c := range s.Text() {
			p := cellauto.Point{X: x, Y: y}
			if !g.Contains(p) {
				log.Panicf("grid of size %s does not contain point %s", g.Size, p)
			}
			var v uint8
			switch c {
			case ' ':
				v = wireworld.StateEmpty
			case '@':
				v = wireworld.StateHead
			case '~':
				v = wireworld.StateTail
			case '#':
				v = wireworld.StateConductor
			default:
				log.Panicf("invalid character '%c' at %dx%d", c, x, y)
			}
			g.Set(p, v)
		}
		y++
	}
	return g
}

func run(g *wireworld.Game) {
	for step := 0; ; step++ {
		if step%flagSteps == 0 {
			log.Printf("step: %d", step)
			writeImage(g.Grid, step)
		}
		g.Step()
	}
}

func writeImage(g *cellauto.Grid, step int) {
	im := image.NewRGBA(image.Rect(0, 0, g.Size.X, g.Size.Y))
	for y := 0; y < g.Size.Y; y++ {
		for x := 0; x < g.Size.X; x++ {
			var c color.RGBA
			switch g.Get(cellauto.Point{X: x, Y: y}) {
			case wireworld.StateEmpty:
				c = color.RGBA{0, 0, 0, 0xff}
			case wireworld.StateHead:
				c = color.RGBA{0, 0, 0xff, 0xff}
			case wireworld.StateTail:
				c = color.RGBA{0xff, 0, 0, 0xff}
			case wireworld.StateConductor:
				c = color.RGBA{0xff, 0xff, 0, 0xff}
			default:
				c = color.RGBA{0, 0xff, 0, 0xff}
			}
			im.SetRGBA(x, y, c)
		}
	}
	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		log.Panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("out_%010d.png", step), buf.Bytes(), 0o644) //nolint: gosec // Allow all users to read the file.
	if err != nil {
		log.Panic(err)
	}
}
