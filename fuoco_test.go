package fuoco

import (
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {
	height := 400
	width := 400

	stateGrid := MakeStateGrid(height, width)
	elevationGrid := MakeParamGrid(height, width)
	fuelGrid := MakeParamGrid(height, width)
	moistureGrid := MakeParamGrid(height, width)

	SetStateGrid(&stateGrid, Ready)
	SetParamGridBooth(&elevationGrid)
	SetParamGridByReverseElevation(&fuelGrid, elevationGrid)
	SetParamGridByReverseElevation(&moistureGrid, elevationGrid)

	config := FuocoConfig{
		NumCases:             100,
		NumIterations:        2000,
		NumSample:            100,
		NumContours:          30,
		ImageScale:           6,
		Height:               height,
		Width:                width,
		ElevationFunc:        Adjacent,
		MoistureFunc:         Moisture,
		WindFunc:             TrigonometricWind,
		FuelFunc:             LinearFuel,
		BurnoutFunc:          UniformBurnout,
		InitialState:         stateGrid,
		InitialElevation:     elevationGrid,
		InitialFuel:          fuelGrid,
		InitialMoisture:      moistureGrid,
		InitialWindDirection: 3.14,
		InitialWindSpeed:     40.0,
	}

	f := New(config)
	f.Run()
}
