package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	builder := strings.Builder{}
	g.print(&builder, 0)
	return builder.String()
}

func (g *GraphicObject) print(builder *strings.Builder, depth int) {
	builder.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		builder.WriteString(g.Color)
		builder.WriteRune(' ')
	}
	builder.WriteString(g.Name)
	builder.WriteRune('\n')
	for _, child := range g.Children {
		child.print(builder, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{
		Name:     "Circle",
		Color:    color,
		Children: nil,
	}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{
		Name:     "Square",
		Color:    color,
		Children: nil,
	}
}

func main() {
	drawing := GraphicObject{"My Drawing", "Black", nil}
	drawing.Children = []GraphicObject{
		*NewCircle("Red"),
		*NewSquare("Yellow"),
		{
			Name:  "Group",
			Color: "Black",
			Children: []GraphicObject{
				*NewCircle("Blue"),
				*NewSquare("Green"),
			},
		},
	}

	fmt.Println(drawing.String())
}
