package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (h *HtmlElement) String() string {
	return h.string(0)
}

func (h *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	ind := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", ind, h.name))
	if len(h.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(h.text)
		sb.WriteString("\n")
	}
	for _, e := range h.elements {
		sb.WriteString(e.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", ind, h.name))
	return sb.String()
}

type HtmlBuilder struct {
	root     HtmlElement
	rootName string
}

func (h *HtmlBuilder) String() string {
	return h.root.String()
}

func (h *HtmlBuilder) AddChild(childName, childText string) *HtmlBuilder {
	h.root.elements = append(h.root.elements, HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	})
	return h
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func main() {
	elements := []string{"Hello", "World!"}
	builder := NewHtmlBuilder("ul")
	for _, element := range elements {
		builder.AddChild("li", element)
	}
	fmt.Println(builder.String())
}
