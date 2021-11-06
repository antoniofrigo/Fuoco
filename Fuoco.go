package Fuoco

import (
	"errors"
	"fmt"
	"math/rand"
)

type State int

type FuocoGrid [][]Cell

const (
	NoFuel State = iota
	Ready
	Burning
	BurnedOut
)

type Fuoco struct {
	Config  *FuocoConfig
	Counter FuocoGrid
}

type FuocoConfig struct {
	NumIterations  uint
	Height         int
	Width          int
	TopographyFunc ModelFunc
	WeatherFunc    ModelFunc
	FuelFunc       ModelFunc
	BurnoutFunc    ModelFunc
	InitialGrid    *FuocoGrid
}

type Cell struct {
	State         State
	Elevation     int
	Fuel          float64 // Between 0.0 and 1.0
	Moisture      float64 // Between 0.0 and 1.0
	WindDirection float64 // Between 0.0 and 2*PI
	WindSpeed     float64 // Greater than 0.0
}

// Signature for function that propagates state changes
type ModelFunc func(*FuocoGrid, int) float64

type FuocoResult struct {
	G1 FuocoGrid
	G2 FuocoGrid
	ID int
}

func New() (f *Fuoco) {
	f = &Fuoco{}
	return f
}

func (f *Fuoco) Run() error {
	ch := make(chan *FuocoResult)
	num := f.Config.NumIterations
	for i := 0; i < int(num); i++ {
		go runCase(i, ch, f.Config)
	}
	for i := 0; i < int(num); i++ {
		result := <-ch
		PrintGrid(&((*result).G1))
		_ = result
	}
	return nil
}

// Sets the configuration variables
func (f *Fuoco) SetConfig(config *FuocoConfig) error {
	if config.NumIterations == 0 {
		return errors.New("NumIterations must be greater than 0")
	}
	if config.Height == 0 || config.Width == 0 {
		return errors.New("Height and width must be greater than 0")
	}
	if len(*(config.InitialGrid)) != config.Width {
		return errors.New("InitialGrid and Width must have same length")
	}
	if len((*(config.InitialGrid))[0]) != config.Height {
		return errors.New("InitialGrid[] and Height must have same length")
	}
	f.Config = config
	return nil
}

// Runs each individual case of the simulation
func runCase(id int, ch chan *FuocoResult, config *FuocoConfig) {
	fmt.Println("Running case:", id)
	result := FuocoResult{ID: id}
	height := (*config).Height
	width := (*config).Width
	TopographyFunc := (*config).TopographyFunc
	WeatherFunc := (*config).WeatherFunc
	FuelFunc := (*config).FuelFunc
	numIterations := (*config).NumIterations

	result.G1 = make([][]Cell, height)
	result.G2 = make([][]Cell, height)
	for i := 0; i < (*config).Height; i++ {
		result.G1[i] = make([]Cell, width)
		result.G2[i] = make([]Cell, width)
		copy(result.G1[i], (*(*config).InitialGrid)[i])
		copy(result.G2[i], (*(*config).InitialGrid)[i])
	}

	for it := uint(0); it < numIterations; it++ {
		for i := 1; i < height-1; i++ {
			for j := 1; j < width-1; j++ {
				cell := result.G1[i][j]
				// Ignition
				switch cell.State {
				case Ready:
					var p float64 = 1.0
					p *= TopographyFunc(&(result.G1), 0)
					p *= WeatherFunc(&(result.G1), 0)
					p *= FuelFunc(&(result.G1), 0)
					p = 1 - p
					if p > rand.Float64() {
						result.G2[i][j].State = Burning
					}
				case Burning:
					var p float64 = 1.0
					p *= TopographyFunc(&(result.G1), 0)
					p *= WeatherFunc(&(result.G1), 0)
					p *= FuelFunc(&(result.G1), 0)
					p = 1 - p
					if p > rand.Float64() {
						result.G2[i][j].State = BurnedOut
					}
				}
			}
		}
		result.G1 = result.G2
	}
	ch <- &result
}
