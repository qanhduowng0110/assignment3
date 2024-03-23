// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/associatedevent"
	"assignment3/ent/earthquake"
	"assignment3/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssociatedEventUpdate is the builder for updating AssociatedEvent entities.
type AssociatedEventUpdate struct {
	config
	hooks    []Hook
	mutation *AssociatedEventMutation
}

// Where appends a list predicates to the AssociatedEventUpdate builder.
func (aeu *AssociatedEventUpdate) Where(ps ...predicate.AssociatedEvent) *AssociatedEventUpdate {
	aeu.mutation.Where(ps...)
	return aeu
}

// SetEarthquakeID sets the "earthquake_id" field.
func (aeu *AssociatedEventUpdate) SetEarthquakeID(i int) *AssociatedEventUpdate {
	aeu.mutation.SetEarthquakeID(i)
	return aeu
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (aeu *AssociatedEventUpdate) SetNillableEarthquakeID(i *int) *AssociatedEventUpdate {
	if i != nil {
		aeu.SetEarthquakeID(*i)
	}
	return aeu
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (aeu *AssociatedEventUpdate) ClearEarthquakeID() *AssociatedEventUpdate {
	aeu.mutation.ClearEarthquakeID()
	return aeu
}

// SetCreatedAt sets the "created_at" field.
func (aeu *AssociatedEventUpdate) SetCreatedAt(t time.Time) *AssociatedEventUpdate {
	aeu.mutation.SetCreatedAt(t)
	return aeu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aeu *AssociatedEventUpdate) SetNillableCreatedAt(t *time.Time) *AssociatedEventUpdate {
	if t != nil {
		aeu.SetCreatedAt(*t)
	}
	return aeu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (aeu *AssociatedEventUpdate) ClearCreatedAt() *AssociatedEventUpdate {
	aeu.mutation.ClearCreatedAt()
	return aeu
}

// SetUpdatedAt sets the "updated_at" field.
func (aeu *AssociatedEventUpdate) SetUpdatedAt(t time.Time) *AssociatedEventUpdate {
	aeu.mutation.SetUpdatedAt(t)
	return aeu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aeu *AssociatedEventUpdate) SetNillableUpdatedAt(t *time.Time) *AssociatedEventUpdate {
	if t != nil {
		aeu.SetUpdatedAt(*t)
	}
	return aeu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aeu *AssociatedEventUpdate) ClearUpdatedAt() *AssociatedEventUpdate {
	aeu.mutation.ClearUpdatedAt()
	return aeu
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (aeu *AssociatedEventUpdate) SetEarthquake(e *Earthquake) *AssociatedEventUpdate {
	return aeu.SetEarthquakeID(e.ID)
}

// Mutation returns the AssociatedEventMutation object of the builder.
func (aeu *AssociatedEventUpdate) Mutation() *AssociatedEventMutation {
	return aeu.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (aeu *AssociatedEventUpdate) ClearEarthquake() *AssociatedEventUpdate {
	aeu.mutation.ClearEarthquake()
	return aeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aeu *AssociatedEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aeu.sqlSave, aeu.mutation, aeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeu *AssociatedEventUpdate) SaveX(ctx context.Context) int {
	affected, err := aeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aeu *AssociatedEventUpdate) Exec(ctx context.Context) error {
	_, err := aeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeu *AssociatedEventUpdate) ExecX(ctx context.Context) {
	if err := aeu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aeu *AssociatedEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(associatedevent.Table, associatedevent.Columns, sqlgraph.NewFieldSpec(associatedevent.FieldID, field.TypeInt))
	if ps := aeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeu.mutation.CreatedAt(); ok {
		_spec.SetField(associatedevent.FieldCreatedAt, field.TypeTime, value)
	}
	if aeu.mutation.CreatedAtCleared() {
		_spec.ClearField(associatedevent.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := aeu.mutation.UpdatedAt(); ok {
		_spec.SetField(associatedevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if aeu.mutation.UpdatedAtCleared() {
		_spec.ClearField(associatedevent.FieldUpdatedAt, field.TypeTime)
	}
	if aeu.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   associatedevent.EarthquakeTable,
			Columns: []string{associatedevent.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aeu.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   associatedevent.EarthquakeTable,
			Columns: []string{associatedevent.EarthquakeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, aeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{associatedevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aeu.mutation.done = true
	return n, nil
}

// AssociatedEventUpdateOne is the builder for updating a single AssociatedEvent entity.
type AssociatedEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssociatedEventMutation
}

// SetEarthquakeID sets the "earthquake_id" field.
func (aeuo *AssociatedEventUpdateOne) SetEarthquakeID(i int) *AssociatedEventUpdateOne {
	aeuo.mutation.SetEarthquakeID(i)
	return aeuo
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (aeuo *AssociatedEventUpdateOne) SetNillableEarthquakeID(i *int) *AssociatedEventUpdateOne {
	if i != nil {
		aeuo.SetEarthquakeID(*i)
	}
	return aeuo
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (aeuo *AssociatedEventUpdateOne) ClearEarthquakeID() *AssociatedEventUpdateOne {
	aeuo.mutation.ClearEarthquakeID()
	return aeuo
}

// SetCreatedAt sets the "created_at" field.
func (aeuo *AssociatedEventUpdateOne) SetCreatedAt(t time.Time) *AssociatedEventUpdateOne {
	aeuo.mutation.SetCreatedAt(t)
	return aeuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aeuo *AssociatedEventUpdateOne) SetNillableCreatedAt(t *time.Time) *AssociatedEventUpdateOne {
	if t != nil {
		aeuo.SetCreatedAt(*t)
	}
	return aeuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (aeuo *AssociatedEventUpdateOne) ClearCreatedAt() *AssociatedEventUpdateOne {
	aeuo.mutation.ClearCreatedAt()
	return aeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aeuo *AssociatedEventUpdateOne) SetUpdatedAt(t time.Time) *AssociatedEventUpdateOne {
	aeuo.mutation.SetUpdatedAt(t)
	return aeuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aeuo *AssociatedEventUpdateOne) SetNillableUpdatedAt(t *time.Time) *AssociatedEventUpdateOne {
	if t != nil {
		aeuo.SetUpdatedAt(*t)
	}
	return aeuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aeuo *AssociatedEventUpdateOne) ClearUpdatedAt() *AssociatedEventUpdateOne {
	aeuo.mutation.ClearUpdatedAt()
	return aeuo
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (aeuo *AssociatedEventUpdateOne) SetEarthquake(e *Earthquake) *AssociatedEventUpdateOne {
	return aeuo.SetEarthquakeID(e.ID)
}

// Mutation returns the AssociatedEventMutation object of the builder.
func (aeuo *AssociatedEventUpdateOne) Mutation() *AssociatedEventMutation {
	return aeuo.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (aeuo *AssociatedEventUpdateOne) ClearEarthquake() *AssociatedEventUpdateOne {
	aeuo.mutation.ClearEarthquake()
	return aeuo
}

// Where appends a list predicates to the AssociatedEventUpdate builder.
func (aeuo *AssociatedEventUpdateOne) Where(ps ...predicate.AssociatedEvent) *AssociatedEventUpdateOne {
	aeuo.mutation.Where(ps...)
	return aeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aeuo *AssociatedEventUpdateOne) Select(field string, fields ...string) *AssociatedEventUpdateOne {
	aeuo.fields = append([]string{field}, fields...)
	return aeuo
}

// Save executes the query and returns the updated AssociatedEvent entity.
func (aeuo *AssociatedEventUpdateOne) Save(ctx context.Context) (*AssociatedEvent, error) {
	return withHooks(ctx, aeuo.sqlSave, aeuo.mutation, aeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeuo *AssociatedEventUpdateOne) SaveX(ctx context.Context) *AssociatedEvent {
	node, err := aeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aeuo *AssociatedEventUpdateOne) Exec(ctx context.Context) error {
	_, err := aeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeuo *AssociatedEventUpdateOne) ExecX(ctx context.Context) {
	if err := aeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aeuo *AssociatedEventUpdateOne) sqlSave(ctx context.Context) (_node *AssociatedEvent, err error) {
	_spec := sqlgraph.NewUpdateSpec(associatedevent.Table, associatedevent.Columns, sqlgraph.NewFieldSpec(associatedevent.FieldID, field.TypeInt))
	id, ok := aeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AssociatedEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, associatedevent.FieldID)
		for _, f := range fields {
			if !associatedevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != associatedevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeuo.mutation.CreatedAt(); ok {
		_spec.SetField(associatedevent.FieldCreatedAt, field.TypeTime, value)
	}
	if aeuo.mutation.CreatedAtCleared() {
		_spec.ClearField(associatedevent.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := aeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(associatedevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if aeuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(associatedevent.FieldUpdatedAt, field.TypeTime)
	}
	if aeuo.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   associatedevent.EarthquakeTable,
			Columns: []string{associatedevent.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aeuo.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   associatedevent.EarthquakeTable,
			Columns: []string{associatedevent.EarthquakeColumn},
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
	_node = &AssociatedEvent{config: aeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{associatedevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aeuo.mutation.done = true
	return _node, nil
}
