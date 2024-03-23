// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/earthquake"
	"assignment3/ent/type_eathquake"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Type_eathquake is the model entity for the Type_eathquake schema.
type Type_eathquake struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EarthquakeID holds the value of the "earthquake_id" field.
	EarthquakeID int `json:"earthquake_id,omitempty"`
	// TypeEathquake holds the value of the "type_eathquake" field.
	TypeEathquake string `json:"type_eathquake,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Type_eathquakeQuery when eager-loading is set.
	Edges        Type_eathquakeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// Type_eathquakeEdges holds the relations/edges for other nodes in the graph.
type Type_eathquakeEdges struct {
	// Earthquake holds the value of the earthquake edge.
	Earthquake *Earthquake `json:"earthquake,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EarthquakeOrErr returns the Earthquake value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Type_eathquakeEdges) EarthquakeOrErr() (*Earthquake, error) {
	if e.Earthquake != nil {
		return e.Earthquake, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: earthquake.Label}
	}
	return nil, &NotLoadedError{edge: "earthquake"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Type_eathquake) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case type_eathquake.FieldID, type_eathquake.FieldEarthquakeID:
			values[i] = new(sql.NullInt64)
		case type_eathquake.FieldTypeEathquake:
			values[i] = new(sql.NullString)
		case type_eathquake.FieldCreatedAt, type_eathquake.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Type_eathquake fields.
func (te *Type_eathquake) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case type_eathquake.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			te.ID = int(value.Int64)
		case type_eathquake.FieldEarthquakeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field earthquake_id", values[i])
			} else if value.Valid {
				te.EarthquakeID = int(value.Int64)
			}
		case type_eathquake.FieldTypeEathquake:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type_eathquake", values[i])
			} else if value.Valid {
				te.TypeEathquake = value.String
			}
		case type_eathquake.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				te.CreatedAt = value.Time
			}
		case type_eathquake.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				te.UpdatedAt = value.Time
			}
		default:
			te.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Type_eathquake.
// This includes values selected through modifiers, order, etc.
func (te *Type_eathquake) Value(name string) (ent.Value, error) {
	return te.selectValues.Get(name)
}

// QueryEarthquake queries the "earthquake" edge of the Type_eathquake entity.
func (te *Type_eathquake) QueryEarthquake() *EarthquakeQuery {
	return NewTypeEathquakeClient(te.config).QueryEarthquake(te)
}

// Update returns a builder for updating this Type_eathquake.
// Note that you need to call Type_eathquake.Unwrap() before calling this method if this Type_eathquake
// was returned from a transaction, and the transaction was committed or rolled back.
func (te *Type_eathquake) Update() *TypeEathquakeUpdateOne {
	return NewTypeEathquakeClient(te.config).UpdateOne(te)
}

// Unwrap unwraps the Type_eathquake entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (te *Type_eathquake) Unwrap() *Type_eathquake {
	_tx, ok := te.config.driver.(*txDriver)
	if !ok {
		panic("ent: Type_eathquake is not a transactional entity")
	}
	te.config.driver = _tx.drv
	return te
}

// String implements the fmt.Stringer.
func (te *Type_eathquake) String() string {
	var builder strings.Builder
	builder.WriteString("Type_eathquake(")
	builder.WriteString(fmt.Sprintf("id=%v, ", te.ID))
	builder.WriteString("earthquake_id=")
	builder.WriteString(fmt.Sprintf("%v", te.EarthquakeID))
	builder.WriteString(", ")
	builder.WriteString("type_eathquake=")
	builder.WriteString(te.TypeEathquake)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(te.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(te.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Type_eathquakes is a parsable slice of Type_eathquake.
type Type_eathquakes []*Type_eathquake
