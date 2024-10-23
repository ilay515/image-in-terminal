package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

var ASCII_HEIGHT int = 3

func get_image_ratio(img image.Image) float64 {
	return float64(img.Bounds().Size().X) / float64(img.Bounds().Size().Y)
}

func create_pixelated_image(img image.Image, new_width int) *image.RGBA {
	ratio := get_image_ratio(img)
	new_height := int(float64(new_width) / ratio / float64(ASCII_HEIGHT))
	rect := image.Rect(0, 0, new_width, new_height)
	pixelated_image := image.NewRGBA(rect)

	return pixelated_image
}

func compress_pixels_block(original_image image.Image, startx int, starty int, block_size int) color.Color {
	var rSum, gSum, bSum, aSum uint32
	var count int

	// Iterate over the specified block of pixels
	for y := starty; y < starty+block_size*ASCII_HEIGHT; y++ {
		for x := startx; x < startx+block_size; x++ {
			pixelColor := original_image.At(x, y) // Get the pixel color

			// Extract RGBA values from the pixel color
			r, g, b, a := pixelColor.RGBA()

			// Accumulate the RGBA values
			rSum += r
			gSum += g
			bSum += b
			aSum += a
			count++
		}
	}

	// Calculate the average RGBA values
	averageR := uint8(rSum / uint32(count) >> 8) // Right shift to convert to 0-255 range
	averageG := uint8(gSum / uint32(count) >> 8)
	averageB := uint8(bSum / uint32(count) >> 8)
	averageA := uint8(aSum / uint32(count) >> 8)

	// Return the average color as a color.RGBA
	return color.RGBA{averageR, averageG, averageB, averageA}
}

func pixelate_image(img image.Image, new_width int) *image.RGBA {
	pixelated_image := create_pixelated_image(img, new_width)
	background := color.RGBA{0, 0xFF, 0, 0xCC}
	draw.Draw(pixelated_image, pixelated_image.Rect.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	block_size := img.Bounds().Size().X / new_width
	bounds := pixelated_image.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := compress_pixels_block(img, x*block_size, y*block_size*ASCII_HEIGHT, block_size)
			pixelated_image.Set(x, y, color)
		}
	}

	save_jpeg_image(pixelated_image, "text.jpg")

	return pixelated_image
}

func main() {
	clear_screen()
	terminal := *create_terminal(194, 47)

	img, err := read_jpeg_image("nighthawks.jpg")
	if err != nil {
		log.Fatal(err)
	}

	pixelated_image := pixelate_image(img, 194)

	print_image(pixelated_image, terminal)

	// // Get image bounds
	// bounds := img.Bounds()

	// // Loop over each pixel
	// for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
	// 	for x := bounds.Min.X; x < bounds.Max.X; x++ {
	// 		pixelColor := img.At(x, y).(color.Color)
	// 		_ = pixelColor

	// 		// Extract RGBA values from the pixel color
	// 		// r, g, b, a := pixelColor.RGBA()

	// 		fmt.Print("#")
	// 		// Print the pixel values (scaled down to 8-bit)
	// 		// fmt.Printf("Pixel at (%d, %d): R=%d, G=%d, B=%d, A=%d\n", x, y, r>>8, g>>8, b>>8, a>>8)
	// 	}
	// 	fmt.Printf("")
	// }
}
