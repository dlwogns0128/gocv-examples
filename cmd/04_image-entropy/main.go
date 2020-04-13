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
)

func main() {
	urlString := "https://imgnews.pstatic.net/image/421/2020/04/10/0004578166_001_20200410130433899.jpg?type=w647"
	img := jhModules.UrlToImage(urlString)

	entropy := jhModules.CalcImageEntropy(img, 32)
	fmt.Println("Entropy: ", entropy)
}
