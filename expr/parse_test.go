package expr

import (
	"testing"

	"github.com/GitterRemote/gcond/condition"
)

func NewContextWithValues(values map[string]interface{}) Context {
	return condition.NewContextWithValues(values)
}

func TestParseJSONObjExp(t *testing.T) {
	var tests = []struct {
		in       *jsonExpr
		expected bool
	}{
		{
			&jsonExpr{
				Type:   "cmd",
				Name:   "And",
				Values: []interface{}{true, true},
			},
			true,
		},
		{
			&jsonExpr{
				Type:   "cmd",
				Name:   "And",
				Values: []interface{}{false, true},
			},
			false,
		},
		{
			&jsonExpr{
				Type: "cmd",
				Name: "And",
				Values: []interface{}{
					true,
					map[string]interface{}{
						"type":   "cmd",
						"name":   "And",
						"values": []interface{}{true, true},
					},
				},
			},
			true,
		},
		{
			&jsonExpr{
				Type: "cmd",
				Name: "And",
				Values: []interface{}{
					true,
					map[string]interface{}{
						"type":   "cmd",
						"name":   "And",
						"values": []interface{}{true, false},
					},
				},
			},
			false,
		},
	}

	parser := NewParser()
	for _, test := range tests {
		expr, err := parser.parseJSONExpr(test.in)
		if err != nil {
			t.Error("parse error:", err)
			return
		}
		result := expr.(ObjExp)(nil)
		if result.(bool) != test.expected {
			t.Error("error result")
			return
		}
	}
}

type testCommand struct {
	*BuiltInCommand
}

func (c *testCommand) TestAnd(ctx Context, other bool) bool {
	return ctx.Value("TestAnd").(bool) && other
}

func TestParseJSONExprData(t *testing.T) {
	parser := NewParser()
	cmd := &testCommand{NewBuiltInCommand("TestAnd")}
	definedParser := NewParserWithCommand(cmd)
	var tests = []struct {
		in       map[string]interface{}
		ctx      Context
		parser   *Parser
		expected bool
	}{
		{
			map[string]interface{}{
				"type":   "ctx",
				"name":   "TestContext",
				"values": []interface{}{},
			},
			NewContextWithValues(map[string]interface{}{"test": true}),
			parser,
			true,
		},
		{
			map[string]interface{}{
				"type":   "ctx",
				"name":   "TestContext",
				"values": []interface{}{},
			},
			NewContextWithValues(map[string]interface{}{"test": false}),
			parser,
			false,
		},
		{
			map[string]interface{}{
				"type": "cmd",
				"name": "And",
				"values": []interface{}{
					true,
					map[string]interface{}{
						"type": "ctx",
						"name": "TestContext",
					},
				},
			},
			NewContextWithValues(map[string]interface{}{"test": true}),
			parser,
			true,
		},
		{
			map[string]interface{}{
				"name":   "TestAnd",
				"values": []interface{}{true},
			},
			NewContextWithValues(map[string]interface{}{"TestAnd": true}),
			definedParser,
			true,
		},
		{
			map[string]interface{}{
				"name":   "TestAnd",
				"values": []interface{}{false},
			},
			NewContextWithValues(map[string]interface{}{"TestAnd": true}),
			definedParser,
			false,
		},
	}

	for _, test := range tests {
		expr, err := test.parser.ParseJSONExprData(test.in)
		if err != nil {
			t.Error("parse error:", err)
			return
		}
		result := expr.(ObjExp)(test.ctx)
		if result.(bool) != test.expected {
			t.Error("error result")
			return
		}
	}
}
