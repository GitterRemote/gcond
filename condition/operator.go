package condition

import (
	"fmt"
	"sort"
)

var operators = map[string]interface{}{}

type operator struct{}

func (o *operator) And(values ...bool) bool {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}

var op = &operator{}

type conditionBasedOp interface {
	newExpFromBoolExps(boolExps ...BoolExp) (BoolExp, error)
}

type conditionBasedExpImp struct {
	conditionBasedOp
}

func (o *conditionBasedExpImp) NewExpFromValues(values []interface{}) (BoolExp, error) {
	boolExps := make([]BoolExp, len(values))
	var e BoolExp
	for i, vInterface := range values {
		switch v := vInterface.(type) {
		case BoolExp:
			e = v
		default:
			return nil, ParseError(fmt.Sprintf("unknown type %v", v))
		}
		boolExps[i] = e
	}
	return o.newExpFromBoolExps(boolExps...)
}

type andOperator struct {
	conditionBasedExpImp
}

func (o *andOperator) newExpFromBoolExps(boolExps ...BoolExp) (BoolExp, error) {
	return func(ctx Context) bool {
		for _, exp := range boolExps {
			rv := exp(ctx)
			if !rv {
				return false
			}
		}
		return true
	}, nil
}

type orOperator struct {
	conditionBasedExpImp
}

// newOrExp create an Or expression based on conditions
func (o *orOperator) newExpFromBoolExps(boolExps ...BoolExp) (BoolExp, error) {
	return func(ctx Context) bool {
		for _, exp := range boolExps {
			rv := exp(ctx)
			if rv {
				return true
			}
		}
		return false
	}, nil
}

type notOperator struct {
	conditionBasedExpImp
}

func (o *notOperator) newExpFromBoolExps(boolExps ...BoolExp) (BoolExp, error) {
	if len(boolExps) != 1 {
		return nil, fmt.Errorf("conditions numer error")
	}
	exp := boolExps[0]
	return func(ctx Context) bool {
		return !exp(ctx)
	}, nil
}

type equalOperator struct {
	operator
}

// newFromString create a string equal compare bool expression
func (o *equalOperator) newFromString(a string, b string) BoolExp {
	return func(ctx Context) bool {
		return a == b
	}
}

func (o *equalOperator) newFromStringExp(a stringExp, b string) BoolExp {
	return func(ctx Context) bool {
		return a(ctx) == b
	}
}

func (o *equalOperator) NewExpFromValues(values []interface{}) (BoolExp, error) {
	if len(values) != 2 {
		return nil, fmt.Errorf("jsonObj.Values len not 2")
	}
	a := values[0]
	b := values[1]
	switch v := a.(type) {
	case string:
		return o.newFromString(v, b.(string)), nil
	case stringExp:
		return o.newFromStringExp(v, b.(string)), nil
	default:
		return nil, fmt.Errorf("equalOperator can't recognize type %v", v)
	}
}

type inOperator struct {
	operator
}

// newInExp create a string contains bool expression
func (o *inOperator) newFromString(item string, items []string) BoolExp {
	//refer: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
	sort.Strings(items)
	return func(ctx Context) bool {
		i := sort.SearchStrings(items, item)
		return i < len(items) && items[i] == item
	}
}

func (o *inOperator) newFromStringExp(itemExp stringExp, items []string) BoolExp {
	//refer: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
	sort.Strings(items)
	return func(ctx Context) bool {
		item := itemExp(ctx)
		i := sort.SearchStrings(items, item)
		return i < len(items) && items[i] == item
	}
}

func (o *inOperator) NewExpFromValues(values []interface{}) (BoolExp, error) {
	if len(values) < 2 {
		return nil, fmt.Errorf("inOperator must have jsonObj.Values equal or more than 2")
	}
	first := values[0]
	items, err := stoa(values[1:])
	if err != nil {
		return nil, err
	}
	switch v := first.(type) {
	case string:
		return o.newFromString(v, items), nil
	case stringExp:
		return o.newFromStringExp(v, items), nil
	default:
		return nil, fmt.Errorf("equalOperator can't recognize type %v", v)
	}
}
