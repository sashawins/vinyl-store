// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"vinyl-store/ent/predicate"
	"vinyl-store/ent/vinyl"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VinylDelete is the builder for deleting a Vinyl entity.
type VinylDelete struct {
	config
	hooks    []Hook
	mutation *VinylMutation
}

// Where appends a list predicates to the VinylDelete builder.
func (vd *VinylDelete) Where(ps ...predicate.Vinyl) *VinylDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VinylDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vd.sqlExec, vd.mutation, vd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VinylDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VinylDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vinyl.Table, sqlgraph.NewFieldSpec(vinyl.FieldID, field.TypeUUID))
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vd.mutation.done = true
	return affected, err
}

// VinylDeleteOne is the builder for deleting a single Vinyl entity.
type VinylDeleteOne struct {
	vd *VinylDelete
}

// Where appends a list predicates to the VinylDelete builder.
func (vdo *VinylDeleteOne) Where(ps ...predicate.Vinyl) *VinylDeleteOne {
	vdo.vd.mutation.Where(ps...)
	return vdo
}

// Exec executes the deletion query.
func (vdo *VinylDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vinyl.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VinylDeleteOne) ExecX(ctx context.Context) {
	if err := vdo.Exec(ctx); err != nil {
		panic(err)
	}
}
