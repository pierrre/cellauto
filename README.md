# Cell Auto

Cellular automaton library for Go (Golang).

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/cellauto.svg)](https://pkg.go.dev/github.com/pierrre/cellauto)

## Features

- [Grid](https://pkg.go.dev/github.com/pierrre/cellauto#Grid) of squares with multiple states
- [Point](https://pkg.go.dev/github.com/pierrre/cellauto#Point) coordinates and [neighbors](https://pkg.go.dev/github.com/pierrre/cellauto#Grid.Neighbors)
- Customizable [Rule](https://pkg.go.dev/github.com/pierrre/cellauto#Rule) functions
- [Game](https://pkg.go.dev/github.com/pierrre/cellauto#Game) with context-aware [Step](https://pkg.go.dev/github.com/pierrre/cellauto#Game.Step)
- [Game of Life](https://pkg.go.dev/github.com/pierrre/cellauto/gameoflife) variant (Conway's rules)
- [Wireworld](https://pkg.go.dev/github.com/pierrre/cellauto/wireworld) variant with optimized stepping
- Commands: `cmd/gameoflife` (interactive terminal) and `cmd/wireworld` (PNG output)

## Usage

```bash
# Module install
go get github.com/pierrre/cellauto@latest

# Local build
make build
./build/gameoflife
./build/wireworld cmd/wireworld/primes.wi

# Remote install
go install github.com/pierrre/cellauto/cmd/gameoflife@latest
go install github.com/pierrre/cellauto/cmd/wireworld@latest
```

### Commands

- `gameoflife` - interactive terminal simulation (termbox); press any key to exit
- `wireworld` - runs a Wireworld simulation and writes PNG frames

The `wireworld` command accepts a `.wi` file as its first argument (default: `primes.wi`) and a `-steps` flag (default: `100`) controlling how often a PNG frame is written.
Frames are written as `out_XXXXXXXXXX.png` (zero-padded step number).

## .wi file format

The `.wi` file format describes a Wireworld grid.

1. First line: `width height` (two integers separated by a space)
2. Following lines: one line per row, where each character represents a cell state:
   - ` ` (space) - empty (background)
   - `#` - conductor (wire)
   - `@` - electron head
   - `~` - electron tail

The number of rows must match `height`, and each row maps to x coordinates starting at 0.
An example file is provided at `cmd/wireworld/primes.wi`.
