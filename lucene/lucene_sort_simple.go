package lucene

// SortSimple struct
type SortSimple struct {
	Type    string `json:"type"`
	Field   string `json:"field"`
	Reverse bool   `json:"reverse"`
}

// Lucene to match Rule interface
func (v *SortSimple) Lucene() bool {
	return true
}

// ByField query
func ByField(ref string, reverse ...bool) Rule {
	var o bool
	for _, b := range reverse {
		o = b
		break
	}
	return &SortSimple{
		Type:    "simple",
		Field:   ref,
		Reverse: o,
	}
}
