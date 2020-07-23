package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Author struct {
	ent.Schema
}

func (Author) Mixin() []ent.Mixin {
	return []ent.Mixin{
		GeneralMixin{},
		TimeMixin{},
	}
}

func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.String("fullName"),
		field.Int8("age"),
	}
}

func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("films", Film.Type).Ref("authors"),
	}
}