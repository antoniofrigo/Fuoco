package fuoco

func One(_ *[][]State, _ *[][]int, _ int, _ int) float64 {
	return 1.0
}

func Spontaneous(_ *[][]State, _ *[][]int, _ int, _ int) float64 {
	return 0.7937
}

func SpontaneousBurnout(_ *[][]State, _ *[][]int, _ int, _ int) float64 {
	return 0.5
}

func LinearIgnition(g *[][]State, _ *[][]int, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if (*g)[i][j] == Burning {
				p += 1.0 / 8.0
			}
		}
	}
	return 1 - p
}

// func LinearElevationIgnition(g *[][]int, _ int, a int, b int) float64 {
// 	alpha := 10
// 	p := 0.0
// 	for i := a - 1; i <= a+1; i++ {
// 		for j := b - 1; j <= b+1; j++ {
// 			if i == a && j == b {
// 				continue
// 			}
// 			if (*g)[i][j].State == Burning {
// 				diff := (*g)[a][b].Elevation - (*g)[i][j].Elevation
// 				if 0 <= diff && diff <= alpha {
// 					p += 1.0 / 8.0
// 				} else if -alpha <= diff && diff <= 0 {
// 					p += 1.0 / 16.0
// 				}
// 			}
// 		}
// 	}
// 	return 1 - p
// }

// func LinearFuelIgnition(fuel *[][]int, _ int, a int, b int) float64 {
// 	for i := a - 1; i <= a+1; i++ {
// 		for j := b - 1; j <= b+1; j++ {
// 			if i == a && j == b {
// 				continue
// 			}
// 			if (*g)[i][j].State == Burning {
// 				return (*g)[a][b].fuel
// 			}
// 		}
// 	}
// 	return 1
// }

// // No index sanitation performed
// func LinearFuelBurnout(fuel *[][]int, _ int, a int, b int) float64 {
// 	return 1 - (*fuel)[a][b]/10
// }
