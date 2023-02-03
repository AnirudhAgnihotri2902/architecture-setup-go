package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Registers struct {
	ent.Schema
}

// Fields of the User.
func (Registers) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password"),
	}
}

// Edges of the User.
func (Registers) Edges() []ent.Edge {
	return nil
}
