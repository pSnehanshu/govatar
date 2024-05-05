package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Email holds the schema definition for the Email entity.
type Email struct {
	ent.Schema
}

// Fields of the Email.
func (Email) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Match(regexp.MustCompile(`^[\w-\.+]+@([\w-]+\.)+\w+$`)).Unique(),
		field.Bool("is_verified").Default(false),
		field.String("shasum").Unique().Immutable(),
		field.Time("created_at").Default(time.Now),
		field.String("user_id"),
	}
}

// Edges of the Email.
func (Email) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("avatar", Avatar.Type).
			Unique().
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.From("user", User.Type).
			Ref("emails").
			Field("user_id").
			Unique().
			Required(),
	}
}
