// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"employees/ent/employee"
	"employees/ent/title"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TitleCreate is the builder for creating a Title entity.
type TitleCreate struct {
	config
	mutation *TitleMutation
	hooks    []Hook
}

// SetEmployeeID sets the "employee_id" field.
func (tc *TitleCreate) SetEmployeeID(s string) *TitleCreate {
	tc.mutation.SetEmployeeID(s)
	return tc
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (tc *TitleCreate) SetNillableEmployeeID(s *string) *TitleCreate {
	if s != nil {
		tc.SetEmployeeID(*s)
	}
	return tc
}

// SetTitle sets the "title" field.
func (tc *TitleCreate) SetTitle(s string) *TitleCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetFromDate sets the "from_date" field.
func (tc *TitleCreate) SetFromDate(t time.Time) *TitleCreate {
	tc.mutation.SetFromDate(t)
	return tc
}

// SetToDate sets the "to_date" field.
func (tc *TitleCreate) SetToDate(t time.Time) *TitleCreate {
	tc.mutation.SetToDate(t)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TitleCreate) SetCreatedAt(t time.Time) *TitleCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TitleCreate) SetNillableCreatedAt(t *time.Time) *TitleCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TitleCreate) SetUpdatedAt(t time.Time) *TitleCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TitleCreate) SetNillableUpdatedAt(t *time.Time) *TitleCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TitleCreate) SetID(s string) *TitleCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (tc *TitleCreate) SetEmployee(e *Employee) *TitleCreate {
	return tc.SetEmployeeID(e.ID)
}

// Mutation returns the TitleMutation object of the builder.
func (tc *TitleCreate) Mutation() *TitleMutation {
	return tc.mutation
}

// Save creates the Title in the database.
func (tc *TitleCreate) Save(ctx context.Context) (*Title, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TitleCreate) SaveX(ctx context.Context) *Title {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TitleCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TitleCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TitleCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Title.title"`)}
	}
	if _, ok := tc.mutation.FromDate(); !ok {
		return &ValidationError{Name: "from_date", err: errors.New(`ent: missing required field "Title.from_date"`)}
	}
	if _, ok := tc.mutation.ToDate(); !ok {
		return &ValidationError{Name: "to_date", err: errors.New(`ent: missing required field "Title.to_date"`)}
	}
	return nil
}

func (tc *TitleCreate) sqlSave(ctx context.Context) (*Title, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Title.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TitleCreate) createSpec() (*Title, *sqlgraph.CreateSpec) {
	var (
		_node = &Title{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(title.Table, sqlgraph.NewFieldSpec(title.FieldID, field.TypeString))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(title.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.FromDate(); ok {
		_spec.SetField(title.FieldFromDate, field.TypeTime, value)
		_node.FromDate = value
	}
	if value, ok := tc.mutation.ToDate(); ok {
		_spec.SetField(title.FieldToDate, field.TypeTime, value)
		_node.ToDate = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(title.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(title.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   title.EmployeeTable,
			Columns: []string{title.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EmployeeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TitleCreateBulk is the builder for creating many Title entities in bulk.
type TitleCreateBulk struct {
	config
	err      error
	builders []*TitleCreate
}

// Save creates the Title entities in the database.
func (tcb *TitleCreateBulk) Save(ctx context.Context) ([]*Title, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Title, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TitleMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TitleCreateBulk) SaveX(ctx context.Context) []*Title {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TitleCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TitleCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
