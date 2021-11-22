package Fuoco

import (
	"fmt"
	"strings"
)

const (
	header   = "<svg viewBox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\">\n"
	polyline = "\t<polyline points=\"%s\" stroke=\"%s\"/>\n"
	points   = "%d,%d %d,%d %d,%d %d,%d"
	tail     = "</svg>\n"
)

func colorScale(value, maxValue int) string {
	r := int(255.0 * value / (1.0 * maxValue))
	g := r
	b := r
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

func GenerateSVG(frame [][]int) string {
	var s strings.Builder
	width := len(frame)
	height := len(frame[0])

	maxValue := 0
	for _, row := range frame {
		for _, value := range row {
			if maxValue < value {
				maxValue = value
			}
		}
	}

	s.WriteString(fmt.Sprintf(header, width, height))
	for i, row := range frame {
		for j, value := range row {
			p := fmt.Sprintf(points, i, j, i+1, j, i+1, j+1, i+1, j)
			color := colorScale(value, maxValue)
			polylineElement := fmt.Sprintf(polyline, p, color)
			s.WriteString(polylineElement)
		}
	}
	s.WriteString(tail)
	return s.String()
}
