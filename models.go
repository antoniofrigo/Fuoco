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
	return 1.0 / 100.0
}

func LinearFuel(state [][]State, fuel [][]int, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if state[i][j] == Burning {
				p += 1.0 / 8.0
			}
		}
	}
	return float64(fuel[a-1][b-1]) * p / 100.0
}

func Moisture(state [][]State, moisture [][]int, a int, b int) float64 {
	p := 0.0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if state[i][j] == Burning {
				p += float64(moisture[a-1][b-1]) / 800.0
			}
		}
	}
	return p
}

func Adjacent(state [][]State, _ [][]int, a int, b int) float64 {
	count := 0
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if i == a && j == b {
				continue
			}
			if state[i][j] == Burning {
				count++
			}
		}
	}

	return 1 - math.Exp(-8.0*float64(count))
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
	return p
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
				p = math.Max(math.Cos(theta)+0.05, p)
			}
		}
	}
	return p
}

func PointedWind(state [][]State, speed float64, angle float64, a int, b int) float64 {
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
				p = math.Max(1/(50*math.Sin(theta)+1), p)
			}
		}
	}
	return p
}
