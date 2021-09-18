package models

import "math/rand"

func BurnoutRandom(G *Grid, H *Grid) error {
	r := rand.New(rand.NewSource(130))
	g := *G
	for i := 1; i < len(g)-1; i++ {
		for j := 1; j < len(g[0])-1; j++ {
			if g[i][j].State == Burning && r.Intn(3) == 1 {
				g[i][j].State = BurnedOut
			}
		}
	}
	return nil
}
