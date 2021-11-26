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
	SetParamGridParabolicCylinder(&elevationGrid)
	SetParamGrid(&fuelGrid, 100)
	SetParamGridByReverseElevation(&moistureGrid, elevationGrid)

	config := FuocoConfig{
		NumCases:             100,
		NumIterations:        200,
		NumSample:            10,
		NumContours:          20,
		ImageScale:           6,
		Height:               height,
		Width:                width,
		ElevationFunc:        Adjacent,
		MoistureFunc:         Moisture,
		WindFunc:             TrigonometricWind,
		FuelFunc:             OneParam,
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
