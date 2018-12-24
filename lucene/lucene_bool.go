package lucene

// BooleanQuery struct
type BooleanQuery struct {
	Type   string `json:"type"`
	Must   []Rule `json:"must,omitempty"`
	Should []Rule `json:"should,omitempty"`
	Not    []Rule `json:"not,omitempty"`
}

// Lucene to match Rule interface
func (v *BooleanQuery) Lucene() bool {
	return true
}

// BooleanMust query
func BooleanMust(rs ...interface{}) Rule {
	v := &BooleanQuery{
		Type: "boolean",
		Must: make([]Rule, 0),
	}
	for _, r := range rs {
		if o, ok := r.(Rule); ok && o != nil {
			v.Must = append(v.Must, o)
		}
	}
	return v
}

// BooleanShould query
func BooleanShould(rs ...interface{}) Rule {
	v := &BooleanQuery{
		Type:   "boolean",
		Should: make([]Rule, 0),
	}
	for _, r := range rs {
		if o, ok := r.(Rule); ok && o != nil {
			v.Should = append(v.Should, o)
		}
	}
	return v
}

// BooleanNot query
func BooleanNot(rs ...interface{}) Rule {
	v := &BooleanQuery{
		Type: "boolean",
		Not:  make([]Rule, 0),
	}
	for _, r := range rs {
		if o, ok := r.(Rule); ok && o != nil {
			v.Not = append(v.Not, o)
		}
	}
	return v
}
