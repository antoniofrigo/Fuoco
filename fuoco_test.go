package Fuoco

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {
	f := New()
	height := 20
	width := 20
	grid := MakeInitialGrid(height, width)
	SetConstantFuel(&grid, 0.7)
	SetLinearElevation(&grid, 100)
	SetStateReady(&grid)
	for _, row := range grid {
		var s strings.Builder
		for _, cell := range row {
			s.WriteString(strconv.Itoa(cell.Elevation) + " ")
		}
	}

	config := FuocoConfig{
		NumCases:       10,
		NumIterations:  20,
		Sampling:       2,
		Height:         height,
		Width:          width,
		InitialGrid:    &grid,
		TopographyFunc: LinearElevationIgnition,
		WeatherFunc:    One,
		FuelFunc:       One,
		BurnoutFunc:    LinearFuelBurnout,
	}
	err := f.SetConfig(&config)
	if err != nil {
		log.Fatal(err)
	}
	stats, err := f.Run()
	_ = err
	PrintStats(&stats)
}
