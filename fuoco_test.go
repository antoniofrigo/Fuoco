package fuoco

import (
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {
	height := 500
	width := 500

	stateGrid := MakeStateGrid(height, width)
	elevationGrid := MakeParamGrid(height, width)
	fuelGrid := MakeParamGrid(height, width)
	moistureGrid := MakeParamGrid(height, width)

	SetStateGrid(&stateGrid, Ready)
	SetParamGridBooth(&elevationGrid)
	SetParamGridByReverseElevation(&fuelGrid, elevationGrid)
	SetParamGridByReverseElevation(&moistureGrid, elevationGrid)

	config := FuocoConfig{
		NumCases:             10,
		NumIterations:        500,
		NumSample:            20,
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
