package main

import (
	"image"
	"os"
)

const width = 600
const height = 600

var baseViewport = Viewport{1, 1, 1}
var cube = Model{[]Vertex{
	{1, 1, 1},
	{-1, 1, 1},
	{-1, -1, 1},
	{1, -1, 1},
	{1, 1, -1},
	{-1, 1, -1},
	{-1, -1, -1},
	{1, -1, -1},
},
	[]Triangle{
		{0, 1, 2, Red},
		{0, 2, 3, Red},
		{4, 0, 3, Green},
		{4, 3, 7, Green},
		{5, 4, 7, Blue},
		{5, 7, 6, Blue},
		{1, 5, 6, Yellow},
		{1, 6, 2, Yellow},
		{4, 5, 1, Magenta},
		{4, 1, 0, Magenta},
		{2, 6, 7, Cyan},
		{2, 7, 3, Cyan},
	}}

func main() {
	drawLines()
	drawTriangles()
	drawCubeByLines()
	drawCube()
	drawCubesScene()
}

func save(c Canvas, name string) {
	f, err := os.Create("renders/" + name + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c.Save(f)
}

func drawLines() {
	canvasLines := NewCanvas(width, height)
	canvasLines.DrawLine(image.Point{-200, -100}, image.Point{240, 120}, Red)
	canvasLines.DrawLine(image.Point{-50, -200}, image.Point{60, 240}, Green)
	canvasLines.DrawLine(image.Point{-70, -50}, image.Point{-70, 50}, Blue)
	canvasLines.DrawLine(image.Point{-100, 250}, image.Point{100, 250}, Black)
	canvasLines.DrawLine(image.Point{150, 150}, image.Point{150, 150}, Red)
	canvasLines.DrawLine(image.Point{100, 100}, image.Point{-100, -100}, Green)
	save(canvasLines, "lines")
}

func drawTriangles() {
	canvasTriangles := NewCanvas(width, height)
	canvasTriangles.DrawShadedTriange(image.Point{-200, -250}, image.Point{200, 50}, image.Point{20, 250}, 0.3, 0.1, 1.0, Green)
	canvasTriangles.DrawFramedTriangle(image.Point{-200, -250}, image.Point{200, 50}, image.Point{20, 250}, Black)
	save(canvasTriangles, "triangles")
}

func drawCubeByLines() {
	canvasCube := NewCanvas(width, height)
	vAf := Vertex{-2, -0.5, 5}
	vBf := Vertex{-2, 0.5, 5}
	vCf := Vertex{-1, 0.5, 5}
	vDf := Vertex{-1, -0.5, 5}
	vAb := Vertex{-2, -0.5, 6}
	vBb := Vertex{-2, 0.5, 6}
	vCb := Vertex{-1, 0.5, 6}
	vDb := Vertex{-1, -0.5, 6}
	canvasCube.DrawLine(canvasCube.ProjectVertex(vAf, baseViewport), canvasCube.ProjectVertex(vBf, baseViewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vBf, baseViewport), canvasCube.ProjectVertex(vCf, baseViewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vCf, baseViewport), canvasCube.ProjectVertex(vDf, baseViewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vDf, baseViewport), canvasCube.ProjectVertex(vAf, baseViewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vAb, baseViewport), canvasCube.ProjectVertex(vBb, baseViewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vBb, baseViewport), canvasCube.ProjectVertex(vCb, baseViewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vCb, baseViewport), canvasCube.ProjectVertex(vDb, baseViewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vDb, baseViewport), canvasCube.ProjectVertex(vAb, baseViewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vAf, baseViewport), canvasCube.ProjectVertex(vAb, baseViewport), Green)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vBf, baseViewport), canvasCube.ProjectVertex(vBb, baseViewport), Green)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vCf, baseViewport), canvasCube.ProjectVertex(vCb, baseViewport), Green)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vDf, baseViewport), canvasCube.ProjectVertex(vDb, baseViewport), Green)
	save(canvasCube, "cube_by_lines")
}

func drawCube() {
	vertices := cube.Vertices
	triangles := cube.Triangles

	newVertices := make([]Vertex, len(vertices))
	for i := range vertices {
		newVertices[i] = Vertex{vertices[i].X - 1.5, vertices[i].Y, vertices[i].Z + 7}
	}

	canvasCube := NewCanvas(width, height)
	canvasCube.RenderObject(newVertices, triangles, baseViewport)
	save(canvasCube, "cube")
}

func drawCubesScene() {
	insts := []Instance{
		{cube, Vertex{-1.5, 0, 7}},
		{cube, Vertex{1.25, 2, 7.5}},
	}

	canvasCubes := NewCanvas(width, height)
	canvasCubes.RenderScene(insts, baseViewport)
	save(canvasCubes, "cubes_scene")
}
