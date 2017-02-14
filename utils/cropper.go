package utils

import (
	"github.com/BurntSushi/graphics-go/graphics"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

func DoImageHandler(url string, newdx int) {
	src, err := LoadImage("." + url)
	if err != nil {
		log.Fatal(err)
	}

	dst := image.NewRGBA(image.Rect(640, 640, 200, 200))
	err = graphics.Thumbnail(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	filen := strings.Replace(url, ".", "-cropper.", -1)
	file, err := os.Create("." + filen)
	defer file.Close()

	err = jpeg.Encode(file, dst, &jpeg.Options{100})

}

func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}
