package models

import (
	"bytes"
	"fmt"
)

// Integers representing the different states of a fire.
const (
	NoFuel    = iota
	Ready     = iota
	Burning   = iota
	BurnedOut = iota
	Buffer    = iota
)

// A "cell" within our grid. The State is given by one of the enum variants
// , elevation is in meters, and fuel level is between 1 and 100.
type Element struct {
	State     int
	Elevation int
	FuelLevel uint8
}

// Grid of Elements containing state. We maintain a ring of buffer
// cells around the perimeter
type Grid [][]Element

// Creates a new grid with a given height and width in cells, with
// a ring of buffer cells around them.
func NewGrid(height uint8, width uint8) *Grid {
	g := make(Grid, height+2)
	for i := range g {
		g[i] = make([]Element, width+2)
	}

	for i := 0; i < len(g); i++ {
		g[i][0].State = Buffer
		g[i][len(g[0])-1].State = Buffer
	}
	for i := 0; i < len(g[0]); i++ {
		g[0][i].State = Buffer
		g[len(g[0])-1][i].State = Buffer
	}

	return &g
}

func (g *Grid) PrintState(msg string) {
	fmt.Println(msg)

	for i := range *g {
		var buffer bytes.Buffer
		for j := range (*g)[i] {
			c := (*g)[i][j].State
			if c == NoFuel {
				buffer.WriteString("_")
			} else if c == Burning {
				buffer.WriteString("#")
			} else if c == BurnedOut {
				buffer.WriteString(".")
			} else if c == Buffer {
				buffer.WriteString("x")
			} else if c == Ready {
				buffer.WriteString("%")
			}
		}
		fmt.Println(buffer.String())
	}
}
