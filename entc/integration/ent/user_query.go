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
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.User
	// eager-loading edges.
	withCard      *CardQuery
	withPets      *PetQuery
	withFiles     *FileQuery
	withGroups    *GroupQuery
	withFriends   *UserQuery
	withFollowers *UserQuery
	withFollowing *UserQuery
	withTeam      *PetQuery
	withSpouse    *UserQuery
	withChildren  *UserQuery
	withParent    *UserQuery
	withFKs       bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.offset = &offset
	return uq
}

// Order adds an order step to the query.
func (uq *UserQuery) Order(o ...Order) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryCard chains the current query on the card edge.
func (uq *UserQuery) QueryCard() *CardQuery {
	query := &CardQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(card.Table, card.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, user.CardTable, user.CardColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryPets chains the current query on the pets edge.
func (uq *UserQuery) QueryPets() *PetQuery {
	query := &PetQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(pet.Table, pet.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, user.PetsTable, user.PetsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryFiles chains the current query on the files edge.
func (uq *UserQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(file.Table, file.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, user.FilesTable, user.FilesColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryGroups chains the current query on the groups edge.
func (uq *UserQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(group.Table, group.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, user.GroupsTable, user.GroupsPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryFriends chains the current query on the friends edge.
func (uq *UserQuery) QueryFriends() *UserQuery {
	query := &UserQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, user.FriendsTable, user.FriendsPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryFollowers chains the current query on the followers edge.
func (uq *UserQuery) QueryFollowers() *UserQuery {
	query := &UserQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, user.FollowersTable, user.FollowersPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryFollowing chains the current query on the following edge.
func (uq *UserQuery) QueryFollowing() *UserQuery {
	query := &UserQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, user.FollowingTable, user.FollowingPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryTeam chains the current query on the team edge.
func (uq *UserQuery) QueryTeam() *PetQuery {
	query := &PetQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(pet.Table, pet.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, user.TeamTable, user.TeamColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QuerySpouse chains the current query on the spouse edge.
func (uq *UserQuery) QuerySpouse() *UserQuery {
	query := &UserQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, user.SpouseTable, user.SpouseColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryChildren chains the current query on the children edge.
func (uq *UserQuery) QueryChildren() *UserQuery {
	query := &UserQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, user.ChildrenTable, user.ChildrenColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// QueryParent chains the current query on the parent edge.
func (uq *UserQuery) QueryParent() *UserQuery {
	query := &UserQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, uq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, user.ParentTable, user.ParentColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// First returns the first User entity in the query. Returns *NotFoundError when no user was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	us, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return us[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	u, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return u
}

// FirstID returns the first User id in the query. Returns *NotFoundError when no id was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstXID(ctx context.Context) string {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only User entity in the query, returns an error if not exactly one entity was returned.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	us, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(us) {
	case 1:
		return us[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	u, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// OnlyID returns the only User id in the query, returns an error if not exactly one id was returned.
func (uq *UserQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyXID(ctx context.Context) string {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	us, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return us
}

// IDs executes the query and returns a list of User ids.
func (uq *UserQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []string {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	return &UserQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]Order{}, uq.order...),
		unique:     append([]string{}, uq.unique...),
		predicates: append([]predicate.User{}, uq.predicates...),
		// clone intermediate query.
		sql: uq.sql.Clone(),
	}
}

//  WithCard tells the query-builder to eager-loads the nodes that are connected to
// the "card" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithCard(opts ...func(*CardQuery)) *UserQuery {
	query := &CardQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withCard = query
	return uq
}

//  WithPets tells the query-builder to eager-loads the nodes that are connected to
// the "pets" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithPets(opts ...func(*PetQuery)) *UserQuery {
	query := &PetQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withPets = query
	return uq
}

//  WithFiles tells the query-builder to eager-loads the nodes that are connected to
// the "files" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithFiles(opts ...func(*FileQuery)) *UserQuery {
	query := &FileQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withFiles = query
	return uq
}

//  WithGroups tells the query-builder to eager-loads the nodes that are connected to
// the "groups" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithGroups(opts ...func(*GroupQuery)) *UserQuery {
	query := &GroupQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withGroups = query
	return uq
}

//  WithFriends tells the query-builder to eager-loads the nodes that are connected to
// the "friends" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithFriends(opts ...func(*UserQuery)) *UserQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withFriends = query
	return uq
}

//  WithFollowers tells the query-builder to eager-loads the nodes that are connected to
// the "followers" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithFollowers(opts ...func(*UserQuery)) *UserQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withFollowers = query
	return uq
}

//  WithFollowing tells the query-builder to eager-loads the nodes that are connected to
// the "following" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithFollowing(opts ...func(*UserQuery)) *UserQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withFollowing = query
	return uq
}

//  WithTeam tells the query-builder to eager-loads the nodes that are connected to
// the "team" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithTeam(opts ...func(*PetQuery)) *UserQuery {
	query := &PetQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withTeam = query
	return uq
}

//  WithSpouse tells the query-builder to eager-loads the nodes that are connected to
// the "spouse" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithSpouse(opts ...func(*UserQuery)) *UserQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withSpouse = query
	return uq
}

//  WithChildren tells the query-builder to eager-loads the nodes that are connected to
// the "children" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithChildren(opts ...func(*UserQuery)) *UserQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withChildren = query
	return uq
}

//  WithParent tells the query-builder to eager-loads the nodes that are connected to
// the "parent" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserQuery) WithParent(opts ...func(*UserQuery)) *UserQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withParent = query
	return uq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OptionalInt int `json:"optional_int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldOptionalInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	group := &UserGroupBy{config: uq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = uq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		OptionalInt int `json:"optional_int,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldOptionalInt).
//		Scan(ctx, &v)
//
func (uq *UserQuery) Select(field string, fields ...string) *UserSelect {
	selector := &UserSelect{config: uq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = uq.sqlQuery()
	return selector
}

func (uq *UserQuery) sqlAll(ctx context.Context) ([]*User, error) {
	var (
		nodes   []*User = []*User{}
		withFKs         = uq.withFKs
		_spec           = uq.querySpec()
	)
	if uq.withSpouse != nil || uq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, user.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &User{config: uq.config}
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
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uq.withCard; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*User)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Card(func(s *sql.Selector) {
			s.Where(sql.InValues(user.CardColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.owner_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "owner_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "owner_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Card = n
		}
	}

	if query := uq.withPets; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*User)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Pet(func(s *sql.Selector) {
			s.Where(sql.InValues(user.PetsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.owner_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "owner_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "owner_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Pets = append(node.Edges.Pets, n)
		}
	}

	if query := uq.withFiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*User)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.File(func(s *sql.Selector) {
			s.Where(sql.InValues(user.FilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.owner_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "owner_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "owner_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Files = append(node.Edges.Files, n)
		}
	}

	if query := uq.withGroups; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*User, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*User)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   user.GroupsTable,
				Columns: user.GroupsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(user.GroupsPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, uq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "groups": %v`, err)
		}
		query.Where(group.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "groups" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Groups = append(nodes[i].Edges.Groups, n)
			}
		}
	}

	if query := uq.withFriends; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*User, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*User)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   user.FriendsTable,
				Columns: user.FriendsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(user.FriendsPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, uq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "friends": %v`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "friends" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Friends = append(nodes[i].Edges.Friends, n)
			}
		}
	}

	if query := uq.withFollowers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*User, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*User)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   user.FollowersTable,
				Columns: user.FollowersPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(user.FollowersPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, uq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "followers": %v`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "followers" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Followers = append(nodes[i].Edges.Followers, n)
			}
		}
	}

	if query := uq.withFollowing; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*User, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []string
			edges   = make(map[string][]*User)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   user.FollowingTable,
				Columns: user.FollowingPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(user.FollowingPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := strconv.FormatInt(eout.Int64, 10)
				inValue := strconv.FormatInt(ein.Int64, 10)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, uq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "following": %v`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "following" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Following = append(nodes[i].Edges.Following, n)
			}
		}
	}

	if query := uq.withTeam; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*User)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Pet(func(s *sql.Selector) {
			s.Where(sql.InValues(user.TeamColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.team_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "team_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "team_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Team = n
		}
	}

	if query := uq.withSpouse; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*User)
		for i := range nodes {
			if fk := nodes[i].user_spouse_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_spouse_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Spouse = n
			}
		}
	}

	if query := uq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*User)
		for i := range nodes {
			id, err := strconv.Atoi(nodes[i].ID)
			if err != nil {
				return nil, err
			}
			fks = append(fks, id)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(user.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.parent_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "parent_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := uq.withParent; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*User)
		for i := range nodes {
			if fk := nodes[i].parent_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	return nodes, nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	selector := builder.Select(t1.Columns(user.Columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(user.Columns...)...)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the builder for group-by User entities.
type UserGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...Aggregate) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scan the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ugb *UserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ugb *UserGroupBy) StringsX(ctx context.Context) []string {
	v, err := ugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ugb *UserGroupBy) IntsX(ctx context.Context) []int {
	v, err := ugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ugb *UserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ugb *UserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ugb.sqlQuery().Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UserGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql
	columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
	columns = append(columns, ugb.fields...)
	for _, fn := range ugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ugb.fields...)
}

// UserSelect is the builder for select fields of User entities.
type UserSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v interface{}) error {
	return us.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (us *UserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := us.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (us *UserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (us *UserSelect) StringsX(ctx context.Context) []string {
	v, err := us.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (us *UserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (us *UserSelect) IntsX(ctx context.Context) []int {
	v, err := us.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (us *UserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (us *UserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := us.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (us *UserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (us *UserSelect) BoolsX(ctx context.Context) []bool {
	v, err := us.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (us *UserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sqlQuery().Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (us *UserSelect) sqlQuery() sql.Querier {
	selector := us.sql
	selector.Select(selector.Columns(us.fields...)...)
	return selector
}
