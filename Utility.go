package Fuoco

import (
	"errors"
	"fmt"
	"strings"
)

func MakeInitialGrid(width, height int) FuocoGrid {
	var grid FuocoGrid = make([][]Cell, width)
	for i := 0; i < height; i++ {
		grid[i] = make([]Cell, width)
	}
	return grid
}

func SetConstantFuel(grid *FuocoGrid, value float64) error {
	if 0.0 > value && value > 1.0 {
		return errors.New("Fuel level out of range: must be between 0 and 1")
	}

	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[0]); j++ {
			(*grid)[i][j].Fuel = value
		}
	}
	return nil
}

// Sets the elevation in meters of a given grid
func SetConstantElevation(grid *FuocoGrid, value int) error {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[0]); j++ {
			(*grid)[i][j].Elevation = value
		}
	}
	return nil
}

func SetLinearElevation(grid *FuocoGrid, value int) error {
	dx := value / (len(*grid))
	dy := value / (len((*grid)[0]))
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[0]); j++ {
			(*grid)[i][j].Elevation = i*dx + j*dy
		}
	}
	return nil
}

// Sets the elevation in meters of a given cell
func SetStateReady(grid *FuocoGrid) error {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[0]); j++ {
			(*grid)[i][j].State = Ready
		}
	}
	return nil
}

func PrintGrid(grid *FuocoGrid) {
	for i := 0; i < len(*grid); i++ {
		var s strings.Builder
		for j := 0; j < len((*grid)[0]); j++ {
			var c byte
			switch (*grid)[i][j].State {
			case NoFuel:
				c = '_'
			case Ready:
				c = '%'
			case Burning:
				c = '#'
			case BurnedOut:
				c = '.'
			}
			s.WriteByte(c)
		}
		fmt.Println(s.String())
	}
}
