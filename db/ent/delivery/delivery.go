// Code generated by entc, DO NOT EDIT.

package delivery

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the delivery type in the database.
	Label = "delivery"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "street"
	// FieldZipcode holds the string denoting the zipcode field in the database.
	FieldZipcode = "zipcode"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"

	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"

	// Table holds the table name of the delivery in the database.
	Table = "delivery"
	// OrderTable is the table the holds the order relation/edge.
	OrderTable = "orders"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "delivery_id"
)

// Columns holds all SQL columns for delivery fields.
var Columns = []string{
	FieldID,
	FieldCity,
	FieldStreet,
	FieldZipcode,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusREADY Status = "READY"
	StatusCOMP  Status = "COMP"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusREADY, StatusCOMP:
		return nil
	default:
		return fmt.Errorf("delivery: invalid enum value for status field: %q", s)
	}
}
