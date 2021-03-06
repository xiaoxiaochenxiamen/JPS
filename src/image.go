package jps

import (
	"fmt"
	"image"
	"os"
)

import _ "image/png"

type MapData map[int]bool

func NewMapData() MapData {
	return make(MapData)
}

func openImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()
	img, _, _ := image.Decode(f)
	return img
}

func parseImage(img image.Image) MapData {
	max := uint32(65536 - 1) // 2^16-1

	bounds := img.Bounds()
	fmt.Printf("width = %v, height = %v\n", bounds.Max.Y, bounds.Max.X)
	map_data := NewMapData()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			key := x + bounds.Max.X*y
			if r == max && g == max && b == max && a == max {
				map_data[key] = true
				//fmt.Printf(".")
			} else {
				map_data[key] = false
				//fmt.Printf("#")
			}
		}
		//fmt.Println("")
	}
	return map_data
}

func GetMapFromImage(filename string) MapData {
	img := openImage(filename)
	if img == nil {
		return nil
	}
	return parseImage(img)
}
