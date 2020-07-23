package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Film struct {
	ent.Schema
}

func (Film) Mixin() []ent.Mixin {
	return []ent.Mixin{
		GeneralMixin{},
		TimeMixin{},
	}
}

func (Film) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description"),
	}
}

func (Film) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("authors", Author.Type),
	}
}