package main

import (
	"image"
	"os"
)

const width = 600
const height = 600

func main() {
	canvasLines := NewCanvas(width, height)
	canvasLines.DrawLine(image.Point{-200, -100}, image.Point{240, 120}, Red)
	canvasLines.DrawLine(image.Point{-50, -200}, image.Point{60, 240}, Green)
	canvasLines.DrawLine(image.Point{-70, -50}, image.Point{-70, 50}, Blue)
	canvasLines.DrawLine(image.Point{-100, 250}, image.Point{100, 250}, Black)
	canvasLines.DrawLine(image.Point{150, 150}, image.Point{150, 150}, Red)
	canvasLines.DrawLine(image.Point{100, 100}, image.Point{-100, -100}, Green)
	save(canvasLines, "lines")

	canvasTriangles := NewCanvas(width, height)
	canvasTriangles.DrawShadedTriange(image.Point{-200, -250}, image.Point{200, 50}, image.Point{20, 250}, 0.3, 0.1, 1.0, Green)
	canvasTriangles.DrawFramedTriangle(image.Point{-200, -250}, image.Point{200, 50}, image.Point{20, 250}, Black)
	save(canvasTriangles, "triangles")

	viewport := Viewport{1, 1, 1}

	canvasCube := NewCanvas(width, height)
	vAf := Vertex{-2, -0.5, 5}
	vBf := Vertex{-2, 0.5, 5}
	vCf := Vertex{-1, 0.5, 5}
	vDf := Vertex{-1, -0.5, 5}
	vAb := Vertex{-2, -0.5, 6}
	vBb := Vertex{-2, 0.5, 6}
	vCb := Vertex{-1, 0.5, 6}
	vDb := Vertex{-1, -0.5, 6}
	canvasCube.DrawLine(canvasCube.ProjectVertex(vAf, viewport), canvasCube.ProjectVertex(vBf, viewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vBf, viewport), canvasCube.ProjectVertex(vCf, viewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vCf, viewport), canvasCube.ProjectVertex(vDf, viewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vDf, viewport), canvasCube.ProjectVertex(vAf, viewport), Blue)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vAb, viewport), canvasCube.ProjectVertex(vBb, viewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vBb, viewport), canvasCube.ProjectVertex(vCb, viewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vCb, viewport), canvasCube.ProjectVertex(vDb, viewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vDb, viewport), canvasCube.ProjectVertex(vAb, viewport), Red)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vAf, viewport), canvasCube.ProjectVertex(vAb, viewport), Green)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vBf, viewport), canvasCube.ProjectVertex(vBb, viewport), Green)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vCf, viewport), canvasCube.ProjectVertex(vCb, viewport), Green)
	canvasCube.DrawLine(canvasCube.ProjectVertex(vDf, viewport), canvasCube.ProjectVertex(vDb, viewport), Green)
	save(canvasCube, "cube")
}

func save(c Canvas, name string) {
	f, err := os.Create("renders/" + name + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c.Save(f)
}
