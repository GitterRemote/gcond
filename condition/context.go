package condition

// SimpleContext implments get/set condition results
type SimpleContext struct {
	Context
	conditionResults map[int]bool
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

// NewContext create a SimpleContext
func NewContext() Context {
	return &SimpleContext{conditionResults: make(map[int]bool)}
}
