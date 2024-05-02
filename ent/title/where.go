// Code generated by ent, DO NOT EDIT.

package title

import (
	"employees/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Title {
	return predicate.Title(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Title {
	return predicate.Title(sql.FieldContainsFold(FieldID, id))
}

// EmployeeID applies equality check predicate on the "employee_id" field. It's identical to EmployeeIDEQ.
func EmployeeID(v string) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldEmployeeID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldTitle, v))
}

// FromDate applies equality check predicate on the "from_date" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldFromDate, v))
}

// ToDate applies equality check predicate on the "to_date" field. It's identical to ToDateEQ.
func ToDate(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldToDate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldUpdatedAt, v))
}

// EmployeeIDEQ applies the EQ predicate on the "employee_id" field.
func EmployeeIDEQ(v string) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldEmployeeID, v))
}

// EmployeeIDNEQ applies the NEQ predicate on the "employee_id" field.
func EmployeeIDNEQ(v string) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldEmployeeID, v))
}

// EmployeeIDIn applies the In predicate on the "employee_id" field.
func EmployeeIDIn(vs ...string) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldEmployeeID, vs...))
}

// EmployeeIDNotIn applies the NotIn predicate on the "employee_id" field.
func EmployeeIDNotIn(vs ...string) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldEmployeeID, vs...))
}

// EmployeeIDGT applies the GT predicate on the "employee_id" field.
func EmployeeIDGT(v string) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldEmployeeID, v))
}

// EmployeeIDGTE applies the GTE predicate on the "employee_id" field.
func EmployeeIDGTE(v string) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldEmployeeID, v))
}

// EmployeeIDLT applies the LT predicate on the "employee_id" field.
func EmployeeIDLT(v string) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldEmployeeID, v))
}

// EmployeeIDLTE applies the LTE predicate on the "employee_id" field.
func EmployeeIDLTE(v string) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldEmployeeID, v))
}

// EmployeeIDContains applies the Contains predicate on the "employee_id" field.
func EmployeeIDContains(v string) predicate.Title {
	return predicate.Title(sql.FieldContains(FieldEmployeeID, v))
}

// EmployeeIDHasPrefix applies the HasPrefix predicate on the "employee_id" field.
func EmployeeIDHasPrefix(v string) predicate.Title {
	return predicate.Title(sql.FieldHasPrefix(FieldEmployeeID, v))
}

// EmployeeIDHasSuffix applies the HasSuffix predicate on the "employee_id" field.
func EmployeeIDHasSuffix(v string) predicate.Title {
	return predicate.Title(sql.FieldHasSuffix(FieldEmployeeID, v))
}

// EmployeeIDIsNil applies the IsNil predicate on the "employee_id" field.
func EmployeeIDIsNil() predicate.Title {
	return predicate.Title(sql.FieldIsNull(FieldEmployeeID))
}

// EmployeeIDNotNil applies the NotNil predicate on the "employee_id" field.
func EmployeeIDNotNil() predicate.Title {
	return predicate.Title(sql.FieldNotNull(FieldEmployeeID))
}

// EmployeeIDEqualFold applies the EqualFold predicate on the "employee_id" field.
func EmployeeIDEqualFold(v string) predicate.Title {
	return predicate.Title(sql.FieldEqualFold(FieldEmployeeID, v))
}

// EmployeeIDContainsFold applies the ContainsFold predicate on the "employee_id" field.
func EmployeeIDContainsFold(v string) predicate.Title {
	return predicate.Title(sql.FieldContainsFold(FieldEmployeeID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Title {
	return predicate.Title(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Title {
	return predicate.Title(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Title {
	return predicate.Title(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Title {
	return predicate.Title(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Title {
	return predicate.Title(sql.FieldContainsFold(FieldTitle, v))
}

// FromDateEQ applies the EQ predicate on the "from_date" field.
func FromDateEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldFromDate, v))
}

// FromDateNEQ applies the NEQ predicate on the "from_date" field.
func FromDateNEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldFromDate, v))
}

// FromDateIn applies the In predicate on the "from_date" field.
func FromDateIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldFromDate, vs...))
}

// FromDateNotIn applies the NotIn predicate on the "from_date" field.
func FromDateNotIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldFromDate, vs...))
}

// FromDateGT applies the GT predicate on the "from_date" field.
func FromDateGT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldFromDate, v))
}

// FromDateGTE applies the GTE predicate on the "from_date" field.
func FromDateGTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldFromDate, v))
}

// FromDateLT applies the LT predicate on the "from_date" field.
func FromDateLT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldFromDate, v))
}

// FromDateLTE applies the LTE predicate on the "from_date" field.
func FromDateLTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldFromDate, v))
}

// ToDateEQ applies the EQ predicate on the "to_date" field.
func ToDateEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldToDate, v))
}

// ToDateNEQ applies the NEQ predicate on the "to_date" field.
func ToDateNEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldToDate, v))
}

// ToDateIn applies the In predicate on the "to_date" field.
func ToDateIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldToDate, vs...))
}

// ToDateNotIn applies the NotIn predicate on the "to_date" field.
func ToDateNotIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldToDate, vs...))
}

// ToDateGT applies the GT predicate on the "to_date" field.
func ToDateGT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldToDate, v))
}

// ToDateGTE applies the GTE predicate on the "to_date" field.
func ToDateGTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldToDate, v))
}

// ToDateLT applies the LT predicate on the "to_date" field.
func ToDateLT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldToDate, v))
}

// ToDateLTE applies the LTE predicate on the "to_date" field.
func ToDateLTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldToDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Title {
	return predicate.Title(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Title {
	return predicate.Title(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Title {
	return predicate.Title(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Title {
	return predicate.Title(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Title {
	return predicate.Title(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Title {
	return predicate.Title(sql.FieldNotNull(FieldUpdatedAt))
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Title {
	return predicate.Title(func(s *sql.Selector) {
		step := newEmployeeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Title) predicate.Title {
	return predicate.Title(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Title) predicate.Title {
	return predicate.Title(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Title) predicate.Title {
	return predicate.Title(sql.NotPredicates(p))
}
