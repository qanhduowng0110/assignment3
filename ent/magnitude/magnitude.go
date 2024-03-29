// Code generated by ent, DO NOT EDIT.

package magnitude

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the magnitude type in the database.
	Label = "magnitude"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMagnitudeValue holds the string denoting the magnitude_value field in the database.
	FieldMagnitudeValue = "magnitude_value"
	// FieldMagnitudeType holds the string denoting the magnitude_type field in the database.
	FieldMagnitudeType = "magnitude_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeEarthquakes holds the string denoting the earthquakes edge name in mutations.
	EdgeEarthquakes = "earthquakes"
	// Table holds the table name of the magnitude in the database.
	Table = "Magnitudes"
	// EarthquakesTable is the table that holds the earthquakes relation/edge.
	EarthquakesTable = "Earthquakes"
	// EarthquakesInverseTable is the table name for the Earthquake entity.
	// It exists in this package in order to avoid circular dependency with the "earthquake" package.
	EarthquakesInverseTable = "Earthquakes"
	// EarthquakesColumn is the table column denoting the earthquakes relation/edge.
	EarthquakesColumn = "magitude_id"
)

// Columns holds all SQL columns for magnitude fields.
var Columns = []string{
	FieldID,
	FieldMagnitudeValue,
	FieldMagnitudeType,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Magnitude queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMagnitudeValue orders the results by the magnitude_value field.
func ByMagnitudeValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMagnitudeValue, opts...).ToFunc()
}

// ByMagnitudeType orders the results by the magnitude_type field.
func ByMagnitudeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMagnitudeType, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByEarthquakesCount orders the results by earthquakes count.
func ByEarthquakesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEarthquakesStep(), opts...)
	}
}

// ByEarthquakes orders the results by earthquakes terms.
func ByEarthquakes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEarthquakesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEarthquakesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EarthquakesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EarthquakesTable, EarthquakesColumn),
	)
}
