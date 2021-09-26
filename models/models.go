package models

import "time"

type EnvModel func(
	g *Grid,
	i, j, a, b int,
	time time.Time,
	state int,
) float64
