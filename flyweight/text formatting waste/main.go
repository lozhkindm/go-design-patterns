package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func (f *FormattedText) String() string {
	builder := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		char := f.plainText[i]
		if f.capitalize[i] {
			builder.WriteRune(unicode.ToUpper(rune(char)))
		} else {
			builder.WriteRune(rune(char))
		}
	}
	return builder.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{
		plainText:  plainText,
		capitalize: make([]bool, len(plainText)),
	}
}

func main() {
	text := "This is a brave new world"
	formatted := NewFormattedText(text)
	formatted.Capitalize(10, 15)
	fmt.Println(formatted.String())
}
