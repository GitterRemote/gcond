package condition

import "github.com/GitterRemote/gcond/expr"

// Context is the data to be evaluated by the Condition
type Context interface {
	expr.Context
}

// Condition can be evaluted to be ture or false
type Condition interface {
	GetID() interface{}
	Evaluate(ctx Context) bool
}

// BoolExpr is expression under Condition
type BoolExpr func(ctx Context) bool

// New a configuration.Condition
func New(id int, expr BoolExpr) Condition {
	return &idCondition{id, expr}
}

// NewNaked new a condition only with expression
func NewNaked(expr BoolExpr) Condition {
	id := 0
	return New(id, expr)
}

// NewTrueCondition always evaluted to true
func NewTrueCondition() Condition {
	return New(0, func(ctx Context) bool {
		return true
	})
}
