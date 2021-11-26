package fuoco

import (
	"fmt"
	"math"
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

func SetParamGridSlide(paramGrid *[][]int) {
	for i, row := range *paramGrid {
		for j, _ := range row {
			(*paramGrid)[i][j] = i + j
		}
	}
}

func SetParamGridCircular(paramGrid *[][]int) {
	height := len(*paramGrid)
	width := len((*paramGrid)[0])
	for i, row := range *paramGrid {
		for j, _ := range row {
			(*paramGrid)[i][j] = (i-height/2)*(i-height/2) + (j-width/2)*(j-width/2)
		}
	}
}

func SetParamGridParabolicCylinder(paramGrid *[][]int) {
	height := len(*paramGrid)
	width := len((*paramGrid)[0])
	for i, row := range *paramGrid {
		for j, _ := range row {
			x := float64(i - height/2)
			y := float64(j - width/2)
			(*paramGrid)[i][j] = int((x + y) * (x + y + 1) / 100)
		}
	}
}

// Sets parameter between 0 and 100 with 100 at the bottom and
// 0 at the highest elevation with a linear gradient
func SetParamGridByReverseElevation(paramGrid *[][]int, elevation [][]int) {
	minElevation := math.MaxInt64
	maxElevation := math.MinInt64
	for _, row := range elevation {
		for _, value := range row {
			if value < minElevation {
				minElevation = value
			} else if value > maxElevation {
				maxElevation = value
			}
		}
	}

	m := 100.0 / float64(maxElevation-minElevation)
	for i, row := range *paramGrid {
		for j, _ := range row {
			(*paramGrid)[i][j] = int(100.0 - m*float64(elevation[i][j]-minElevation))
		}
	}
}

func SetParamGridValley(paramGrid *[][]int) {
	height := len(*paramGrid)
	width := len((*paramGrid)[0])
	for i, row := range *paramGrid {
		for j, _ := range row {
			x := float64(i - height/2)
			y := float64(j - width/2)
			(*paramGrid)[i][j] = int((x*x + 10*y*y) / (x*x + 100))
		}
	}
}

func SetParamGridBooth(paramGrid *[][]int) {
	height := len(*paramGrid)
	width := len((*paramGrid)[0])
	for i, row := range *paramGrid {
		for j, _ := range row {
			x := float64(i - height/2)
			y := float64(j - width/2)
			(*paramGrid)[i][j] = int(math.Pow((x+2*y-7), 2) + math.Pow((2*x+y-5), 2))
		}
	}
	ScaleParamGrid(paramGrid)
}

// Scales parameter grid down to the range 0 to 100
func ScaleParamGrid(paramGrid *[][]int) {
	maxValue, minValue := 0, math.MaxInt64
	for i, row := range *paramGrid {
		for j, _ := range row {
			if maxValue < (*paramGrid)[i][j] {
				maxValue = (*paramGrid)[i][j]
			} else if minValue > (*paramGrid)[i][j] {
				minValue = (*paramGrid)[i][j]
			}
		}
	}
	maxValue += minValue

	for i, row := range *paramGrid {
		for j, _ := range row {
			(*paramGrid)[i][j] = int(100.0 / float64(maxValue) * float64((*paramGrid)[i][j]+minValue))
		}
	}

}
