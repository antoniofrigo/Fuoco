package Fuoco

import (
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestFuocoConfiguration(t *testing.T) {
	f := New()
	height := 100
	width := 100
	grid := MakeInitialGrid(height, width)
	SetConstantFuel(&grid, 0.7)
	SetConstantElevation(&grid, 1000)
	SetStateReady(&grid)
	for _, row := range grid {
		var s strings.Builder
		for _, cell := range row {
			s.WriteString(strconv.Itoa(cell.Elevation) + " ")
		}
	}

	config := FuocoConfig{
		NumCases:       100,
		NumIterations:  100,
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
	img := GeneratePNG(stats.Frames[49])
	file, err := os.Create("../../test.png")
	png.Encode(file, img)
	file.Close()
}
