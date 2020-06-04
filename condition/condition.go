package condition

// Context is the data to be evaluated by the Condition
type Context interface {
	GetConditionResult(conditionID int) (result bool, ok bool)
	SetConditionResult(conditionID int, result bool) (err error)
	Value(key interface{}) interface{}
}

// BoolExp is expression under Condition
type BoolExp func(ctx Context) bool
type stringExp func(ctx Context) string
type intExp func(ctx Context) int

type expRegistry map[string]interface{}

// Condition implement configuration.Condition
type Condition struct {
	ID  int
	exp BoolExp
}

// Evaluate the condition to a suc or fail value
func (c *Condition) Evaluate(ctx Context) (rv bool) {
	if c.ID != 0 {
		if rv, ok := ctx.GetConditionResult(c.ID); ok {
			return rv
		}
		rv = c.exp(ctx)
		ctx.SetConditionResult(c.ID, rv)
	} else {
		rv = c.exp(ctx)
	}
	return
}

// New a configuration.Condition
func New(id int, exp BoolExp) *Condition {
	return &Condition{id, exp}
}

// NewTrueCondition always evaluted to true
func NewTrueCondition() *Condition {
	return New(0, func(ctx Context) bool {
		return true
	})
}
