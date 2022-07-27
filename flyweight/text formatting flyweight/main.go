package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return t.Start <= position && position <= t.End
}

type FormattedText struct {
	plainText  string
	formatting []*TextRange
}

func (f *FormattedText) Range(start, end int) *TextRange {
	rng := &TextRange{
		Start:      start,
		End:        end,
		Capitalize: false,
		Bold:       false,
		Italic:     false,
	}
	f.formatting = append(f.formatting, rng)
	return rng
}

func (f *FormattedText) String() string {
	builder := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		char := f.plainText[i]
		for _, format := range f.formatting {
			if format.Covers(i) && format.Capitalize {
				char = uint8(unicode.ToUpper(rune(char)))
			}
		}
		builder.WriteRune(rune(char))
	}
	return builder.String()
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText: plainText}
}

func main() {
	text := "This is a brave new world"
	formatted := NewFormattedText(text)
	formatted.Range(10, 15).Capitalize = true
	fmt.Println(formatted.String())
}
