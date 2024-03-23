// Code generated by ent, DO NOT EDIT.

package type_eathquake

import (
	"assignment3/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLTE(FieldID, id))
}

// EarthquakeID applies equality check predicate on the "earthquake_id" field. It's identical to EarthquakeIDEQ.
func EarthquakeID(v int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldEarthquakeID, v))
}

// TypeEathquake applies equality check predicate on the "type_eathquake" field. It's identical to TypeEathquakeEQ.
func TypeEathquake(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldTypeEathquake, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldUpdatedAt, v))
}

// EarthquakeIDEQ applies the EQ predicate on the "earthquake_id" field.
func EarthquakeIDEQ(v int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldEarthquakeID, v))
}

// EarthquakeIDNEQ applies the NEQ predicate on the "earthquake_id" field.
func EarthquakeIDNEQ(v int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNEQ(FieldEarthquakeID, v))
}

// EarthquakeIDIn applies the In predicate on the "earthquake_id" field.
func EarthquakeIDIn(vs ...int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIn(FieldEarthquakeID, vs...))
}

// EarthquakeIDNotIn applies the NotIn predicate on the "earthquake_id" field.
func EarthquakeIDNotIn(vs ...int) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotIn(FieldEarthquakeID, vs...))
}

// EarthquakeIDIsNil applies the IsNil predicate on the "earthquake_id" field.
func EarthquakeIDIsNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIsNull(FieldEarthquakeID))
}

// EarthquakeIDNotNil applies the NotNil predicate on the "earthquake_id" field.
func EarthquakeIDNotNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotNull(FieldEarthquakeID))
}

// TypeEathquakeEQ applies the EQ predicate on the "type_eathquake" field.
func TypeEathquakeEQ(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldTypeEathquake, v))
}

// TypeEathquakeNEQ applies the NEQ predicate on the "type_eathquake" field.
func TypeEathquakeNEQ(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNEQ(FieldTypeEathquake, v))
}

// TypeEathquakeIn applies the In predicate on the "type_eathquake" field.
func TypeEathquakeIn(vs ...string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIn(FieldTypeEathquake, vs...))
}

// TypeEathquakeNotIn applies the NotIn predicate on the "type_eathquake" field.
func TypeEathquakeNotIn(vs ...string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotIn(FieldTypeEathquake, vs...))
}

// TypeEathquakeGT applies the GT predicate on the "type_eathquake" field.
func TypeEathquakeGT(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGT(FieldTypeEathquake, v))
}

// TypeEathquakeGTE applies the GTE predicate on the "type_eathquake" field.
func TypeEathquakeGTE(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGTE(FieldTypeEathquake, v))
}

// TypeEathquakeLT applies the LT predicate on the "type_eathquake" field.
func TypeEathquakeLT(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLT(FieldTypeEathquake, v))
}

// TypeEathquakeLTE applies the LTE predicate on the "type_eathquake" field.
func TypeEathquakeLTE(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLTE(FieldTypeEathquake, v))
}

// TypeEathquakeContains applies the Contains predicate on the "type_eathquake" field.
func TypeEathquakeContains(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldContains(FieldTypeEathquake, v))
}

// TypeEathquakeHasPrefix applies the HasPrefix predicate on the "type_eathquake" field.
func TypeEathquakeHasPrefix(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldHasPrefix(FieldTypeEathquake, v))
}

// TypeEathquakeHasSuffix applies the HasSuffix predicate on the "type_eathquake" field.
func TypeEathquakeHasSuffix(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldHasSuffix(FieldTypeEathquake, v))
}

// TypeEathquakeIsNil applies the IsNil predicate on the "type_eathquake" field.
func TypeEathquakeIsNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIsNull(FieldTypeEathquake))
}

// TypeEathquakeNotNil applies the NotNil predicate on the "type_eathquake" field.
func TypeEathquakeNotNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotNull(FieldTypeEathquake))
}

// TypeEathquakeEqualFold applies the EqualFold predicate on the "type_eathquake" field.
func TypeEathquakeEqualFold(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEqualFold(FieldTypeEathquake, v))
}

// TypeEathquakeContainsFold applies the ContainsFold predicate on the "type_eathquake" field.
func TypeEathquakeContainsFold(v string) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldContainsFold(FieldTypeEathquake, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.FieldNotNull(FieldUpdatedAt))
}

// HasEarthquake applies the HasEdge predicate on the "earthquake" edge.
func HasEarthquake() predicate.Type_eathquake {
	return predicate.Type_eathquake(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EarthquakeTable, EarthquakeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEarthquakeWith applies the HasEdge predicate on the "earthquake" edge with a given conditions (other predicates).
func HasEarthquakeWith(preds ...predicate.Earthquake) predicate.Type_eathquake {
	return predicate.Type_eathquake(func(s *sql.Selector) {
		step := newEarthquakeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Type_eathquake) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Type_eathquake) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Type_eathquake) predicate.Type_eathquake {
	return predicate.Type_eathquake(sql.NotPredicates(p))
}