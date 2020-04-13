package jhModules

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

// CalcImageHistogram calculate histogram of image(gray-scale)
//
// Parameters:
//     src - source image matrix
//     Returns histogram matrix
func CalcImageHistogram(src gocv.Mat) gocv.Mat {
	img := src.Clone()
	defer img.Close()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	hist := gocv.NewMat()
	mask := gocv.NewMat()
	defer mask.Close()

	channels := []int{0}
	size := []int{256}
	ranges := []float64{0, 256}
	acc := false

	gocv.CalcHist([]gocv.Mat{gray}, channels, mask, &hist, size, ranges, acc)

	return hist
}

// DrawHistogram process image to show easily in NewWindow
//
// Parameters:
//     hist - source image matrix
//     Returns image matrix
func DrawHistogram(hist gocv.Mat) gocv.Mat {

	dHist := gocv.NewMat()
	defer dHist.Close()

	//set matrix size to be shown
	histW := 512
	histH := 400
	size := hist.Size()[0]
	binW := int(float64(histW) / float64(size))
	histImage := gocv.NewMatWithSize(512, 400, gocv.MatTypeCV8U)

	gocv.Normalize(hist, &dHist, 0, float64(histImage.Rows()), gocv.NormMinMax) // normalize to show easily

	for idx := 1; idx < size; idx++ {
		gocv.Line(&histImage, image.Point{binW * (idx - 1), histH - int(dHist.GetFloatAt(idx-1, 0))}, image.Point{binW * (idx), histH - int(dHist.GetFloatAt(idx, 0))}, color.RGBA{255, 255, 255, 0}, 2)
	}

	return histImage
}
