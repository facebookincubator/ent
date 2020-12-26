// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/migrate/entv2/customtype"
	"github.com/facebook/ent/entc/integration/migrate/entv2/predicate"
	"github.com/facebook/ent/schema/field"
)

// CustomTypeQuery is the builder for querying CustomType entities.
type CustomTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.CustomType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ctq *CustomTypeQuery) Where(ps ...predicate.CustomType) *CustomTypeQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit adds a limit step to the query.
func (ctq *CustomTypeQuery) Limit(limit int) *CustomTypeQuery {
	ctq.limit = &limit
	return ctq
}

// Offset adds an offset step to the query.
func (ctq *CustomTypeQuery) Offset(offset int) *CustomTypeQuery {
	ctq.offset = &offset
	return ctq
}

// Order adds an order step to the query.
func (ctq *CustomTypeQuery) Order(o ...OrderFunc) *CustomTypeQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// First returns the first CustomType entity in the query. Returns *NotFoundError when no customtype was found.
func (ctq *CustomTypeQuery) First(ctx context.Context) (*CustomType, error) {
	nodes, err := ctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *CustomTypeQuery) FirstX(ctx context.Context) *CustomType {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomType id in the query. Returns *NotFoundError when no id was found.
func (ctq *CustomTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *CustomTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only CustomType entity in the query, returns an error if not exactly one entity was returned.
func (ctq *CustomTypeQuery) Only(ctx context.Context) (*CustomType, error) {
	nodes, err := ctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customtype.Label}
	default:
		return nil, &NotSingularError{customtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *CustomTypeQuery) OnlyX(ctx context.Context) *CustomType {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only CustomType id in the query, returns an error if not exactly one id was returned.
func (ctq *CustomTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = &NotSingularError{customtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *CustomTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomTypes.
func (ctq *CustomTypeQuery) All(ctx context.Context) ([]*CustomType, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ctq *CustomTypeQuery) AllX(ctx context.Context) []*CustomType {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomType ids.
func (ctq *CustomTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ctq.Select(customtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *CustomTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *CustomTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *CustomTypeQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *CustomTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *CustomTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *CustomTypeQuery) Clone() *CustomTypeQuery {
	if ctq == nil {
		return nil
	}
	return &CustomTypeQuery{
		config:     ctq.config,
		limit:      ctq.limit,
		offset:     ctq.offset,
		order:      append([]OrderFunc{}, ctq.order...),
		predicates: append([]predicate.CustomType{}, ctq.predicates...),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Custom string `json:"custom,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomType.Query().
//		GroupBy(customtype.FieldCustom).
//		Aggregate(entv2.Count()).
//		Scan(ctx, &v)
//
func (ctq *CustomTypeQuery) GroupBy(field string, fields ...string) *CustomTypeGroupBy {
	group := &CustomTypeGroupBy{config: ctq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Custom string `json:"custom,omitempty"`
//	}
//
//	client.CustomType.Query().
//		Select(customtype.FieldCustom).
//		Scan(ctx, &v)
//
func (ctq *CustomTypeQuery) Select(field string, fields ...string) *CustomTypeSelect {
	selector := &CustomTypeSelect{config: ctq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ctq.sqlQuery(), nil
	}
	return selector
}

func (ctq *CustomTypeQuery) prepareQuery(ctx context.Context) error {
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *CustomTypeQuery) sqlAll(ctx context.Context) ([]*CustomType, error) {
	var (
		nodes = []*CustomType{}
		_spec = ctq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &CustomType{config: ctq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("entv2: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ctq *CustomTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *CustomTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ctq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entv2: check existence: %v", err)
	}
	return n > 0, nil
}

func (ctq *CustomTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customtype.Table,
			Columns: customtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		},
		From:   ctq.sql,
		Unique: true,
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, customtype.ValidColumn)
			}
		}
	}
	return _spec
}

func (ctq *CustomTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(customtype.Table)
	selector := builder.Select(t1.Columns(customtype.Columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(customtype.Columns...)...)
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector, customtype.ValidColumn)
	}
	if offset := ctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CustomTypeGroupBy is the builder for group-by CustomType entities.
type CustomTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *CustomTypeGroupBy) Aggregate(fns ...AggregateFunc) *CustomTypeGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ctgb *CustomTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ctgb.path(ctx)
	if err != nil {
		return err
	}
	ctgb.sql = query
	return ctgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ctgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ctgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ctgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ctgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ctgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ctgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ctgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ctgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ctgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ctgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ctgb.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ctgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ctgb *CustomTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ctgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ctgb *CustomTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ctgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ctgb *CustomTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ctgb.fields {
		if !customtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ctgb *CustomTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ctgb.sql
	columns := make([]string, 0, len(ctgb.fields)+len(ctgb.fns))
	columns = append(columns, ctgb.fields...)
	for _, fn := range ctgb.fns {
		columns = append(columns, fn(selector, customtype.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ctgb.fields...)
}

// CustomTypeSelect is the builder for select fields of CustomType entities.
type CustomTypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (cts *CustomTypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := cts.path(ctx)
	if err != nil {
		return err
	}
	cts.sql = query
	return cts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cts *CustomTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cts *CustomTypeSelect) StringsX(ctx context.Context) []string {
	v, err := cts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cts *CustomTypeSelect) StringX(ctx context.Context) string {
	v, err := cts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cts *CustomTypeSelect) IntsX(ctx context.Context) []int {
	v, err := cts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cts *CustomTypeSelect) IntX(ctx context.Context) int {
	v, err := cts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cts *CustomTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cts *CustomTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := cts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cts.fields) > 1 {
		return nil, errors.New("entv2: CustomTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cts *CustomTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := cts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (cts *CustomTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customtype.Label}
	default:
		err = fmt.Errorf("entv2: CustomTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cts *CustomTypeSelect) BoolX(ctx context.Context) bool {
	v, err := cts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cts *CustomTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cts.fields {
		if !customtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := cts.sqlQuery().Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cts *CustomTypeSelect) sqlQuery() sql.Querier {
	selector := cts.sql
	selector.Select(selector.Columns(cts.fields...)...)
	return selector
}
