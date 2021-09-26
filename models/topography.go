package models

import "time"

func InitialTopographyUniform(g *Grid, _ *Grid) error {
	for i := range *g {
		for j := range (*g)[i] {
			(*g)[i][j].Elevation = 10
		}
	}
	return nil
}

func TopographyUniform(_ *Grid,
	_,
	_,
	_,
	_ int,
	_ *time.Time,
	state int,
) float64 {
	if state == Burning {
		return 0.3
	}
	return 0.0
}
