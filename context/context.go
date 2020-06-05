package context

// SimpleContext implments get/set condition results
type SimpleContext struct {
	conditionResults map[interface{}]interface{}
	values           map[string]interface{}
}

// ExistsError is a err occurred when setConditionResult
type ExistsError string

func (e ExistsError) Error() string {
	return string(e)
}

// TmpValue get result of condition pre-evaluted
func (s *SimpleContext) TmpValue(key interface{}) (result interface{}, ok bool) {
	result, ok = s.conditionResults[key]
	return
}

// SetTmpValue cache a result of evaluted condition
func (s *SimpleContext) SetTmpValue(key interface{}, result interface{}) (err error) {
	_, ok := s.conditionResults[key]
	if ok {
		err = ExistsError("Condition result exists")
		return
	}
	s.conditionResults[key] = result
	return
}

// Value get value from context, only string key supported now
func (s *SimpleContext) Value(key interface{}) interface{} {
	val, _ := s.values[key.(string)]
	return val
}

// New create a SimpleContext
func New() *SimpleContext {
	return &SimpleContext{conditionResults: make(map[interface{}]interface{})}
}

// NewContextWithValues create a SimpleContext with values
func NewContextWithValues(values map[string]interface{}) *SimpleContext {
	return &SimpleContext{conditionResults: make(map[interface{}]interface{}), values: values}
}
