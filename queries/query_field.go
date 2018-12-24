package query

type CQLField interface {
	Key() string
	Val() interface{}
	IsCounter() bool
	IsIncrement() bool
}

type cqlField struct {
	key string
	val interface{}
}

func (v *cqlField) Key() string {
	return v.key
}

func (v *cqlField) Val() interface{} {
	switch o := v.val.(type) {
	case CQLCounterChange:
		return o.Difference()
	default:
		return v.val
	}
}

func (v *cqlField) IsCounter() bool {
	switch v.val.(type) {
	case CQLCounterChange:
		return true
	default:
		return false
	}
}

func (v *cqlField) IsIncrement() bool {
	switch o := v.val.(type) {
	case CQLCounterChange:
		return o.Increment()
	default:
		return false
	}
}

// Field creates new CQLField
func Field(key string, val interface{}) CQLField {
	return &cqlField{key, val}
}
