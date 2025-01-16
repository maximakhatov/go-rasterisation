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

func (c Canvas) DrawShadedTriange(p0, p1, p2 image.Point, h0, h1, h2 float64, col color.RGBA) {
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

	h12 := interpolateFloat(p1.Y, h1, p2.Y, h2)
	h02 := interpolateFloat(p0.Y, h0, p2.Y, h2)
	h01 := interpolateFloat(p0.Y, h0, p1.Y, h1)

	x01 = x01[:len(x01)-1]
	x012 := append(x01, x12...)

	h01 = h01[:len(h01)-1]
	h012 := append(h01, h12...)

	m := len(x012) / 2
	var xLeft, xRight []int
	var hLeft, hRight []float64
	if x02[m] < x012[m] {
		xLeft = x02
		xRight = x012
		hLeft = h02
		hRight = h012
	} else {
		xLeft = x012
		xRight = x02
		hLeft = h012
		hRight = h02
	}
	for y := p0.Y; y <= p2.Y; y++ {
		xl := xLeft[y-p0.Y]
		xr := xRight[y-p0.Y]
		hSegment := interpolateFloat(xl, hLeft[y-p0.Y], xr, hRight[y-p0.Y])
		for x := xLeft[y-p0.Y]; x <= xRight[y-p0.Y]; x++ {
			shadedCol := MultColor(col, hSegment[x-xl])
			c.PutPixel(image.Point{x, y}, shadedCol)
		}
	}
}

func (c Canvas) RenderObject(vertices []Vertex, triangles []Triangle, viewport Viewport) {
	projected := make([]image.Point, len(vertices))
	for i, v := range vertices {
		projected[i] = c.ProjectVertex(v, viewport)
	}
	for _, t := range triangles {
		c.DrawFramedTriangle(projected[t.v0], projected[t.v1], projected[t.v2], t.color)
	}
}

func (c Canvas) RenderInstance(inst Instance, viewport Viewport, transform Mat4x4) {
	projected := make([]image.Point, len(inst.Vertices))
	for i, v := range inst.Vertices {
		v4 := Vertex4{v, 1}
		projected[i] = c.ProjectVertex(transform.MultiplyMV(v4).Vertex, viewport)
	}
	for _, t := range inst.Triangles {
		c.DrawFramedTriangle(projected[t.v0], projected[t.v1], projected[t.v2], t.color)
	}
}

func (c Canvas) RenderScene(insts []Instance, viewport Viewport, cam Camera) {
	cameraTranslation := MakeTranslationMatrix(Vertex{-cam.Position.X, -cam.Position.Y, -cam.Position.Z})
	cameraMatrix := cam.Orientation.Transpose().MultiplyMM4(cameraTranslation)

	for _, inst := range insts {
		transform := cameraMatrix.MultiplyMM4(inst.Transform())
		c.RenderInstance(inst, viewport, transform)
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

func interpolateFloat(i0 int, d0 float64, i1 int, d1 float64) []float64 {
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
