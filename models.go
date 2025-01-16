package main

import (
	"image/color"
)

type Viewport struct {
	VW, VH, D float64
}

type Camera struct {
	Position    Vertex
	Orientation Mat4x4
}

type Vertex struct {
	X, Y, Z float64
}

type Vertex4 struct {
	Vertex
	W float64
}

type Triangle struct {
	v0, v1, v2 int
	color      color.RGBA
}

type Model struct {
	Vertices  []Vertex
	Triangles []Triangle
}

type Instance struct {
	Model
	Position    Vertex
	Orientation Mat4x4
	Scale       float64
}

func (i Instance) Transform() Mat4x4 {
	return MakeTranslationMatrix(i.Position).MultiplyMM4(i.Orientation.MultiplyMM4(MakeScalingMatrix(i.Scale)))
}
