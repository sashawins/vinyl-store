// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"vinyl-store/ent/orderitem"
	"vinyl-store/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrderItemQuery is the builder for querying OrderItem entities.
type OrderItemQuery struct {
	config
	ctx        *QueryContext
	order      []orderitem.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderItemQuery builder.
func (oiq *OrderItemQuery) Where(ps ...predicate.OrderItem) *OrderItemQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit the number of records to be returned by this query.
func (oiq *OrderItemQuery) Limit(limit int) *OrderItemQuery {
	oiq.ctx.Limit = &limit
	return oiq
}

// Offset to start from.
func (oiq *OrderItemQuery) Offset(offset int) *OrderItemQuery {
	oiq.ctx.Offset = &offset
	return oiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oiq *OrderItemQuery) Unique(unique bool) *OrderItemQuery {
	oiq.ctx.Unique = &unique
	return oiq
}

// Order specifies how the records should be ordered.
func (oiq *OrderItemQuery) Order(o ...orderitem.OrderOption) *OrderItemQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// First returns the first OrderItem entity from the query.
// Returns a *NotFoundError when no OrderItem was found.
func (oiq *OrderItemQuery) First(ctx context.Context) (*OrderItem, error) {
	nodes, err := oiq.Limit(1).All(setContextOp(ctx, oiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrderItemQuery) FirstX(ctx context.Context) *OrderItem {
	node, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderItem ID from the query.
// Returns a *NotFoundError when no OrderItem ID was found.
func (oiq *OrderItemQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oiq.Limit(1).IDs(setContextOp(ctx, oiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oiq *OrderItemQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderItem entity is found.
// Returns a *NotFoundError when no OrderItem entities are found.
func (oiq *OrderItemQuery) Only(ctx context.Context) (*OrderItem, error) {
	nodes, err := oiq.Limit(2).All(setContextOp(ctx, oiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderitem.Label}
	default:
		return nil, &NotSingularError{orderitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrderItemQuery) OnlyX(ctx context.Context) *OrderItem {
	node, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderItem ID in the query.
// Returns a *NotSingularError when more than one OrderItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (oiq *OrderItemQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oiq.Limit(2).IDs(setContextOp(ctx, oiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = &NotSingularError{orderitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oiq *OrderItemQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderItems.
func (oiq *OrderItemQuery) All(ctx context.Context) ([]*OrderItem, error) {
	ctx = setContextOp(ctx, oiq.ctx, ent.OpQueryAll)
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderItem, *OrderItemQuery]()
	return withInterceptors[[]*OrderItem](ctx, oiq, qr, oiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrderItemQuery) AllX(ctx context.Context) []*OrderItem {
	nodes, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderItem IDs.
func (oiq *OrderItemQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if oiq.ctx.Unique == nil && oiq.path != nil {
		oiq.Unique(true)
	}
	ctx = setContextOp(ctx, oiq.ctx, ent.OpQueryIDs)
	if err = oiq.Select(orderitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrderItemQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrderItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oiq.ctx, ent.OpQueryCount)
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oiq, querierCount[*OrderItemQuery](), oiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oiq *OrderItemQuery) CountX(ctx context.Context) int {
	count, err := oiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oiq *OrderItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oiq.ctx, ent.OpQueryExist)
	switch _, err := oiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrderItemQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrderItemQuery) Clone() *OrderItemQuery {
	if oiq == nil {
		return nil
	}
	return &OrderItemQuery{
		config:     oiq.config,
		ctx:        oiq.ctx.Clone(),
		order:      append([]orderitem.OrderOption{}, oiq.order...),
		inters:     append([]Interceptor{}, oiq.inters...),
		predicates: append([]predicate.OrderItem{}, oiq.predicates...),
		// clone intermediate query.
		sql:  oiq.sql.Clone(),
		path: oiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID uuid.UUID `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderItem.Query().
//		GroupBy(orderitem.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oiq *OrderItemQuery) GroupBy(field string, fields ...string) *OrderItemGroupBy {
	oiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderItemGroupBy{build: oiq}
	grbuild.flds = &oiq.ctx.Fields
	grbuild.label = orderitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID uuid.UUID `json:"order_id,omitempty"`
//	}
//
//	client.OrderItem.Query().
//		Select(orderitem.FieldOrderID).
//		Scan(ctx, &v)
func (oiq *OrderItemQuery) Select(fields ...string) *OrderItemSelect {
	oiq.ctx.Fields = append(oiq.ctx.Fields, fields...)
	sbuild := &OrderItemSelect{OrderItemQuery: oiq}
	sbuild.label = orderitem.Label
	sbuild.flds, sbuild.scan = &oiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderItemSelect configured with the given aggregations.
func (oiq *OrderItemQuery) Aggregate(fns ...AggregateFunc) *OrderItemSelect {
	return oiq.Select().Aggregate(fns...)
}

func (oiq *OrderItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oiq); err != nil {
				return err
			}
		}
	}
	for _, f := range oiq.ctx.Fields {
		if !orderitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrderItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderItem, error) {
	var (
		nodes = []*OrderItem{}
		_spec = oiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderItem{config: oiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oiq *OrderItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	_spec.Node.Columns = oiq.ctx.Fields
	if len(oiq.ctx.Fields) > 0 {
		_spec.Unique = oiq.ctx.Unique != nil && *oiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrderItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeUUID))
	_spec.From = oiq.sql
	if unique := oiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oiq.path != nil {
		_spec.Unique = true
	}
	if fields := oiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for i := range fields {
			if fields[i] != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oiq *OrderItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orderitem.Table)
	columns := oiq.ctx.Fields
	if len(columns) == 0 {
		columns = orderitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oiq.ctx.Unique != nil && *oiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderItemGroupBy is the group-by builder for OrderItem entities.
type OrderItemGroupBy struct {
	selector
	build *OrderItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrderItemGroupBy) Aggregate(fns ...AggregateFunc) *OrderItemGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the selector query and scans the result into the given value.
func (oigb *OrderItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oigb.build.ctx, ent.OpQueryGroupBy)
	if err := oigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderItemQuery, *OrderItemGroupBy](ctx, oigb.build, oigb, oigb.build.inters, v)
}

func (oigb *OrderItemGroupBy) sqlScan(ctx context.Context, root *OrderItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oigb.fns))
	for _, fn := range oigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oigb.flds)+len(oigb.fns))
		for _, f := range *oigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderItemSelect is the builder for selecting fields of OrderItem entities.
type OrderItemSelect struct {
	*OrderItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ois *OrderItemSelect) Aggregate(fns ...AggregateFunc) *OrderItemSelect {
	ois.fns = append(ois.fns, fns...)
	return ois
}

// Scan applies the selector query and scans the result into the given value.
func (ois *OrderItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ois.ctx, ent.OpQuerySelect)
	if err := ois.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderItemQuery, *OrderItemSelect](ctx, ois.OrderItemQuery, ois, ois.inters, v)
}

func (ois *OrderItemSelect) sqlScan(ctx context.Context, root *OrderItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ois.fns))
	for _, fn := range ois.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ois.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
