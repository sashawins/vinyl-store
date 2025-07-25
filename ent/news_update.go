// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"vinyl-store/ent/news"
	"vinyl-store/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsUpdate is the builder for updating News entities.
type NewsUpdate struct {
	config
	hooks    []Hook
	mutation *NewsMutation
}

// Where appends a list predicates to the NewsUpdate builder.
func (nu *NewsUpdate) Where(ps ...predicate.News) *NewsUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetTitle sets the "title" field.
func (nu *NewsUpdate) SetTitle(s string) *NewsUpdate {
	nu.mutation.SetTitle(s)
	return nu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nu *NewsUpdate) SetNillableTitle(s *string) *NewsUpdate {
	if s != nil {
		nu.SetTitle(*s)
	}
	return nu
}

// SetContent sets the "content" field.
func (nu *NewsUpdate) SetContent(s string) *NewsUpdate {
	nu.mutation.SetContent(s)
	return nu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (nu *NewsUpdate) SetNillableContent(s *string) *NewsUpdate {
	if s != nil {
		nu.SetContent(*s)
	}
	return nu
}

// SetPublishedAt sets the "published_at" field.
func (nu *NewsUpdate) SetPublishedAt(t time.Time) *NewsUpdate {
	nu.mutation.SetPublishedAt(t)
	return nu
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (nu *NewsUpdate) SetNillablePublishedAt(t *time.Time) *NewsUpdate {
	if t != nil {
		nu.SetPublishedAt(*t)
	}
	return nu
}

// Mutation returns the NewsMutation object of the builder.
func (nu *NewsUpdate) Mutation() *NewsMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NewsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NewsUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NewsUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NewsUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NewsUpdate) check() error {
	if v, ok := nu.mutation.Title(); ok {
		if err := news.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "News.title": %w`, err)}
		}
	}
	if v, ok := nu.mutation.Content(); ok {
		if err := news.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "News.content": %w`, err)}
		}
	}
	return nil
}

func (nu *NewsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(news.Table, news.Columns, sqlgraph.NewFieldSpec(news.FieldID, field.TypeUUID))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Title(); ok {
		_spec.SetField(news.FieldTitle, field.TypeString, value)
	}
	if value, ok := nu.mutation.Content(); ok {
		_spec.SetField(news.FieldContent, field.TypeString, value)
	}
	if value, ok := nu.mutation.PublishedAt(); ok {
		_spec.SetField(news.FieldPublishedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{news.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NewsUpdateOne is the builder for updating a single News entity.
type NewsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewsMutation
}

// SetTitle sets the "title" field.
func (nuo *NewsUpdateOne) SetTitle(s string) *NewsUpdateOne {
	nuo.mutation.SetTitle(s)
	return nuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nuo *NewsUpdateOne) SetNillableTitle(s *string) *NewsUpdateOne {
	if s != nil {
		nuo.SetTitle(*s)
	}
	return nuo
}

// SetContent sets the "content" field.
func (nuo *NewsUpdateOne) SetContent(s string) *NewsUpdateOne {
	nuo.mutation.SetContent(s)
	return nuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (nuo *NewsUpdateOne) SetNillableContent(s *string) *NewsUpdateOne {
	if s != nil {
		nuo.SetContent(*s)
	}
	return nuo
}

// SetPublishedAt sets the "published_at" field.
func (nuo *NewsUpdateOne) SetPublishedAt(t time.Time) *NewsUpdateOne {
	nuo.mutation.SetPublishedAt(t)
	return nuo
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (nuo *NewsUpdateOne) SetNillablePublishedAt(t *time.Time) *NewsUpdateOne {
	if t != nil {
		nuo.SetPublishedAt(*t)
	}
	return nuo
}

// Mutation returns the NewsMutation object of the builder.
func (nuo *NewsUpdateOne) Mutation() *NewsMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NewsUpdate builder.
func (nuo *NewsUpdateOne) Where(ps ...predicate.News) *NewsUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NewsUpdateOne) Select(field string, fields ...string) *NewsUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated News entity.
func (nuo *NewsUpdateOne) Save(ctx context.Context) (*News, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NewsUpdateOne) SaveX(ctx context.Context) *News {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NewsUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NewsUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NewsUpdateOne) check() error {
	if v, ok := nuo.mutation.Title(); ok {
		if err := news.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "News.title": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.Content(); ok {
		if err := news.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "News.content": %w`, err)}
		}
	}
	return nil
}

func (nuo *NewsUpdateOne) sqlSave(ctx context.Context) (_node *News, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(news.Table, news.Columns, sqlgraph.NewFieldSpec(news.FieldID, field.TypeUUID))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "News.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, news.FieldID)
		for _, f := range fields {
			if !news.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != news.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Title(); ok {
		_spec.SetField(news.FieldTitle, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Content(); ok {
		_spec.SetField(news.FieldContent, field.TypeString, value)
	}
	if value, ok := nuo.mutation.PublishedAt(); ok {
		_spec.SetField(news.FieldPublishedAt, field.TypeTime, value)
	}
	_node = &News{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{news.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
