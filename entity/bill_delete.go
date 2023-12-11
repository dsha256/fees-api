// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dsha256/feesapi/entity/bill"
	"github.com/dsha256/feesapi/entity/predicate"
)

// BillDelete is the builder for deleting a Bill entity.
type BillDelete struct {
	config
	hooks    []Hook
	mutation *BillMutation
}

// Where appends a list predicates to the BillDelete builder.
func (bd *BillDelete) Where(ps ...predicate.Bill) *BillDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BillDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BillDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BillDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bill.Table, sqlgraph.NewFieldSpec(bill.FieldID, field.TypeString))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BillDeleteOne is the builder for deleting a single Bill entity.
type BillDeleteOne struct {
	bd *BillDelete
}

// Where appends a list predicates to the BillDelete builder.
func (bdo *BillDeleteOne) Where(ps ...predicate.Bill) *BillDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BillDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bill.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BillDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}
