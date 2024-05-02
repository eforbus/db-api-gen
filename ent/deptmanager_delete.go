// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"employees/ent/deptmanager"
	"employees/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeptManagerDelete is the builder for deleting a DeptManager entity.
type DeptManagerDelete struct {
	config
	hooks    []Hook
	mutation *DeptManagerMutation
}

// Where appends a list predicates to the DeptManagerDelete builder.
func (dmd *DeptManagerDelete) Where(ps ...predicate.DeptManager) *DeptManagerDelete {
	dmd.mutation.Where(ps...)
	return dmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dmd *DeptManagerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dmd.sqlExec, dmd.mutation, dmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dmd *DeptManagerDelete) ExecX(ctx context.Context) int {
	n, err := dmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dmd *DeptManagerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(deptmanager.Table, sqlgraph.NewFieldSpec(deptmanager.FieldID, field.TypeString))
	if ps := dmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dmd.mutation.done = true
	return affected, err
}

// DeptManagerDeleteOne is the builder for deleting a single DeptManager entity.
type DeptManagerDeleteOne struct {
	dmd *DeptManagerDelete
}

// Where appends a list predicates to the DeptManagerDelete builder.
func (dmdo *DeptManagerDeleteOne) Where(ps ...predicate.DeptManager) *DeptManagerDeleteOne {
	dmdo.dmd.mutation.Where(ps...)
	return dmdo
}

// Exec executes the deletion query.
func (dmdo *DeptManagerDeleteOne) Exec(ctx context.Context) error {
	n, err := dmdo.dmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deptmanager.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dmdo *DeptManagerDeleteOne) ExecX(ctx context.Context) {
	if err := dmdo.Exec(ctx); err != nil {
		panic(err)
	}
}