package main

import (
	"time"

	"github.com/antoniofrigo/FUOCO/fire"
	"github.com/antoniofrigo/FUOCO/io"
)

func main() {
	c := io.Config{
		NumIterations: 10000,
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(100000000),
		Topography:    "uniform",
		IgnitionModel: "random",
		BurnoutModel:  "random",
		WeatherModel:  "uniform",
		FuelModel:     "uniform",
	}
	_ = fire.NewFire(c)
}
