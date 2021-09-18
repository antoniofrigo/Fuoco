package models

func FuelUniform(G *Grid, _ *Grid) error {
	g := *G
	for i := 1; i < len(g)-1; i++ {
		for j := 1; j < len(g[0])-1; j++ {
			g[i][j].State = Ready
		}
	}
	return nil
}
