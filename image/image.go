package image

import (
	"image/png"
	"os"

	"github.com/nfnt/resize"
	"github.com/qeesung/image2ascii/convert"
)

func Image2Ascii() string {

	file, _ := os.Open("logo2.png")
	img, _ := png.Decode(file)

	// Set size of image
	size := 30

	// Aspect ratio correction factor (experiment with 0.5–0.6)
	aspectRatio := 0.45
	newHeight := uint(float64(size) * aspectRatio)

	// Resize image
	resized := resize.Resize(uint(size), newHeight, img, resize.Lanczos3)

	converter := convert.NewImageConverter()
	ascii := converter.Image2ASCIIString(resized, &convert.Options{
		Colored: true,
	})
	return ascii
}
