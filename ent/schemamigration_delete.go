// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/predicate"
	"assignment3/ent/schemamigration"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SchemaMigrationDelete is the builder for deleting a SchemaMigration entity.
type SchemaMigrationDelete struct {
	config
	hooks    []Hook
	mutation *SchemaMigrationMutation
}

// Where appends a list predicates to the SchemaMigrationDelete builder.
func (smd *SchemaMigrationDelete) Where(ps ...predicate.SchemaMigration) *SchemaMigrationDelete {
	smd.mutation.Where(ps...)
	return smd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smd *SchemaMigrationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, smd.sqlExec, smd.mutation, smd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (smd *SchemaMigrationDelete) ExecX(ctx context.Context) int {
	n, err := smd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smd *SchemaMigrationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(schemamigration.Table, sqlgraph.NewFieldSpec(schemamigration.FieldID, field.TypeInt))
	if ps := smd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, smd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	smd.mutation.done = true
	return affected, err
}

// SchemaMigrationDeleteOne is the builder for deleting a single SchemaMigration entity.
type SchemaMigrationDeleteOne struct {
	smd *SchemaMigrationDelete
}

// Where appends a list predicates to the SchemaMigrationDelete builder.
func (smdo *SchemaMigrationDeleteOne) Where(ps ...predicate.SchemaMigration) *SchemaMigrationDeleteOne {
	smdo.smd.mutation.Where(ps...)
	return smdo
}

// Exec executes the deletion query.
func (smdo *SchemaMigrationDeleteOne) Exec(ctx context.Context) error {
	n, err := smdo.smd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{schemamigration.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smdo *SchemaMigrationDeleteOne) ExecX(ctx context.Context) {
	if err := smdo.Exec(ctx); err != nil {
		panic(err)
	}
}