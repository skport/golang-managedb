package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(30)",
			}),
		field.String("last_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(30)",
			}),
		field.String("mail").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(50)",
			}),
		field.Int("usergroup_id"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
