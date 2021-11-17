package Fuoco

import "testing"

func TestLinearIgnition(t *testing.T) {
	grid := MakeInitialGrid(10, 10)
	SetStateReady(&grid)
	grid[5][5].State = Burning
	for i := 4; i <= 6; i++ {
		for j := 4; j <= 6; j++ {
			if i == 5 && j == 5 {
				continue
			}
			result := LinearIgnition(&grid, 0, 4, 4)
			if result != 1-1.0/8.0 {
				t.Error("LinearIgnition expected 1.0/8.0, received: ", result)
			}
		}
	}
}
