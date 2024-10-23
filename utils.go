package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func read_jpeg_image(filename string) (image image.Image, error error) {
	// Open the image file
	file, err := os.Open(filename) // Replace with your image file path
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image (supports JPEG, PNG, etc., based on import)
	img, err := jpeg.Decode(file)
	return img, err
}

func save_jpeg_image(img image.Image, filename string) {
	out, err := os.Create(filename)
	var opt jpeg.Options
	opt.Quality = 80
	err = jpeg.Encode(out, img, &opt)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
