package expr

import (
	"fmt"
	"reflect"
)

// ObjExp is an expression to evaluated to any interface{}
type ObjExp func(ctx Context) interface{}
type objMethodExp func(ctx Context) []reflect.Value

// NewObjExp create a object expression return obj as evaluted result, used in test
func NewObjExp(obj interface{}) ObjExp {
	return func(ctx Context) interface{} {
		return obj
	}
}

// evaluateValues used to evaluate expression in values by context
func evaluateValues(ctx Context, values []interface{}) (results []interface{}) {
	results = make([]interface{}, len(values))
	for i, value := range values {
		switch v := value.(type) {
		case ObjExp:
			results[i] = v(ctx)
		case contextValue:
			results[i] = ctx
		case reflect.Value:
			results[i] = v.Interface()
		default:
			results[i] = v
		}

	}
	return
}

// TODO: improve the evaluate
func evaluateValuesToReflectValue(ctx Context, values []interface{}) (results []reflect.Value) {
	args := evaluateValues(ctx, values)
	results = make([]reflect.Value, len(args))
	for i, v := range args {
		results[i] = reflect.ValueOf(v)
	}
	return
}

type method string

func (m *method) name() string {
	return string(*m)
}

// refer: https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func (m *method) newObjMethodExpFromObjExpMethod(exp ObjExp, values ...interface{}) (objMethodExp, error) {
	// TODO: check ObjExp's obj has the method in advance
	return func(ctx Context) []reflect.Value {
		inputs := evaluateValuesToReflectValue(ctx, values)
		return reflect.ValueOf(exp(ctx)).MethodByName(m.name()).Call(inputs)
	}, nil
}

func (m *method) newObjMethodExpFromObjMethod(obj interface{}, values ...interface{}) (objMethodExp, error) {
	method := reflect.ValueOf(obj).MethodByName(m.name())
	// TODO: check inputs number match the method args number
	if !method.IsValid() {
		return nil, fmt.Errorf("no method %s found on obj %v", m.name(), obj)
	}
	return func(ctx Context) []reflect.Value {
		inputs := evaluateValuesToReflectValue(ctx, values)
		return method.Call(inputs)
	}, nil
}

// Only One result value allow command(object method) return
func (m *method) newObjExpFromObjMethodExp(exp objMethodExp) (ObjExp, error) {
	return func(ctx Context) interface{} {
		reflectValues := exp(ctx)
		if len(reflectValues) != 1 {
			panic(fmt.Sprintf("objExp only allow one return value, but %d returned", len(reflectValues)))
		}
		return reflectValues[0].Interface()
	}, nil
}

// NewObjExpFromObjExpMethod new an expression from method of an ObjExp
func (m *method) NewObjExpFromObjExpMethod(exp ObjExp, values ...interface{}) (ObjExp, error) {
	objMethodExp, err := m.newObjMethodExpFromObjExpMethod(exp, values...)
	if err != nil {
		return nil, err
	}
	return m.newObjExpFromObjMethodExp(objMethodExp)
}

// NewObjExpFromObjMethod new an expression from method of an Obj
func (m *method) NewObjExpFromObjMethod(obj interface{}, values ...interface{}) (ObjExp, error) {
	objMethodExp, err := m.newObjMethodExpFromObjMethod(obj, values...)
	if err != nil {
		return nil, err
	}
	return m.newObjExpFromObjMethodExp(objMethodExp)
}
