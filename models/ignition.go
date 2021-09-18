package models

import (
	"math/rand"
)

// Random ignition
func IgnitionRandom(G *Grid, H *Grid) error {
	r := rand.New(rand.NewSource(132))
	g := *G
	h := *H

	for i := 1; i < len(g)-1; i++ {
		for j := 1; j < len(g[0])-1; j++ {
			h[i][j] = g[i][j]
		}
	}

	for i := 1; i < len(g)-1; i++ {
		for j := 1; j < len(g[0])-1; j++ {
			// Only consider Ready cells
			if g[i][j].State != Ready {
				continue
			}

			if g[i-1][j].State == Burning || g[i+1][j].State == Burning ||
				g[i][j-1].State == Burning || g[i][j+1].State == Burning ||
				g[i-1][j-1].State == Burning || g[i+1][j+1].State == Burning ||
				g[i+1][j-1].State == Burning || g[i-1][j+1].State == Burning {
				if r.Intn(2) == 1 {
					h[i][j].State = Burning
				}
			}
		}
	}

	*G, *H = h, g

	return nil
}
