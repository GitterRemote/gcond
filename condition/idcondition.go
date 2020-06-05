package condition

// idCondition implement configuration.idCondition
type idCondition struct {
	id  int
	exp BoolExpr
}

// GetID return the id of Condition
func (c *idCondition) GetID() interface{} {
	if c.id == 0 {
		return &c.exp
	}
	return c.id
}

// Evaluate the condition to a suc or fail value
func (c *idCondition) Evaluate(ctx Context) bool {
	return c.exp(ctx)
}
