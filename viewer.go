package fuoco

import (
	"image"
	"image/color"
)

const imageSize = 1000

func GeneratePNG(frame [][]int) image.Image {
	scale := 1000 / len(frame)
	width := scale * len(frame)
	height := scale * len(frame[0])

	maxValue := 0
	for _, row := range frame {
		for _, value := range row {
			if maxValue < value {
				maxValue = value
			}
		}
	}

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for i, row := range frame {
		for j, value := range row {
			v := uint8(255.0 * value / (1.0 * maxValue))
			for a := 0; a < scale; a++ {
				for b := 0; b < scale; b++ {
					img.Set(scale*i+a, scale*j+b, color.NRGBA{
						R: v,
						G: v,
						B: v,
						A: 255,
					})
				}
			}
		}
	}

	return img
}

// Generates contour lines based on data.
func marchingSquares(scale int, numContours int, data [][]int) image.Image {
	minValue := 10000
	maxValue := 0
	for _, row := range data {
		for _, value := range row {
			if value < minValue {
				minValue = value
			} else if value > maxValue {
				maxValue = value
			}
		}
	}
	interval := (maxValue - minValue) / numContours

	tmp := make([][]int, len(data))
	for idx, _ := range tmp {
		tmp[idx] = make([]int, len(data[idx]))
	}

	width := scale * len(data)
	height := scale * len(data[0])
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for contour := minValue; contour < maxValue; contour += interval {

	}

	return img
}

func generateElevationMask() {}
