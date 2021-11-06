package Fuoco

import (
	"log"
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {
	f := New()
	height := 100
	width := 100
	grid := MakeInitialGrid(height, width)
	SetConstantFuel(&grid, 0.5)
	SetConstantElevation(&grid, 100)
	SetStateReady(&grid)

	config := FuocoConfig{
		NumIterations:  10,
		Height:         height,
		Width:          width,
		InitialGrid:    &grid,
		TopographyFunc: Constant,
		WeatherFunc:    Constant,
		FuelFunc:       Constant,
	}
	err := f.SetConfig(&config)
	if err != nil {
		log.Fatalf("Unable to create configuration")
	}
	err = f.Run()
	_ = err
}
