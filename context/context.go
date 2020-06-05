package context

// SimpleContext implments get/set condition results
type SimpleContext struct {
	conditionResults map[int]bool
	values           map[string]interface{}
}

// ExistsError is a err occurred when setConditionResult
type ExistsError string

func (e ExistsError) Error() string {
	return string(e)
}

// GetConditionResult get result of condition pre-evaluted
func (s *SimpleContext) GetConditionResult(conditionID int) (result bool, ok bool) {
	result, ok = s.conditionResults[conditionID]
	return
}

// SetConditionResult cache a result of evaluted condition
func (s *SimpleContext) SetConditionResult(conditionID int, result bool) (err error) {
	_, ok := s.conditionResults[conditionID]
	if ok {
		err = ExistsError("Condition result exists")
		return
	}
	s.conditionResults[conditionID] = result
	return
}

// Value get value from context, only string key supported now
func (s *SimpleContext) Value(key interface{}) interface{} {
	val, _ := s.values[key.(string)]
	return val
}

// NewContext create a SimpleContext
func NewContext() *SimpleContext {
	return &SimpleContext{conditionResults: make(map[int]bool)}
}

// NewContextWithValues create a SimpleContext with values
func NewContextWithValues(values map[string]interface{}) *SimpleContext {
	return &SimpleContext{conditionResults: make(map[int]bool), values: values}
}
