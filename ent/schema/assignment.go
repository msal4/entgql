package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Assignment holds the schema definition for the Assignment entity.
type Assignment struct {
	ent.Schema
}

// Fields of the Assignment.
func (Assignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bool("is_exam").Default(false),
		field.Time("due_date").Annotations(entgql.OrderField("DUE_DATE")),
		field.Int64("duration").GoType(time.Duration(0)).Optional().
			Annotations(entgql.OrderField("DURATION"), entgql.Type("Duration")),
		field.Time("deleted_at").Optional().Nillable(),
	}
}
