package condition

import "testing"

func fakeBoolExpr(ctx Context) bool {
	return true
}

func TestGetID(t *testing.T) {
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
