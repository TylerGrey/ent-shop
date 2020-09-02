package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Config ...
func (Member) Config() ent.Config {
	return ent.Config{
		Table: "member",
	}
}

// Mixin ...
func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AddressMixin{},
		TimeMixin{},
	}
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type).
			StorageKey(edge.Column("member_id")),
	}
}
