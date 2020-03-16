// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupInfoQuery is the builder for querying GroupInfo entities.
type GroupInfoQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.GroupInfo
	// eager-loading edges.
	withGroups *GroupQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (giq *GroupInfoQuery) Where(ps ...predicate.GroupInfo) *GroupInfoQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit adds a limit step to the query.
func (giq *GroupInfoQuery) Limit(limit int) *GroupInfoQuery {
	giq.limit = &limit
	return giq
}

// Offset adds an offset step to the query.
func (giq *GroupInfoQuery) Offset(offset int) *GroupInfoQuery {
	giq.offset = &offset
	return giq
}

// Order adds an order step to the query.
func (giq *GroupInfoQuery) Order(o ...Order) *GroupInfoQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// QueryGroups chains the current query on the groups edge.
func (giq *GroupInfoQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: giq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(groupinfo.Table, groupinfo.FieldID, giq.sqlQuery()),
		sqlgraph.To(group.Table, group.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, groupinfo.GroupsTable, groupinfo.GroupsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(giq.driver.Dialect(), step)
	return query
}

// First returns the first GroupInfo entity in the query. Returns *NotFoundError when no groupinfo was found.
func (giq *GroupInfoQuery) First(ctx context.Context) (*GroupInfo, error) {
	gis, err := giq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gis) == 0 {
		return nil, &NotFoundError{groupinfo.Label}
	}
	return gis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstX(ctx context.Context) *GroupInfo {
	gi, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gi
}

// FirstID returns the first GroupInfo id in the query. Returns *NotFoundError when no id was found.
func (giq *GroupInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = giq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstXID(ctx context.Context) int {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GroupInfo entity in the query, returns an error if not exactly one entity was returned.
func (giq *GroupInfoQuery) Only(ctx context.Context) (*GroupInfo, error) {
	gis, err := giq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gis) {
	case 1:
		return gis[0], nil
	case 0:
		return nil, &NotFoundError{groupinfo.Label}
	default:
		return nil, &NotSingularError{groupinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyX(ctx context.Context) *GroupInfo {
	gi, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gi
}

// OnlyID returns the only GroupInfo id in the query, returns an error if not exactly one id was returned.
func (giq *GroupInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = giq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = &NotSingularError{groupinfo.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyXID(ctx context.Context) int {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupInfos.
func (giq *GroupInfoQuery) All(ctx context.Context) ([]*GroupInfo, error) {
	return giq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (giq *GroupInfoQuery) AllX(ctx context.Context) []*GroupInfo {
	gis, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gis
}

// IDs executes the query and returns a list of GroupInfo ids.
func (giq *GroupInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := giq.Select(groupinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GroupInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GroupInfoQuery) Count(ctx context.Context) (int, error) {
	return giq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (giq *GroupInfoQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GroupInfoQuery) Exist(ctx context.Context) (bool, error) {
	return giq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GroupInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GroupInfoQuery) Clone() *GroupInfoQuery {
	return &GroupInfoQuery{
		config:     giq.config,
		limit:      giq.limit,
		offset:     giq.offset,
		order:      append([]Order{}, giq.order...),
		unique:     append([]string{}, giq.unique...),
		predicates: append([]predicate.GroupInfo{}, giq.predicates...),
		// clone intermediate query.
		sql: giq.sql.Clone(),
	}
}

//  WithGroups tells the query-builder to eager-loads the nodes that are connected to
// the "groups" edge. The optional arguments used to configure the query builder of the edge.
func (giq *GroupInfoQuery) WithGroups(opts ...func(*GroupQuery)) *GroupInfoQuery {
	query := &GroupQuery{config: giq.config}
	for _, opt := range opts {
		opt(query)
	}
	giq.withGroups = query
	return giq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		GroupBy(groupinfo.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) GroupBy(field string, fields ...string) *GroupInfoGroupBy {
	group := &GroupInfoGroupBy{config: giq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = giq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		Select(groupinfo.FieldDesc).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) Select(field string, fields ...string) *GroupInfoSelect {
	selector := &GroupInfoSelect{config: giq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = giq.sqlQuery()
	return selector
}

func (giq *GroupInfoQuery) sqlAll(ctx context.Context) ([]*GroupInfo, error) {
	var (
		nodes       = []*GroupInfo{}
		_spec       = giq.querySpec()
		loadedTypes = [1]bool{
			giq.withGroups != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &GroupInfo{config: giq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, giq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := giq.withGroups; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GroupInfo)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Group(func(s *sql.Selector) {
			s.Where(sql.InValues(groupinfo.GroupsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.group_info
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "group_info" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_info" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Groups = append(node.Edges.Groups, n)
		}
	}

	return nodes, nil
}

func (giq *GroupInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := giq.querySpec()
	return sqlgraph.CountNodes(ctx, giq.driver, _spec)
}

func (giq *GroupInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := giq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (giq *GroupInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupinfo.Table,
			Columns: groupinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		},
		From:   giq.sql,
		Unique: true,
	}
	if ps := giq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := giq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := giq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := giq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (giq *GroupInfoQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(giq.driver.Dialect())
	t1 := builder.Table(groupinfo.Table)
	selector := builder.Select(t1.Columns(groupinfo.Columns...)...).From(t1)
	if giq.sql != nil {
		selector = giq.sql
		selector.Select(selector.Columns(groupinfo.Columns...)...)
	}
	for _, p := range giq.predicates {
		p(selector)
	}
	for _, p := range giq.order {
		p(selector)
	}
	if offset := giq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := giq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupInfoGroupBy is the builder for group-by GroupInfo entities.
type GroupInfoGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GroupInfoGroupBy) Aggregate(fns ...Aggregate) *GroupInfoGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the group-by query and scan the result into the given value.
func (gigb *GroupInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	return gigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := gigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := gigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gigb *GroupInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gigb.sqlQuery().Query()
	if err := gigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gigb *GroupInfoGroupBy) sqlQuery() *sql.Selector {
	selector := gigb.sql
	columns := make([]string, 0, len(gigb.fields)+len(gigb.fns))
	columns = append(columns, gigb.fields...)
	for _, fn := range gigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(gigb.fields...)
}

// GroupInfoSelect is the builder for select fields of GroupInfo entities.
type GroupInfoSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (gis *GroupInfoSelect) Scan(ctx context.Context, v interface{}) error {
	return gis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gis *GroupInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gis *GroupInfoSelect) StringsX(ctx context.Context) []string {
	v, err := gis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gis *GroupInfoSelect) IntsX(ctx context.Context) []int {
	v, err := gis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gis *GroupInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gis *GroupInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := gis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gis *GroupInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gis.sqlQuery().Query()
	if err := gis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gis *GroupInfoSelect) sqlQuery() sql.Querier {
	selector := gis.sql
	selector.Select(selector.Columns(gis.fields...)...)
	return selector
}
