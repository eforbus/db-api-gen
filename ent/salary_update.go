// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"employees/ent/employee"
	"employees/ent/predicate"
	"employees/ent/salary"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SalaryUpdate is the builder for updating Salary entities.
type SalaryUpdate struct {
	config
	hooks    []Hook
	mutation *SalaryMutation
}

// Where appends a list predicates to the SalaryUpdate builder.
func (su *SalaryUpdate) Where(ps ...predicate.Salary) *SalaryUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetEmployeeID sets the "employee_id" field.
func (su *SalaryUpdate) SetEmployeeID(s string) *SalaryUpdate {
	su.mutation.SetEmployeeID(s)
	return su
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (su *SalaryUpdate) SetNillableEmployeeID(s *string) *SalaryUpdate {
	if s != nil {
		su.SetEmployeeID(*s)
	}
	return su
}

// ClearEmployeeID clears the value of the "employee_id" field.
func (su *SalaryUpdate) ClearEmployeeID() *SalaryUpdate {
	su.mutation.ClearEmployeeID()
	return su
}

// SetSalary sets the "salary" field.
func (su *SalaryUpdate) SetSalary(i int32) *SalaryUpdate {
	su.mutation.ResetSalary()
	su.mutation.SetSalary(i)
	return su
}

// AddSalary adds i to the "salary" field.
func (su *SalaryUpdate) AddSalary(i int32) *SalaryUpdate {
	su.mutation.AddSalary(i)
	return su
}

// SetFromDate sets the "from_date" field.
func (su *SalaryUpdate) SetFromDate(t time.Time) *SalaryUpdate {
	su.mutation.SetFromDate(t)
	return su
}

// SetToDate sets the "to_date" field.
func (su *SalaryUpdate) SetToDate(t time.Time) *SalaryUpdate {
	su.mutation.SetToDate(t)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SalaryUpdate) SetCreatedAt(t time.Time) *SalaryUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SalaryUpdate) SetNillableCreatedAt(t *time.Time) *SalaryUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// ClearCreatedAt clears the value of the "created_at" field.
func (su *SalaryUpdate) ClearCreatedAt() *SalaryUpdate {
	su.mutation.ClearCreatedAt()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SalaryUpdate) SetUpdatedAt(t time.Time) *SalaryUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (su *SalaryUpdate) SetNillableUpdatedAt(t *time.Time) *SalaryUpdate {
	if t != nil {
		su.SetUpdatedAt(*t)
	}
	return su
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (su *SalaryUpdate) ClearUpdatedAt() *SalaryUpdate {
	su.mutation.ClearUpdatedAt()
	return su
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (su *SalaryUpdate) SetEmployee(e *Employee) *SalaryUpdate {
	return su.SetEmployeeID(e.ID)
}

// Mutation returns the SalaryMutation object of the builder.
func (su *SalaryUpdate) Mutation() *SalaryMutation {
	return su.mutation
}

// ClearEmployee clears the "employee" edge to the Employee entity.
func (su *SalaryUpdate) ClearEmployee() *SalaryUpdate {
	su.mutation.ClearEmployee()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SalaryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SalaryUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SalaryUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SalaryUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SalaryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(salary.Table, salary.Columns, sqlgraph.NewFieldSpec(salary.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Salary(); ok {
		_spec.SetField(salary.FieldSalary, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedSalary(); ok {
		_spec.AddField(salary.FieldSalary, field.TypeInt32, value)
	}
	if value, ok := su.mutation.FromDate(); ok {
		_spec.SetField(salary.FieldFromDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.ToDate(); ok {
		_spec.SetField(salary.FieldToDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(salary.FieldCreatedAt, field.TypeTime, value)
	}
	if su.mutation.CreatedAtCleared() {
		_spec.ClearField(salary.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(salary.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UpdatedAtCleared() {
		_spec.ClearField(salary.FieldUpdatedAt, field.TypeTime)
	}
	if su.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{salary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SalaryUpdateOne is the builder for updating a single Salary entity.
type SalaryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SalaryMutation
}

// SetEmployeeID sets the "employee_id" field.
func (suo *SalaryUpdateOne) SetEmployeeID(s string) *SalaryUpdateOne {
	suo.mutation.SetEmployeeID(s)
	return suo
}

// SetNillableEmployeeID sets the "employee_id" field if the given value is not nil.
func (suo *SalaryUpdateOne) SetNillableEmployeeID(s *string) *SalaryUpdateOne {
	if s != nil {
		suo.SetEmployeeID(*s)
	}
	return suo
}

// ClearEmployeeID clears the value of the "employee_id" field.
func (suo *SalaryUpdateOne) ClearEmployeeID() *SalaryUpdateOne {
	suo.mutation.ClearEmployeeID()
	return suo
}

// SetSalary sets the "salary" field.
func (suo *SalaryUpdateOne) SetSalary(i int32) *SalaryUpdateOne {
	suo.mutation.ResetSalary()
	suo.mutation.SetSalary(i)
	return suo
}

// AddSalary adds i to the "salary" field.
func (suo *SalaryUpdateOne) AddSalary(i int32) *SalaryUpdateOne {
	suo.mutation.AddSalary(i)
	return suo
}

// SetFromDate sets the "from_date" field.
func (suo *SalaryUpdateOne) SetFromDate(t time.Time) *SalaryUpdateOne {
	suo.mutation.SetFromDate(t)
	return suo
}

// SetToDate sets the "to_date" field.
func (suo *SalaryUpdateOne) SetToDate(t time.Time) *SalaryUpdateOne {
	suo.mutation.SetToDate(t)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SalaryUpdateOne) SetCreatedAt(t time.Time) *SalaryUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SalaryUpdateOne) SetNillableCreatedAt(t *time.Time) *SalaryUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (suo *SalaryUpdateOne) ClearCreatedAt() *SalaryUpdateOne {
	suo.mutation.ClearCreatedAt()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SalaryUpdateOne) SetUpdatedAt(t time.Time) *SalaryUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suo *SalaryUpdateOne) SetNillableUpdatedAt(t *time.Time) *SalaryUpdateOne {
	if t != nil {
		suo.SetUpdatedAt(*t)
	}
	return suo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suo *SalaryUpdateOne) ClearUpdatedAt() *SalaryUpdateOne {
	suo.mutation.ClearUpdatedAt()
	return suo
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (suo *SalaryUpdateOne) SetEmployee(e *Employee) *SalaryUpdateOne {
	return suo.SetEmployeeID(e.ID)
}

// Mutation returns the SalaryMutation object of the builder.
func (suo *SalaryUpdateOne) Mutation() *SalaryMutation {
	return suo.mutation
}

// ClearEmployee clears the "employee" edge to the Employee entity.
func (suo *SalaryUpdateOne) ClearEmployee() *SalaryUpdateOne {
	suo.mutation.ClearEmployee()
	return suo
}

// Where appends a list predicates to the SalaryUpdate builder.
func (suo *SalaryUpdateOne) Where(ps ...predicate.Salary) *SalaryUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SalaryUpdateOne) Select(field string, fields ...string) *SalaryUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Salary entity.
func (suo *SalaryUpdateOne) Save(ctx context.Context) (*Salary, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SalaryUpdateOne) SaveX(ctx context.Context) *Salary {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SalaryUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SalaryUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SalaryUpdateOne) sqlSave(ctx context.Context) (_node *Salary, err error) {
	_spec := sqlgraph.NewUpdateSpec(salary.Table, salary.Columns, sqlgraph.NewFieldSpec(salary.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Salary.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, salary.FieldID)
		for _, f := range fields {
			if !salary.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != salary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Salary(); ok {
		_spec.SetField(salary.FieldSalary, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedSalary(); ok {
		_spec.AddField(salary.FieldSalary, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.FromDate(); ok {
		_spec.SetField(salary.FieldFromDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.ToDate(); ok {
		_spec.SetField(salary.FieldToDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(salary.FieldCreatedAt, field.TypeTime, value)
	}
	if suo.mutation.CreatedAtCleared() {
		_spec.ClearField(salary.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(salary.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UpdatedAtCleared() {
		_spec.ClearField(salary.FieldUpdatedAt, field.TypeTime)
	}
	if suo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   salary.EmployeeTable,
			Columns: []string{salary.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Salary{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{salary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
