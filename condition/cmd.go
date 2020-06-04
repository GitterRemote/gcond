package condition

import (
	"math"
)

// Command includes all funcs used in expression
type Command interface {
	AddCtxCmdName(string)
	HasCtxCmdName(string) bool
}

// BuiltInCommand implements some operator commands
type BuiltInCommand struct {
	operator
	ctxCommandNames map[string]bool
}

// NewBuiltInCommand create a BuiltInCommand
func NewBuiltInCommand(ctxCmdNames ...string) *BuiltInCommand {
	ctxCommandNames := make(map[string]bool)
	for _, name := range ctxCmdNames {
		ctxCommandNames[name] = true
	}
	return &BuiltInCommand{ctxCommandNames: ctxCommandNames}
}

// AddCtxCmdName add name of func as a CtxCommand in Command
func (c *BuiltInCommand) AddCtxCmdName(name string) {
	c.ctxCommandNames[name] = true
}

// HasCtxCmdName check a cmd name is a CtxCommand or not
func (c *BuiltInCommand) HasCtxCmdName(name string) bool {
	_, ok := c.ctxCommandNames[name]
	return ok
}

// CtxCommand includes all funcs used in expression but require Context as first parameter
type CtxCommand struct{}

// TestContext is a command used to test context
func (c *CtxCommand) TestContext(ctx Context) bool {
	return ctx.Value("test").(bool)
}

// Mod do modulus for integer
func (c *BuiltInCommand) Mod(x, y int) int {
	rv := math.Mod(float64(x), float64(y))
	return int(rv)
}

// Lenable defines the objects that have length
type Lenable interface {
	Len() int
}

// Len call object's Len() method
func (c *BuiltInCommand) Len(o Lenable) int {
	return o.Len()
}

// Condition create an ObjExp value by loading from condition ID
func (c *BuiltInCommand) Condition(conditionID int) ObjExp {
	// TODO: load condition first
	var condition *Condition
	if condition == nil {
		panic("not implemented")
	}
	return func(ctx Context) interface{} {
		return condition.Evaluate(ctx)
	}
}
