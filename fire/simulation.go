package fire

import (
	"math/rand"
	"time"

	"github.com/antoniofrigo/FUOCO/io"
	"github.com/antoniofrigo/FUOCO/models"
)

type EnvModel func(
	g *models.Grid,
	i, j, a, b int,
	time *time.Time,
	state int,
) float64

type Model func(g *models.Grid, h *models.Grid)
type InitModel func(g *models.Grid)
type StateModel func(g *models.Grid,
	h *models.Grid,
	topography, weather EnvModel)

type Fire struct {
	Grid             models.Grid
	NewGrid          models.Grid
	Config           io.Config
	IgnitionModel    StateModel
	BurnoutModel     Model
	WeatherModel     EnvModel
	TopographyModel  EnvModel
	InitialFuelModel InitModel
}

func NewFire(c io.Config) *Fire {
	s := new(Fire)
	g := models.NewGrid(400, 400)
	h := models.NewGrid(400, 400)
	s.Grid = *g
	s.NewGrid = *h

	s.Config = c
	// Modify state
	s.SetInitialGridFuel(models.InitialFuelUniform)

	// Set propagation models
	s.SetTopographyModel(models.TopographyUniform)

	// Actual running
	s.Ignite(128, 128)
	s.Run()
	g.PrintState("")

	return s
}

// Sets the callback function that determines whether a cell is ignited
func (s *Fire) SetIgnitionModel(f StateModel) {
	(*s).IgnitionModel = f
}

// Sets the callback function for whether a burning cell burns out
func (s *Fire) SetBurnoutModel(f Model) {
	(*s).BurnoutModel = f
}

// Sets the callback function for defining the time-dependent weather model
func (s *Fire) SetWeatherModel(f EnvModel) {
	(*s).WeatherModel = f
}

func (s *Fire) SetTopographyModel(f EnvModel) {
	(*s).TopographyModel = f
}

// Sets the function for defining grid topography. This modifies the grid.
// func (s *Fire) SetInitialGridTopography(f EnvModel) {}

// Sets the function for defining the initial fuel levels. This
// modifies the grid.
func (s *Fire) SetInitialGridFuel(f InitModel) {
	f(&(*s).Grid)
}

func (s *Fire) Ignite(i uint8, j uint8) {
	s.Grid[i][j].State = models.Burning
}

func (s *Fire) Run() {
	step := 0
	g := (*s).Grid
	h := (*s).NewGrid

	topography := (*s).TopographyModel
	// weather := (*s).WeatherModel
	t := time.Now()

	r := rand.New(rand.NewSource(1233))
	for step < (*s).Config.NumIterations {
		num_fires := 0
		// Burnout
		for i := 1; i < len(g)-1; i++ {
			for j := 1; j < len(g[0])-1; j++ {
				if g[i][j].State == models.Burning && 0.3 > r.Float64() {
					g[i][j].State = models.BurnedOut
				}
				h[i][j] = g[i][j]
			}
		}

		// Ignition
		for i := 1; i < len(g)-1; i++ {
			for j := 1; j < len(g[0])-1; j++ {
				if g[i][j].State == models.Burning {
					num_fires++
				}
				if g[i][j].State != models.Ready {
					continue
				}
				var p float64 = 1.0

				for a := i - 1; a < i+2; a++ {
					for b := j - 1; b < j+2; b++ {
						if a == i && b == j {
							continue
						}
						p *= (1.0 - topography(&g, i, j, a, b, &t, g[a][b].State))
					}
				}

				if 1-p > r.Float64() {
					h[i][j].State = models.Burning
				}
			}
		}
		g, h = h, g
		step++

		if num_fires == 0 {
			break
		}
	}
}
