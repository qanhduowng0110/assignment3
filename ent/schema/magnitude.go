// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Magnitude struct {
	ent.Schema
}

func (Magnitude) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Float("magnitude_value").Optional(), field.String("magnitude_type").Optional(), field.Time("created_at").Optional(), field.Time("updated_at").Optional()}
}
func (Magnitude) Edges() []ent.Edge {
	return []ent.Edge{edge.To("earthquakes", Earthquake.Type)}
}
func (Magnitude) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Magnitudes"}}
}
