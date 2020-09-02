package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// AddressMixin holds the schema definition for the AddressMixin entity.
type AddressMixin struct {
	mixin.Schema
}

// Fields of the AddressMixin.
func (AddressMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("city").Optional(),
		field.String("street").Optional(),
		field.String("zipcode").Optional(),
	}
}

// Edges of the AddressMixin.
func (AddressMixin) Edges() []ent.Edge {
	return nil
}
