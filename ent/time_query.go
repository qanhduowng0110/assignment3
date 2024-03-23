// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/earthquake"
	"assignment3/ent/predicate"
	enttime "assignment3/ent/time"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimeQuery is the builder for querying Time entities.
type TimeQuery struct {
	config
	ctx             *QueryContext
	order           []enttime.OrderOption
	inters          []Interceptor
	predicates      []predicate.Time
	withEarthquakes *EarthquakeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TimeQuery builder.
func (tq *TimeQuery) Where(ps ...predicate.Time) *TimeQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TimeQuery) Limit(limit int) *TimeQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TimeQuery) Offset(offset int) *TimeQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TimeQuery) Unique(unique bool) *TimeQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TimeQuery) Order(o ...enttime.OrderOption) *TimeQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryEarthquakes chains the current query on the "earthquakes" edge.
func (tq *TimeQuery) QueryEarthquakes() *EarthquakeQuery {
	query := (&EarthquakeClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(enttime.Table, enttime.FieldID, selector),
			sqlgraph.To(earthquake.Table, earthquake.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, enttime.EarthquakesTable, enttime.EarthquakesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Time entity from the query.
// Returns a *NotFoundError when no Time was found.
func (tq *TimeQuery) First(ctx context.Context) (*Time, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enttime.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TimeQuery) FirstX(ctx context.Context) *Time {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Time ID from the query.
// Returns a *NotFoundError when no Time ID was found.
func (tq *TimeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enttime.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TimeQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Time entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Time entity is found.
// Returns a *NotFoundError when no Time entities are found.
func (tq *TimeQuery) Only(ctx context.Context) (*Time, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enttime.Label}
	default:
		return nil, &NotSingularError{enttime.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TimeQuery) OnlyX(ctx context.Context) *Time {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Time ID in the query.
// Returns a *NotSingularError when more than one Time ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TimeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enttime.Label}
	default:
		err = &NotSingularError{enttime.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TimeQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Times.
func (tq *TimeQuery) All(ctx context.Context) ([]*Time, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Time, *TimeQuery]()
	return withInterceptors[[]*Time](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TimeQuery) AllX(ctx context.Context) []*Time {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Time IDs.
func (tq *TimeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(enttime.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TimeQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TimeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TimeQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TimeQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TimeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TimeQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TimeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TimeQuery) Clone() *TimeQuery {
	if tq == nil {
		return nil
	}
	return &TimeQuery{
		config:          tq.config,
		ctx:             tq.ctx.Clone(),
		order:           append([]enttime.OrderOption{}, tq.order...),
		inters:          append([]Interceptor{}, tq.inters...),
		predicates:      append([]predicate.Time{}, tq.predicates...),
		withEarthquakes: tq.withEarthquakes.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithEarthquakes tells the query-builder to eager-load the nodes that are connected to
// the "earthquakes" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimeQuery) WithEarthquakes(opts ...func(*EarthquakeQuery)) *TimeQuery {
	query := (&EarthquakeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withEarthquakes = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DateTime time.Time `json:"date_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Time.Query().
//		GroupBy(enttime.FieldDateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TimeQuery) GroupBy(field string, fields ...string) *TimeGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TimeGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = enttime.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DateTime time.Time `json:"date_time,omitempty"`
//	}
//
//	client.Time.Query().
//		Select(enttime.FieldDateTime).
//		Scan(ctx, &v)
func (tq *TimeQuery) Select(fields ...string) *TimeSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TimeSelect{TimeQuery: tq}
	sbuild.label = enttime.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TimeSelect configured with the given aggregations.
func (tq *TimeQuery) Aggregate(fns ...AggregateFunc) *TimeSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TimeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !enttime.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TimeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Time, error) {
	var (
		nodes       = []*Time{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withEarthquakes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Time).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Time{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withEarthquakes; query != nil {
		if err := tq.loadEarthquakes(ctx, query, nodes,
			func(n *Time) { n.Edges.Earthquakes = []*Earthquake{} },
			func(n *Time, e *Earthquake) { n.Edges.Earthquakes = append(n.Edges.Earthquakes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TimeQuery) loadEarthquakes(ctx context.Context, query *EarthquakeQuery, nodes []*Time, init func(*Time), assign func(*Time, *Earthquake)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Time)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(earthquake.FieldTimeID)
	}
	query.Where(predicate.Earthquake(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(enttime.EarthquakesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TimeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "time_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TimeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TimeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(enttime.Table, enttime.Columns, sqlgraph.NewFieldSpec(enttime.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enttime.FieldID)
		for i := range fields {
			if fields[i] != enttime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TimeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(enttime.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = enttime.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TimeGroupBy is the group-by builder for Time entities.
type TimeGroupBy struct {
	selector
	build *TimeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TimeGroupBy) Aggregate(fns ...AggregateFunc) *TimeGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TimeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimeQuery, *TimeGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TimeGroupBy) sqlScan(ctx context.Context, root *TimeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TimeSelect is the builder for selecting fields of Time entities.
type TimeSelect struct {
	*TimeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TimeSelect) Aggregate(fns ...AggregateFunc) *TimeSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TimeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimeQuery, *TimeSelect](ctx, ts.TimeQuery, ts, ts.inters, v)
}

func (ts *TimeSelect) sqlScan(ctx context.Context, root *TimeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
