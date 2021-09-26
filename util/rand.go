package util

// A linear congruence pseudo-random generator
func LCGenerator(a, c int) func() float64 {
	r := 11231
	return func() float64 {
		r = (a*r + c) % a
		return float64(r) / float64(a)
	}
}
