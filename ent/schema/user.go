// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.String("username").Optional(), field.String("password").Optional(), field.Time("created_at").Optional(), field.Int("role_id").Optional()}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{edge.From("role", Role.Type).Ref("users").Unique().Field("role_id")}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Users"}}
}
