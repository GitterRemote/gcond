package condition

import "github.com/GitterRemote/gcond/expr"

// Loader is used to load condition by id
type Loader interface {
	LoadByID(id int) *idCondition
}

type dbLoader struct {
}

func loadFromStringEqualExpr(stringExprValue expr.Value, other string) *idCondition {
	panic("")
}

// LoadByID load condition by id from database
func (l *dbLoader) LoadByID(id int) *idCondition {
	panic("")
}
