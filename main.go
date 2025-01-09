package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type point struct {
	x, y int
}

const width = 600
const height = 600

var img = image.NewRGBA(image.Rect(0, 0, width, height))

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}
var black = color.RGBA{0, 0, 0, 255}
var white = color.RGBA{255, 255, 255, 255}

func main() {
	prepareCanvas()
	drawLine(point{-200, -100}, point{240, 120}, red)
	drawLine(point{-50, -200}, point{60, 240}, green)
	drawLine(point{-70, -50}, point{-70, 50}, blue)
	drawLine(point{-100, 250}, point{100, 250}, black)
	drawLine(point{150, 150}, point{150, 150}, red)
	save("lines")
}

func prepareCanvas() {
	for x := range width {
		for y := range height {
			img.Set(x, y, white)
		}
	}
}

func save(name string) {
	f, err := os.Create("renders/" + name + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func putPixel(p point, col color.RGBA) {
	x := width/2 + p.x
	y := height/2 - p.y + 1
	img.Set(x, y, col)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func drawLine(p0, p1 point, col color.RGBA) {
	if abs(p1.x-p0.x) > abs(p1.y-p0.y) {
		if p0.x > p1.x {
			p0, p1 = p1, p0
		}
		ys := interpolate(p0.x, p0.y, p1.x, p1.y)
		for x := p0.x; x <= p1.x; x++ {
			putPixel(point{x, int(ys[x-p0.x])}, col)
		}
	} else {
		if p0.y > p1.y {
			p0, p1 = p1, p0
		}
		xs := interpolate(p0.y, p0.x, p1.y, p1.x)
		for y := p0.y; y <= p1.y; y++ {
			putPixel(point{int(xs[y-p0.y]), y}, col)
		}
	}
}

func interpolate(i0, d0, i1, d1 int) []float64 {
	values := make([]float64, i1-i0+1)
	d := float64(d0)

	if i1 == i0 {
		for i := range values {
			values[i] = d
		}
	} else {
		a := float64(d1-d0) / float64(i1-i0)

		for i := i0; i <= i1; i++ {
			values[i-i0] = d
			d = d + a
		}
	}

	return values
}
