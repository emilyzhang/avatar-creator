package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
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
		"#373dfa",
		"",
		"",
		"",
		"#3297a8",
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
				c := overlay(imageColor, newColor)
				newImg.Set(x, y, c)
			}
		}
	}
	return newImg
}

// hexColor returns an HTML hex-representation of c. The alpha channel is dropped
// and precision is truncated to 8 bits per channel
func hexColor(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

func overlay(base color.Color, top color.Color) color.Color {
	baseR, baseG, baseB, baseA := base.RGBA()
	topR, topG, topB, topA := top.RGBA()
	newColor := color.RGBA{
		R: overlayValue(baseR, topR, baseA, topA),
		G: overlayValue(baseG, topG, baseA, topA),
		B: overlayValue(baseB, topB, baseA, topA),
		A: overlayValue(baseA, topA, baseA, topA),
	}
	// fmt.Println("color1:", base, "color2:", top, "new", newColor)
	return newColor
}

func overlayValue(baseValue, topValue, baseOpacity, topOpacity uint32) uint8 {
	b := float64(baseValue) / float64(baseOpacity)
	t := float64(topValue) / float64(topOpacity)
	if b <= 0.5 {
		return uint8((int)(math.Min((2*b*t)*255, 255)))
	}
	return uint8((int)(math.Min(((1 - 2*(1-b)*(1-t)) * 255), 255)))
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
