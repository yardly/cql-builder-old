package lucene

// MatchRule struct
type MatchRule struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lucene to match Rule interface
func (v *MatchRule) Lucene() bool {
	return true
}

// Match query
func Match(ref string, val interface{}) Rule {
	return &MatchRule{
		Type:  "match",
		Field: ref,
		Value: val,
	}
}
