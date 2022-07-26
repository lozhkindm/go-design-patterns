package main

import "fmt"

type Buffer struct {
	width, height int
	buffer        []rune
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{
		width:  width,
		height: height,
		buffer: make([]rune, width*height),
	}
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func NewConsole() *Console {
	buffer := NewBuffer(200, 150)
	return &Console{
		buffers:   []*Buffer{buffer},
		viewports: []*Viewport{NewViewport(buffer)},
		offset:    0,
	}
}

func main() {
	console := NewConsole()
	char := console.GetCharacterAt(1)
	fmt.Println(char)
}
