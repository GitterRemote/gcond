package configuration

import "github.com/GitterRemote/gcond/condition"

type Context interface {
	condition.Context
}

// Condition can be evaluated based on the context
type Condition interface {
	Evaluate(ctx condition.Context) bool
}

// Configuration contains expression and result
type Configuration struct {
	ID        int
	Condition Condition
	Result    string
}
