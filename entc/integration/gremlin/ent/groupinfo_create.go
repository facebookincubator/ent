// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/group"
	"entgo.io/ent/entc/integration/gremlin/ent/groupinfo"
)

// GroupInfoCreate is the builder for creating a GroupInfo entity.
type GroupInfoCreate struct {
	config
	mutation         *GroupInfoMutation
	hooks            []Hook
	constraintFields []string
}

// SetDesc sets the "desc" field.
func (gic *GroupInfoCreate) SetDesc(s string) *GroupInfoCreate {
	gic.mutation.SetDesc(s)
	return gic
}

// SetMaxUsers sets the "max_users" field.
func (gic *GroupInfoCreate) SetMaxUsers(i int) *GroupInfoCreate {
	gic.mutation.SetMaxUsers(i)
	return gic
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (gic *GroupInfoCreate) SetNillableMaxUsers(i *int) *GroupInfoCreate {
	if i != nil {
		gic.SetMaxUsers(*i)
	}
	return gic
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (gic *GroupInfoCreate) AddGroupIDs(ids ...string) *GroupInfoCreate {
	gic.mutation.AddGroupIDs(ids...)
	return gic
}

// AddGroups adds the "groups" edges to the Group entity.
func (gic *GroupInfoCreate) AddGroups(g ...*Group) *GroupInfoCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gic.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (gic *GroupInfoCreate) Mutation() *GroupInfoMutation {
	return gic.mutation
}

// Save creates the GroupInfo in the database.
func (gic *GroupInfoCreate) Save(ctx context.Context) (*GroupInfo, error) {
	var (
		err  error
		node *GroupInfo
	)
	gic.defaults()
	if len(gic.hooks) == 0 {
		if err = gic.check(); err != nil {
			return nil, err
		}
		node, err = gic.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gic.check(); err != nil {
				return nil, err
			}
			gic.mutation = mutation
			node, err = gic.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gic.hooks) - 1; i >= 0; i-- {
			mut = gic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GroupInfoCreate) SaveX(ctx context.Context) *GroupInfo {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gic *GroupInfoCreate) defaults() {
	if _, ok := gic.mutation.MaxUsers(); !ok {
		v := groupinfo.DefaultMaxUsers
		gic.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gic *GroupInfoCreate) check() error {
	if _, ok := gic.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New("ent: missing required field \"desc\"")}
	}
	if _, ok := gic.mutation.MaxUsers(); !ok {
		return &ValidationError{Name: "max_users", err: errors.New("ent: missing required field \"max_users\"")}
	}
	return nil
}

func (gic *GroupInfoCreate) gremlinSave(ctx context.Context) (*GroupInfo, error) {
	res := &gremlin.Response{}
	query, bindings := gic.gremlin().Query()
	if err := gic.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gi := &GroupInfo{config: gic.config}
	if err := gi.FromResponse(res); err != nil {
		return nil, err
	}
	return gi, nil
}

func (gic *GroupInfoCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(groupinfo.Label)
	if value, ok := gic.mutation.Desc(); ok {
		v.Property(dsl.Single, groupinfo.FieldDesc, value)
	}
	if value, ok := gic.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, value)
	}
	for _, id := range gic.mutation.GroupsIDs() {
		v.AddE(group.InfoLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.InfoLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(groupinfo.Label, group.InfoLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// GroupInfoCreateBulk is the builder for creating many GroupInfo entities in bulk.
type GroupInfoCreateBulk struct {
	config
	builders              []*GroupInfoCreate
	batchConstraintFields []string
}
