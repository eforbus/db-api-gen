// Code generated by ent, DO NOT EDIT.

package ent

import (
	"employees/ent/employee"
	"employees/ent/title"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Title is the model entity for the Title schema.
type Title struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// EmployeeID holds the value of the "employee_id" field.
	EmployeeID string `json:"employee_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// FromDate holds the value of the "from_date" field.
	FromDate time.Time `json:"from_date,omitempty"`
	// ToDate holds the value of the "to_date" field.
	ToDate time.Time `json:"to_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TitleQuery when eager-loading is set.
	Edges        TitleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TitleEdges holds the relations/edges for other nodes in the graph.
type TitleEdges struct {
	// Employee holds the value of the employee edge.
	Employee *Employee `json:"employee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TitleEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[0] {
		if e.Employee == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Title) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case title.FieldID, title.FieldEmployeeID, title.FieldTitle:
			values[i] = new(sql.NullString)
		case title.FieldFromDate, title.FieldToDate, title.FieldCreatedAt, title.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Title fields.
func (t *Title) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case title.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case title.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field employee_id", values[i])
			} else if value.Valid {
				t.EmployeeID = value.String
			}
		case title.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case title.FieldFromDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from_date", values[i])
			} else if value.Valid {
				t.FromDate = value.Time
			}
		case title.FieldToDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field to_date", values[i])
			} else if value.Valid {
				t.ToDate = value.Time
			}
		case title.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case title.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Title.
// This includes values selected through modifiers, order, etc.
func (t *Title) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryEmployee queries the "employee" edge of the Title entity.
func (t *Title) QueryEmployee() *EmployeeQuery {
	return NewTitleClient(t.config).QueryEmployee(t)
}

// Update returns a builder for updating this Title.
// Note that you need to call Title.Unwrap() before calling this method if this Title
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Title) Update() *TitleUpdateOne {
	return NewTitleClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Title entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Title) Unwrap() *Title {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Title is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Title) String() string {
	var builder strings.Builder
	builder.WriteString("Title(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("employee_id=")
	builder.WriteString(t.EmployeeID)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("from_date=")
	builder.WriteString(t.FromDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("to_date=")
	builder.WriteString(t.ToDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Titles is a parsable slice of Title.
type Titles []*Title
