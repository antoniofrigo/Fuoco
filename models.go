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
