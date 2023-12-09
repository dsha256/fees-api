package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.Enum("currency").
			Values("USD").
			Default("USD"),
		field.Int64("total").
			Default(0),
		field.Bool("is_open").
			Default(true),
		field.Time("closed_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("line_items", LineItem.Type),
	}
}
