package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserSegment holds the schema definition for the UserSegment entity.
type UserSegment struct {
	ent.Schema
}

// Fields of the UserSegment.
func (UserSegment) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.String("segment"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UserSegment.
func (UserSegment) Edges() []ent.Edge {
	return nil
}
