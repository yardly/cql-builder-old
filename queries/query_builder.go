package query

import (
	"bytes"
	"text/template"
)

type CQLQuery interface {
	Type() string
	Table() string
	Fields() []CQLField
	Where(...interface{}) CQLQuery
	IfNotExist() CQLQuery
	Limit(int) CQLQuery
	Compile() (string, []interface{})
}

type cqlQuery struct {
	action     string
	table      string
	fields     []CQLField
	where      []CQLWhere
	ifNotExist bool
	limit      int
}

func (v *cqlQuery) Type() string {
	return v.action
}

func (v *cqlQuery) Table() string {
	return v.table
}

func (v *cqlQuery) Fields() []CQLField {
	return v.fields
}

func (v *cqlQuery) Where(fs ...interface{}) CQLQuery {
	for _, f := range fs {
		if o, ok := f.(CQLWhere); ok && o != nil {
			v.where = append(v.where, o)
		}
	}
	return v
}

func (v *cqlQuery) IfNotExist() CQLQuery {
	v.ifNotExist = true
	return v
}

func (v *cqlQuery) Limit(l int) CQLQuery {
	v.limit = l
	return v
}

func (v *cqlQuery) Compile() (string, []interface{}) {
	var tmpl string
	var bufs bytes.Buffer
	var vals []interface{}
	switch v.Type() {
	case "INSERT":
		tmpl = InsertTemplate
	case "UPDATE":
		tmpl = UpdateTemplate
	case "DELETE":
		tmpl = DeleteTemplate
	case "SELECT":
		tmpl = SelectTemplate
	}
	p := &parameters{
		Table:      v.table,
		Fields:     v.fields,
		Where:      v.where,
		IfNotExist: v.ifNotExist,
		Limit:      v.limit,
	}
	t, _ := template.
		New("CQL").
		Funcs(template.FuncMap{"add": add, "minus": minus}).
		Parse(tmpl)
	if err := t.Execute(&bufs, p); err != nil {
		return "", nil
	}
	for _, field := range v.fields {
		if v.Type() != "SELECT" && v.Type() != "DELETE" {
			vals = append(vals, field.Val())
		}
	}
	for _, where := range v.where {
		vals = append(vals, where.Val())
	}
	return string(bufs.Bytes()), vals
}
