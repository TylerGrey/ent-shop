package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Config ...
func (Item) Config() ent.Config {
	return ent.Config{
		Table: "item",
	}
}

// Mixin ...
func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int32("price"),
		field.Int32("stockQuantity"),
		field.Enum("dtype").
			Values("ALBUM", "BOOK", "MOVIE"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("items"),
	}
}
