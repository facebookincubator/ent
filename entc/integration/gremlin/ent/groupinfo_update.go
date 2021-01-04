// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebook/ent/entc/integration/gremlin/ent/group"
	"github.com/facebook/ent/entc/integration/gremlin/ent/groupinfo"
	"github.com/facebook/ent/entc/integration/gremlin/ent/predicate"
)

// GroupInfoUpdate is the builder for updating GroupInfo entities.
type GroupInfoUpdate struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// Where adds a new predicate for the GroupInfoUpdate builder.
func (giu *GroupInfoUpdate) Where(ps ...predicate.GroupInfo) *GroupInfoUpdate {
	giu.mutation.predicates = append(giu.mutation.predicates, ps...)
	return giu
}

// SetDesc sets the "desc" field.
func (giu *GroupInfoUpdate) SetDesc(s string) *GroupInfoUpdate {
	giu.mutation.SetDesc(s)
	return giu
}

// SetMaxUsers sets the "max_users" field.
func (giu *GroupInfoUpdate) SetMaxUsers(i int) *GroupInfoUpdate {
	giu.mutation.ResetMaxUsers()
	giu.mutation.SetMaxUsers(i)
	return giu
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (giu *GroupInfoUpdate) SetNillableMaxUsers(i *int) *GroupInfoUpdate {
	if i != nil {
		giu.SetMaxUsers(*i)
	}
	return giu
}

// AddMaxUsers adds i to the "max_users" field.
func (giu *GroupInfoUpdate) AddMaxUsers(i int) *GroupInfoUpdate {
	giu.mutation.AddMaxUsers(i)
	return giu
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (giu *GroupInfoUpdate) AddGroupIDs(ids ...string) *GroupInfoUpdate {
	giu.mutation.AddGroupIDs(ids...)
	return giu
}

// AddGroups adds the "groups" edges to the Group entity.
func (giu *GroupInfoUpdate) AddGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (giu *GroupInfoUpdate) Mutation() *GroupInfoMutation {
	return giu.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (giu *GroupInfoUpdate) ClearGroups() *GroupInfoUpdate {
	giu.mutation.ClearGroups()
	return giu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (giu *GroupInfoUpdate) RemoveGroupIDs(ids ...string) *GroupInfoUpdate {
	giu.mutation.RemoveGroupIDs(ids...)
	return giu
}

// RemoveGroups removes "groups" edges to Group entities.
func (giu *GroupInfoUpdate) RemoveGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (giu *GroupInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(giu.hooks) == 0 {
		affected, err = giu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			giu.mutation = mutation
			affected, err = giu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(giu.hooks) - 1; i >= 0; i-- {
			mut = giu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, giu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (giu *GroupInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := giu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (giu *GroupInfoUpdate) Exec(ctx context.Context) error {
	_, err := giu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giu *GroupInfoUpdate) ExecX(ctx context.Context) {
	if err := giu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giu *GroupInfoUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := giu.gremlin().Query()
	if err := giu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (giu *GroupInfoUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V().HasLabel(groupinfo.Label)
	for _, p := range giu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := giu.mutation.Desc(); ok {
		v.Property(dsl.Single, groupinfo.FieldDesc, value)
	}
	if value, ok := giu.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, value)
	}
	if value, ok := giu.mutation.AddedMaxUsers(); ok {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, __.Union(__.Values(groupinfo.FieldMaxUsers), __.Constant(value)).Sum())
	}
	for _, id := range giu.mutation.RemovedGroupsIDs() {
		tr := rv.Clone().InE(group.InfoLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range giu.mutation.GroupsIDs() {
		v.AddE(group.InfoLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.InfoLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(groupinfo.Label, group.InfoLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// GroupInfoUpdateOne is the builder for updating a single GroupInfo entity.
type GroupInfoUpdateOne struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// SetDesc sets the "desc" field.
func (giuo *GroupInfoUpdateOne) SetDesc(s string) *GroupInfoUpdateOne {
	giuo.mutation.SetDesc(s)
	return giuo
}

// SetMaxUsers sets the "max_users" field.
func (giuo *GroupInfoUpdateOne) SetMaxUsers(i int) *GroupInfoUpdateOne {
	giuo.mutation.ResetMaxUsers()
	giuo.mutation.SetMaxUsers(i)
	return giuo
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (giuo *GroupInfoUpdateOne) SetNillableMaxUsers(i *int) *GroupInfoUpdateOne {
	if i != nil {
		giuo.SetMaxUsers(*i)
	}
	return giuo
}

// AddMaxUsers adds i to the "max_users" field.
func (giuo *GroupInfoUpdateOne) AddMaxUsers(i int) *GroupInfoUpdateOne {
	giuo.mutation.AddMaxUsers(i)
	return giuo
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (giuo *GroupInfoUpdateOne) AddGroupIDs(ids ...string) *GroupInfoUpdateOne {
	giuo.mutation.AddGroupIDs(ids...)
	return giuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (giuo *GroupInfoUpdateOne) AddGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (giuo *GroupInfoUpdateOne) Mutation() *GroupInfoMutation {
	return giuo.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (giuo *GroupInfoUpdateOne) ClearGroups() *GroupInfoUpdateOne {
	giuo.mutation.ClearGroups()
	return giuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (giuo *GroupInfoUpdateOne) RemoveGroupIDs(ids ...string) *GroupInfoUpdateOne {
	giuo.mutation.RemoveGroupIDs(ids...)
	return giuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (giuo *GroupInfoUpdateOne) RemoveGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated GroupInfo entity.
func (giuo *GroupInfoUpdateOne) Save(ctx context.Context) (*GroupInfo, error) {
	var (
		err  error
		node *GroupInfo
	)
	if len(giuo.hooks) == 0 {
		node, err = giuo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			giuo.mutation = mutation
			node, err = giuo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(giuo.hooks) - 1; i >= 0; i-- {
			mut = giuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, giuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) SaveX(ctx context.Context) *GroupInfo {
	node, err := giuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (giuo *GroupInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := giuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) ExecX(ctx context.Context) {
	if err := giuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giuo *GroupInfoUpdateOne) gremlinSave(ctx context.Context) (*GroupInfo, error) {
	res := &gremlin.Response{}
	id, ok := giuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GroupInfo.ID for update")}
	}
	query, bindings := giuo.gremlin(id).Query()
	if err := giuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gi := &GroupInfo{config: giuo.config}
	if err := gi.FromResponse(res); err != nil {
		return nil, err
	}
	return gi, nil
}

func (giuo *GroupInfoUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := giuo.mutation.Desc(); ok {
		v.Property(dsl.Single, groupinfo.FieldDesc, value)
	}
	if value, ok := giuo.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, value)
	}
	if value, ok := giuo.mutation.AddedMaxUsers(); ok {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, __.Union(__.Values(groupinfo.FieldMaxUsers), __.Constant(value)).Sum())
	}
	for _, id := range giuo.mutation.RemovedGroupsIDs() {
		tr := rv.Clone().InE(group.InfoLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range giuo.mutation.GroupsIDs() {
		v.AddE(group.InfoLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.InfoLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(groupinfo.Label, group.InfoLabel, id)),
		})
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
