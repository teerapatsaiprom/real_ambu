// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/ambu/app/ent/car"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CarNo holds the value of the "car_no" field.
	CarNo string `json:"car_no,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarQuery when eager-loading is set.
	Edges CarEdges `json:"edges"`
}

// CarEdges holds the relations/edges for other nodes in the graph.
type CarEdges struct {
	// CarPredicament holds the value of the car_predicament edge.
	CarPredicament []*Predicament
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CarPredicamentOrErr returns the CarPredicament value or an error if the edge
// was not loaded in eager-loading.
func (e CarEdges) CarPredicamentOrErr() ([]*Predicament, error) {
	if e.loadedTypes[0] {
		return e.CarPredicament, nil
	}
	return nil, &NotLoadedError{edge: "car_predicament"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // car_no
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(values ...interface{}) error {
	if m, n := len(values), len(car.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field car_no", values[0])
	} else if value.Valid {
		c.CarNo = value.String
	}
	return nil
}

// QueryCarPredicament queries the car_predicament edge of the Car.
func (c *Car) QueryCarPredicament() *PredicamentQuery {
	return (&CarClient{config: c.config}).QueryCarPredicament(c)
}

// Update returns a builder for updating this Car.
// Note that, you need to call Car.Unwrap() before calling this method, if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", car_no=")
	builder.WriteString(c.CarNo)
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

func (c Cars) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
