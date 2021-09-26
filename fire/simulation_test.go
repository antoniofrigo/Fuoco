package fire

import (
	"testing"
	"time"

	"github.com/antoniofrigo/FUOCO/io"
)

func TestSimulation(t *testing.T) {
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
	_ = NewFire(c)
}
