package main

import (
	"image"
	"log"

	"github.com/alexflint/go-arg"
)

func handle_user_input() image.Image {
	var args struct {
		URL       string
		ImagePath string `arg:"positional"`
	}
	arg.MustParse(&args)

	if args.URL != "" && !file_exists(args.ImagePath) {
		get_image_from_url(args.URL, args.ImagePath)
	}

	img, err := read_jpeg_image(args.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func main() {
	img := handle_user_input()

	screen_width, screen_height := get_screen_size()
	screen_height -= 2 // for extra lines
	// clear_screen()
	pixelated_image := pixelate_image(img, screen_width, screen_height)

	display_image(pixelated_image)
}
