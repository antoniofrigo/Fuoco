package fuoco

import (
	"fmt"
	"image"
	"log"
	"math/rand"
)

type State int

const (
	NoFuel State = iota
	Ready
	Burning
	BurnedOut
)

type ParamFunc func(state [][]State, param [][]int, i int, j int) float64
type WindFunc func(state [][]State, speed float64, angle float64, i int, j int) float64
type BurnoutFunc func(state [][]State, fuel [][]int, i int, j int) float64

type FuocoConfig struct {
	NumCases             uint
	NumIterations        uint
	NumSample            int // Number of samples
	NumContours          int
	Height               int // Height of grid
	Width                int // Width of grid
	ImageScale           int // Output image has dim. height * scale, width * scale
	ElevationFunc        ParamFunc
	MoistureFunc         ParamFunc
	FuelFunc             ParamFunc
	WindFunc             WindFunc
	BurnoutFunc          BurnoutFunc
	InitialState         [][]State
	InitialElevation     [][]int
	InitialFuel          [][]int // Between 0 and 100
	InitialMoisture      [][]int // Between 0 and 100
	InitialWindDirection float64 // Between 0.0 and 2*PI
	InitialWindSpeed     float64 // Greater than 0.0
}

// The main Fuoco object
type Fuoco struct {
	FuocoConfig
	Frames      [][][]int
	Images      []image.Image
	MoistureImg image.Image
	freqSample  int
}

// Result of an individual case
type CaseResult struct {
	ID     int
	Frames [][][]State // Each frame is a 2D array of [][]int
	Count  int
}

func New(config FuocoConfig) *Fuoco {
	f := Fuoco{FuocoConfig: config}
	f.Frames = make([][][]int, f.NumSample)
	for idx, _ := range f.Frames {
		f.Frames[idx] = make([][]int, f.Height)
		for i, _ := range f.Frames[idx] {
			f.Frames[idx][i] = make([]int, f.Width)
		}
	}

	return &f
}

func (f Fuoco) Run() {
	f.freqSample = int(f.NumIterations) / f.NumSample

	ch := make(chan *CaseResult)
	for i := 0; i < int(f.NumCases); i++ {
		go f.runCase(i, ch)
	}
	results := make([]*CaseResult, f.NumCases)
	for i := 0; i < int(f.NumCases); i++ {
		results[i] = <-ch
	}

	// Generates the counts
	for s := 0; s < int(f.NumCases); s++ {
		for idx := 0; idx < f.NumSample; idx++ {
			frame := (*(results[s])).Frames[idx]
			for i, _ := range frame {
				for j, value := range frame[i] {
					if value == Burning || value == BurnedOut {
						f.Frames[idx][i][j]++
					}
				}
			}
		}
	}

	f.generateImages()
	for idx, img := range f.Images {
		s := "/tmp/test/" + fmt.Sprint(idx) + ".png"
		err := f.saveImage(s, img)
		if err != nil {
			log.Fatal(err)
		}
	}

	s := "/tmp/test/moisture.png"
	err := f.saveImage(s, f.MoistureImg)
	_ = err
}

// Runs each individual case of the simulation
func (f Fuoco) runCase(id int, ch chan *CaseResult) {
	fmt.Println("Running case:", id)
	result := CaseResult{ID: id}

	// Unpack config
	height := f.Height
	width := f.Width
	numSample := f.NumSample
	numIterations := f.NumIterations

	ElevationFunc := f.ElevationFunc
	MoistureFunc := f.MoistureFunc
	WindFunc := f.WindFunc
	FuelFunc := f.FuelFunc
	BurnoutFunc := f.BurnoutFunc

	// Allocate frames samples
	result.Frames = make([][][]State, numSample)
	for s := 0; s < numSample; s++ {
		result.Frames[s] = make([][]State, height)
		for i, _ := range result.Frames[s] {
			result.Frames[s][i] = make([]State, width)
		}
	}

	// Setup the propagation grids. Note that there will be a
	// set of border cells of width 1 around the actual grid.
	G1 := make([][]State, height+2)
	G2 := make([][]State, height+2)
	for i := 0; i < height+2; i++ {
		G1[i] = make([]State, width+2)
		G2[i] = make([]State, width+2)
	}
	for i := 1; i < height+1; i++ {
		for j := 1; j < width+1; j++ {
			G1[i][j] = f.InitialState[i-1][j-1]
			G2[i][j] = f.InitialState[i-1][j-1]
		}
	}
	r := rand.New(rand.NewSource(int64(71 * id)))
	sample := 0

	// Ignition
	G1[(height+1)/2][(width+1)/2] = Burning
	G2[(height+1)/2][(width+1)/2] = Burning

	for it := uint(0); it < numIterations; it++ {
		if it%uint(f.freqSample) == 0 {
			for i := 1; i < height+1; i++ {
				for j := 1; j < width+1; j++ {
					result.Frames[sample][i-1][j-1] = G1[i][j]
				}
			}
			result.Count++
			sample++
		}
		// Update to the next timestep G1 -> G2
		for i := 1; i < height+1; i++ {
			for j := 1; j < width+1; j++ {
				state := G1[i][j]
				switch state {
				case Ready:
					var p float64 = 1.0
					p *= ElevationFunc(G1, f.InitialElevation, i, j)
					p *= MoistureFunc(G1, f.InitialMoisture, i, j)
					p *= FuelFunc(G1, f.InitialFuel, i, j)
					p *= WindFunc(G1, f.InitialWindSpeed, f.InitialWindDirection, i, j)
					p = 1 - p
					if p > r.Float64() {
						G2[i][j] = Burning
					}
				case Burning:
					var p float64 = 1.0
					p *= BurnoutFunc(G1, nil, i, j)
					p = 1 - p
					if p > r.Float64() {
						G2[i][j] = BurnedOut
					}
				}
			}
		}

		// Copy new grid into the old one
		for i, row := range G1 {
			for j, _ := range row {
				G1[i][j] = G2[i][j]
			}
		}
	}
	ch <- &result
}
