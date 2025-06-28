package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Vinyl struct {
	ent.Schema
}

func (Vinyl) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(), // уникальный идентификатор

		field.String("title").NotEmpty(),
		field.Float("price").Positive(),
		field.Int32("stock_count").NonNegative(),
		field.String("cover_url").Optional(),

		field.UUID("artist_id", uuid.UUID{}),
		field.UUID("genre_id", uuid.UUID{}),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}
