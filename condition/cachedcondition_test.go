package condition

import (
	"testing"

	"github.com/GitterRemote/gcond/context"
)

func TestCachedCondition_Evaluate(t *testing.T) {

	c := cachedCondition{newFirstTrueIDCondition()}

	ctx := context.NewContext()
	rv := c.Evaluate(ctx)
	if rv != true {
		t.Fatal("result error")
	}
	rv = c.Evaluate(ctx)
	if rv != true {
		t.Fatal("cached fail")
	}

	ctx = context.NewContext()
	rv = c.Evaluate(ctx)
	if rv != false {
		t.Fatal("cached error")
	}
}
