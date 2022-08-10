package main

import (
	"fmt"
	"strings"
)

type Expression interface{}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(exp Expression, builder *strings.Builder) {
	if double, ok := exp.(*DoubleExpression); ok {
		builder.WriteString(fmt.Sprintf("%g", double.value))
	} else if addition, ok := exp.(*AdditionExpression); ok {
		builder.WriteRune('(')
		Print(addition.left, builder)
		builder.WriteRune('+')
		Print(addition.right, builder)
		builder.WriteRune(')')
	}
}

func main() {
	// 1+(2+3)
	exp := &AdditionExpression{
		left: &DoubleExpression{value: 1},
		right: &AdditionExpression{
			left:  &DoubleExpression{value: 2},
			right: &DoubleExpression{value: 3},
		},
	}
	builder := strings.Builder{}
	Print(exp, &builder)
	fmt.Println(builder.String())
}
