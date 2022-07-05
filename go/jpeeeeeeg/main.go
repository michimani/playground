package main

import (
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter for path of image is required.")
		return
	}

	path := args[1]

	img, err := getImageInfo(path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Width: %d Height: %d\n", img.Bounds().Dx(), img.Bounds().Dy())
	}

	orientation, err := getImageOrientation(path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Exif.Orientation: %d\n", orientation)
	}
}

func getImageInfo(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return img, nil
}

func getImageOrientation(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	e, err := exif.Decode(file)
	if err != nil {
		return 0, err
	}

	tag, err := e.Get(exif.Orientation)
	if err != nil {
		return 0, err
	}

	o, err := tag.Int(0)
	if err != nil {
		return 0, err
	}

	return o, nil
}
