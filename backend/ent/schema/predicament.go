package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Predicament holds the schema definition for the Predicament entity.
type Predicament struct {
	ent.Schema
}

// Fields of the Predicament.
func (Predicament) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Added_Time").Default(time.Now),
	}
}

// Edges of the Predicament.
func (Predicament) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Car", Car.Type).Ref("car_predicament").Unique(),
		edge.From("Statuscar", Statuscar.Type).Ref("statuscar_predicament").Unique(),
		edge.From("Staff", Staff.Type).Ref("staff_predicament").Unique(),
		edge.From("User", User.Type).Ref("user_predicament").Unique(),
	}
}
