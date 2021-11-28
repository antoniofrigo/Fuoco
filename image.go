package fuoco

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var RLDiag = [4]float64{0.0, 0.5, 0.5, 1.0}
var LUDiag = [4]float64{0.5, 0.0, 1.0, 0.5}
var LLDiag = [4]float64{0.0, 0.5, 0.5, 0.0}
var RUDiag = [4]float64{0.5, 1.0, 1.0, 0.5}
var Horizontal = [4]float64{0.5, 0.0, 0.5, 1.0}
var Vertical = [4]float64{0.0, 0.5, 1.0, 0.5}

type imageWrapper struct {
	Id    int
	Image image.Image
}

// Generate the sampled images
func (f *Fuoco) generateImages() error {
	(*f).Images = make([]image.Image, (*f).NumSample)
	ch := make(chan imageWrapper)
	for idx, _ := range (*f).Images {
		go (*f).generatePNG((*f).Frames[idx], idx, ch)
	}

	for idx, _ := range (*f).Images {
		result := <-ch
		(*f).Images[result.Id] = result.Image
		_ = idx
	}

	return nil
}

// Save images to disk
func (f Fuoco) saveImage(name string, img image.Image) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	if err := png.Encode(file, img); err != nil {
		file.Close()
		return err
	}
	return nil
}

// Generate an individual PNG given some frame of values. These are
// scaled up by whatever the ImageScale factor is
func (f Fuoco) generatePNG(frame [][]int, idx int, ch chan imageWrapper) {
	scale := f.ImageScale
	height := scale * f.Height
	width := scale * f.Width
	maxValue := 0
	for _, row := range frame {
		for _, value := range row {
			if maxValue < value {
				maxValue = value
			}
		}
	}

	img := image.NewNRGBA(image.Rect(0, 0, height, width))
	for i, row := range frame {
		for j, value := range row {
			v := uint8(200.0 * value / (1.0 * maxValue))
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
	f.generateElevationMask(img)
	ch <- imageWrapper{Id: idx, Image: img}
}

// Generate a generic image with a given color, defined by
// the rscale, gscale, and bscale factors
func (f Fuoco) generateImg(frame [][]int, rscale, gscale, bscale float64) image.Image {
	scale := f.ImageScale
	height := scale * f.Height
	width := scale * f.Width
	maxValue := 0
	for _, row := range frame {
		for _, value := range row {
			if maxValue < value {
				maxValue = value
			}
		}
	}

	img := image.NewNRGBA(image.Rect(0, 0, height, width))
	for i, row := range frame {
		for j, value := range row {
			v := 200.0 * float64(value) / (1.0 * float64(maxValue))
			for a := 0; a < scale; a++ {
				for b := 0; b < scale; b++ {
					img.Set(scale*i+a, scale*j+b, color.NRGBA{
						R: uint8(rscale * v),
						G: uint8(gscale * v),
						B: uint8(bscale * v),
						A: 255,
					})
				}
			}
		}
	}
	f.generateElevationMask(img)
	return img
}

// Generates contour lines based on data.
func (f Fuoco) generateElevationMask(img *image.NRGBA) {
	isovalueMask := make([][]bool, f.Height)

	for idx, _ := range isovalueMask {
		isovalueMask[idx] = make([]bool, f.Width)
	}

	minElev := math.MaxInt32
	maxElev := math.MinInt32
	for _, row := range f.InitialElevation {
		for _, value := range row {
			if minElev > value {
				minElev = value
			} else if maxElev < value {
				maxElev = value
			}
		}
	}

	dContour := (maxElev - minElev) / f.NumContours
	if dContour == 0 {
		dContour = 1
	}

	for idx := 0; idx < f.NumContours; idx++ {
		contour := minElev + idx*dContour
		isovalueMask = f.generateIsovalueMask(isovalueMask, contour)
		for i := 0; i < f.Height-1; i++ {
			for j := 0; j < f.Width-1; j++ {
				contourType := f.contourType(isovalueMask, i, j)
				f.drawContourSegment(contourType, img, i, j)
			}
		}
	}

}

// Generates the isovalue mask for marching squares
func (f Fuoco) generateIsovalueMask(isovalueMask [][]bool, contour int) [][]bool {
	for i, _ := range f.InitialElevation {
		for j, value := range f.InitialElevation[i] {
			if contour <= value {
				isovalueMask[i][j] = true
			} else {
				isovalueMask[i][j] = false
			}
		}
	}
	return isovalueMask
}

// Determines the necessary contour type for four points, with
// i,j in the bottom left
func (f Fuoco) contourType(isovalueMask [][]bool, i, j int) int {
	res := 0
	if isovalueMask[i][j] == true {
		res += 1
	}
	if isovalueMask[i][j+1] == true {
		res += 2
	}
	if isovalueMask[i+1][j+1] == true {
		res += 4
	}
	if isovalueMask[i+1][j] == true {
		res += 8
	}
	return res
}

// Helper function to get the coordinates for the line endpoints
// when generating the contours
func getCoordinates(i, j int, shifts [4]float64, scale int) (int, int, int, int) {
	x0 := int(float64(scale) * (float64(i) + shifts[0] - 0.5))
	y0 := int(float64(scale) * (float64(j) + shifts[1] - 0.5))
	x1 := int(float64(scale) * (float64(i) + shifts[2] - 0.5))
	y1 := int(float64(scale) * (float64(j) + shifts[3] - 0.5))
	return x0, y0, x1, y1
}

// Generates contour lines via marching squares
// https://en.wikipedia.org/wiki/Marching_squares
func (f Fuoco) drawContourSegment(contourType int, img *image.NRGBA, i, j int) {
	scale := f.ImageScale
	switch contourType {
	case 0:
		return
	case 1:
		x0, y0, x1, y1 := getCoordinates(i, j, LLDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 2: // Problem
		x0, y0, x1, y1 := getCoordinates(i, j, RLDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 3:
		x0, y0, x1, y1 := getCoordinates(i, j, Horizontal, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 4:
		x0, y0, x1, y1 := getCoordinates(i, j, RUDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 5:
		x0, y0, x1, y1 := getCoordinates(i, j, LUDiag, scale)
		x2, y2, x3, y3 := getCoordinates(i, j, RLDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
		f.drawLine(img, x2, y2, x3, y3)
	case 6:
		x0, y0, x1, y1 := getCoordinates(i, j, Vertical, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 7: // Problem
		x0, y0, x1, y1 := getCoordinates(i, j, LUDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 8: // Problem
		x0, y0, x1, y1 := getCoordinates(i, j, LUDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 9:
		x0, y0, x1, y1 := getCoordinates(i, j, Vertical, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 10:
		x0, y0, x1, y1 := getCoordinates(i, j, RUDiag, scale)
		x2, y2, x3, y3 := getCoordinates(i, j, LLDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
		f.drawLine(img, x2, y2, x3, y3)
	case 11:
		x0, y0, x1, y1 := getCoordinates(i, j, RUDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 12:
		x0, y0, x1, y1 := getCoordinates(i, j, Horizontal, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 13: // Problem
		x0, y0, x1, y1 := getCoordinates(i, j, RLDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 14:
		x0, y0, x1, y1 := getCoordinates(i, j, LLDiag, scale)
		f.drawLine(img, x0, y0, x1, y1)
	case 15:
		return
	}
}

// Helper function to draw a straight line from
// (x0, y0) to (x1, y1)
func (f Fuoco) drawLine(img *image.NRGBA, x0, y0, x1, y1 int) {
	if x0 == x1 && y0 == y1 {
		return
	}
	if x1 == x0 {
		for y := y0; y <= y1; y++ {
			img.Set(x0, y, color.NRGBA{
				R: 0,
				G: 255,
				B: 0,
				A: 255,
			})
		}
	} else {
		m := (y1 - y0) / (x1 - x0)
		for x := x0; x <= x1; x++ {
			img.Set(x, m*(x-x0)+y0, color.NRGBA{
				R: 0,
				G: 255,
				B: 0,
				A: 255,
			})
		}
	}
}
