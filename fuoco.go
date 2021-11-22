package fuoco

import (
	"fmt"
)

type State int

const (
	NoFuel State = iota
	Ready
	Burning
	BurnedOut
)

type FuocoConfig struct {
	NumCases             uint
	NumIterations        uint
	NumSamples           int // Number of samples
	Height               int // Height of grid
	Width                int // Width of grid
	TopographyFunc       ModelFunc
	WeatherFunc          ModelFunc
	FuelFunc             ModelFunc
	BurnoutFunc          ModelFunc
	InitialState         [][]State
	InitialElevation     [][]int
	InitialFuel          [][]int // Between 0 and 100
	InitialMoisture      [][]int // Between 0 and 100
	InitialWindDirection float32 // Between 0.0 and 2*PI
	InitialWindSpeed     float32 // Greater than 0.0
}

// The main Fuoco object
type Fuoco struct {
	FuocoConfig

	freqSamples int
}

// Signature for function that propagates state changes.
// This returns the probability that no ignition will occur due
// to this quantity.
type ModelFunc func(state *[][]State, param *[][]int, i int, j int) float64

// Result of an individual case
type CaseResult struct {
	ID     int
	Frames [][][]State // Each frame is a 2D array of [][]int
	Count  int
}

func New(config FuocoConfig) (f *Fuoco) {
	f = &Fuoco{FuocoConfig: config}
	return f
}

func (f Fuoco) Run() {
	f.freqSamples = int(f.NumIterations) / f.NumSamples

	ch := make(chan *CaseResult)
	for i := 0; i < int(f.NumCases); i++ {
		go f.runCase(i, ch)
	}
	results := make([]*CaseResult, f.NumCases)
	for i := 0; i < int(f.NumCases); i++ {
		results[i] = <-ch
	}
}

// Runs each individual case of the simulation
func (f Fuoco) runCase(id int, ch chan *CaseResult) {
	fmt.Println("HERE")
	// fmt.Println("Running case:", id)
	result := CaseResult{ID: id}

	// Unpack config
	height := f.Height
	width := f.Width
	// TopographyFunc := f.TopographyFunc
	// WeatherFunc := f.WeatherFunc
	// FuelFunc := f.FuelFunc
	// BurnoutFunc := f.BurnoutFunc
	numIterations := f.NumIterations
	numSamples := f.NumSamples

	// Allocate frames samples
	result.Frames = make([][][]State, numSamples)
	for s := 0; s < numSamples; s++ {
		result.Frames[s] = make([][]State, height)
		for i, _ := range result.Frames[s] {
			result.Frames[s][i] = make([]State, width)
		}
	}

	// Setup the propagation grids. Note that there will be a
	// set of border cells of width 1 around the actual grid
	G1 := make([][]State, height+2)
	G2 := make([][]State, height+2)
	for i := 0; i < height+2; i++ {
		G1[i] = make([]State, width+2)
		G2[i] = make([]State, width+2)
	}
	for i := 1; i < width+1; i++ {
		for j := 1; j < height+1; j++ {
			G1[i][j] = f.InitialState[i-1][j-1]
			G2[i][j] = f.InitialState[i-1][j-1]
		}
	}
	// r := rand.New(rand.NewSource(int64(71 * id)))
	// sample := 0

	// Ignition
	G1[(width+1)/2][(height+1)/2] = Burning

	for it := uint(0); it < numIterations; it++ {
		// 	if it%uint(sampling) == 0 {
		// 		for i := 0; i < width; i++ {
		// 			for j := 0; j < height; j++ {
		// 				result.Timeline[sample][i][j] = G1[i][j]
		// 			}
		// 		}
		// 		result.Count++
		// 		sample++
		// 	}

		// 	// Update to the next timestep G1 -> G2
		// 	for i := 1; i < height-1; i++ {
		// 		for j := 1; j < width-1; j++ {
		// 			cell := G1[i][j]
		// 			switch cell.State {
		// 			case Ready:
		// 				var p float64 = 1.0
		// 				p *= TopographyFunc(&(G1), 0, i, j)
		// 				p *= WeatherFunc(&(G1), 0, i, j)
		// 				p *= FuelFunc(&(G1), 0, i, j)
		// 				p = 1 - p
		// 				if p > r.Float64() {
		// 					G2[i][j].State = Burning
		// 				}
		// 			case Burning:
		// 				var p float64 = 1.0
		// 				p *= BurnoutFunc(&(G1), 0, i, j)
		// 				p = 1 - p
		// 				if p > r.Float64() {
		// 					G2[i][j].State = BurnedOut
		// 				}
		// 			}
		// 		}
		// 	}

		// 	// Copy new grid into the old one
		// 	for i := 0; i < width; i++ {
		// 		for j := 0; j < height; j++ {
		// 			G1[i][j] = G2[i][j]
		// 		}
		// 	}
	}
	ch <- &result
}
