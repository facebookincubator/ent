// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/card"
	"github.com/facebook/ent/entc/integration/ent/spec"
	"github.com/facebook/ent/entc/integration/ent/user"
	"github.com/facebook/ent/schema/field"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *CardCreate) SetCreateTime(t time.Time) *CardCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableCreateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CardCreate) SetUpdateTime(t time.Time) *CardCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetNumber sets the "number" field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.mutation.SetNumber(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CardCreate) SetName(s string) *CardCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CardCreate) SetNillableName(s *string) *CardCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CardCreate) SetOwnerID(id int) *CardCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(id *int) *CardCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (cc *CardCreate) AddSpecIDs(ids ...int) *CardCreate {
	cc.mutation.AddSpecIDs(ids...)
	return cc
}

// AddSpec adds the "spec" edges to the Spec entity.
func (cc *CardCreate) AddSpec(s ...*Spec) *CardCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSpecIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	var (
		err  error
		node *Card
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
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
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := card.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := card.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := cc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New("ent: missing required field \"number\"")}
	}
	if v, ok := cc.mutation.Number(); ok {
		if err := card.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf("ent: validator failed for field \"number\": %w", err)}
		}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table:  card.Table,
			Schema: cc.CardSchema,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: card.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: card.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: card.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Number(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: card.FieldNumber,
		})
		_node.Number = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: card.FieldName,
		})
		_node.Name = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: spec.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	builders []*CardCreate
}

// Save creates the Card entities in the database.
func (ccb *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Card, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
