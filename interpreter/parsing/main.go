package main

import (
	"fmt"
	"strconv"
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

const (
	Addition Operation = iota
	Subtraction
)

type Operation int
type TokenType int

type Element interface {
	Value() int
}

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() - b.Right.Value()
	default:
		panic("unsupported operation")
	}
}

type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

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

func Parse(tokens []Token) Element {
	result := &BinaryOperation{}
	haveLeft := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case Int:
			v, _ := strconv.Atoi(token.Text)
			if !haveLeft {
				result.Left = NewInteger(v)
				haveLeft = true
			} else {
				result.Right = NewInteger(v)
			}
		case Plus:
			result.Type = Addition
		case Minus:
			result.Type = Subtraction
		case Lparen:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == Rparen {
					break
				}
			}
			var subexp []Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}
			element := Parse(subexp)
			if !haveLeft {
				result.Left = element
				haveLeft = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return result
}

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)

	parsed := Parse(tokens)
	fmt.Printf("%s = %d\n", input, parsed.Value())
}
