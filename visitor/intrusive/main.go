package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Print(builder *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(builder *strings.Builder) {
	builder.WriteRune('(')
	a.left.Print(builder)
	builder.WriteRune('+')
	a.right.Print(builder)
	builder.WriteRune(')')
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
	exp.Print(&builder)
	fmt.Println(builder.String())
}
