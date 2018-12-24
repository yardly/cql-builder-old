package query

type parameters struct {
	Table      string
	Fields     []CQLField
	Where      []CQLWhere
	IfNotExist bool
	Limit      int
}
