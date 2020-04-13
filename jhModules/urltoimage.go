package jhModules

import (
	"io/ioutil"
	"log"
	"net/http"

	"gocv.io/x/gocv"
)

// UrlToImage read url string and convert to image based on MAT which is basic type on OpenCV
//
// Parameters:
//     url - image link
//     Returns the image matrix
func UrlToImage(url string) gocv.Mat {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//HTTP Error 403: Forbidden problem -> set header!
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
	if img.Empty() {
		log.Fatal(err)
	}

	return img
}
