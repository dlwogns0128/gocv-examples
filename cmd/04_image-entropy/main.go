// What it does:
//
// This example calculates entorpy of image
//
// How to run:
//
// 		go run ./cmd/image-entropy/main.go
//

package main

import (
	"fmt"
	"gocv-examples/jhModules"
	"image"

	"gocv.io/x/gocv"
)

func main() {
	urlString := "https://imgnews.pstatic.net/image/421/2020/04/10/0004578166_001_20200410130433899.jpg?type=w647"
	img := jhModules.UrlToImage(urlString)

	height := img.Size()[0]
	width := img.Size()[1]

	var interpolation gocv.InterpolationFlags
	if height < 640 || width < 480 {
		interpolation = gocv.InterpolationCubic
	} else {
		interpolation = gocv.InterpolationArea
	}

	gocv.Resize(img, &img, image.Point{640, 480}, 0, 0, interpolation) //640, 480

	entropy := jhModules.CalcImageEntropy(img, 32)
	fmt.Println("Entropy: ", entropy)
}
