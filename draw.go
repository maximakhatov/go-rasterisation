package main

import (
	"image"
	"image/color"
)

func (c Canvas) DrawLine(p0, p1 image.Point, col color.RGBA) {
	if Abs(p1.X-p0.X) > Abs(p1.Y-p0.Y) {
		if p0.X > p1.X {
			p0, p1 = p1, p0
		}
		ys := interpolate(p0.X, p0.Y, p1.X, p1.Y)
		for x := p0.X; x <= p1.X; x++ {
			c.PutPixel(image.Point{x, int(ys[x-p0.X])}, col)
		}
	} else {
		if p0.Y > p1.Y {
			p0, p1 = p1, p0
		}
		xs := interpolate(p0.Y, p0.X, p1.Y, p1.X)
		for y := p0.Y; y <= p1.Y; y++ {
			c.PutPixel(image.Point{int(xs[y-p0.Y]), y}, col)
		}
	}
}

func (c Canvas) DrawFramedTriangle(p0, p1, p2 image.Point, col color.RGBA) {
	c.DrawLine(p0, p1, col)
	c.DrawLine(p1, p2, col)
	c.DrawLine(p0, p2, col)
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
