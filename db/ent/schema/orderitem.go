package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Mixin ...
func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("orderPrice"),
		field.Int32("count"),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("orderItems").
			Unique(),
	}
}
