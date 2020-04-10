package jhModules

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func CalcImageHistogram(src gocv.Mat) gocv.Mat {
	img := gocv.NewMat()
	defer img.Close()

	img = src.Clone()
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)
	hist := gocv.NewMat()

	gocv.CalcHist([]gocv.Mat{gray}, []int{0}, gocv.NewMat(), &hist, []int{256}, []float64{0, 256}, false)

	return hist
}

func DrawHistogram(hist gocv.Mat) gocv.Mat {

	dHist := gocv.NewMat()
	defer dHist.Close()

	histW := 512
	histH := 400
	size := hist.Size()[0]
	binW := int(float64(histW) / float64(size))
	histImage := gocv.NewMatWithSize(512, 400, gocv.MatTypeCV8U)

	gocv.Normalize(hist, &dHist, 0, float64(histImage.Rows()), gocv.NormMinMax)

	for idx := 1; idx < size; idx++ {
		gocv.Line(&histImage, image.Point{binW * (idx - 1), histH - int(dHist.GetFloatAt(idx-1, 0))}, image.Point{binW * (idx), histH - int(dHist.GetFloatAt(idx, 0))}, color.RGBA{255, 255, 255, 0}, 2)
	}

	return histImage
}
