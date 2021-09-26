package io

import "time"

// TODO: Add config validation
type Config struct {
	NumIterations int
	StartTime     time.Time
	EndTime       time.Time
	FuelModel     string
	UpperLeft     int
	BottomRight   int
	Topography    string
	IgnitionModel string
	BurnoutModel  string
	WeatherModel  string
	WeatherParams WeatherParams
}

type WeatherParams struct {
	Direction int
	Speed     int
}
