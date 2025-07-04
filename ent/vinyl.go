// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"vinyl-store/ent/vinyl"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Vinyl is the model entity for the Vinyl schema.
type Vinyl struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// StockCount holds the value of the "stock_count" field.
	StockCount int32 `json:"stock_count,omitempty"`
	// CoverURL holds the value of the "cover_url" field.
	CoverURL string `json:"cover_url,omitempty"`
	// ArtistID holds the value of the "artist_id" field.
	ArtistID uuid.UUID `json:"artist_id,omitempty"`
	// GenreID holds the value of the "genre_id" field.
	GenreID uuid.UUID `json:"genre_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vinyl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vinyl.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case vinyl.FieldStockCount:
			values[i] = new(sql.NullInt64)
		case vinyl.FieldTitle, vinyl.FieldCoverURL:
			values[i] = new(sql.NullString)
		case vinyl.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case vinyl.FieldID, vinyl.FieldArtistID, vinyl.FieldGenreID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vinyl fields.
func (v *Vinyl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vinyl.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				v.ID = *value
			}
		case vinyl.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case vinyl.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				v.Price = value.Float64
			}
		case vinyl.FieldStockCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stock_count", values[i])
			} else if value.Valid {
				v.StockCount = int32(value.Int64)
			}
		case vinyl.FieldCoverURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_url", values[i])
			} else if value.Valid {
				v.CoverURL = value.String
			}
		case vinyl.FieldArtistID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field artist_id", values[i])
			} else if value != nil {
				v.ArtistID = *value
			}
		case vinyl.FieldGenreID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field genre_id", values[i])
			} else if value != nil {
				v.GenreID = *value
			}
		case vinyl.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vinyl.
// This includes values selected through modifiers, order, etc.
func (v *Vinyl) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// Update returns a builder for updating this Vinyl.
// Note that you need to call Vinyl.Unwrap() before calling this method if this Vinyl
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vinyl) Update() *VinylUpdateOne {
	return NewVinylClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vinyl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vinyl) Unwrap() *Vinyl {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vinyl is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vinyl) String() string {
	var builder strings.Builder
	builder.WriteString("Vinyl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("title=")
	builder.WriteString(v.Title)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", v.Price))
	builder.WriteString(", ")
	builder.WriteString("stock_count=")
	builder.WriteString(fmt.Sprintf("%v", v.StockCount))
	builder.WriteString(", ")
	builder.WriteString("cover_url=")
	builder.WriteString(v.CoverURL)
	builder.WriteString(", ")
	builder.WriteString("artist_id=")
	builder.WriteString(fmt.Sprintf("%v", v.ArtistID))
	builder.WriteString(", ")
	builder.WriteString("genre_id=")
	builder.WriteString(fmt.Sprintf("%v", v.GenreID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Vinyls is a parsable slice of Vinyl.
type Vinyls []*Vinyl
