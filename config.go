package Fuoco

import "errors"

type FuocoConfig struct {
	NumCases       uint
	NumIterations  uint
	Height         int
	Width          int
	Sampling       int // Sample every N iterations
	TopographyFunc ModelFunc
	WeatherFunc    ModelFunc
	FuelFunc       ModelFunc
	BurnoutFunc    ModelFunc
	InitialGrid    *FuocoGrid
}

// Sets the configuration variables
func (f *Fuoco) SetConfig(config *FuocoConfig) error {
	err := f.validateConfig(config)
	if err != nil {
		return err
	}
	f.Config = config
	return nil
}

// Validate the configuration.
// TODO: Make this nicer.
func (f *Fuoco) validateConfig(config *FuocoConfig) error {
	if config.NumIterations == 0 {
		return errors.New("NumIterations must be greater than 0")
	}
	if config.Height == 0 || config.Width == 0 {
		return errors.New("Height and width must be greater than 0")
	}
	if len(*(config.InitialGrid)) != config.Width {
		return errors.New("InitialGrid and Width must have same length")
	}
	if len((*(config.InitialGrid))[0]) != config.Height {
		return errors.New("InitialGrid[] and Height must have same length")
	}
	if config.TopographyFunc == nil {
		return errors.New("TopographyFunc must be defined")
	}
	if config.WeatherFunc == nil {
		return errors.New("WeatherFunc must be defined")
	}
	if config.FuelFunc == nil {
		return errors.New("FuelFunc must be defined")
	}
	if config.FuelFunc == nil {
		return errors.New("BurnoutFunc must be defined")
	}
	if config.Sampling <= 0 {
		return errors.New("Sampling must be specified")
	}
	return nil
}
