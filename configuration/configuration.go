package configuration

import "github.com/GitterRemote/gcond/condition"

// Context contains the data to be used during evaluation
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
	Result    interface{}
}

// New a configuration
func New(condition Condition, result interface{}) *Configuration {
	return &Configuration{0, condition, result}
}
