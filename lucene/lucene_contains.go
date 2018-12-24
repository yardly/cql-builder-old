package lucene

import "reflect"

// ContainsQuery struct
type ContainsQuery struct {
	Type   string        `json:"type"`
	Field  string        `json:"field"`
	Values []interface{} `json:"values"`
}

// Lucene to match Rule interface
func (v *ContainsQuery) Lucene() bool {
	return true
}

// Contains query
func Contains(ref string, values interface{}) Rule {
	s := reflect.ValueOf(values)
	if s.Kind() != reflect.Slice {
		return nil
	}
	if s.Len() == 0 {
		return nil
	}
	arr := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		arr[i] = s.Index(i).Interface()
	}
	return &ContainsQuery{
		Type:   "contains",
		Field:  ref,
		Values: arr,
	}
}
