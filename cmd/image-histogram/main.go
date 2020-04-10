// What it does:
//
// This example calculates histogram of gray-scale image and show the histogram
//
// How to run:
//
// 		go run ./cmd/image-histogram/main.go
//

package main

import (
	"gocv-examples/jhModules"

	"gocv.io/x/gocv"
)

func main() {
	urlString := "https://imgnews.pstatic.net/image/421/2020/04/10/0004578166_001_20200410130433899.jpg?type=w647"
	img := jhModules.UrlToImage(urlString)

	hist := jhModules.CalcImageHistogram(img)
	hImg := jhModules.DrawHistogram(hist)

	window := gocv.NewWindow("test")

	height := img.Size()[0]
	width := img.Size()[1]
	window.ResizeWindow(width, height)
	window.IMShow(hImg)
	window.WaitKey(0)
}
