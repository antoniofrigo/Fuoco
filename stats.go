package fuoco

// type FuocoStats struct {
// 	Frames [][][]int
// 	Count  int
// }

// func GenerateStats(results [](*FuocoResult), width int, height int) FuocoStats {
// 	var stats FuocoStats = FuocoStats{}
// 	numSamples := len((*(results[0])).Timeline)
// 	stats.Frames = make([][][]int, numSamples)

// 	for it := 0; it < numSamples; it++ {
// 		stats.Frames[it] = make([][]int, width)
// 		for i := 0; i < width; i++ {
// 			stats.Frames[it][i] = make([]int, height)
// 		}
// 	}

// 	for _, result := range results {
// 		for it := 0; it < numSamples; it++ {
// 			for i := 0; i < width; i++ {
// 				for j := 0; j < height; j++ {
// 					state := (*result).Timeline[it][i][j].State
// 					if state == BurnedOut || state == Burning {
// 						stats.Frames[it][i][j]++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return stats
// }

// func PrintStats(stats *FuocoStats) {
// 	frames := (*stats).Frames
// 	for it, frame := range frames {
// 		fmt.Println("Case: " + strconv.Itoa(it))
// 		for _, line := range frame {
// 			fmt.Println(line)
// 		}
// 	}
// }
