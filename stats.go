package Fuoco

import (
	"fmt"
	"strconv"
)

type FuocoStats struct {
	Counts [][][]int
	Count  int
}

func GenerateStats(results [](*FuocoResult), width int, height int) FuocoStats {
	var stats FuocoStats = FuocoStats{}
	numSamples := len((*(results[0])).Timeline)
	stats.Counts = make([][][]int, numSamples)

	for it := 0; it < numSamples; it++ {
		stats.Counts[it] = make([][]int, width)
		for i := 0; i < width; i++ {
			stats.Counts[it][i] = make([]int, height)
		}
	}

	for _, result := range results {
		for it := 0; it < numSamples; it++ {
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					state := (*result).Timeline[it][i][j].State
					if state == BurnedOut || state == Burning {
						stats.Counts[it][i][j]++
					}
				}
			}
		}
	}
	return stats
}

func PrintStats(stats *FuocoStats) {
	counts := (*stats).Counts
	for it, frame := range counts {
		fmt.Println("Case: " + strconv.Itoa(it))
		for _, line := range frame {
			fmt.Println(line)
		}
	}
}
