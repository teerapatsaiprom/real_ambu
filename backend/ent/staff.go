// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/ambu/app/ent/staff"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Staff is the model entity for the Staff schema.
type Staff struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StaffEmail holds the value of the "staff_email" field.
	StaffEmail string `json:"staff_email,omitempty"`
	// StaffName holds the value of the "staff_name" field.
	StaffName string `json:"staff_name,omitempty"`
	// StaffPassword holds the value of the "staff_password" field.
	StaffPassword string `json:"staff_password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StaffQuery when eager-loading is set.
	Edges StaffEdges `json:"edges"`
}

// StaffEdges holds the relations/edges for other nodes in the graph.
type StaffEdges struct {
	// StaffPredicament holds the value of the staff_predicament edge.
	StaffPredicament []*Predicament
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StaffPredicamentOrErr returns the StaffPredicament value or an error if the edge
// was not loaded in eager-loading.
func (e StaffEdges) StaffPredicamentOrErr() ([]*Predicament, error) {
	if e.loadedTypes[0] {
		return e.StaffPredicament, nil
	}
	return nil, &NotLoadedError{edge: "staff_predicament"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Staff) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // staff_email
		&sql.NullString{}, // staff_name
		&sql.NullString{}, // staff_password
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Staff fields.
func (s *Staff) assignValues(values ...interface{}) error {
	if m, n := len(values), len(staff.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field staff_email", values[0])
	} else if value.Valid {
		s.StaffEmail = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field staff_name", values[1])
	} else if value.Valid {
		s.StaffName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field staff_password", values[2])
	} else if value.Valid {
		s.StaffPassword = value.String
	}
	return nil
}

// QueryStaffPredicament queries the staff_predicament edge of the Staff.
func (s *Staff) QueryStaffPredicament() *PredicamentQuery {
	return (&StaffClient{config: s.config}).QueryStaffPredicament(s)
}

// Update returns a builder for updating this Staff.
// Note that, you need to call Staff.Unwrap() before calling this method, if this Staff
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Staff) Update() *StaffUpdateOne {
	return (&StaffClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Staff) Unwrap() *Staff {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Staff is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Staff) String() string {
	var builder strings.Builder
	builder.WriteString("Staff(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", staff_email=")
	builder.WriteString(s.StaffEmail)
	builder.WriteString(", staff_name=")
	builder.WriteString(s.StaffName)
	builder.WriteString(", staff_password=")
	builder.WriteString(s.StaffPassword)
	builder.WriteByte(')')
	return builder.String()
}

// Staffs is a parsable slice of Staff.
type Staffs []*Staff

func (s Staffs) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
