package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Delivery holds the schema definition for the Delivery entity.
type Delivery struct {
	ent.Schema
}

// Config ...
func (Delivery) Config() ent.Config {
	return ent.Config{
		Table: "delivery",
	}
}

// Mixin ...
func (Delivery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AddressMixin{},
		TimeMixin{},
	}
}

// Fields of the Delivery.
func (Delivery) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("READY", "COMP"),
	}
}

// Edges of the Delivery.
func (Delivery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			StorageKey(edge.Column("delivery_id")).
			Unique(),
	}
}
