package util

import (
	_ "github.com/oov/psd"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func GetImageSize(path string) (int, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	im, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0, err
	}
	return im.Width, im.Height, nil
}
