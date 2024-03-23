// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/associatedevent"
	"assignment3/ent/earthquake"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssociatedEventCreate is the builder for creating a AssociatedEvent entity.
type AssociatedEventCreate struct {
	config
	mutation *AssociatedEventMutation
	hooks    []Hook
}

// SetEarthquakeID sets the "earthquake_id" field.
func (aec *AssociatedEventCreate) SetEarthquakeID(i int) *AssociatedEventCreate {
	aec.mutation.SetEarthquakeID(i)
	return aec
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (aec *AssociatedEventCreate) SetNillableEarthquakeID(i *int) *AssociatedEventCreate {
	if i != nil {
		aec.SetEarthquakeID(*i)
	}
	return aec
}

// SetCreatedAt sets the "created_at" field.
func (aec *AssociatedEventCreate) SetCreatedAt(t time.Time) *AssociatedEventCreate {
	aec.mutation.SetCreatedAt(t)
	return aec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aec *AssociatedEventCreate) SetNillableCreatedAt(t *time.Time) *AssociatedEventCreate {
	if t != nil {
		aec.SetCreatedAt(*t)
	}
	return aec
}

// SetUpdatedAt sets the "updated_at" field.
func (aec *AssociatedEventCreate) SetUpdatedAt(t time.Time) *AssociatedEventCreate {
	aec.mutation.SetUpdatedAt(t)
	return aec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aec *AssociatedEventCreate) SetNillableUpdatedAt(t *time.Time) *AssociatedEventCreate {
	if t != nil {
		aec.SetUpdatedAt(*t)
	}
	return aec
}

// SetID sets the "id" field.
func (aec *AssociatedEventCreate) SetID(i int) *AssociatedEventCreate {
	aec.mutation.SetID(i)
	return aec
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (aec *AssociatedEventCreate) SetEarthquake(e *Earthquake) *AssociatedEventCreate {
	return aec.SetEarthquakeID(e.ID)
}

// Mutation returns the AssociatedEventMutation object of the builder.
func (aec *AssociatedEventCreate) Mutation() *AssociatedEventMutation {
	return aec.mutation
}

// Save creates the AssociatedEvent in the database.
func (aec *AssociatedEventCreate) Save(ctx context.Context) (*AssociatedEvent, error) {
	return withHooks(ctx, aec.sqlSave, aec.mutation, aec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aec *AssociatedEventCreate) SaveX(ctx context.Context) *AssociatedEvent {
	v, err := aec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aec *AssociatedEventCreate) Exec(ctx context.Context) error {
	_, err := aec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aec *AssociatedEventCreate) ExecX(ctx context.Context) {
	if err := aec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aec *AssociatedEventCreate) check() error {
	return nil
}

func (aec *AssociatedEventCreate) sqlSave(ctx context.Context) (*AssociatedEvent, error) {
	if err := aec.check(); err != nil {
		return nil, err
	}
	_node, _spec := aec.createSpec()
	if err := sqlgraph.CreateNode(ctx, aec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	aec.mutation.id = &_node.ID
	aec.mutation.done = true
	return _node, nil
}

func (aec *AssociatedEventCreate) createSpec() (*AssociatedEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &AssociatedEvent{config: aec.config}
		_spec = sqlgraph.NewCreateSpec(associatedevent.Table, sqlgraph.NewFieldSpec(associatedevent.FieldID, field.TypeInt))
	)
	if id, ok := aec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aec.mutation.CreatedAt(); ok {
		_spec.SetField(associatedevent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aec.mutation.UpdatedAt(); ok {
		_spec.SetField(associatedevent.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := aec.mutation.EarthquakeIDs(); len(nodes) > 0 {
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
		_node.EarthquakeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssociatedEventCreateBulk is the builder for creating many AssociatedEvent entities in bulk.
type AssociatedEventCreateBulk struct {
	config
	err      error
	builders []*AssociatedEventCreate
}

// Save creates the AssociatedEvent entities in the database.
func (aecb *AssociatedEventCreateBulk) Save(ctx context.Context) ([]*AssociatedEvent, error) {
	if aecb.err != nil {
		return nil, aecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aecb.builders))
	nodes := make([]*AssociatedEvent, len(aecb.builders))
	mutators := make([]Mutator, len(aecb.builders))
	for i := range aecb.builders {
		func(i int, root context.Context) {
			builder := aecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssociatedEventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aecb *AssociatedEventCreateBulk) SaveX(ctx context.Context) []*AssociatedEvent {
	v, err := aecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aecb *AssociatedEventCreateBulk) Exec(ctx context.Context) error {
	_, err := aecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aecb *AssociatedEventCreateBulk) ExecX(ctx context.Context) {
	if err := aecb.Exec(ctx); err != nil {
		panic(err)
	}
}
