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
			c.PutPixel(image.Point{x, ys[x-p0.X]}, col)
		}
	} else {
		if p0.Y > p1.Y {
			p0, p1 = p1, p0
		}
		xs := interpolate(p0.Y, p0.X, p1.Y, p1.X)
		for y := p0.Y; y <= p1.Y; y++ {
			c.PutPixel(image.Point{xs[y-p0.Y], y}, col)
		}
	}
}

func (c Canvas) DrawFramedTriangle(p0, p1, p2 image.Point, col color.RGBA) {
	c.DrawLine(p0, p1, col)
	c.DrawLine(p1, p2, col)
	c.DrawLine(p0, p2, col)
}

func (c Canvas) DrawFilledTriange(p0, p1, p2 image.Point, col color.RGBA) {
	if p1.Y < p0.Y {
		p1, p0 = p0, p1
	}
	if p2.Y < p0.Y {
		p2, p0 = p0, p2
	}
	if p2.Y < p1.Y {
		p2, p1 = p1, p2
	}

	x01 := interpolate(p0.Y, p0.X, p1.Y, p1.X)
	x12 := interpolate(p1.Y, p1.X, p2.Y, p2.X)
	x02 := interpolate(p0.Y, p0.X, p2.Y, p2.X)

	x01 = x01[:len(x01)-1]
	x012 := append(x01, x12...)

	m := len(x012) / 2
	var xLeft, xRight []int
	if x02[m] < x012[m] {
		xLeft = x02
		xRight = x012
	} else {
		xLeft = x012
		xRight = x02
	}
	for y := p0.Y; y <= p2.Y; y++ {
		for x := xLeft[y-p0.Y]; x <= xRight[y-p0.Y]; x++ {
			c.PutPixel(image.Point{x, y}, col)
		}
	}
}

func interpolate(i0, d0, i1, d1 int) []int {
	values := make([]int, i1-i0+1)
	d := float64(d0)

	if i1 == i0 {
		for i := range values {
			values[i] = int(d)
		}
	} else {
		a := float64(d1-d0) / float64(i1-i0)

		for i := i0; i <= i1; i++ {
			values[i-i0] = int(d)
			d = d + a
		}
	}

	return values
}
