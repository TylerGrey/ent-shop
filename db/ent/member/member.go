// Code generated by entc, DO NOT EDIT.

package member

import (
	"time"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"

	// Table holds the table name of the member in the database.
	Table = "member"
	// OrdersTable is the table the holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "member_id"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldCity,
	FieldStreet,
	FieldZipcode,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)