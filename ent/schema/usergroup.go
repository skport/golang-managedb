package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id"),
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(30)",
			}),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return nil
}
