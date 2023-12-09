package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LineItem holds the schema definition for the LineItem entity.
type LineItem struct {
	ent.Schema
}

// Fields of the LineItem.
func (LineItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("name").
			MinLen(1).
			MaxLen(50),
		field.Int64("price").
			Min(0),
		field.Int64("quantity").
			Min(1),
		field.Time("added_at").
			Default(time.Now),
	}
}

// Edges of the LineItem.
func (LineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Bill.Type).
			Ref("line_items").
			Unique(),
	}
}
