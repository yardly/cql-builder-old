package lucene

import (
	"bytes"
	"strings"
)

// WildcardQuery struct
type WildcardQuery struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lucene to match Rule interface
func (v *WildcardQuery) Lucene() bool {
	return true
}

// Wildcard query
func Wildcard(ref string, val interface{}) Rule {
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: val,
	}
}

// AnyWildcard query
func AnyWildcard(ref string, val string) Rule {
	var b bytes.Buffer
	b.WriteRune('*')
	b.WriteString(val)
	b.WriteRune('*')
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}

// AutoWildcard query
func AutoWildcard(ref string, val string) Rule {
	var b bytes.Buffer
	var res = strings.Join(
		strings.Split(val, " "),
		"*",
	)
	if len(res) > 0 && []rune(res)[0] != '*' {
		b.WriteRune('*')
	}
	b.WriteString(res)
	if s := b.String(); len(s) > 0 && []rune(s)[len(s)-1] != '*' {
		b.WriteRune('*')
	}
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}

// PrefixWildcard query
func PrefixWildcard(ref string, val string) Rule {
	var b bytes.Buffer
	b.WriteString(val)
	b.WriteRune('*')
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}

// PrefixWildcardIf query
func PrefixWildcardIf(condition bool, ref string, val string) Rule {
	if condition {
		return PrefixWildcard(ref, val)
	}
	return nil
}

// SuffixWildcard query
func SuffixWildcard(ref string, val string) Rule {
	var b bytes.Buffer
	b.WriteRune('*')
	b.WriteString(val)
	return &WildcardQuery{
		Type:  "wildcard",
		Field: ref,
		Value: b.String(),
	}
}
