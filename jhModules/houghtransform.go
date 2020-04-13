package jhModules

import (
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

// HoughTransform extract edges using canny edge detector and find lines from image(gray-scale)
//
// Parameters:
//     src - source image matrix
//     Returns image matrix containing lines
func HoughTransform(src gocv.Mat) gocv.Mat {
	img := gocv.NewMat()
	img = src.Clone()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	edges := gocv.NewMat()
	defer edges.Close()

	var weakThreshold float32 = 300
	var strongThreshold float32 = 400
	gocv.Canny(gray, &edges, weakThreshold, strongThreshold)

	lines := gocv.NewMat()
	defer lines.Close()

	var rho float32 = 1
	var theta float32 = math.Pi / 180
	var threshold int = 100
	var minLineLength float32 = 100
	var maxLineGap float32 = 5
	if !edges.Empty() {
		gocv.HoughLinesPWithParams(edges, &lines, rho, theta, threshold, minLineLength, maxLineGap)
	}

	if !lines.Empty() {
		for idx := 0; idx < lines.Rows(); idx++ {
			line := lines.GetVeciAt(idx, 0)
			gocv.Line(&img, image.Point{int(line[0]), int(line[1])}, image.Point{int(line[2]), int(line[3])}, color.RGBA{0, 255, 0, 0}, 2)
		}
	} else {
		//if there are no lines, return Empty Mat
		img = gocv.NewMat()
	}

	return img
}
