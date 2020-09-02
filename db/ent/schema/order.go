package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Config ...
func (Order) Config() ent.Config {
	return ent.Config{
		Table: "orders",
	}
}

// Mixin ...
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("orderDate"),
		field.Enum("status").Values("ORDER", "CANCEL"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orderItems", OrderItem.Type).
			StorageKey(edge.Column("order_id")),
		edge.From("delivery", Delivery.Type).
			Ref("order").
			Unique(),
	}
}
