package jhModules

import (
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

//직선 검출
func HoughTransform(src gocv.Mat) gocv.Mat {
	img := gocv.NewMat()
	img = src.Clone()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(gray, &edges, 300, 400)

	lines := gocv.NewMat()
	defer lines.Close()
	if !edges.Empty() {
		gocv.HoughLinesPWithParams(edges, &lines, 1, math.Pi/180, 100, 100, 5)
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
