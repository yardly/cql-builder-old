package query

// CQLCounterChange interface
type CQLCounterChange interface {
	Difference() int
	Increment() bool
}

type cqlCounterChange struct {
	diff      int
	increment bool
}

func (v *cqlCounterChange) Difference() int {
	return v.diff
}

func (v *cqlCounterChange) Increment() bool {
	return v.increment
}

// CounterChange instance
func CounterChange(diff int) CQLCounterChange {
	var n int
	var i bool
	if diff > 0 {
		i = true
		n = diff
	}
	if diff < 0 {
		n = -diff
	}
	return &cqlCounterChange{n, i}
}
