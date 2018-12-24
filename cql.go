package cql

import (
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/samuelngs/cql-builder/lucene"
	"github.com/samuelngs/cql-builder/queries"
)

type CQL interface {
	Batch(...interface{}) CQL
	Prepare(query.CQLQuery) CQL
	Unlogged() CQL
	Compile() (string, []interface{})
	Iter() *gocql.Iter
	Exec() error
}

type cql struct {
	s        *gocql.Session
	q        []query.CQLQuery
	unlogged bool
}

func (v *cql) Batch(qs ...interface{}) CQL {
	for _, q := range qs {
		if c, ok := q.(query.CQLQuery); ok && c != nil {
			v.q = append(v.q, c)
		}
	}
	return v
}

func (v *cql) Prepare(q query.CQLQuery) CQL {
	if q != nil {
		v.q = append(v.q, q)
	}
	return v
}

func (v *cql) Unlogged() CQL {
	v.unlogged = true
	return v
}

func (v *cql) Compile() (string, []interface{}) {
	var queries []string
	var parameters []interface{}
	if len(v.q) > 1 {
		if v.unlogged {
			queries = append(queries, "BEGIN UNLOGGED BATCH")
		} else {
			queries = append(queries, "BEGIN BATCH")
		}
	}
	for _, q := range v.q {
		s, v := q.Compile()
		queries = append(queries, s)
		parameters = append(parameters, v...)
	}
	if len(v.q) > 1 {
		queries = append(queries, "APPLY BATCH")
	}
	return strings.Join(queries, ""), parameters
}

func (v *cql) Iter() *gocql.Iter {
	query, parameters := v.Compile()
	if v.s != nil {
		return v.s.
			Query(query, parameters...).
			Consistency(gocql.LocalQuorum).
			Iter()
	}
	return nil
}

func (v *cql) Exec() error {
	query, parameters := v.Compile()
	if v.s != nil {
		return v.s.
			Query(query, parameters...).
			Consistency(gocql.LocalQuorum).
			Exec()
	}
	return nil
}

func Connect(s *gocql.Session) CQL {
	return &cql{s: s}
}

func Table(s string) query.CQLTable {
	return query.Table(s)
}

func Field(k string, v ...interface{}) query.CQLField {
	var value interface{}
	if len(v) > 0 {
		value = v[0]
	}
	return query.Field(k, value)
}

func Lucene(os ...lucene.Option) query.CQLWhere {
	opts := lucene.NewOptions(os...)
	return query.Where("eq", opts.Column, opts.Lucene)
}

func Token(v interface{}) query.CQLToken {
	return query.Token(v)
}

func Increment(n int) query.CQLCounterChange {
	var i int
	switch {
	case n > 0:
		i = n
	case n < 0:
		i = -n
	}
	return query.CounterChange(i)
}

func Decrement(n int) query.CQLCounterChange {
	var i int
	switch {
	case n > 0:
		i = -n
	case n < 0:
		i = n
	}
	return query.CounterChange(i)
}

func Eq(k, v interface{}) query.CQLWhere {
	return query.Where("eq", k, v)
}

func Lt(k, v interface{}) query.CQLWhere {
	return query.Where("lt", k, v)
}

func Lte(k, v interface{}) query.CQLWhere {
	return query.Where("lte", k, v)
}

func Gt(k, v interface{}) query.CQLWhere {
	return query.Where("gt", k, v)
}

func Gte(k, v interface{}) query.CQLWhere {
	return query.Where("gte", k, v)
}

func In(k, v interface{}) query.CQLWhere {
	return query.Where("in", k, v)
}

func If(condition bool, o interface{}, e ...interface{}) interface{} {
	if condition {
		return o
	}
	if len(e) > 0 {
		return e[0]
	}
	return nil
}

func FieldNonZero(k string, v ...interface{}) interface{} {
	var value interface{}
	if len(v) > 0 {
		value = v[0]
	}
	if value == nil {
		return nil
	}
	if o, ok := value.(int); ok && o == 0 {
		return nil
	}
	if o, ok := value.(int32); ok && o == 0 {
		return nil
	}
	if o, ok := value.(int64); ok && o == 0 {
		return nil
	}
	if o, ok := value.(uint32); ok && o == 0 {
		return nil
	}
	if o, ok := value.(uint64); ok && o == 0 {
		return nil
	}
	if o, ok := value.(float32); ok && o == 0.0 {
		return nil
	}
	if o, ok := value.(float64); ok && o == 0.0 {
		return nil
	}
	if o, ok := value.(string); ok && o == "" {
		return nil
	}
	if o, ok := value.(time.Time); ok && o.IsZero() {
		return nil
	}
	return query.Field(k, value)
}
