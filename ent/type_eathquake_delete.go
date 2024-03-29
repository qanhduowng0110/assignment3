// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/predicate"
	"assignment3/ent/type_eathquake"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TypeEathquakeDelete is the builder for deleting a Type_eathquake entity.
type TypeEathquakeDelete struct {
	config
	hooks    []Hook
	mutation *TypeEathquakeMutation
}

// Where appends a list predicates to the TypeEathquakeDelete builder.
func (ted *TypeEathquakeDelete) Where(ps ...predicate.Type_eathquake) *TypeEathquakeDelete {
	ted.mutation.Where(ps...)
	return ted
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ted *TypeEathquakeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ted.sqlExec, ted.mutation, ted.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ted *TypeEathquakeDelete) ExecX(ctx context.Context) int {
	n, err := ted.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ted *TypeEathquakeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(type_eathquake.Table, sqlgraph.NewFieldSpec(type_eathquake.FieldID, field.TypeInt))
	if ps := ted.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ted.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ted.mutation.done = true
	return affected, err
}

// TypeEathquakeDeleteOne is the builder for deleting a single Type_eathquake entity.
type TypeEathquakeDeleteOne struct {
	ted *TypeEathquakeDelete
}

// Where appends a list predicates to the TypeEathquakeDelete builder.
func (tedo *TypeEathquakeDeleteOne) Where(ps ...predicate.Type_eathquake) *TypeEathquakeDeleteOne {
	tedo.ted.mutation.Where(ps...)
	return tedo
}

// Exec executes the deletion query.
func (tedo *TypeEathquakeDeleteOne) Exec(ctx context.Context) error {
	n, err := tedo.ted.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{type_eathquake.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tedo *TypeEathquakeDeleteOne) ExecX(ctx context.Context) {
	if err := tedo.Exec(ctx); err != nil {
		panic(err)
	}
}
