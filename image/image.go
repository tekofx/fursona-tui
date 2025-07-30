package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/qeesung/image2ascii/convert"
	"github.com/tekofx/fursona-tui/config"
	"github.com/tekofx/fursona-tui/style"
)

func Image2Ascii() string {
	configPath := config.GetConfigPath()
	var img image.Image

	// Find the first .png or .jpg file in configPath
	found := false
	filepath.WalkDir(configPath, func(path string, d fs.DirEntry, err error) error {
		if found || err != nil || d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			file, openErr := os.Open(path)
			if openErr != nil {
				return nil
			}
			defer file.Close()
			if ext == ".png" {
				img, err = png.Decode(file)
			} else {
				img, err = jpeg.Decode(file)
			}
			found = (err == nil)
			return nil
		}
		return nil
	})

	if !found || img == nil {
		fmt.Println(style.Error.Render("No image found in config path. Add a JPG or PNG file in", configPath))
		os.Exit(0)
	}

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
