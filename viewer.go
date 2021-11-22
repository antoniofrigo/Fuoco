package Fuoco

import (
	"image"
	"image/color"
)

const scale = 10

func GeneratePNG(frame [][]int) image.Image {
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
