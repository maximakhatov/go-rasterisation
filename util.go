package main

import "image/color"

var (
	Red   = color.RGBA{255, 0, 0, 255}
	Green = color.RGBA{0, 255, 0, 255}
	Blue  = color.RGBA{0, 0, 255, 255}
	Black = color.RGBA{0, 0, 0, 255}
	White = color.RGBA{255, 255, 255, 255}
)

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
