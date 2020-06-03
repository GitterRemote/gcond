package configuration

import (
	"testing"

	"github.com/GitterRemote/gcond/condition"
)

type trueCondition struct {
}

func (c *trueCondition) Evaluate(ctx condition.Context) bool {
	return true
}

var trueConf = Configuration{1, &trueCondition{}, "1"}

func TestEvaluate(t *testing.T) {
	ctx := condition.NewContext()
	g := Group([]Configuration{trueConf})
	conf := g.Evaluate(ctx)
	if conf.ID != trueConf.ID {
		t.Error("evalute error")
	}
}
