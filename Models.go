package Fuoco

func One(_ *FuocoGrid, _ int, _ int, _ int) float64 {
	return 1.0
}

func Spontaneous(_ *FuocoGrid, _ int, _ int, _ int) float64 {
	return 0.7937
}

func SpontaneousBurnout(_ *FuocoGrid, _ int, _ int, _ int) float64 {
	return 0.5
}

func LinearIgnition(g *FuocoGrid, _ int, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if (*g)[i][j].State == Burning {
				p += 1.0 / 8.0
			}
		}
	}
	return 1 - p
}

func LinearFuelIgnition(g *FuocoGrid, _ int, a int, b int) float64 {
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if (*g)[i][j].State == Burning {
				return (*g)[a][b].Fuel
			}
		}
	}
	return 1
}

// No index sanitation performed
func LinearFuelBurnout(g *FuocoGrid, _ int, a int, b int) float64 {
	return 1 - (*g)[a][b].Fuel/10
}
