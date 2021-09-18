package io

import "time"

// TODO: Add config validation
type Config struct {
	NumIterations int
	StartTime     time.Time
	EndTime       time.Time
	Topography    string
	IgnitionModel string
	BurnoutModel  string
	WeatherModel  string
	FuelModel     string
	UpperLeft     int
	BottomRight   int
}
