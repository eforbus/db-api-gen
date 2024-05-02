// Code generated by ent, DO NOT EDIT.

package deptemployee

import (
	"employees/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldContainsFold(FieldID, id))
}

// EmployeeID applies equality check predicate on the "employee_id" field. It's identical to EmployeeIDEQ.
func EmployeeID(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldEmployeeID, v))
}

// DepartmentID applies equality check predicate on the "department_id" field. It's identical to DepartmentIDEQ.
func DepartmentID(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldDepartmentID, v))
}

// FromDate applies equality check predicate on the "from_date" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldFromDate, v))
}

// ToDate applies equality check predicate on the "to_date" field. It's identical to ToDateEQ.
func ToDate(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldToDate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldUpdatedAt, v))
}

// EmployeeIDEQ applies the EQ predicate on the "employee_id" field.
func EmployeeIDEQ(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldEmployeeID, v))
}

// EmployeeIDNEQ applies the NEQ predicate on the "employee_id" field.
func EmployeeIDNEQ(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldEmployeeID, v))
}

// EmployeeIDIn applies the In predicate on the "employee_id" field.
func EmployeeIDIn(vs ...string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldEmployeeID, vs...))
}

// EmployeeIDNotIn applies the NotIn predicate on the "employee_id" field.
func EmployeeIDNotIn(vs ...string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldEmployeeID, vs...))
}

// EmployeeIDGT applies the GT predicate on the "employee_id" field.
func EmployeeIDGT(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldEmployeeID, v))
}

// EmployeeIDGTE applies the GTE predicate on the "employee_id" field.
func EmployeeIDGTE(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldEmployeeID, v))
}

// EmployeeIDLT applies the LT predicate on the "employee_id" field.
func EmployeeIDLT(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldEmployeeID, v))
}

// EmployeeIDLTE applies the LTE predicate on the "employee_id" field.
func EmployeeIDLTE(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldEmployeeID, v))
}

// EmployeeIDContains applies the Contains predicate on the "employee_id" field.
func EmployeeIDContains(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldContains(FieldEmployeeID, v))
}

// EmployeeIDHasPrefix applies the HasPrefix predicate on the "employee_id" field.
func EmployeeIDHasPrefix(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldHasPrefix(FieldEmployeeID, v))
}

// EmployeeIDHasSuffix applies the HasSuffix predicate on the "employee_id" field.
func EmployeeIDHasSuffix(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldHasSuffix(FieldEmployeeID, v))
}

// EmployeeIDIsNil applies the IsNil predicate on the "employee_id" field.
func EmployeeIDIsNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIsNull(FieldEmployeeID))
}

// EmployeeIDNotNil applies the NotNil predicate on the "employee_id" field.
func EmployeeIDNotNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotNull(FieldEmployeeID))
}

// EmployeeIDEqualFold applies the EqualFold predicate on the "employee_id" field.
func EmployeeIDEqualFold(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEqualFold(FieldEmployeeID, v))
}

// EmployeeIDContainsFold applies the ContainsFold predicate on the "employee_id" field.
func EmployeeIDContainsFold(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldContainsFold(FieldEmployeeID, v))
}

// DepartmentIDEQ applies the EQ predicate on the "department_id" field.
func DepartmentIDEQ(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldDepartmentID, v))
}

// DepartmentIDNEQ applies the NEQ predicate on the "department_id" field.
func DepartmentIDNEQ(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldDepartmentID, v))
}

// DepartmentIDIn applies the In predicate on the "department_id" field.
func DepartmentIDIn(vs ...string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldDepartmentID, vs...))
}

// DepartmentIDNotIn applies the NotIn predicate on the "department_id" field.
func DepartmentIDNotIn(vs ...string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldDepartmentID, vs...))
}

// DepartmentIDGT applies the GT predicate on the "department_id" field.
func DepartmentIDGT(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldDepartmentID, v))
}

// DepartmentIDGTE applies the GTE predicate on the "department_id" field.
func DepartmentIDGTE(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldDepartmentID, v))
}

// DepartmentIDLT applies the LT predicate on the "department_id" field.
func DepartmentIDLT(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldDepartmentID, v))
}

// DepartmentIDLTE applies the LTE predicate on the "department_id" field.
func DepartmentIDLTE(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldDepartmentID, v))
}

// DepartmentIDContains applies the Contains predicate on the "department_id" field.
func DepartmentIDContains(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldContains(FieldDepartmentID, v))
}

// DepartmentIDHasPrefix applies the HasPrefix predicate on the "department_id" field.
func DepartmentIDHasPrefix(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldHasPrefix(FieldDepartmentID, v))
}

// DepartmentIDHasSuffix applies the HasSuffix predicate on the "department_id" field.
func DepartmentIDHasSuffix(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldHasSuffix(FieldDepartmentID, v))
}

// DepartmentIDIsNil applies the IsNil predicate on the "department_id" field.
func DepartmentIDIsNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIsNull(FieldDepartmentID))
}

// DepartmentIDNotNil applies the NotNil predicate on the "department_id" field.
func DepartmentIDNotNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotNull(FieldDepartmentID))
}

// DepartmentIDEqualFold applies the EqualFold predicate on the "department_id" field.
func DepartmentIDEqualFold(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEqualFold(FieldDepartmentID, v))
}

// DepartmentIDContainsFold applies the ContainsFold predicate on the "department_id" field.
func DepartmentIDContainsFold(v string) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldContainsFold(FieldDepartmentID, v))
}

// FromDateEQ applies the EQ predicate on the "from_date" field.
func FromDateEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldFromDate, v))
}

// FromDateNEQ applies the NEQ predicate on the "from_date" field.
func FromDateNEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldFromDate, v))
}

// FromDateIn applies the In predicate on the "from_date" field.
func FromDateIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldFromDate, vs...))
}

// FromDateNotIn applies the NotIn predicate on the "from_date" field.
func FromDateNotIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldFromDate, vs...))
}

// FromDateGT applies the GT predicate on the "from_date" field.
func FromDateGT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldFromDate, v))
}

// FromDateGTE applies the GTE predicate on the "from_date" field.
func FromDateGTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldFromDate, v))
}

// FromDateLT applies the LT predicate on the "from_date" field.
func FromDateLT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldFromDate, v))
}

// FromDateLTE applies the LTE predicate on the "from_date" field.
func FromDateLTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldFromDate, v))
}

// ToDateEQ applies the EQ predicate on the "to_date" field.
func ToDateEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldToDate, v))
}

// ToDateNEQ applies the NEQ predicate on the "to_date" field.
func ToDateNEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldToDate, v))
}

// ToDateIn applies the In predicate on the "to_date" field.
func ToDateIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldToDate, vs...))
}

// ToDateNotIn applies the NotIn predicate on the "to_date" field.
func ToDateNotIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldToDate, vs...))
}

// ToDateGT applies the GT predicate on the "to_date" field.
func ToDateGT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldToDate, v))
}

// ToDateGTE applies the GTE predicate on the "to_date" field.
func ToDateGTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldToDate, v))
}

// ToDateLT applies the LT predicate on the "to_date" field.
func ToDateLT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldToDate, v))
}

// ToDateLTE applies the LTE predicate on the "to_date" field.
func ToDateLTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldToDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.FieldNotNull(FieldUpdatedAt))
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.DeptEmployee {
	return predicate.DeptEmployee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.DeptEmployee {
	return predicate.DeptEmployee(func(s *sql.Selector) {
		step := newEmployeeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.DeptEmployee {
	return predicate.DeptEmployee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.DeptEmployee {
	return predicate.DeptEmployee(func(s *sql.Selector) {
		step := newDepartmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeptEmployee) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeptEmployee) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeptEmployee) predicate.DeptEmployee {
	return predicate.DeptEmployee(sql.NotPredicates(p))
}
