package condition

import "github.com/GitterRemote/gcond/expr"

// ExprCommand is command expand expr.BuiltInCommand to add condition loader
type ExprCommand struct {
	*expr.BuiltInCommand
}

// Condition create an ObjExp value by loading from condition ID
func (c *ExprCommand) Condition(conditionID int) expr.ObjExpr {
	// TODO: load condition first
	var cond Condition
	if cond == nil {
		panic("not implemented")
	}
	return func(ctx expr.Context) (rv interface{}) {
		conditionID := cond.GetID()
		if conditionID != nil {
			if rvInterface, ok := ctx.TmpValue(conditionID); ok {
				return rvInterface.(bool)
			}
			rv = cond.Evaluate(ctx)
			err := ctx.SetTmpValue(conditionID, rv)
			if err != nil {
				panic(err)
			}
		} else {
			rv = cond.Evaluate(ctx)
		}
		return
	}
}
