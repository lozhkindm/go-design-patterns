package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type TokenType int

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
	var result []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Type: Plus, Text: "+"})
		case '-':
			result = append(result, Token{Type: Minus, Text: "-"})
		case '(':
			result = append(result, Token{Type: Lparen, Text: "("})
		case ')':
			result = append(result, Token{Type: Rparen, Text: ")"})
		default:
			builder := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					builder.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, Token{Type: Int, Text: builder.String()})
					i--
					break
				}
			}
		}
	}

	return result
}

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)
}
