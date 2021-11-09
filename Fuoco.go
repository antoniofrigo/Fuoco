package Fuoco

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
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
	NumCases       uint
	NumIterations  uint
	Height         int
	Width          int
	Sampling       int // Sample every N iterations
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

// Signature for function that propagates state changes.
// This returns the probability that no ignition will occur due
// to this quantity.
type ModelFunc func(g *FuocoGrid, t int, i int, j int) float64

type FuocoResult struct {
	ID       int
	Timeline []FuocoGrid
	Count    int
	G1       FuocoGrid
	G2       FuocoGrid
}

func New() (f *Fuoco) {
	f = &Fuoco{}
	return f
}

func (f *Fuoco) Run() error {
	ch := make(chan *FuocoResult)
	num := f.Config.NumCases
	for i := 0; i < int(num); i++ {
		go runCase(i, ch, f.Config)
	}
	results := make([](*FuocoResult), num)
	for i := 0; i < int(num); i++ {
		result := <-ch
		results[i] = result
	}
	// stats := GenerateStats(results, f.Config.Width, f.Config.Height)
	// PrintStats(&stats)
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
	if config.TopographyFunc == nil {
		return errors.New("TopographyFunc must be defined")
	}
	if config.WeatherFunc == nil {
		return errors.New("WeatherFunc must be defined")
	}
	if config.FuelFunc == nil {
		return errors.New("FuelFunc must be defined")
	}
	if config.FuelFunc == nil {
		return errors.New("BurnoutFunc must be defined")
	}
	if config.Sampling <= 0 {
		return errors.New("Sampling must be specified")
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
	BurnoutFunc := (*config).BurnoutFunc
	numIterations := (*config).NumIterations
	sampling := float64((*config).Sampling)

	numSamples := int(math.Ceil(float64((*config).NumIterations) / sampling))
	result.Timeline = make([]FuocoGrid, numSamples)
	for s := 0; s < numSamples; s++ {
		result.Timeline[s] = make([][]Cell, height)
		for i := 0; i < height; i++ {
			result.Timeline[s][i] = make([]Cell, width)
		}
	}

	result.G1 = make([][]Cell, height)
	result.G2 = make([][]Cell, height)
	for i := 0; i < (*config).Height; i++ {
		result.G1[i] = make([]Cell, width)
		result.G2[i] = make([]Cell, width)
		copy(result.G1[i], (*(*config).InitialGrid)[i])
		copy(result.G2[i], (*(*config).InitialGrid)[i])
	}

	r := rand.New(rand.NewSource(int64(7 * id)))
	sample := 0
	result.G1[width/2][height/2].State = Burning

	for it := uint(0); it < numIterations; it++ {
		fmt.Println("Iteration: " + strconv.Itoa(int(it)))
		PrintGrid(&result.G1)
		if it%uint(sampling) == 0 {
			copy(result.Timeline[sample], result.G1)
			result.Count++
			sample++
		}

		for i := 1; i < height-1; i++ {
			for j := 1; j < width-1; j++ {
				cell := result.G1[i][j]
				switch cell.State {
				case Ready:
					var p float64 = 1.0
					p *= TopographyFunc(&(result.G1), 0, i, j)
					p *= WeatherFunc(&(result.G1), 0, i, j)
					p *= FuelFunc(&(result.G1), 0, i, j)
					p = 1 - p
					if p > r.Float64() {
						result.G2[i][j].State = Burning
					}
				case Burning:
					var p float64 = 1.0
					p *= BurnoutFunc(&(result.G1), 0, i, j)
					p = 1 - p
					if p > r.Float64() {
						result.G2[i][j].State = BurnedOut
					}
				}
			}
		}

		copy(result.G1, result.G2)
	}
	ch <- &result
}
