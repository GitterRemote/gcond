package condition

import "github.com/GitterRemote/gcond/expr"

// Loader is used to load condition by id
type Loader interface {
	LoadByID(id int) *idCondition
}

type dbLoader struct {
}

func loadFromJSONExpr(jsonExpr expr.JSONExpr) Condition {
	panic("")
}

// LoadByID load condition by id from database
func (l *dbLoader) LoadByID(id int) Condition {
	panic("")
}
