package condition

import (
	"fmt"
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
func NewBuiltInCommand() *BuiltInCommand {
	return &BuiltInCommand{ctxCommandNames: make(map[string]bool)}
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

var cmds = map[string]cmdDeprecated{
	"mod":       &modCMD{},
	"len":       &lenCMD{},
	"condition": &conditionCMD{},
}

type cmdDeprecated interface {
	NewExpFromValues(values []interface{}) (interface{}, error)
}

type modCMD struct {
	cmdDeprecated
}

func (c *modCMD) execute(x, y int) int {
	rv := math.Mod(float64(x), float64(y))
	return int(rv)
}

func (c *modCMD) newExp(e intExp, y int) intExp {
	return func(ctx Context) int {
		x := e(ctx)
		return c.execute(x, y)
	}
}

func (c *modCMD) NewExpFromValues(values []interface{}) (interface{}, error) {
	if len(values) != 2 {
		return nil, fmt.Errorf("mod need values number to be 2")
	}
	return c.newExp(values[0].(intExp), values[1].(int)), nil
}

type lenCMD struct {
	cmdDeprecated
}

// Lenable defines the objects that have length
type Lenable interface {
	Len() int
}

type lenableExp func(ctx Context) Lenable

func (c *lenCMD) execute(o Lenable) int {
	return o.Len()
}

func (c *lenCMD) newExp(e lenableExp) intExp {
	return func(ctx Context) int {
		return c.execute(e(ctx))
	}
}

func (c *lenCMD) NewExpFromValues(values []interface{}) (interface{}, error) {
	if len(values) != 1 {
		return nil, &ParseValueLengthError{1}
	}
	return c.newExp(values[0].(lenableExp)), nil
}

type conditionCMD struct {
	cmdDeprecated
}

func (o *conditionCMD) NewExp(c Condition) (BoolExp, error) {
	return func(ctx Context) bool {
		return c.Evaluate(ctx)
	}, nil
}

func (o *conditionCMD) NewExpFromValues(values []interface{}) (interface{}, error) {
	var c Condition
	switch v := values[0].(type) {
	case Condition:
		c = v
	case int:
		return nil, fmt.Errorf("not implemented error")
	default:
		return nil, fmt.Errorf("Unknown type %v", v)
	}
	return o.NewExp(c)
}
