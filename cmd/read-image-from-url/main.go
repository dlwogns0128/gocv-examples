// What it does:
//
// This example reads image from url by using http get and gocv's IMDecode modules
//
// How to run:
//
// 		go run ./cmd/read-image-from-url/main.go
//

package main

import (
	"gocv-examples/jhModules"

	"gocv.io/x/gocv"
)

func main() {
	urlString := "https://imgnews.pstatic.net/image/421/2020/04/10/0004578166_001_20200410130433899.jpg?type=w647"
	img := jhModules.UrlToImage(urlString)

	window := gocv.NewWindow("test")

	height := img.Size()[0]
	width := img.Size()[1]
	window.ResizeWindow(width, height)
	window.IMShow(img)
	window.WaitKey(0)
}
