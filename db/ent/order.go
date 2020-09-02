// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/TylerGrey/ent-shop/db/ent/delivery"
	"github.com/TylerGrey/ent-shop/db/ent/order"
	"github.com/facebook/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// OrderDate holds the value of the "orderDate" field.
	OrderDate time.Time `json:"orderDate,omitempty"`
	// Status holds the value of the "status" field.
	Status order.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges       OrderEdges `json:"edges"`
	delivery_id *int
	member_id   *int
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// OrderItems holds the value of the orderItems edge.
	OrderItems []*OrderItem
	// Delivery holds the value of the delivery edge.
	Delivery *Delivery
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrderItemsOrErr returns the OrderItems value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) OrderItemsOrErr() ([]*OrderItem, error) {
	if e.loadedTypes[0] {
		return e.OrderItems, nil
	}
	return nil, &NotLoadedError{edge: "orderItems"}
}

// DeliveryOrErr returns the Delivery value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) DeliveryOrErr() (*Delivery, error) {
	if e.loadedTypes[1] {
		if e.Delivery == nil {
			// The edge delivery was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: delivery.Label}
		}
		return e.Delivery, nil
	}
	return nil, &NotLoadedError{edge: "delivery"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // orderDate
		&sql.NullString{}, // status
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Order) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // delivery_id
		&sql.NullInt64{}, // member_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(values ...interface{}) error {
	if m, n := len(values), len(order.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		o.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		o.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field orderDate", values[2])
	} else if value.Valid {
		o.OrderDate = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[3])
	} else if value.Valid {
		o.Status = order.Status(value.String)
	}
	values = values[4:]
	if len(values) == len(order.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field delivery_id", value)
		} else if value.Valid {
			o.delivery_id = new(int)
			*o.delivery_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field member_id", value)
		} else if value.Valid {
			o.member_id = new(int)
			*o.member_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOrderItems queries the orderItems edge of the Order.
func (o *Order) QueryOrderItems() *OrderItemQuery {
	return (&OrderClient{config: o.config}).QueryOrderItems(o)
}

// QueryDelivery queries the delivery edge of the Order.
func (o *Order) QueryDelivery() *DeliveryQuery {
	return (&OrderClient{config: o.config}).QueryDelivery(o)
}

// Update returns a builder for updating this Order.
// Note that, you need to call Order.Unwrap() before calling this method, if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", orderDate=")
	builder.WriteString(o.OrderDate.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}