package fuoco

import (
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {

	height := 100
	width := 150

	stateGrid := MakeStateGrid(height, width)
	elevationGrid := MakeParamGrid(height, width)
	fuelGrid := MakeParamGrid(height, width)
	moistureGrid := MakeParamGrid(height, width)

	SetStateGrid(&stateGrid, Ready)
	SetParamGrid(&elevationGrid, 100)
	SetParamGrid(&fuelGrid, 100)
	SetParamGrid(&moistureGrid, 100)

	config := FuocoConfig{
		NumCases:             11,
		NumIterations:        100,
		NumSamples:           10,
		Height:               height,
		Width:                width,
		TopographyFunc:       Spontaneous,
		WeatherFunc:          Spontaneous,
		FuelFunc:             Spontaneous,
		BurnoutFunc:          Spontaneous,
		InitialState:         stateGrid,
		InitialElevation:     elevationGrid,
		InitialFuel:          fuelGrid,
		InitialMoisture:      moistureGrid,
		InitialWindDirection: 1.2,
		InitialWindSpeed:     40.0,
	}

	f := New(config)
	f.Run()
	_ = f
}
