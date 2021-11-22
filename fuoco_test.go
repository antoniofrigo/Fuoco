package fuoco

import (
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {

	height := 10
	width := 15

	stateGrid := MakeStateGrid(height, width)
	elevationGrid := MakeParamGrid(height, width)
	fuelGrid := MakeParamGrid(height, width)
	moistureGrid := MakeParamGrid(height, width)

	SetStateGrid(&stateGrid, Ready)
	SetParamGrid(&elevationGrid, 100)
	SetParamGrid(&fuelGrid, 100)
	SetParamGrid(&moistureGrid, 100)

	config := FuocoConfig{
		NumCases:             1,
		NumIterations:        20,
		NumSample:            1,
		Height:               height,
		Width:                width,
		TopographyFunc:       LinearIgnition,
		WeatherFunc:          One,
		FuelFunc:             One,
		BurnoutFunc:          One,
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
