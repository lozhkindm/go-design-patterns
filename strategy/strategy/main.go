package main

import (
	"fmt"
	"strings"
)

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type OutputFormat int

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(_ *strings.Builder) {}

func (m *MarkdownListStrategy) End(_ *strings.Builder) {}

func (m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("* %s\n", item))
}

type HtmlListStrategy struct {
}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("\t<li>%s</li>\n", item))
}

type TextProcessor struct {
	builder  strings.Builder
	strategy ListStrategy
}

func (t *TextProcessor) SetOutputFormat(format OutputFormat) {
	switch format {
	case Markdown:
		t.strategy = &MarkdownListStrategy{}
	case Html:
		t.strategy = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	t.strategy.Start(&t.builder)
	for _, item := range items {
		t.strategy.AddListItem(&t.builder, item)
	}
	t.strategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

func NewTextProcessor(strategy ListStrategy) *TextProcessor {
	return &TextProcessor{
		builder:  strings.Builder{},
		strategy: strategy,
	}
}

func main() {
	processor := NewTextProcessor(&MarkdownListStrategy{})
	processor.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(processor)

	processor.Reset()

	processor.SetOutputFormat(Html)
	processor.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(processor)
}
