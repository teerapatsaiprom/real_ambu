// Code generated by entc, DO NOT EDIT.

package predicament

import (
	"time"
)

const (
	// Label holds the string label denoting the predicament type in the database.
	Label = "predicament"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddedTime holds the string denoting the added_time field in the database.
	FieldAddedTime = "added_time"

	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "Car"
	// EdgeStatuscar holds the string denoting the statuscar edge name in mutations.
	EdgeStatuscar = "Statuscar"
	// EdgeStaff holds the string denoting the staff edge name in mutations.
	EdgeStaff = "Staff"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "User"

	// Table holds the table name of the predicament in the database.
	Table = "predicaments"
	// CarTable is the table the holds the Car relation/edge.
	CarTable = "predicaments"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the Car relation/edge.
	CarColumn = "carid"
	// StatuscarTable is the table the holds the Statuscar relation/edge.
	StatuscarTable = "predicaments"
	// StatuscarInverseTable is the table name for the Statuscar entity.
	// It exists in this package in order to avoid circular dependency with the "statuscar" package.
	StatuscarInverseTable = "statuscars"
	// StatuscarColumn is the table column denoting the Statuscar relation/edge.
	StatuscarColumn = "status_id"
	// StaffTable is the table the holds the Staff relation/edge.
	StaffTable = "predicaments"
	// StaffInverseTable is the table name for the Staff entity.
	// It exists in this package in order to avoid circular dependency with the "staff" package.
	StaffInverseTable = "staffs"
	// StaffColumn is the table column denoting the Staff relation/edge.
	StaffColumn = "staffid"
	// UserTable is the table the holds the User relation/edge.
	UserTable = "predicaments"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the User relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for predicament fields.
var Columns = []string{
	FieldID,
	FieldAddedTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Predicament type.
var ForeignKeys = []string{
	"carid",
	"staffid",
	"status_id",
	"user_id",
}

var (
	// DefaultAddedTime holds the default value on creation for the Added_Time field.
	DefaultAddedTime func() time.Time
)