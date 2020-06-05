package condition

import (
	"testing"
)

func fakeBoolExpr(ctx Context) bool {
	return true
}

func TestIDCondition_GetID(t *testing.T) {
	c := &idCondition{id: 0, exp: fakeBoolExpr}
	if c.GetID() == nil {
		t.Fatal("GetID return nil")
	}
	c2 := c
	if c.GetID() != c2.GetID() {
		t.Fatal("id not the same")
	}

	c3 := &idCondition{id: 1, exp: fakeBoolExpr}
	// refer: https://stackoverflow.com/questions/41498385/interface-integer-comparison-in-golang
	if c3.GetID() != 1 {
		t.Fatal("id not the setting")
	}
}

func newFirstTrueBoolExpr() BoolExpr {
	counter := new(int)

	return func(ctx Context) bool {
		if *counter == 0 {
			*counter++
			return true
		}
		return false
	}

}

func newFirstTrueIDCondition() *idCondition {
	c := &idCondition{
		0, newFirstTrueBoolExpr(),
	}
	return c
}

func TestIDCondition_Evaluate(t *testing.T) {

	c := newFirstTrueIDCondition()
	rv := c.Evaluate(nil)
	if rv != true {
		t.Fatal("result error")
	}
	rv = c.Evaluate(nil)
	if rv != false {
		t.Fatal("result error")
	}
}
