// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeIndication holds the string denoting the indication edge name in mutations.
	EdgeIndication = "indication"
	// EdgeIndicated holds the string denoting the indicated edge name in mutations.
	EdgeIndicated = "indicated"
	// Table holds the table name of the user in the database.
	Table = "users"
	// IndicationTable is the table that holds the indication relation/edge.
	IndicationTable = "users"
	// IndicationColumn is the table column denoting the indication relation/edge.
	IndicationColumn = "user_indicated"
	// IndicatedTable is the table that holds the indicated relation/edge.
	IndicatedTable = "users"
	// IndicatedColumn is the table column denoting the indicated relation/edge.
	IndicatedColumn = "user_indicated"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_indicated",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIndicationCount orders the results by indication count.
func ByIndicationCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIndicationStep(), opts...)
	}
}

// ByIndication orders the results by indication terms.
func ByIndication(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIndicationStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIndicatedField orders the results by indicated field.
func ByIndicatedField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIndicatedStep(), sql.OrderByField(field, opts...))
	}
}
func newIndicationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, IndicationTable, IndicationColumn),
	)
}
func newIndicatedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, IndicatedTable, IndicatedColumn),
	)
}
