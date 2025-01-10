package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Canvas struct {
	width, height int
	img           *image.RGBA
}

func NewCanvas(width, height int) Canvas {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := range width {
		for y := range height {
			img.Set(x, y, White)
		}
	}
	return Canvas{width, height, img}
}

func (c Canvas) PutPixel(p image.Point, col color.RGBA) {
	x := width/2 + p.X
	y := height/2 - p.Y + 1
	c.img.Set(x, y, col)
}

func (c Canvas) Save(f *os.File) {
	png.Encode(f, c.img)
}
