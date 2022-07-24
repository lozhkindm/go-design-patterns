package main

import "fmt"

// Circle, Square
// Raster, Vector

// RasterCircle, VectorCircle, RasterSquare, VectorSquare

type Renderer interface {
	RenderCircle(radius float32)
}

type RasterRenderer struct {
	// some fields
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("(raster) rendering a circle of radius", radius)
}

type VectorRenderer struct {
	// some fields
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("(vector) rendering a circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func main() {
	//raster := RasterRenderer{}
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
