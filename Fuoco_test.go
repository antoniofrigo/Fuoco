package Fuoco

import (
	"log"
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {
	f := New()
	height := 20
	width := 20
	grid := MakeInitialGrid(height, width)
	SetConstantFuel(&grid, 0.7)
	SetConstantElevation(&grid, 100)
	SetStateReady(&grid)

	config := FuocoConfig{
		NumCases:       10,
		NumIterations:  20,
		Sampling:       2,
		Height:         height,
		Width:          width,
		InitialGrid:    &grid,
		TopographyFunc: LinearIgnition,
		WeatherFunc:    One,
		FuelFunc:       One,
		BurnoutFunc:    LinearFuelBurnout,
	}
	err := f.SetConfig(&config)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Run()
	_ = err
}
