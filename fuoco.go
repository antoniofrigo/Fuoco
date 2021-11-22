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

// The main Fuoco object
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

// Result of an individual case
type FuocoResult struct {
	ID       int
	Timeline []FuocoGrid
	Count    int
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
	var G1 FuocoGrid = make([][]Cell, width)
	var G2 FuocoGrid = make([][]Cell, width)
	for i := 0; i < width; i++ {
		G1[i] = make([]Cell, height)
		G2[i] = make([]Cell, height)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			G1[i][j] = (*(*config).InitialGrid)[i][j]
			G2[i][j] = (*(*config).InitialGrid)[i][j]
		}
	}
	r := rand.New(rand.NewSource(int64(71 * id)))
	sample := 0

	// Ignition
	G1[width/2][height/2].State = Burning

	for it := uint(0); it < numIterations; it++ {
		if it%uint(sampling) == 0 {
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					result.Timeline[sample][i][j] = G1[i][j]
				}
			}
			result.Count++
			sample++
		}

		// Update to the next timestep G1 -> G2
		for i := 1; i < height-1; i++ {
			for j := 1; j < width-1; j++ {
				cell := G1[i][j]
				switch cell.State {
				case Ready:
					var p float64 = 1.0
					p *= TopographyFunc(&(G1), 0, i, j)
					p *= WeatherFunc(&(G1), 0, i, j)
					p *= FuelFunc(&(G1), 0, i, j)
					p = 1 - p
					if p > r.Float64() {
						G2[i][j].State = Burning
					}
				case Burning:
					var p float64 = 1.0
					p *= BurnoutFunc(&(G1), 0, i, j)
					p = 1 - p
					if p > r.Float64() {
						G2[i][j].State = BurnedOut
					}
				}
			}
		}

		// Copy new grid into the old one
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				G1[i][j] = G2[i][j]
			}
		}
	}
	ch <- &result
}
