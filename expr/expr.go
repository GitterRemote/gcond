package expr

// Context interface required in expr package
type Context interface {
	Value(key interface{}) interface{}
	TmpValue(key interface{}) (result interface{}, ok bool)
	SetTmpValue(key interface{}, result interface{}) (err error)
}
