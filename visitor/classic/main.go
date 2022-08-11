package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(exp *DoubleExpression)
	VisitAdditionExpression(exp *AdditionExpression)
}

type Expression interface {
	Accept(visitor ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	builder strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression(exp *DoubleExpression) {
	e.builder.WriteString(fmt.Sprintf("%g", exp.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(exp *AdditionExpression) {
	e.builder.WriteRune('(')
	exp.left.Accept(e)
	e.builder.WriteRune('+')
	exp.right.Accept(e)
	e.builder.WriteRune(')')
}

func (e *ExpressionPrinter) String() string {
	return e.builder.String()
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{builder: strings.Builder{}}
}

type Evaluator struct {
	result float64
}

func (e *Evaluator) VisitDoubleExpression(exp *DoubleExpression) {
	e.result = exp.value
}

func (e *Evaluator) VisitAdditionExpression(exp *AdditionExpression) {
	exp.left.Accept(e)
	res := e.result
	exp.right.Accept(e)
	res += e.result
	e.result = res
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
	printer := NewExpressionPrinter()
	exp.Accept(printer)
	fmt.Println(printer)

	evaluator := &Evaluator{}
	exp.Accept(evaluator)
	fmt.Println(evaluator.result)
}
