// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Rustixir/go-challenge/ent/predicate"
	"github.com/Rustixir/go-challenge/ent/usersegment"
)

// UserSegmentQuery is the builder for querying UserSegment entities.
type UserSegmentQuery struct {
	config
	ctx        *QueryContext
	order      []usersegment.OrderOption
	inters     []Interceptor
	predicates []predicate.UserSegment
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserSegmentQuery builder.
func (usq *UserSegmentQuery) Where(ps ...predicate.UserSegment) *UserSegmentQuery {
	usq.predicates = append(usq.predicates, ps...)
	return usq
}

// Limit the number of records to be returned by this query.
func (usq *UserSegmentQuery) Limit(limit int) *UserSegmentQuery {
	usq.ctx.Limit = &limit
	return usq
}

// Offset to start from.
func (usq *UserSegmentQuery) Offset(offset int) *UserSegmentQuery {
	usq.ctx.Offset = &offset
	return usq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (usq *UserSegmentQuery) Unique(unique bool) *UserSegmentQuery {
	usq.ctx.Unique = &unique
	return usq
}

// Order specifies how the records should be ordered.
func (usq *UserSegmentQuery) Order(o ...usersegment.OrderOption) *UserSegmentQuery {
	usq.order = append(usq.order, o...)
	return usq
}

// First returns the first UserSegment entity from the query.
// Returns a *NotFoundError when no UserSegment was found.
func (usq *UserSegmentQuery) First(ctx context.Context) (*UserSegment, error) {
	nodes, err := usq.Limit(1).All(setContextOp(ctx, usq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usersegment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (usq *UserSegmentQuery) FirstX(ctx context.Context) *UserSegment {
	node, err := usq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserSegment ID from the query.
// Returns a *NotFoundError when no UserSegment ID was found.
func (usq *UserSegmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(1).IDs(setContextOp(ctx, usq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usersegment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (usq *UserSegmentQuery) FirstIDX(ctx context.Context) int {
	id, err := usq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserSegment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserSegment entity is found.
// Returns a *NotFoundError when no UserSegment entities are found.
func (usq *UserSegmentQuery) Only(ctx context.Context) (*UserSegment, error) {
	nodes, err := usq.Limit(2).All(setContextOp(ctx, usq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usersegment.Label}
	default:
		return nil, &NotSingularError{usersegment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (usq *UserSegmentQuery) OnlyX(ctx context.Context) *UserSegment {
	node, err := usq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserSegment ID in the query.
// Returns a *NotSingularError when more than one UserSegment ID is found.
// Returns a *NotFoundError when no entities are found.
func (usq *UserSegmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(2).IDs(setContextOp(ctx, usq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usersegment.Label}
	default:
		err = &NotSingularError{usersegment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (usq *UserSegmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := usq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserSegments.
func (usq *UserSegmentQuery) All(ctx context.Context) ([]*UserSegment, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryAll)
	if err := usq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserSegment, *UserSegmentQuery]()
	return withInterceptors[[]*UserSegment](ctx, usq, qr, usq.inters)
}

// AllX is like All, but panics if an error occurs.
func (usq *UserSegmentQuery) AllX(ctx context.Context) []*UserSegment {
	nodes, err := usq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserSegment IDs.
func (usq *UserSegmentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if usq.ctx.Unique == nil && usq.path != nil {
		usq.Unique(true)
	}
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryIDs)
	if err = usq.Select(usersegment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (usq *UserSegmentQuery) IDsX(ctx context.Context) []int {
	ids, err := usq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (usq *UserSegmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryCount)
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, usq, querierCount[*UserSegmentQuery](), usq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (usq *UserSegmentQuery) CountX(ctx context.Context) int {
	count, err := usq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (usq *UserSegmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryExist)
	switch _, err := usq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (usq *UserSegmentQuery) ExistX(ctx context.Context) bool {
	exist, err := usq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserSegmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (usq *UserSegmentQuery) Clone() *UserSegmentQuery {
	if usq == nil {
		return nil
	}
	return &UserSegmentQuery{
		config:     usq.config,
		ctx:        usq.ctx.Clone(),
		order:      append([]usersegment.OrderOption{}, usq.order...),
		inters:     append([]Interceptor{}, usq.inters...),
		predicates: append([]predicate.UserSegment{}, usq.predicates...),
		// clone intermediate query.
		sql:  usq.sql.Clone(),
		path: usq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserSegment.Query().
//		GroupBy(usersegment.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (usq *UserSegmentQuery) GroupBy(field string, fields ...string) *UserSegmentGroupBy {
	usq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserSegmentGroupBy{build: usq}
	grbuild.flds = &usq.ctx.Fields
	grbuild.label = usersegment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.UserSegment.Query().
//		Select(usersegment.FieldUserID).
//		Scan(ctx, &v)
func (usq *UserSegmentQuery) Select(fields ...string) *UserSegmentSelect {
	usq.ctx.Fields = append(usq.ctx.Fields, fields...)
	sbuild := &UserSegmentSelect{UserSegmentQuery: usq}
	sbuild.label = usersegment.Label
	sbuild.flds, sbuild.scan = &usq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSegmentSelect configured with the given aggregations.
func (usq *UserSegmentQuery) Aggregate(fns ...AggregateFunc) *UserSegmentSelect {
	return usq.Select().Aggregate(fns...)
}

func (usq *UserSegmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range usq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, usq); err != nil {
				return err
			}
		}
	}
	for _, f := range usq.ctx.Fields {
		if !usersegment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if usq.path != nil {
		prev, err := usq.path(ctx)
		if err != nil {
			return err
		}
		usq.sql = prev
	}
	return nil
}

func (usq *UserSegmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserSegment, error) {
	var (
		nodes = []*UserSegment{}
		_spec = usq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserSegment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserSegment{config: usq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, usq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (usq *UserSegmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := usq.querySpec()
	_spec.Node.Columns = usq.ctx.Fields
	if len(usq.ctx.Fields) > 0 {
		_spec.Unique = usq.ctx.Unique != nil && *usq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, usq.driver, _spec)
}

func (usq *UserSegmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usersegment.Table, usersegment.Columns, sqlgraph.NewFieldSpec(usersegment.FieldID, field.TypeInt))
	_spec.From = usq.sql
	if unique := usq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if usq.path != nil {
		_spec.Unique = true
	}
	if fields := usq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersegment.FieldID)
		for i := range fields {
			if fields[i] != usersegment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := usq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := usq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := usq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := usq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (usq *UserSegmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(usq.driver.Dialect())
	t1 := builder.Table(usersegment.Table)
	columns := usq.ctx.Fields
	if len(columns) == 0 {
		columns = usersegment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if usq.sql != nil {
		selector = usq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if usq.ctx.Unique != nil && *usq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range usq.predicates {
		p(selector)
	}
	for _, p := range usq.order {
		p(selector)
	}
	if offset := usq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := usq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserSegmentGroupBy is the group-by builder for UserSegment entities.
type UserSegmentGroupBy struct {
	selector
	build *UserSegmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (usgb *UserSegmentGroupBy) Aggregate(fns ...AggregateFunc) *UserSegmentGroupBy {
	usgb.fns = append(usgb.fns, fns...)
	return usgb
}

// Scan applies the selector query and scans the result into the given value.
func (usgb *UserSegmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, usgb.build.ctx, ent.OpQueryGroupBy)
	if err := usgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSegmentQuery, *UserSegmentGroupBy](ctx, usgb.build, usgb, usgb.build.inters, v)
}

func (usgb *UserSegmentGroupBy) sqlScan(ctx context.Context, root *UserSegmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(usgb.fns))
	for _, fn := range usgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*usgb.flds)+len(usgb.fns))
		for _, f := range *usgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*usgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := usgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSegmentSelect is the builder for selecting fields of UserSegment entities.
type UserSegmentSelect struct {
	*UserSegmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uss *UserSegmentSelect) Aggregate(fns ...AggregateFunc) *UserSegmentSelect {
	uss.fns = append(uss.fns, fns...)
	return uss
}

// Scan applies the selector query and scans the result into the given value.
func (uss *UserSegmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uss.ctx, ent.OpQuerySelect)
	if err := uss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSegmentQuery, *UserSegmentSelect](ctx, uss.UserSegmentQuery, uss, uss.inters, v)
}

func (uss *UserSegmentSelect) sqlScan(ctx context.Context, root *UserSegmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uss.fns))
	for _, fn := range uss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}