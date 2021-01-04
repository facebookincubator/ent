// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebook/ent/entc/integration/gremlin/ent/comment"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetUniqueInt sets the "unique_int" field.
func (cc *CommentCreate) SetUniqueInt(i int) *CommentCreate {
	cc.mutation.SetUniqueInt(i)
	return cc
}

// SetUniqueFloat sets the "unique_float" field.
func (cc *CommentCreate) SetUniqueFloat(f float64) *CommentCreate {
	cc.mutation.SetUniqueFloat(f)
	return cc
}

// SetNillableInt sets the "nillable_int" field.
func (cc *CommentCreate) SetNillableInt(i int) *CommentCreate {
	cc.mutation.SetNillableInt(i)
	return cc
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cc *CommentCreate) SetNillableNillableInt(i *int) *CommentCreate {
	if i != nil {
		cc.SetNillableInt(*i)
	}
	return cc
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.UniqueInt(); !ok {
		return &ValidationError{Name: "unique_int", err: errors.New("ent: missing required field \"unique_int\"")}
	}
	if _, ok := cc.mutation.UniqueFloat(); !ok {
		return &ValidationError{Name: "unique_float", err: errors.New("ent: missing required field \"unique_float\"")}
	}
	return nil
}

func (cc *CommentCreate) gremlinSave(ctx context.Context) (*Comment, error) {
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Comment{config: cc.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CommentCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(comment.Label)
	if value, ok := cc.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := cc.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := cc.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
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

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	builders []*CommentCreate
}
