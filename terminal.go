package main

import (
	"fmt"
	"image"

	"github.com/fatih/color"
)

func print_image(img image.Image) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// fmt.Printf("Pixel at (%d, %d): R=%d, G=%d, B=%d", x, y, r, g, b)
			color.RGB(int(r>>8), int(g>>8), int(b>>8)).Print("#")
		}
		fmt.Printf("")
	}
}
