package lucene

// PhraseQuery struct
type PhraseQuery struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lucene to match Rule interface
func (v *PhraseQuery) Lucene() bool {
	return true
}

// Phrase query
func Phrase(ref string, val interface{}) Rule {
	return &PhraseQuery{
		Type:  "phrase",
		Field: ref,
		Value: val,
	}
}
