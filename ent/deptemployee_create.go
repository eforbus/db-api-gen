// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"employees/ent/department"
	"employees/ent/deptemployee"
	"employees/ent/employee"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeptEmployeeCreate is the builder for creating a DeptEmployee entity.
type DeptEmployeeCreate struct {
	config
	mutation *DeptEmployeeMutation
	hooks    []Hook
}

// SetEmployeeID sets the "employee_id" field.
func (dec *DeptEmployeeCreate) SetEmployeeID(s string) *DeptEmployeeCreate {
	dec.mutation.SetEmployeeID(s)
	return dec
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (dec *DeptEmployeeCreate) SetNillableEmployeeID(s *string) *DeptEmployeeCreate {
	if s != nil {
		dec.SetEmployeeID(*s)
	}
	return dec
}

// SetDepartmentID sets the "department_id" field.
func (dec *DeptEmployeeCreate) SetDepartmentID(s string) *DeptEmployeeCreate {
	dec.mutation.SetDepartmentID(s)
	return dec
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (dec *DeptEmployeeCreate) SetNillableDepartmentID(s *string) *DeptEmployeeCreate {
	if s != nil {
		dec.SetDepartmentID(*s)
	}
	return dec
}

// SetFromDate sets the "from_date" field.
func (dec *DeptEmployeeCreate) SetFromDate(t time.Time) *DeptEmployeeCreate {
	dec.mutation.SetFromDate(t)
	return dec
}

// SetToDate sets the "to_date" field.
func (dec *DeptEmployeeCreate) SetToDate(t time.Time) *DeptEmployeeCreate {
	dec.mutation.SetToDate(t)
	return dec
}

// SetCreatedAt sets the "created_at" field.
func (dec *DeptEmployeeCreate) SetCreatedAt(t time.Time) *DeptEmployeeCreate {
	dec.mutation.SetCreatedAt(t)
	return dec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dec *DeptEmployeeCreate) SetNillableCreatedAt(t *time.Time) *DeptEmployeeCreate {
	if t != nil {
		dec.SetCreatedAt(*t)
	}
	return dec
}

// SetUpdatedAt sets the "updated_at" field.
func (dec *DeptEmployeeCreate) SetUpdatedAt(t time.Time) *DeptEmployeeCreate {
	dec.mutation.SetUpdatedAt(t)
	return dec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dec *DeptEmployeeCreate) SetNillableUpdatedAt(t *time.Time) *DeptEmployeeCreate {
	if t != nil {
		dec.SetUpdatedAt(*t)
	}
	return dec
}

// SetID sets the "id" field.
func (dec *DeptEmployeeCreate) SetID(s string) *DeptEmployeeCreate {
	dec.mutation.SetID(s)
	return dec
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (dec *DeptEmployeeCreate) SetEmployee(e *Employee) *DeptEmployeeCreate {
	return dec.SetEmployeeID(e.ID)
}

// SetDepartment sets the "department" edge to the Department entity.
func (dec *DeptEmployeeCreate) SetDepartment(d *Department) *DeptEmployeeCreate {
	return dec.SetDepartmentID(d.ID)
}

// Mutation returns the DeptEmployeeMutation object of the builder.
func (dec *DeptEmployeeCreate) Mutation() *DeptEmployeeMutation {
	return dec.mutation
}

// Save creates the DeptEmployee in the database.
func (dec *DeptEmployeeCreate) Save(ctx context.Context) (*DeptEmployee, error) {
	return withHooks(ctx, dec.sqlSave, dec.mutation, dec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dec *DeptEmployeeCreate) SaveX(ctx context.Context) *DeptEmployee {
	v, err := dec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dec *DeptEmployeeCreate) Exec(ctx context.Context) error {
	_, err := dec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dec *DeptEmployeeCreate) ExecX(ctx context.Context) {
	if err := dec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dec *DeptEmployeeCreate) check() error {
	if _, ok := dec.mutation.FromDate(); !ok {
		return &ValidationError{Name: "from_date", err: errors.New(`ent: missing required field "DeptEmployee.from_date"`)}
	}
	if _, ok := dec.mutation.ToDate(); !ok {
		return &ValidationError{Name: "to_date", err: errors.New(`ent: missing required field "DeptEmployee.to_date"`)}
	}
	return nil
}

func (dec *DeptEmployeeCreate) sqlSave(ctx context.Context) (*DeptEmployee, error) {
	if err := dec.check(); err != nil {
		return nil, err
	}
	_node, _spec := dec.createSpec()
	if err := sqlgraph.CreateNode(ctx, dec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DeptEmployee.ID type: %T", _spec.ID.Value)
		}
	}
	dec.mutation.id = &_node.ID
	dec.mutation.done = true
	return _node, nil
}

func (dec *DeptEmployeeCreate) createSpec() (*DeptEmployee, *sqlgraph.CreateSpec) {
	var (
		_node = &DeptEmployee{config: dec.config}
		_spec = sqlgraph.NewCreateSpec(deptemployee.Table, sqlgraph.NewFieldSpec(deptemployee.FieldID, field.TypeString))
	)
	if id, ok := dec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dec.mutation.FromDate(); ok {
		_spec.SetField(deptemployee.FieldFromDate, field.TypeTime, value)
		_node.FromDate = value
	}
	if value, ok := dec.mutation.ToDate(); ok {
		_spec.SetField(deptemployee.FieldToDate, field.TypeTime, value)
		_node.ToDate = value
	}
	if value, ok := dec.mutation.CreatedAt(); ok {
		_spec.SetField(deptemployee.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dec.mutation.UpdatedAt(); ok {
		_spec.SetField(deptemployee.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := dec.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deptemployee.EmployeeTable,
			Columns: []string{deptemployee.EmployeeColumn},
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
	if nodes := dec.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deptemployee.DepartmentTable,
			Columns: []string{deptemployee.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DepartmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeptEmployeeCreateBulk is the builder for creating many DeptEmployee entities in bulk.
type DeptEmployeeCreateBulk struct {
	config
	err      error
	builders []*DeptEmployeeCreate
}

// Save creates the DeptEmployee entities in the database.
func (decb *DeptEmployeeCreateBulk) Save(ctx context.Context) ([]*DeptEmployee, error) {
	if decb.err != nil {
		return nil, decb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(decb.builders))
	nodes := make([]*DeptEmployee, len(decb.builders))
	mutators := make([]Mutator, len(decb.builders))
	for i := range decb.builders {
		func(i int, root context.Context) {
			builder := decb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeptEmployeeMutation)
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
					_, err = mutators[i+1].Mutate(root, decb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, decb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, decb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (decb *DeptEmployeeCreateBulk) SaveX(ctx context.Context) []*DeptEmployee {
	v, err := decb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (decb *DeptEmployeeCreateBulk) Exec(ctx context.Context) error {
	_, err := decb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (decb *DeptEmployeeCreateBulk) ExecX(ctx context.Context) {
	if err := decb.Exec(ctx); err != nil {
		panic(err)
	}
}
