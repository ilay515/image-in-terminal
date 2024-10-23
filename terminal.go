package main

import (
	"fmt"
	"image"

	"github.com/fatih/color"
)

type Terminal struct {
	width  int
	height int
}

func clear_screen() {
	fmt.Printf("\033c")
}

func create_terminal(width int, height int) *Terminal {
	return &Terminal{
		width:  width,
		height: height,
	}
}

func print_image(img image.Image, terminal Terminal) {
	for y := 0; y < terminal.height-1; y++ {
		for x := 0; x < terminal.width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// fmt.Printf("Pixel at (%d, %d): R=%d, G=%d, B=%d", x, y, r, g, b)
			color.RGB(int(r>>8), int(g>>8), int(b>>8)).Print("#")
		}
		fmt.Printf("")
	}
}
