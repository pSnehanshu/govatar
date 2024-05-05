package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Avatar holds the schema definition for the Avatar entity.
type Avatar struct {
	ent.Schema
}

// Fields of the Avatar.
func (Avatar) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.String("loc"),
		field.Enum("rating").Values("G", "PG", "R", "X").Default("G"),
		field.UUID("email_id", uuid.UUID{}),
	}
}

// Edges of the Avatar.
func (Avatar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("email", Email.Type).
			Ref("avatar").
			Field("email_id").
			Unique().
			Required(),
	}
}
