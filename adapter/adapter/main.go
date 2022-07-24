package main

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{Lines: []Line{
		{X1: 0, Y1: 0, X2: width, Y2: 0},
		{X1: 0, Y1: 0, X2: 0, Y2: height},
		{X1: width, Y1: 0, X2: width, Y2: height},
		{X1: 0, Y1: height, X2: width, Y2: height},
	}}
}

// ↑↑↑ the interface we're given

// ↓↓↓ the interface we have

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	builder := strings.Builder{}
	for _, line := range data {
		builder.WriteString(string(line))
		builder.WriteRune('\n')
	}
	return builder.String()
}

// adapter
type vectorToRasterAdapter struct {
	points []Point
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{X: left, Y: y})
		}
	} else if dy == 0 {
		for x := 0; x <= right; x++ {
			v.points = append(v.points, Point{X: x, Y: top})
		}
	}

}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func NewVectorToRasterAdapter(vector *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vector.Lines {
		adapter.addLine(line)
	}
	return &adapter
}

func main() {
	rectangle := NewRectangle(6, 4)
	adapter := NewVectorToRasterAdapter(rectangle)
	fmt.Print(DrawPoints(adapter))
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
