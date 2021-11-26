package fuoco

import (
	"math"
)

func OneParam(_ [][]State, _ [][]int, _ int, _ int) float64 {
	return 1.0
}

func SpontaneousParam(_ [][]State, _ [][]int, _ int, _ int) float64 {
	return 0.7937
}

func SpontaneousBurnout(_ [][]State, _ [][]int, _ int, _ int) float64 {
	return 0.5
}

func UniformBurnout(_ [][]State, _ [][]int, _ int, _ int) float64 {
	return 1 - 1/10.0
}

func Moisture(state [][]State, moisture [][]int, a int, b int) float64 {
	p := 0.0
	if state[a][b] == Burning {
		p = float64(moisture[a][b]) / 100.0
	}
	return 1 - p
}

func Adjacent(g [][]State, _ [][]int, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if g[i][j] == Burning {
				p += 1.0 / 8.0
			}
		}
	}
	return 1 - p
}

func AdjacentBiasElevation(g [][]State, _ [][]int, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if g[i][j] == Burning {
				p += 1.0 / 8.0
			}
		}
	}
	return 1 - p
}

func OneWind(_ [][]State, _ float64, _ float64, _ int, _ int) float64 {
	return 1.0
}

func TrigonometricWind(state [][]State, speed float64, angle float64, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if state[i][j] == Burning {
				dx := i - a
				dy := j - b
				theta := math.Atan(float64(dy) / float64(dx))
				if dy < 0 {
					theta += math.Pi
				}
				theta -= angle
				p += math.Cos(theta) / 8
			}
		}
	}
	return 1 - p
}
