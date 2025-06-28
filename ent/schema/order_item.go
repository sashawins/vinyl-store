package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("order_id", uuid.UUID{}),
		field.UUID("vinyl_id", uuid.UUID{}),
		field.Int("quantity").Default(1),
		field.Float("unit_price"),
	}
}
