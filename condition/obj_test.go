package condition

import (
	"reflect"
	"testing"
	"time"
)

func TestEvaluateValues(t *testing.T) {
	values := make([]interface{}, 2)
	values[0] = 1
	values[1] = reflect.ValueOf(1)
	results := evaluateValues(nil, values)
	if v, ok := results[1].(int); !ok {
		t.Error("evaluate error, reflect value unpack fail", results, v)
		return
	}
	if v, ok := values[1].(int); ok {
		t.Error("evaluate error, values has been modified", results, v)
		return
	}
}

func TestNewObjMethodExpFromObjExpMethod(t *testing.T) {
	obj := time.Now()
	m := method("Weekday")
	expr, err := m.newObjMethodExpFromObjExpMethod(NewObjExp(obj))
	if err != nil {
		t.Errorf("NewExpFromObjMethod error %v", err)
		return
	}
	values := expr(nil)
	if len(values) != 1 {
		t.Errorf("values error of NewExpFromObjMethod %v", values)
		return
	}
	if v := values[0].Int(); v > 7 {
		t.Errorf("value error of NewExpFromObjMethod %v", values)
		return
	}
}

func TestNewObjMethodExpFromObjMethod(t *testing.T) {
	obj := time.Now()

	invalidMethod := method("Weekda")
	expr, err := invalidMethod.newObjMethodExpFromObjMethod(obj)
	if err == nil {
		t.Errorf("new expression error: invalidMethod create suc: %v", err)
	}

	m := method("Weekday")
	expr, err = m.newObjMethodExpFromObjMethod(obj)
	if err != nil {
		t.Errorf("new expression error: %v", err)
		return
	}
	values := expr(nil)
	if len(values) != 1 {
		t.Errorf("return values error %v", values)
		return
	}
	if v := values[0].Int(); v > 7 {
		t.Errorf("return value error %v", values)
		return
	}
}
