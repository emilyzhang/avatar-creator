package main

import (
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
	var imgs []image.Image
	for _, path := range filepaths {
		imgs = append(imgs, decodePNG(path))
	}

	newPNG := combine(imgs)
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

func combine(imgs []image.Image) *image.RGBA {
	offset := image.Pt(0, 0)
	b := imgs[0].Bounds()
	newPNG := image.NewRGBA(b)
	draw.Draw(newPNG, b, &image.Uniform{color.White}, image.ZP, draw.Src)
	for _, img := range imgs {
		draw.Draw(newPNG, img.Bounds().Add(offset), img, image.ZP, draw.Over)
	}
	return newPNG
}
