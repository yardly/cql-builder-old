package query

import "fmt"

type CQLWhere interface {
	IsEq() bool
	IsLt() bool
	IsLte() bool
	IsGt() bool
	IsGte() bool
	IsIn() bool
	Type() string
	Key() string
	Val() interface{}
	Parameter() string
}

type cqlWhere struct {
	typ string
	key string
	val interface{}
}

func (v *cqlWhere) Type() string {
	return v.typ
}

func (v *cqlWhere) IsEq() bool {
	return v.typ == "eq"
}

func (v *cqlWhere) IsLt() bool {
	return v.typ == "lt"
}

func (v *cqlWhere) IsLte() bool {
	return v.typ == "lte"
}

func (v *cqlWhere) IsGt() bool {
	return v.typ == "gt"
}

func (v *cqlWhere) IsGte() bool {
	return v.typ == "gte"
}

func (v *cqlWhere) IsIn() bool {
	return v.typ == "in"
}

func (v *cqlWhere) Key() string {
	return v.key
}

func (v *cqlWhere) Val() interface{} {
	switch o := v.val.(type) {
	case CQLToken:
		return o.Value()
	default:
		return v.val
	}
}

func (v *cqlWhere) Parameter() string {
	switch o := v.val.(type) {
	case CQLToken:
		return o.Token(1)
	default:
		return "?"
	}
}

func Where(typ string, key, val interface{}) CQLWhere {
	switch v := key.(type) {
	case CQLToken:
		return &cqlWhere{typ, v.Token(0), val}
	case string:
		return &cqlWhere{typ, v, val}
	default:
		return &cqlWhere{typ, fmt.Sprintf("%v", key), val}
	}
}
