package main

import (
	"math"
)

type Mat4x4 struct {
	Data [][]float64
}

var Identity4x4 = Mat4x4{[][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}

func MakeOYRotationMatrix(degrees float64) Mat4x4 {
	cos := math.Cos(degrees * math.Pi / 180)
	sin := math.Sin(degrees * math.Pi / 180)

	return Mat4x4{[][]float64{
		{cos, 0, -sin, 0},
		{0, 1, 0, 0},
		{sin, 0, cos, 0},
		{0, 0, 0, 1},
	}}
}

func MakeTranslationMatrix(transform Vertex) Mat4x4 {
	return Mat4x4{[][]float64{
		{1, 0, 0, transform.X},
		{0, 1, 0, transform.Y},
		{0, 0, 1, transform.Z},
		{0, 0, 0, 1},
	}}
}

func MakeScalingMatrix(scale float64) Mat4x4 {
	return Mat4x4{[][]float64{
		{scale, 0, 0, 0},
		{0, scale, 0, 0},
		{0, 0, scale, 0},
		{0, 0, 0, 1},
	}}
}

func (m Mat4x4) MultiplyMV(v Vertex4) Vertex4 {
	result := make([]float64, 4)
	vec := []float64{v.X, v.Y, v.Z, v.W}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i] += m.Data[i][j] * vec[j]
		}
	}
	return Vertex4{Vertex{result[0], result[1], result[2]}, result[3]}
}

func (ma Mat4x4) MultiplyMM4(mb Mat4x4) Mat4x4 {
	result := [][]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result[i][j] += ma.Data[i][k] * mb.Data[k][j]
			}
		}
	}

	return Mat4x4{result}
}

func (m Mat4x4) Transpose() Mat4x4 {
	result := [][]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = m.Data[j][i]
		}
	}

	return Mat4x4{result}
}
