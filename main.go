package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	filepaths := []string{
		"layers/face_outline.png",
		"layers/eye_color_1.png",
		"layers/eye_color_2.png",
		"layers/eye_color_3.png",
		"layers/eye_color_4.png",
		"layers/eye_outline.png",
		"layers/hair_color.png",
		"layers/hair_outline.png",
		"layers/lips_outline.png",
	}
	colors := []string{
		"",
		"#3297a8",
		"#0b6878",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	var imgs []image.Image
	for _, path := range filepaths {
		imgs = append(imgs, decodePNG(path))
	}

	newPNG := combine(imgs, colors)
	savePNG, err := os.Create("result.png")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	png.Encode(savePNG, newPNG)
	defer savePNG.Close()

}

func decodePNG(filepath string) image.Image {
	i, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	img, err := png.Decode(i)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer i.Close()
	return img
}

func combine(imgs []image.Image, colors []string) *image.RGBA {
	offset := image.Pt(0, 0)
	b := imgs[0].Bounds()
	newPNG := image.NewRGBA(b)
	draw.Draw(newPNG, b, &image.Uniform{color.White}, image.Point{}, draw.Src)
	for i, img := range imgs {
		img := swapColor(img, colors[i])
		draw.Draw(newPNG, img.Bounds().Add(offset), img, image.Point{}, draw.Over)
	}
	return newPNG
}

func swapColor(img image.Image, hexColor string) image.Image {
	if hexColor == "" {
		return img
	}
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)
	newColor, err := parseHexColorFast(hexColor)
	if err != nil {
		log.Fatalf("Could not parse hex color: %s", err)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := img.At(x, y)
			rr, gg, bb, aa := imageColor.RGBA()
			if rr != 0 || gg != 0 || bb != 0 || aa != 0 {
				// TODO: how do we do overlays? soft light? etc?
				fmt.Println(x, y, imageColor)
				// grayscale below
				// r := math.Pow(float64(rr), 2.2)
				// g := math.Pow(float64(gg), 2.2)
				// b := math.Pow(float64(bb), 2.2)
				// m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
				// Y := uint16(m + 0.5)
				// grayColor := color.Gray{uint8(Y >> 8)}
				newImg.Set(x, y, newColor)
			}
		}
	}
	return newImg
}

// below stolen from https://stackoverflow.com/a/54200713
var errInvalidFormat = errors.New("invalid format")

func parseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
