package main

import (
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Image string `arg:"positional"`
	}

	arg.MustParse(&args)
	fmt.Println("Input:", args.Image)

	screen_width, screen_height := get_screen_size()
	screen_height -= 2 // for extra lines
	// clear_screen()

	img, err := read_jpeg_image(fmt.Sprintf("test-images/%s", args.Image))
	if err != nil {
		log.Fatal(err)
	}

	pixelated_image := pixelate_image(img, screen_width, screen_height)

	display_image(pixelated_image)
}
