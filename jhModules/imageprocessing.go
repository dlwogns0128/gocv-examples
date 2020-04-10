package jhModules

import (
	"image"
	"math"

	"gocv.io/x/gocv"
)

//entropy 계산
func CalcImageEntropy(src gocv.Mat, stride int) float64 {

	gray := gocv.NewMat()
	if src.Channels() != 1 {
		gocv.CvtColor(src, &gray, gocv.ColorBGRToGray)
	} else {
		gray = src.Clone()
	}

	shape := gray.Size()
	row := shape[0]
	col := shape[1]
	numBlock := int(row * col / (stride * stride))
	size := float64(stride * stride)

	var entropy float64
	entropy = 0

	for i := 0; i < row; i += stride {
		for j := 0; j < col; j += stride {
			//image 에서 x는 col, y는 row!
			croppedMat := gray.Region(image.Rect(j, i, j+stride, i+stride))
			subImg := croppedMat.Clone()
			hist := gocv.NewMat()
			defer croppedMat.Close()
			defer subImg.Close()
			defer hist.Close()
			gocv.CalcHist([]gocv.Mat{subImg}, []int{0}, gocv.NewMat(), &hist, []int{256}, []float64{0, 256}, false)

			for k := 0; k < 256; k++ {
				bin := float64(hist.GetFloatAt(k, 0))
				if bin > 0 {
					entropy += bin / size * math.Log2(size/bin)
				}
			}
		}
	}
	return entropy / float64(numBlock)
}
