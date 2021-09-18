package models

func TopographyUniform(g *Grid, _ *Grid) error {
	for i := range *g {
		for j := range (*g)[i] {
			(*g)[i][j].Elevation = 10
		}
	}
	return nil
}
