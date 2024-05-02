// Code generated by ent, DO NOT EDIT.

package department

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeDeptEmployees holds the string denoting the dept_employees edge name in mutations.
	EdgeDeptEmployees = "dept_employees"
	// EdgeDeptManagers holds the string denoting the dept_managers edge name in mutations.
	EdgeDeptManagers = "dept_managers"
	// Table holds the table name of the department in the database.
	Table = "department"
	// DeptEmployeesTable is the table that holds the dept_employees relation/edge.
	DeptEmployeesTable = "dept_employee"
	// DeptEmployeesInverseTable is the table name for the DeptEmployee entity.
	// It exists in this package in order to avoid circular dependency with the "deptemployee" package.
	DeptEmployeesInverseTable = "dept_employee"
	// DeptEmployeesColumn is the table column denoting the dept_employees relation/edge.
	DeptEmployeesColumn = "department_id"
	// DeptManagersTable is the table that holds the dept_managers relation/edge.
	DeptManagersTable = "dept_manager"
	// DeptManagersInverseTable is the table name for the DeptManager entity.
	// It exists in this package in order to avoid circular dependency with the "deptmanager" package.
	DeptManagersInverseTable = "dept_manager"
	// DeptManagersColumn is the table column denoting the dept_managers relation/edge.
	DeptManagersColumn = "department_id"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Department queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeptEmployeesCount orders the results by dept_employees count.
func ByDeptEmployeesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeptEmployeesStep(), opts...)
	}
}

// ByDeptEmployees orders the results by dept_employees terms.
func ByDeptEmployees(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeptEmployeesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeptManagersCount orders the results by dept_managers count.
func ByDeptManagersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeptManagersStep(), opts...)
	}
}

// ByDeptManagers orders the results by dept_managers terms.
func ByDeptManagers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeptManagersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDeptEmployeesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeptEmployeesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeptEmployeesTable, DeptEmployeesColumn),
	)
}
func newDeptManagersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeptManagersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeptManagersTable, DeptManagersColumn),
	)
}
