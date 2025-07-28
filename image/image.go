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

	// Get original dimensions
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Aspect ratio correction factor (experiment with 0.5–0.6)
	aspectRatio := 0.5
	newHeight := uint(float64(height) * aspectRatio)

	// Resize image
	resized := resize.Resize(uint(width), newHeight, img, resize.Lanczos3)

	converter := convert.NewImageConverter()
	ascii := converter.Image2ASCIIString(resized, &convert.Options{
		Colored:     true,
		FixedHeight: 20,
	})
	return ascii
}
