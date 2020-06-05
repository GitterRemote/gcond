package condition

type cachedCondition struct {
	*idCondition
}

func (c *cachedCondition) Evaluate(ctx Context) (rv bool) {
	conditionID := c.GetID()
	if conditionID != nil {
		if cachedRv, ok := ctx.TmpValue(conditionID); ok {
			return cachedRv.(bool)
		}
		rv = c.idCondition.Evaluate(ctx)
		err := ctx.SetTmpValue(conditionID, rv)
		if err != nil {
			panic(err)
		}
	} else {
		rv = c.idCondition.Evaluate(ctx)
	}
	return
}
