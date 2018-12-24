package query

import (
	"fmt"
)

type CQLToken interface {
	Token(int) string
	Value() interface{}
}

type cqlToken struct {
	val interface{}
}

func (v *cqlToken) Token(i int) string {
	if i == 0 {
		return fmt.Sprintf("token(%v)", v.val)
	}
	return "token(?)"
}

func (v *cqlToken) Value() interface{} {
	return v.val
}

func Token(val interface{}) CQLToken {
	return &cqlToken{val}
}
