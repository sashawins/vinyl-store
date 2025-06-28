package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type News struct {
	ent.Schema
}

func (News) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("title").NotEmpty(),
		field.String("content").NotEmpty(),
		field.Time("published_at").Default(time.Now),
	}
}
