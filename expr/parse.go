package expr

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Value is an expression
// TODO: Consider to use ObjExpr to replace this type
type Value interface{}

type expJSONValues []interface{}

type jsonExpr struct {
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Values expJSONValues `json:"values"`
}

func loadJSONExpr(data map[string]interface{}) (*jsonExpr, error) {
	cmdType := ""
	cmdTypeData, ok := data["type"]
	if ok {
		cmdType = strings.ToLower(cmdTypeData.(string))
	}

	var exprValues expJSONValues
	exprValuesData, ok := data["values"]
	if ok {
		exprValues = expJSONValues(exprValuesData.([]interface{}))
	}
	jsonObj := jsonExpr{
		Name:   data["name"].(string),
		Type:   cmdType,
		Values: exprValues,
	}
	return &jsonObj, nil
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

// Parser is a parser to parse expression of condition
type Parser struct {
	cmd    Command
	ctxCmd *CtxCommand
}

var cmd = NewBuiltInCommand()
var ctxCmd = &CtxCommand{}

// NewParser create a default ExprParser
func NewParser() *Parser {
	return &Parser{cmd, ctxCmd}
}

// NewParserWithCommand create a ExprParser with customised command
func NewParserWithCommand(cmd Command) *Parser {
	return &Parser{cmd, ctxCmd}
}

func (p *Parser) parseValues(values []interface{}) ([]interface{}, error) {
	newValues := make([]interface{}, len(values))
	for i, value := range values {
		// refer: https://golang.org/pkg/encoding/json/#Unmarshal
		switch v := value.(type) {
		case string:
			newValues[i] = v
		case float64:
			newValues[i] = int(v)
		case bool:
			newValues[i] = v
		case map[string]interface{}:
			expValue, err := p.ParseJSONExprData(v)
			if err != nil {
				return nil, err
			}
			newValues[i] = expValue
		default:
			return nil, fmt.Errorf("unknonw type %T of value %v in values", value, value)
		}

	}
	return newValues, nil
}

type contextValue bool

var ctxValue = contextValue(true)

func (p *Parser) parseJSONExpr(jsonObj *jsonExpr) (Value, error) {
	m := method(jsonObj.Name)

	values, err := p.parseValues(jsonObj.Values)
	if err != nil {
		return nil, err
	}

	var exp ObjExp

	switch jsonObj.Type {
	case "", "cmd":
		if p.cmd.HasCtxCmdName(jsonObj.Name) {
			oldValues := values
			values = make([]interface{}, len(oldValues)+1)
			copy(values[1:], oldValues)
			values[0] = ctxValue
		}
		exp, err = m.NewObjExpFromObjMethod(p.cmd, values...)
	case "ctx":
		values, oldValues := make([]interface{}, len(values)+1), values
		copy(values[1:], oldValues)
		values[0] = ctxValue
		exp, err = m.NewObjExpFromObjMethod(p.ctxCmd, values...)
	case "obj":
		return nil, fmt.Errorf("not implemented error of type obj")
	default:
		return nil, fmt.Errorf("unknonw type : %v ", jsonObj.Type)
	}

	if err != nil {
		return nil, err
	}
	return exp, nil
}

// ParseJSONExprData parse a json object configuration of condition expression
// examples:
// {"type": "cmd", "name": "and", "values": []}
func (p *Parser) ParseJSONExprData(data map[string]interface{}) (Value, error) {
	jsonExp, err := loadJSONExpr(data)
	if err != nil {
		return nil, err
	}
	return p.parseJSONExpr(jsonExp)
}

func (p *Parser) parseJSONExprString(expJSONStr string) (Value, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(expJSONStr), &data)
	if err != nil {
		return nil, err
	}
	return p.ParseJSONExprData(data)
}
