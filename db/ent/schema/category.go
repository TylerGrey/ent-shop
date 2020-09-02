package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Config ...
func (Category) Config() ent.Config {
	return ent.Config{
		Table: "category",
	}
}

// Mixin ...
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Category.Type).
			StorageKey(edge.Column("parent_id")).
			From("parent").
			Unique(),
		edge.To("items", Item.Type).
			StorageKey(edge.Table("category_item")),
	}
}
