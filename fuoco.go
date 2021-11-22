package Fuoco

import (
	"fmt"
	"math"
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
	Config *FuocoConfig
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

func (f *Fuoco) Run() (FuocoStats, error) {
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
	stats := GenerateStats(results, f.Config.Width, f.Config.Height)
	return stats, nil
}

// Runs each individual case of the simulation
func runCase(id int, ch chan *FuocoResult, config *FuocoConfig) {
	fmt.Println("Running case:", id)
	result := FuocoResult{ID: id}

	// Unpack config
	height := (*config).Height
	width := (*config).Width
	TopographyFunc := (*config).TopographyFunc
	WeatherFunc := (*config).WeatherFunc
	FuelFunc := (*config).FuelFunc
	BurnoutFunc := (*config).BurnoutFunc
	numIterations := (*config).NumIterations
	sampling := float64((*config).Sampling)

	// Allocate timeline samples
	numSamples := int(math.Ceil(float64((*config).NumIterations) / sampling))
	result.Timeline = make([]FuocoGrid, numSamples)
	for s := 0; s < numSamples; s++ {
		result.Timeline[s] = make([][]Cell, width)
		for i := 0; i < height; i++ {
			result.Timeline[s][i] = make([]Cell, height)
		}
	}

	// Setup the propagation grids
	result.G1 = make([][]Cell, width)
	result.G2 = make([][]Cell, width)
	for i := 0; i < width; i++ {
		result.G1[i] = make([]Cell, height)
		result.G2[i] = make([]Cell, height)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			result.G1[i][j] = (*(*config).InitialGrid)[i][j]
			result.G2[i][j] = (*(*config).InitialGrid)[i][j]
		}
	}
	r := rand.New(rand.NewSource(int64(71 * id)))
	sample := 0

	// Ignition
	result.G1[width/2][height/2].State = Burning

	for it := uint(0); it < numIterations; it++ {
		if it%uint(sampling) == 0 {
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					result.Timeline[sample][i][j] = result.G1[i][j]
				}
			}
			result.Count++
			sample++
		}

		// Update to the next timestep G1 -> G2
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

		// Copy new grid into the old one
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				result.G1[i][j] = result.G2[i][j]
			}
		}
	}
	ch <- &result
}
