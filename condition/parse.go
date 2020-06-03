package condition

import (
	"fmt"
)

type expJSONValues []interface{}

func stoi(values []interface{}) ([]int, error) {
	newValues := make([]int, len(values))
	for i, v := range values {
		newValues[i] = v.(int)
	}
	return newValues, nil
}

func stoa(values []interface{}) ([]string, error) {
	newValues := make([]string, len(values))
	for i, v := range values {
		newValues[i] = v.(string)
	}
	return newValues, nil
}

type jsonExpObj struct {
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Values expJSONValues `json:"values"`
}

// ParseError is a parse error
type ParseError string

func (e ParseError) Error() string {
	return string(e)
}

// ParseValueLengthError is a parse error when command required values number is not the same as passed in
type ParseValueLengthError struct {
	length int
}

func (e *ParseValueLengthError) Error() string {
	return fmt.Sprintf("need values number is %v", e.length)
}

func getExp(expType, name string) (interface{}, error) {
	switch expType {
	case "op":
		cmd, ok := operators[name]
		if !ok {
			return nil, ParseError(fmt.Sprintf("%s not exists", name))
		}
		return cmd, nil
	default:
		return nil, ParseError(fmt.Sprintf("%s type not exists", expType))
	}
}

// parseJSONObjExp parse a json object configuration of condition expression
// examples:
// {"type": "op", "name": "and", "values": []}
func parseJSONObjExp(jsonObj jsonExpObj) BoolExp {
	panic("")
}

func parseJSONExp(jsonExp string) BoolExp {
	panic("")
}
