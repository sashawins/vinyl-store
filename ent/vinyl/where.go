// Code generated by ent, DO NOT EDIT.

package vinyl

import (
	"time"
	"vinyl-store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldTitle, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldPrice, v))
}

// StockCount applies equality check predicate on the "stock_count" field. It's identical to StockCountEQ.
func StockCount(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldStockCount, v))
}

// CoverURL applies equality check predicate on the "cover_url" field. It's identical to CoverURLEQ.
func CoverURL(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldCoverURL, v))
}

// ArtistID applies equality check predicate on the "artist_id" field. It's identical to ArtistIDEQ.
func ArtistID(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldArtistID, v))
}

// GenreID applies equality check predicate on the "genre_id" field. It's identical to GenreIDEQ.
func GenreID(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldGenreID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldCreatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldContainsFold(FieldTitle, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldPrice, v))
}

// StockCountEQ applies the EQ predicate on the "stock_count" field.
func StockCountEQ(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldStockCount, v))
}

// StockCountNEQ applies the NEQ predicate on the "stock_count" field.
func StockCountNEQ(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldStockCount, v))
}

// StockCountIn applies the In predicate on the "stock_count" field.
func StockCountIn(vs ...int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldStockCount, vs...))
}

// StockCountNotIn applies the NotIn predicate on the "stock_count" field.
func StockCountNotIn(vs ...int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldStockCount, vs...))
}

// StockCountGT applies the GT predicate on the "stock_count" field.
func StockCountGT(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldStockCount, v))
}

// StockCountGTE applies the GTE predicate on the "stock_count" field.
func StockCountGTE(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldStockCount, v))
}

// StockCountLT applies the LT predicate on the "stock_count" field.
func StockCountLT(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldStockCount, v))
}

// StockCountLTE applies the LTE predicate on the "stock_count" field.
func StockCountLTE(v int32) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldStockCount, v))
}

// CoverURLEQ applies the EQ predicate on the "cover_url" field.
func CoverURLEQ(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldCoverURL, v))
}

// CoverURLNEQ applies the NEQ predicate on the "cover_url" field.
func CoverURLNEQ(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldCoverURL, v))
}

// CoverURLIn applies the In predicate on the "cover_url" field.
func CoverURLIn(vs ...string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldCoverURL, vs...))
}

// CoverURLNotIn applies the NotIn predicate on the "cover_url" field.
func CoverURLNotIn(vs ...string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldCoverURL, vs...))
}

// CoverURLGT applies the GT predicate on the "cover_url" field.
func CoverURLGT(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldCoverURL, v))
}

// CoverURLGTE applies the GTE predicate on the "cover_url" field.
func CoverURLGTE(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldCoverURL, v))
}

// CoverURLLT applies the LT predicate on the "cover_url" field.
func CoverURLLT(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldCoverURL, v))
}

// CoverURLLTE applies the LTE predicate on the "cover_url" field.
func CoverURLLTE(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldCoverURL, v))
}

// CoverURLContains applies the Contains predicate on the "cover_url" field.
func CoverURLContains(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldContains(FieldCoverURL, v))
}

// CoverURLHasPrefix applies the HasPrefix predicate on the "cover_url" field.
func CoverURLHasPrefix(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldHasPrefix(FieldCoverURL, v))
}

// CoverURLHasSuffix applies the HasSuffix predicate on the "cover_url" field.
func CoverURLHasSuffix(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldHasSuffix(FieldCoverURL, v))
}

// CoverURLIsNil applies the IsNil predicate on the "cover_url" field.
func CoverURLIsNil() predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIsNull(FieldCoverURL))
}

// CoverURLNotNil applies the NotNil predicate on the "cover_url" field.
func CoverURLNotNil() predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotNull(FieldCoverURL))
}

// CoverURLEqualFold applies the EqualFold predicate on the "cover_url" field.
func CoverURLEqualFold(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEqualFold(FieldCoverURL, v))
}

// CoverURLContainsFold applies the ContainsFold predicate on the "cover_url" field.
func CoverURLContainsFold(v string) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldContainsFold(FieldCoverURL, v))
}

// ArtistIDEQ applies the EQ predicate on the "artist_id" field.
func ArtistIDEQ(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldArtistID, v))
}

// ArtistIDNEQ applies the NEQ predicate on the "artist_id" field.
func ArtistIDNEQ(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldArtistID, v))
}

// ArtistIDIn applies the In predicate on the "artist_id" field.
func ArtistIDIn(vs ...uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldArtistID, vs...))
}

// ArtistIDNotIn applies the NotIn predicate on the "artist_id" field.
func ArtistIDNotIn(vs ...uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldArtistID, vs...))
}

// ArtistIDGT applies the GT predicate on the "artist_id" field.
func ArtistIDGT(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldArtistID, v))
}

// ArtistIDGTE applies the GTE predicate on the "artist_id" field.
func ArtistIDGTE(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldArtistID, v))
}

// ArtistIDLT applies the LT predicate on the "artist_id" field.
func ArtistIDLT(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldArtistID, v))
}

// ArtistIDLTE applies the LTE predicate on the "artist_id" field.
func ArtistIDLTE(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldArtistID, v))
}

// GenreIDEQ applies the EQ predicate on the "genre_id" field.
func GenreIDEQ(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldGenreID, v))
}

// GenreIDNEQ applies the NEQ predicate on the "genre_id" field.
func GenreIDNEQ(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldGenreID, v))
}

// GenreIDIn applies the In predicate on the "genre_id" field.
func GenreIDIn(vs ...uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldGenreID, vs...))
}

// GenreIDNotIn applies the NotIn predicate on the "genre_id" field.
func GenreIDNotIn(vs ...uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldGenreID, vs...))
}

// GenreIDGT applies the GT predicate on the "genre_id" field.
func GenreIDGT(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldGenreID, v))
}

// GenreIDGTE applies the GTE predicate on the "genre_id" field.
func GenreIDGTE(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldGenreID, v))
}

// GenreIDLT applies the LT predicate on the "genre_id" field.
func GenreIDLT(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldGenreID, v))
}

// GenreIDLTE applies the LTE predicate on the "genre_id" field.
func GenreIDLTE(v uuid.UUID) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldGenreID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Vinyl {
	return predicate.Vinyl(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vinyl) predicate.Vinyl {
	return predicate.Vinyl(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vinyl) predicate.Vinyl {
	return predicate.Vinyl(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vinyl) predicate.Vinyl {
	return predicate.Vinyl(sql.NotPredicates(p))
}
