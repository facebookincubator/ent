// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebook/ent/entc/integration/gremlin/ent/pet"
	"github.com/facebook/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebook/ent/entc/integration/gremlin/ent/user"
)

// PetQuery is the builder for querying Pet entities.
type PetQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.Pet
	// eager-loading edges.
	withTeam  *UserQuery
	withOwner *UserQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the builder.
func (pq *PetQuery) Where(ps ...predicate.Pet) *PetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PetQuery) Limit(limit int) *PetQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PetQuery) Offset(offset int) *PetQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PetQuery) Order(o ...OrderFunc) *PetQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryTeam chains the current query on the team edge.
func (pq *PetQuery) QueryTeam() *UserQuery {
	query := &UserQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pq.gremlinQuery()
		fromU = gremlin.InE(user.TeamLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the owner edge.
func (pq *PetQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pq.gremlinQuery()
		fromU = gremlin.InE(user.PetsLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first Pet entity in the query. Returns *NotFoundError when no pet was found.
func (pq *PetQuery) First(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PetQuery) FirstX(ctx context.Context) *Pet {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pet id in the query. Returns *NotFoundError when no id was found.
func (pq *PetQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PetQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Pet entity in the query, returns an error if not exactly one entity was returned.
func (pq *PetQuery) Only(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pet.Label}
	default:
		return nil, &NotSingularError{pet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PetQuery) OnlyX(ctx context.Context) *Pet {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Pet id in the query, returns an error if not exactly one id was returned.
func (pq *PetQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = &NotSingularError{pet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PetQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pets.
func (pq *PetQuery) All(ctx context.Context) ([]*Pet, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PetQuery) AllX(ctx context.Context) []*Pet {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pet ids.
func (pq *PetQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := pq.Select(pet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PetQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PetQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PetQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PetQuery) Clone() *PetQuery {
	if pq == nil {
		return nil
	}
	return &PetQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		predicates: append([]predicate.Pet{}, pq.predicates...),
		withTeam:   pq.withTeam.Clone(),
		withOwner:  pq.withOwner.Clone(),
		// clone intermediate query.
		gremlin: pq.gremlin.Clone(),
		path:    pq.path,
	}
}

//  WithTeam tells the query-builder to eager-loads the nodes that are connected to
// the "team" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PetQuery) WithTeam(opts ...func(*UserQuery)) *PetQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withTeam = query
	return pq
}

//  WithOwner tells the query-builder to eager-loads the nodes that are connected to
// the "owner" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PetQuery) WithOwner(opts ...func(*UserQuery)) *PetQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pet.Query().
//		GroupBy(pet.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PetQuery) GroupBy(field string, fields ...string) *PetGroupBy {
	group := &PetGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.gremlinQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Pet.Query().
//		Select(pet.FieldName).
//		Scan(ctx, &v)
//
func (pq *PetQuery) Select(field string, fields ...string) *PetSelect {
	selector := &PetSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.gremlinQuery(), nil
	}
	return selector
}

func (pq *PetQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.gremlin = prev
	}
	return nil
}

func (pq *PetQuery) gremlinAll(ctx context.Context) ([]*Pet, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().ValueMap(true).Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var pes Pets
	if err := pes.FromResponse(res); err != nil {
		return nil, err
	}
	pes.config(pq.config)
	return pes, nil
}

func (pq *PetQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().Count().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pq *PetQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().HasNext().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (pq *PetQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(pet.Label)
	if pq.gremlin != nil {
		v = pq.gremlin.Clone()
	}
	for _, p := range pq.predicates {
		p(v)
	}
	if len(pq.order) > 0 {
		v.Order()
		for _, p := range pq.order {
			p(v)
		}
	}
	switch limit, offset := pq.limit, pq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	v.Dedup()
	return v
}

// PetGroupBy is the builder for group-by Pet entities.
type PetGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PetGroupBy) Aggregate(fns ...AggregateFunc) *PetGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PetGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.gremlin = query
	return pgb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PetGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PetGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PetGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PetGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PetGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PetGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PetGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PetGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PetGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PetGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PetGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PetGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PetGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PetGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := pgb.gremlinQuery().Query()
	if err := pgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(pgb.fields)+len(pgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (pgb *PetGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range pgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range pgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return pgb.gremlin.Group().
		By(__.Values(pgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// PetSelect is the builder for select fields of Pet entities.
type PetSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PetSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ps.path(ctx)
	if err != nil {
		return err
	}
	ps.gremlin = query
	return ps.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PetSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PetSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PetSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ps *PetSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PetSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PetSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PetSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PetSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PetSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PetSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PetSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PetSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PetSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ps *PetSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = fmt.Errorf("ent: PetSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PetSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PetSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(ps.fields) == 1 {
		if ps.fields[0] != pet.FieldID {
			traversal = ps.gremlin.Values(ps.fields...)
		} else {
			traversal = ps.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(ps.fields))
		for i, f := range ps.fields {
			fields[i] = f
		}
		traversal = ps.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ps.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ps.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
