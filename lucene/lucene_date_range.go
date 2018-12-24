package lucene

import "time"

// DateRangeQuery struct
type DateRangeQuery struct {
	Type      string `json:"type"`
	Field     string `json:"field"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	Operation string `json:"operation,omitempty"`
}

// Lucene to match Rule interface
func (v *DateRangeQuery) Lucene() bool {
	return true
}

// BeforeDate query
func BeforeDate(field string, date time.Time, operations ...string) Rule {
	var to = date.Format("2006-01-02")
	return BeforeDateString(field, to, operations...)
}

// BeforeDateString query
func BeforeDateString(field, to string, operations ...string) Rule {
	var operation = "intersects"
	for _, v := range operations {
		operation = v
		break
	}
	return &DateRangeQuery{
		Type:      "date_range",
		Field:     field,
		To:        to,
		Operation: operation,
	}
}

// AfterDate query
func AfterDate(field string, date time.Time, operations ...string) Rule {
	var from = date.Format("2006-01-02")
	return AfterDateString(field, from, operations...)
}

// AfterDateString query
func AfterDateString(field, from string, operations ...string) Rule {
	var operation = "intersects"
	for _, v := range operations {
		operation = v
		break
	}
	return &DateRangeQuery{
		Type:      "date_range",
		Field:     field,
		From:      from,
		Operation: operation,
	}
}

// BetweenDate query
func BetweenDate(field string, ft, tt time.Time, operations ...string) Rule {
	var (
		from = ft.Format("2006-01-02")
		to   = tt.Format("2006-01-02")
	)
	return BetweenDateString(field, from, to, operations...)
}

// BetweenDateString query
func BetweenDateString(field, from, to string, operations ...string) Rule {
	var operation = "intersects"
	for _, v := range operations {
		operation = v
		break
	}
	return &DateRangeQuery{
		Type:      "date_range",
		Field:     field,
		From:      from,
		To:        to,
		Operation: operation,
	}
}
