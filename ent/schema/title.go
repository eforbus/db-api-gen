// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Title struct {
	ent.Schema
}

func (Title) Fields() []ent.Field {
	return []ent.Field{field.String("id"), field.String("employee_id").Optional(), field.String("title"), field.Time("from_date"), field.Time("to_date"), field.Time("created_at").Optional(), field.Time("updated_at").Optional()}
}
func (Title) Edges() []ent.Edge {
	return []ent.Edge{edge.From("employee", Employee.Type).Ref("titles").Unique().Field("employee_id")}
}
func (Title) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "title"}}
}