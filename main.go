package main

import (
	"image"
	"os"
)

const width = 600
const height = 600

func main() {
	c := NewCanvas(width, height)
	c.DrawLine(image.Point{-200, -100}, image.Point{240, 120}, Red)
	c.DrawLine(image.Point{-50, -200}, image.Point{60, 240}, Green)
	c.DrawLine(image.Point{-70, -50}, image.Point{-70, 50}, Blue)
	c.DrawLine(image.Point{-100, 250}, image.Point{100, 250}, Black)
	c.DrawLine(image.Point{150, 150}, image.Point{150, 150}, Red)
	c.DrawLine(image.Point{100, 100}, image.Point{-100, -100}, Green)
	save(c, "lines")

	c2 := NewCanvas(width, height)
	c2.DrawShadedTriange(image.Point{-200, -250}, image.Point{200, 50}, image.Point{20, 250}, 0.3, 0.1, 1.0, Green)
	c2.DrawFramedTriangle(image.Point{-200, -250}, image.Point{200, 50}, image.Point{20, 250}, Black)
	save(c2, "triangles")
}

func save(c Canvas, name string) {
	f, err := os.Create("renders/" + name + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c.Save(f)
}
