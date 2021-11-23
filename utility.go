package fuoco

import (
	"fmt"
	"strings"
)

func printState(stateGrid *[][]State) {
	for _, row := range *stateGrid {
		var s strings.Builder
		for _, value := range row {
			switch value {
			case NoFuel:
				s.WriteString("_")
			case Ready:
				s.WriteString("%")
			case Burning:
				s.WriteString("#")
			case BurnedOut:
				s.WriteString(".")
			}
		}
		fmt.Println(s.String())
	}
}

// Creates an empty state grid
func MakeStateGrid(height, width int) [][]State {
	stateGrid := make([][]State, height)
	for i, _ := range stateGrid {
		stateGrid[i] = make([]State, width)
	}
	return stateGrid
}

// Sets all values of a given grid to a certain state
func SetStateGrid(stateGrid *[][]State, state State) {
	for i, row := range *stateGrid {
		for j, _ := range row {
			(*stateGrid)[i][j] = state
		}
	}
}

// Creates an empty parameter grid
func MakeParamGrid(height, width int) [][]int {
	paramGrid := make([][]int, height)
	for i, _ := range paramGrid {
		paramGrid[i] = make([]int, width)
	}
	return paramGrid
}

func SetParamGrid(paramGrid *[][]int, param int) {
	for i, row := range *paramGrid {
		for j, _ := range row {
			(*paramGrid)[i][j] = param
		}
	}
}
