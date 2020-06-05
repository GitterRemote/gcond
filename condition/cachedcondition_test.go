package condition

import (
	"testing"

	"github.com/GitterRemote/gcond/context"
)

func TestCachedCondition_Evaluate(t *testing.T) {

	c := cachedCondition{newFirstTrueIDCondition()}

	ctx := context.New()
	rv := c.Evaluate(ctx)
	if rv != true {
		t.Fatal("result error")
	}
	rv = c.Evaluate(ctx)
	if rv != true {
		t.Fatal("cached fail")
	}

	ctx = context.New()
	rv = c.Evaluate(ctx)
	if rv != false {
		t.Fatal("cached error")
	}
}
