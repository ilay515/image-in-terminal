package main

import (
	"fmt"
	"image"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func clear_screen() {
	fmt.Printf("\033c")
}

func get_screen_size() (width, height int) {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	return width, height
}

func display_image(img image.Image) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		fmt.Println("")
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			color.RGB(int(r>>8), int(g>>8), int(b>>8)).Print("\u2588")
		}
	}
}
