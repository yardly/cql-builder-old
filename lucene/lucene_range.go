package lucene

// RangeQuery struct
type RangeQuery struct {
	Type         string      `json:"type"`
	Field        string      `json:"field"`
	Lower        interface{} `json:"lower,omitempty"`
	Upper        interface{} `json:"upper,omitempty"`
	IncludeLower bool        `json:"include_lower,omitempty"`
	IncludeUpper bool        `json:"include_upper,omitempty"`
}

// Lucene to match Rule interface
func (v *RangeQuery) Lucene() bool {
	return true
}

// RangeAll query
func RangeAll(ref string, lower, upper interface{}, incl, incu bool) Rule {
	return &RangeQuery{
		Type:         "range",
		Field:        ref,
		Lower:        lower,
		Upper:        upper,
		IncludeLower: incl,
		IncludeUpper: incu,
	}
}

// RangeLower query
func RangeLower(ref string, lower interface{}, incl ...bool) Rule {
	var b bool
	for _, v := range incl {
		b = v
		break
	}
	return &RangeQuery{
		Type:         "range",
		Field:        ref,
		Lower:        lower,
		IncludeLower: b,
	}
}

// RangeUpper query
func RangeUpper(ref string, upper interface{}, incu ...bool) Rule {
	var b bool
	for _, v := range incu {
		b = v
		break
	}
	return &RangeQuery{
		Type:         "range",
		Field:        ref,
		Upper:        upper,
		IncludeUpper: b,
	}
}
