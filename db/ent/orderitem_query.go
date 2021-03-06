// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/TylerGrey/ent-shop/db/ent/order"
	"github.com/TylerGrey/ent-shop/db/ent/orderitem"
	"github.com/TylerGrey/ent-shop/db/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// OrderItemQuery is the builder for querying OrderItem entities.
type OrderItemQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.OrderItem
	// eager-loading edges.
	withOrder *OrderQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (oiq *OrderItemQuery) Where(ps ...predicate.OrderItem) *OrderItemQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit adds a limit step to the query.
func (oiq *OrderItemQuery) Limit(limit int) *OrderItemQuery {
	oiq.limit = &limit
	return oiq
}

// Offset adds an offset step to the query.
func (oiq *OrderItemQuery) Offset(offset int) *OrderItemQuery {
	oiq.offset = &offset
	return oiq
}

// Order adds an order step to the query.
func (oiq *OrderItemQuery) Order(o ...OrderFunc) *OrderItemQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// QueryOrder chains the current query on the order edge.
func (oiq *OrderItemQuery) QueryOrder() *OrderQuery {
	query := &OrderQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderitem.Table, orderitem.FieldID, oiq.sqlQuery()),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderitem.OrderTable, orderitem.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderItem entity in the query. Returns *NotFoundError when no orderitem was found.
func (oiq *OrderItemQuery) First(ctx context.Context) (*OrderItem, error) {
	ois, err := oiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ois) == 0 {
		return nil, &NotFoundError{orderitem.Label}
	}
	return ois[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrderItemQuery) FirstX(ctx context.Context) *OrderItem {
	oi, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return oi
}

// FirstID returns the first OrderItem id in the query. Returns *NotFoundError when no id was found.
func (oiq *OrderItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderitem.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (oiq *OrderItemQuery) FirstXID(ctx context.Context) int {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only OrderItem entity in the query, returns an error if not exactly one entity was returned.
func (oiq *OrderItemQuery) Only(ctx context.Context) (*OrderItem, error) {
	ois, err := oiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ois) {
	case 1:
		return ois[0], nil
	case 0:
		return nil, &NotFoundError{orderitem.Label}
	default:
		return nil, &NotSingularError{orderitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrderItemQuery) OnlyX(ctx context.Context) *OrderItem {
	oi, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return oi
}

// OnlyID returns the only OrderItem id in the query, returns an error if not exactly one id was returned.
func (oiq *OrderItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oiq.Limit(2).IDs(ctx); err != nil {
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
func (oiq *OrderItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderItems.
func (oiq *OrderItemQuery) All(ctx context.Context) ([]*OrderItem, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrderItemQuery) AllX(ctx context.Context) []*OrderItem {
	ois, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ois
}

// IDs executes the query and returns a list of OrderItem ids.
func (oiq *OrderItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oiq.Select(orderitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrderItemQuery) IDsX(ctx context.Context) []int {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrderItemQuery) Count(ctx context.Context) (int, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oiq.sqlCount(ctx)
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
	if err := oiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrderItemQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrderItemQuery) Clone() *OrderItemQuery {
	return &OrderItemQuery{
		config:     oiq.config,
		limit:      oiq.limit,
		offset:     oiq.offset,
		order:      append([]OrderFunc{}, oiq.order...),
		unique:     append([]string{}, oiq.unique...),
		predicates: append([]predicate.OrderItem{}, oiq.predicates...),
		// clone intermediate query.
		sql:  oiq.sql.Clone(),
		path: oiq.path,
	}
}

//  WithOrder tells the query-builder to eager-loads the nodes that are connected to
// the "order" edge. The optional arguments used to configure the query builder of the edge.
func (oiq *OrderItemQuery) WithOrder(opts ...func(*OrderQuery)) *OrderItemQuery {
	query := &OrderQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withOrder = query
	return oiq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderItem.Query().
//		GroupBy(orderitem.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oiq *OrderItemQuery) GroupBy(field string, fields ...string) *OrderItemGroupBy {
	group := &OrderItemGroupBy{config: oiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oiq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OrderItem.Query().
//		Select(orderitem.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (oiq *OrderItemQuery) Select(field string, fields ...string) *OrderItemSelect {
	selector := &OrderItemSelect{config: oiq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oiq.sqlQuery(), nil
	}
	return selector
}

func (oiq *OrderItemQuery) prepareQuery(ctx context.Context) error {
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrderItemQuery) sqlAll(ctx context.Context) ([]*OrderItem, error) {
	var (
		nodes       = []*OrderItem{}
		withFKs     = oiq.withFKs
		_spec       = oiq.querySpec()
		loadedTypes = [1]bool{
			oiq.withOrder != nil,
		}
	)
	if oiq.withOrder != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &OrderItem{config: oiq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oiq.withOrder; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderItem)
		for i := range nodes {
			if fk := nodes[i].order_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(order.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Order = n
			}
		}
	}

	return nodes, nil
}

func (oiq *OrderItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrderItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (oiq *OrderItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderitem.Table,
			Columns: orderitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderitem.FieldID,
			},
		},
		From:   oiq.sql,
		Unique: true,
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.offset; offset != nil {
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

func (oiq *OrderItemQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orderitem.Table)
	selector := builder.Select(t1.Columns(orderitem.Columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(orderitem.Columns...)...)
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderItemGroupBy is the builder for group-by OrderItem entities.
type OrderItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrderItemGroupBy) Aggregate(fns ...AggregateFunc) *OrderItemGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the group-by query and scan the result into the given value.
func (oigb *OrderItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oigb.path(ctx)
	if err != nil {
		return err
	}
	oigb.sql = query
	return oigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oigb *OrderItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oigb *OrderItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := oigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oigb *OrderItemGroupBy) StringX(ctx context.Context) string {
	v, err := oigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oigb *OrderItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := oigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oigb *OrderItemGroupBy) IntX(ctx context.Context) int {
	v, err := oigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oigb *OrderItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oigb *OrderItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oigb.fields) > 1 {
		return nil, errors.New("ent: OrderItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oigb *OrderItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (oigb *OrderItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oigb *OrderItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := oigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oigb *OrderItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oigb.sqlQuery().Query()
	if err := oigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oigb *OrderItemGroupBy) sqlQuery() *sql.Selector {
	selector := oigb.sql
	columns := make([]string, 0, len(oigb.fields)+len(oigb.fns))
	columns = append(columns, oigb.fields...)
	for _, fn := range oigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(oigb.fields...)
}

// OrderItemSelect is the builder for select fields of OrderItem entities.
type OrderItemSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ois *OrderItemSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ois.path(ctx)
	if err != nil {
		return err
	}
	ois.sql = query
	return ois.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ois *OrderItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ois.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ois *OrderItemSelect) StringsX(ctx context.Context) []string {
	v, err := ois.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ois.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ois *OrderItemSelect) StringX(ctx context.Context) string {
	v, err := ois.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ois *OrderItemSelect) IntsX(ctx context.Context) []int {
	v, err := ois.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ois.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ois *OrderItemSelect) IntX(ctx context.Context) int {
	v, err := ois.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ois *OrderItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ois.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ois.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ois *OrderItemSelect) Float64X(ctx context.Context) float64 {
	v, err := ois.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ois.fields) > 1 {
		return nil, errors.New("ent: OrderItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ois.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ois *OrderItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := ois.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ois *OrderItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ois.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderitem.Label}
	default:
		err = fmt.Errorf("ent: OrderItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ois *OrderItemSelect) BoolX(ctx context.Context) bool {
	v, err := ois.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ois *OrderItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ois.sqlQuery().Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ois *OrderItemSelect) sqlQuery() sql.Querier {
	selector := ois.sql
	selector.Select(selector.Columns(ois.fields...)...)
	return selector
}
