package main

import (
	"log"
)

func main() {
	width, height := get_screen_size()
	// clear_screen()

	// img, err := read_jpeg_image("test-images/nighthawks.jpg")
	img, err := read_jpeg_image("test-images/monalisa.jpg")
	if err != nil {
		log.Fatal(err)
	}

	pixelated_image := pixelate_image(img, width, height)

	display_image(pixelated_image)
}
