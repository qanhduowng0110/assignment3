// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AssociatedEvent struct {
	ent.Schema
}

func (AssociatedEvent) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("earthquake_id").Optional(), field.Time("created_at").Optional(), field.Time("updated_at").Optional()}
}
func (AssociatedEvent) Edges() []ent.Edge {
	return []ent.Edge{edge.From("earthquake", Earthquake.Type).Ref("associated_events").Unique().Field("earthquake_id")}
}
func (AssociatedEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "AssociatedEvent"}}
}
