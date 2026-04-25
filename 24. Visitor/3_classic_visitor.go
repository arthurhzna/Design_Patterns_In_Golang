package visitor

import (
	"fmt"
	"strings"
)

// ==================================================
// ExpressionVisitor interface
// ==================================================

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

// ==================================================
// Expression interface
// ==================================================

type Expression interface {
	Accept(ev ExpressionVisitor)
}

// ==================================================
// DoubleExpression
// ==================================================
type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

// ==================================================
// AdditionExpression
// ==================================================
type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

// ==================================================
// ExpressionPrinter
// ==================================================
type ExpressionPrinter struct {
	sb strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
	return e.sb.String()
}

func main() {
	// 1+(2+3)
	e := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	ep.VisitAdditionExpression(e)
	fmt.Println(ep.String())
}

// Addition  (root: 1 + (2+3))
// /        \
// Double(1)    Addition  (inner: 2+3)
// 		/        \
// 	Double(2)   Double(3)

// Classic visitor is a strong fit when…

// vs intrusive
// You have many different operations and you do not want domain
// structs cluttered with methods like Print, Evaluate, and so on.

// vs reflective
// You want to avoid a single function full of type switches,
// and you want behavior per concrete type to live in clear, separate places.
