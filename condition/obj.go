package condition

var objExps = map[string]obj{}

type obj interface {
	NewExpFromValues(values []interface{}) (interface{}, error)
}

type nowObj struct {
	//obj
}

func (o *nowObj) newExp() interface{} {
	return func(ctx Context) int {
	}
	panic("")
}
