// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/earthquake"
	"assignment3/ent/predicate"
	"assignment3/ent/type_eathquake"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TypeEathquakeUpdate is the builder for updating Type_eathquake entities.
type TypeEathquakeUpdate struct {
	config
	hooks    []Hook
	mutation *TypeEathquakeMutation
}

// Where appends a list predicates to the TypeEathquakeUpdate builder.
func (teu *TypeEathquakeUpdate) Where(ps ...predicate.Type_eathquake) *TypeEathquakeUpdate {
	teu.mutation.Where(ps...)
	return teu
}

// SetEarthquakeID sets the "earthquake_id" field.
func (teu *TypeEathquakeUpdate) SetEarthquakeID(i int) *TypeEathquakeUpdate {
	teu.mutation.SetEarthquakeID(i)
	return teu
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (teu *TypeEathquakeUpdate) SetNillableEarthquakeID(i *int) *TypeEathquakeUpdate {
	if i != nil {
		teu.SetEarthquakeID(*i)
	}
	return teu
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (teu *TypeEathquakeUpdate) ClearEarthquakeID() *TypeEathquakeUpdate {
	teu.mutation.ClearEarthquakeID()
	return teu
}

// SetTypeEathquake sets the "type_eathquake" field.
func (teu *TypeEathquakeUpdate) SetTypeEathquake(s string) *TypeEathquakeUpdate {
	teu.mutation.SetTypeEathquake(s)
	return teu
}

// SetNillableTypeEathquake sets the "type_eathquake" field if the given value is not nil.
func (teu *TypeEathquakeUpdate) SetNillableTypeEathquake(s *string) *TypeEathquakeUpdate {
	if s != nil {
		teu.SetTypeEathquake(*s)
	}
	return teu
}

// ClearTypeEathquake clears the value of the "type_eathquake" field.
func (teu *TypeEathquakeUpdate) ClearTypeEathquake() *TypeEathquakeUpdate {
	teu.mutation.ClearTypeEathquake()
	return teu
}

// SetCreatedAt sets the "created_at" field.
func (teu *TypeEathquakeUpdate) SetCreatedAt(t time.Time) *TypeEathquakeUpdate {
	teu.mutation.SetCreatedAt(t)
	return teu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (teu *TypeEathquakeUpdate) SetNillableCreatedAt(t *time.Time) *TypeEathquakeUpdate {
	if t != nil {
		teu.SetCreatedAt(*t)
	}
	return teu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (teu *TypeEathquakeUpdate) ClearCreatedAt() *TypeEathquakeUpdate {
	teu.mutation.ClearCreatedAt()
	return teu
}

// SetUpdatedAt sets the "updated_at" field.
func (teu *TypeEathquakeUpdate) SetUpdatedAt(t time.Time) *TypeEathquakeUpdate {
	teu.mutation.SetUpdatedAt(t)
	return teu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (teu *TypeEathquakeUpdate) SetNillableUpdatedAt(t *time.Time) *TypeEathquakeUpdate {
	if t != nil {
		teu.SetUpdatedAt(*t)
	}
	return teu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (teu *TypeEathquakeUpdate) ClearUpdatedAt() *TypeEathquakeUpdate {
	teu.mutation.ClearUpdatedAt()
	return teu
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (teu *TypeEathquakeUpdate) SetEarthquake(e *Earthquake) *TypeEathquakeUpdate {
	return teu.SetEarthquakeID(e.ID)
}

// Mutation returns the TypeEathquakeMutation object of the builder.
func (teu *TypeEathquakeUpdate) Mutation() *TypeEathquakeMutation {
	return teu.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (teu *TypeEathquakeUpdate) ClearEarthquake() *TypeEathquakeUpdate {
	teu.mutation.ClearEarthquake()
	return teu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (teu *TypeEathquakeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, teu.sqlSave, teu.mutation, teu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (teu *TypeEathquakeUpdate) SaveX(ctx context.Context) int {
	affected, err := teu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (teu *TypeEathquakeUpdate) Exec(ctx context.Context) error {
	_, err := teu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teu *TypeEathquakeUpdate) ExecX(ctx context.Context) {
	if err := teu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (teu *TypeEathquakeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(type_eathquake.Table, type_eathquake.Columns, sqlgraph.NewFieldSpec(type_eathquake.FieldID, field.TypeInt))
	if ps := teu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teu.mutation.TypeEathquake(); ok {
		_spec.SetField(type_eathquake.FieldTypeEathquake, field.TypeString, value)
	}
	if teu.mutation.TypeEathquakeCleared() {
		_spec.ClearField(type_eathquake.FieldTypeEathquake, field.TypeString)
	}
	if value, ok := teu.mutation.CreatedAt(); ok {
		_spec.SetField(type_eathquake.FieldCreatedAt, field.TypeTime, value)
	}
	if teu.mutation.CreatedAtCleared() {
		_spec.ClearField(type_eathquake.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := teu.mutation.UpdatedAt(); ok {
		_spec.SetField(type_eathquake.FieldUpdatedAt, field.TypeTime, value)
	}
	if teu.mutation.UpdatedAtCleared() {
		_spec.ClearField(type_eathquake.FieldUpdatedAt, field.TypeTime)
	}
	if teu.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   type_eathquake.EarthquakeTable,
			Columns: []string{type_eathquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := teu.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   type_eathquake.EarthquakeTable,
			Columns: []string{type_eathquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, teu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{type_eathquake.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	teu.mutation.done = true
	return n, nil
}

// TypeEathquakeUpdateOne is the builder for updating a single Type_eathquake entity.
type TypeEathquakeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TypeEathquakeMutation
}

// SetEarthquakeID sets the "earthquake_id" field.
func (teuo *TypeEathquakeUpdateOne) SetEarthquakeID(i int) *TypeEathquakeUpdateOne {
	teuo.mutation.SetEarthquakeID(i)
	return teuo
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (teuo *TypeEathquakeUpdateOne) SetNillableEarthquakeID(i *int) *TypeEathquakeUpdateOne {
	if i != nil {
		teuo.SetEarthquakeID(*i)
	}
	return teuo
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (teuo *TypeEathquakeUpdateOne) ClearEarthquakeID() *TypeEathquakeUpdateOne {
	teuo.mutation.ClearEarthquakeID()
	return teuo
}

// SetTypeEathquake sets the "type_eathquake" field.
func (teuo *TypeEathquakeUpdateOne) SetTypeEathquake(s string) *TypeEathquakeUpdateOne {
	teuo.mutation.SetTypeEathquake(s)
	return teuo
}

// SetNillableTypeEathquake sets the "type_eathquake" field if the given value is not nil.
func (teuo *TypeEathquakeUpdateOne) SetNillableTypeEathquake(s *string) *TypeEathquakeUpdateOne {
	if s != nil {
		teuo.SetTypeEathquake(*s)
	}
	return teuo
}

// ClearTypeEathquake clears the value of the "type_eathquake" field.
func (teuo *TypeEathquakeUpdateOne) ClearTypeEathquake() *TypeEathquakeUpdateOne {
	teuo.mutation.ClearTypeEathquake()
	return teuo
}

// SetCreatedAt sets the "created_at" field.
func (teuo *TypeEathquakeUpdateOne) SetCreatedAt(t time.Time) *TypeEathquakeUpdateOne {
	teuo.mutation.SetCreatedAt(t)
	return teuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (teuo *TypeEathquakeUpdateOne) SetNillableCreatedAt(t *time.Time) *TypeEathquakeUpdateOne {
	if t != nil {
		teuo.SetCreatedAt(*t)
	}
	return teuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (teuo *TypeEathquakeUpdateOne) ClearCreatedAt() *TypeEathquakeUpdateOne {
	teuo.mutation.ClearCreatedAt()
	return teuo
}

// SetUpdatedAt sets the "updated_at" field.
func (teuo *TypeEathquakeUpdateOne) SetUpdatedAt(t time.Time) *TypeEathquakeUpdateOne {
	teuo.mutation.SetUpdatedAt(t)
	return teuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (teuo *TypeEathquakeUpdateOne) SetNillableUpdatedAt(t *time.Time) *TypeEathquakeUpdateOne {
	if t != nil {
		teuo.SetUpdatedAt(*t)
	}
	return teuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (teuo *TypeEathquakeUpdateOne) ClearUpdatedAt() *TypeEathquakeUpdateOne {
	teuo.mutation.ClearUpdatedAt()
	return teuo
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (teuo *TypeEathquakeUpdateOne) SetEarthquake(e *Earthquake) *TypeEathquakeUpdateOne {
	return teuo.SetEarthquakeID(e.ID)
}

// Mutation returns the TypeEathquakeMutation object of the builder.
func (teuo *TypeEathquakeUpdateOne) Mutation() *TypeEathquakeMutation {
	return teuo.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (teuo *TypeEathquakeUpdateOne) ClearEarthquake() *TypeEathquakeUpdateOne {
	teuo.mutation.ClearEarthquake()
	return teuo
}

// Where appends a list predicates to the TypeEathquakeUpdate builder.
func (teuo *TypeEathquakeUpdateOne) Where(ps ...predicate.Type_eathquake) *TypeEathquakeUpdateOne {
	teuo.mutation.Where(ps...)
	return teuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (teuo *TypeEathquakeUpdateOne) Select(field string, fields ...string) *TypeEathquakeUpdateOne {
	teuo.fields = append([]string{field}, fields...)
	return teuo
}

// Save executes the query and returns the updated Type_eathquake entity.
func (teuo *TypeEathquakeUpdateOne) Save(ctx context.Context) (*Type_eathquake, error) {
	return withHooks(ctx, teuo.sqlSave, teuo.mutation, teuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (teuo *TypeEathquakeUpdateOne) SaveX(ctx context.Context) *Type_eathquake {
	node, err := teuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (teuo *TypeEathquakeUpdateOne) Exec(ctx context.Context) error {
	_, err := teuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teuo *TypeEathquakeUpdateOne) ExecX(ctx context.Context) {
	if err := teuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (teuo *TypeEathquakeUpdateOne) sqlSave(ctx context.Context) (_node *Type_eathquake, err error) {
	_spec := sqlgraph.NewUpdateSpec(type_eathquake.Table, type_eathquake.Columns, sqlgraph.NewFieldSpec(type_eathquake.FieldID, field.TypeInt))
	id, ok := teuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Type_eathquake.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := teuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, type_eathquake.FieldID)
		for _, f := range fields {
			if !type_eathquake.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != type_eathquake.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := teuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teuo.mutation.TypeEathquake(); ok {
		_spec.SetField(type_eathquake.FieldTypeEathquake, field.TypeString, value)
	}
	if teuo.mutation.TypeEathquakeCleared() {
		_spec.ClearField(type_eathquake.FieldTypeEathquake, field.TypeString)
	}
	if value, ok := teuo.mutation.CreatedAt(); ok {
		_spec.SetField(type_eathquake.FieldCreatedAt, field.TypeTime, value)
	}
	if teuo.mutation.CreatedAtCleared() {
		_spec.ClearField(type_eathquake.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := teuo.mutation.UpdatedAt(); ok {
		_spec.SetField(type_eathquake.FieldUpdatedAt, field.TypeTime, value)
	}
	if teuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(type_eathquake.FieldUpdatedAt, field.TypeTime)
	}
	if teuo.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   type_eathquake.EarthquakeTable,
			Columns: []string{type_eathquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := teuo.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   type_eathquake.EarthquakeTable,
			Columns: []string{type_eathquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Type_eathquake{config: teuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, teuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{type_eathquake.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	teuo.mutation.done = true
	return _node, nil
}
