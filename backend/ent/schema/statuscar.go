package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Statuscar holds the schema definition for the Statuscar entity.
type Statuscar struct {
	ent.Schema
}

// Fields of the Statuscar.
func (Statuscar) Fields() []ent.Field {
	return []ent.Field{
		field.String("status_detail").NotEmpty(),
	}
}

// Edges of the Statuscar.
func (Statuscar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statuscar_predicament", Predicament.Type).StorageKey(edge.Column("status_id")),
	}
}
