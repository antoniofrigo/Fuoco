package fire

import (
	"github.com/antoniofrigo/FUOCO/io"
	"github.com/antoniofrigo/FUOCO/models"
)

type Model func(g *models.Grid, h *models.Grid) error

type Fire struct {
	Grid             models.Grid
	NewGrid          models.Grid
	Config           io.Config
	IgnitionModel    Model
	BurnoutModel     Model
	WeatherModel     Model
	InitialFuelModel Model
}

func NewFire(c io.Config) *Fire {
	s := new(Fire)
	g := models.NewGrid(50, 50)
	h := models.NewGrid(50, 50)
	s.Grid = *g
	s.NewGrid = *h
	g.PrintState("")

	s.Config = c
	s.SetGridTopography(models.TopographyUniform)
	s.SetGridInitialFuel(models.FuelUniform)
	s.SetIgnitionModel(models.IgnitionRandom)
	s.SetBurnoutModel(models.BurnoutRandom)
	s.Ignite(25, 25)
	s.Run()
	g.PrintState("")

	return s
}

// Sets the callback function that determines whether a cell is ignited
func (s *Fire) SetIgnitionModel(f Model) {
	(*s).IgnitionModel = f
}

// Sets the callback function for whether a burning cell burns out
func (s *Fire) SetBurnoutModel(f Model) {
	(*s).BurnoutModel = f
}

// Sets the callback function for defining the time-dependent weather model
func (s *Fire) SetWeatherModel(f Model) {}

// Sets the function for defining grid topography. This modifies the grid.
func (s *Fire) SetGridTopography(f Model) error {
	return f(&(*s).Grid, &(*s).NewGrid)
}

// Sets the function for defining the initial fuel levels. This
// modifies the grid.
func (s *Fire) SetGridInitialFuel(f Model) error {
	return f(&(*s).Grid, &(*s).NewGrid)
}

func (s *Fire) Ignite(i uint8, j uint8) {
	s.Grid[i][j].State = models.Burning
}

func (s *Fire) Run() {
	step := 0
	g := (*s).Grid
	h := (*s).NewGrid

	for step < (*s).Config.NumIterations {
		(*s).BurnoutModel(&g, &h)
		(*s).IgnitionModel(&g, &h)
		step++
	}
}
