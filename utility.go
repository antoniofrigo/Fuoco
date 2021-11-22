package fuoco

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

// func MakeInitialGrid(width, height int) [][]int {
// 	var grid [][]int = make([][]Cell, width)
// 	for i := 0; i < height; i++ {
// 		grid[i] = make([]Cell, width)
// 	}
// 	return grid
// }

// func SetConstantFuel(grid *[][]int, value float64) error {
// 	if 0.0 > value && value > 1.0 {
// 		return errors.New("Fuel level out of range: must be between 0 and 1")
// 	}

// 	for i := 0; i < len(*grid); i++ {
// 		for j := 0; j < len((*grid)[0]); j++ {
// 			(*grid)[i][j].Fuel = value
// 		}
// 	}
// 	return nil
// }

// // Sets the elevation in meters of a given grid
// func SetConstantElevation(grid *[][]int, value int) error {
// 	for i := 0; i < len(*grid); i++ {
// 		for j := 0; j < len((*grid)[0]); j++ {
// 			(*grid)[i][j].Elevation = value
// 		}
// 	}
// 	return nil
// }

// func SetLinearElevation(grid *[][]int, value int) error {
// 	dx := value / (len(*grid))
// 	dy := value / (len((*grid)[0]))
// 	for i := 0; i < len(*grid); i++ {
// 		for j := 0; j < len((*grid)[0]); j++ {
// 			(*grid)[i][j].Elevation = i*dx + j*dy
// 		}
// 	}
// 	return nil
// }

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

// func SetValleyElevation(grid *[][]int, value int) error {
// 	dx := value / (len(*grid))
// 	dy := value / (len((*grid)[0]))
// 	for i := 0; i < len(*grid); i++ {
// 		for j := 0; j < len((*grid)[0]); j++ {
// 			(*grid)[i][j].Elevation = abs(i*dx) + abs(j*dy)
// 		}
// 	}
// 	return nil
// }

// // Sets the elevation in meters of a given cell
// func SetStateReady(grid *[][]int) error {
// 	for i := 0; i < len(*grid); i++ {
// 		for j := 0; j < len((*grid)[0]); j++ {
// 			(*grid)[i][j].State = Ready
// 		}
// 	}
// 	return nil
// }

// func PrintGrid(grid *[][]int) {
// 	for i := 0; i < len(*grid); i++ {
// 		var s strings.Builder
// 		for j := 0; j < len((*grid)[0]); j++ {
// 			var c byte
// 			switch (*grid)[i][j].State {
// 			case NoFuel:
// 				c = '_'
// 			case Ready:
// 				c = '%'
// 			case Burning:
// 				c = '#'
// 			case BurnedOut:
// 				c = '.'
// 			}
// 			s.WriteByte(c)
// 		}
// 		fmt.Println(s.String())
// 	}
// }
