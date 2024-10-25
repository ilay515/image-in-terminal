package main

import (
	"image"
	"image/color"
)

var ASCII_HEIGHT float64 = 2.7

func get_image_ratio(img image.Image) float64 {
	return float64(img.Bounds().Size().X) / float64(img.Bounds().Size().Y)
}

func create_pixelated_image(img image.Image, screen_width int, screen_height int) *image.RGBA {
	img_ratio := get_image_ratio(img)
	screen_ratio := float64(screen_width) / (float64(screen_height) * float64(ASCII_HEIGHT))

	if img_ratio > screen_ratio {
		height := int(float64(screen_width) / img_ratio / float64(ASCII_HEIGHT))
		rect := image.Rect(0, 0, screen_width, height)
		return image.NewRGBA(rect)
	}

	width := int(float64(screen_height) * img_ratio * float64(ASCII_HEIGHT))
	rect := image.Rect(0, 0, width, screen_height)

	return image.NewRGBA(rect)
}

func compress_pixels_block(original_image image.Image, startx int, starty int, block_size int) color.Color {
	var rSum, gSum, bSum, aSum uint32
	var count int

	for y := starty; y < starty+block_size*int(ASCII_HEIGHT); y++ {
		for x := startx; x < startx+block_size; x++ {
			pixelColor := original_image.At(x, y)

			r, g, b, a := pixelColor.RGBA()

			rSum += r
			gSum += g
			bSum += b
			aSum += a
			count++
		}
	}

	averageR := uint8(rSum / uint32(count) >> 8)
	averageG := uint8(gSum / uint32(count) >> 8)
	averageB := uint8(bSum / uint32(count) >> 8)
	averageA := uint8(aSum / uint32(count) >> 8)

	return color.RGBA{averageR, averageG, averageB, averageA}
}

func pixelate_image(img image.Image, screen_width int, screen_height int) *image.RGBA {
	pixelated_image := create_pixelated_image(img, screen_width, screen_height)

	block_size := float64(img.Bounds().Size().X) / float64(pixelated_image.Bounds().Size().X)
	bounds := pixelated_image.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := compress_pixels_block(
				img,
				int(float64(x)*block_size),
				int(float64(y)*block_size*float64(ASCII_HEIGHT)),
				int(block_size))
			pixelated_image.Set(x, y, color)
		}
	}

	return pixelated_image
}
