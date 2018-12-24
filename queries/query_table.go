package query

type CQLTable interface {
	Insert(...interface{}) CQLQuery
	Update(...interface{}) CQLQuery
	Delete(...interface{}) CQLQuery
	Select(...interface{}) CQLQuery
}

type cqlTable struct {
	s string
}

func (v *cqlTable) action(n string, fs ...interface{}) CQLQuery {
	var fields []CQLField
	for _, f := range fs {
		if v, ok := f.(CQLField); ok && v != nil {
			fields = append(fields, v)
		}
	}
	return &cqlQuery{
		action: n,
		table:  v.s,
		fields: fields,
	}
}

func (v *cqlTable) Insert(fs ...interface{}) CQLQuery {
	return v.action("INSERT", fs...)
}

func (v *cqlTable) Update(fs ...interface{}) CQLQuery {
	return v.action("UPDATE", fs...)
}

func (v *cqlTable) Delete(fs ...interface{}) CQLQuery {
	return v.action("DELETE", fs...)
}

func (v *cqlTable) Select(fs ...interface{}) CQLQuery {
	return v.action("SELECT", fs...)
}

func Table(s string) CQLTable {
	return &cqlTable{s}
}
