package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math"
)

type Image struct {
	w, h  int
	pixel [][]color.Color
}

func draw(i *Image) {
	var pixel [][]color.Color
	for x := 0; x < i.w; x++ {
		var row []color.Color
		for y := 0; y < i.h; y++ {
			row = append(row, color.RGBA{
				R: uint8((x + y) / 2),
				G: uint8(x * y),
				B: uint8(float64(x) * math.Log(float64(y))),
				A: uint8(x % (y + 1)),
			})
		}
		pixel = append(pixel, row)
	}
	i.pixel = pixel
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: i.w, Y: i.h},
	}
}

func (i *Image) At(x, y int) color.Color {
	return i.pixel[x][y]
}

func main() {
	m := &Image{
		w:     200,
		h:     200,
		pixel: nil,
	}
	draw(m)
	pic.ShowImage(m)
}
