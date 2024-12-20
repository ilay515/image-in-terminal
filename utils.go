package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
)

var IMAGE_FOLDER string = "images"

func file_exists(filename string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s.jpg", IMAGE_FOLDER, filename))
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func read_jpeg_image(filename string) (image image.Image, error error) {
	file, err := os.Open(fmt.Sprintf("%s/%s.jpg", IMAGE_FOLDER, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	return img, err
}

func get_image_from_url(url string, image_name string) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("%s/%s.jpg", IMAGE_FOLDER, image_name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}
