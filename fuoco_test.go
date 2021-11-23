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
		NumCases:             5,
		NumIterations:        20,
		NumSample:            10,
		ImageScale:           5,
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
