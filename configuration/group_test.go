package configuration

import (
	"testing"

	"github.com/GitterRemote/gcond/condition"
	"github.com/GitterRemote/gcond/context"
)

type trueCondition struct {
}

func (c *trueCondition) Evaluate(ctx condition.Context) bool {
	return true
}

func newContext() Context {
	return context.New()
}

var trueConf = Configuration{1, &trueCondition{}, "1"}

func TestEvaluate(t *testing.T) {
	g := group([]Configuration{trueConf})
	conf := g.Evaluate(newContext())
	if conf.ID != trueConf.ID {
		t.Error("evalute error")
	}
}
