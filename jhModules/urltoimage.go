package jhModules

import (
	"image"
	"io/ioutil"
	"log"
	"net/http"

	"gocv.io/x/gocv"
)

func UrlToImage(url string) gocv.Mat {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	img, err := gocv.IMDecode(bodyBytes, gocv.IMReadColor)
	if !img.Empty() {
		height := img.Size()[0]
		width := img.Size()[1]
		var interpolation gocv.InterpolationFlags
		if height < 640 || width < 480 {
			interpolation = gocv.InterpolationCubic
		} else {
			interpolation = gocv.InterpolationArea
		}
		gocv.Resize(img, &img, image.Point{640, 480}, 0, 0, interpolation) //640, 480
	} else {
		log.Fatal(err)
	}

	return img
}
