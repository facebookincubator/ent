// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/info"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// InfoCreate is the builder for creating a Info entity.
type InfoCreate struct {
	config
	mutation *InfoMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (ic *InfoCreate) SetContent(jm json.RawMessage) *InfoCreate {
	ic.mutation.SetContent(jm)
	return ic
}

// SetID sets the "id" field.
func (ic *InfoCreate) SetID(i int) *InfoCreate {
	ic.mutation.SetID(i)
	return ic
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ic *InfoCreate) SetUserID(id int) *InfoCreate {
	ic.mutation.SetUserID(id)
	return ic
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ic *InfoCreate) SetNillableUserID(id *int) *InfoCreate {
	if id != nil {
		ic = ic.SetUserID(*id)
	}
	return ic
}

// SetUser sets the "user" edge to the User entity.
func (ic *InfoCreate) SetUser(u *User) *InfoCreate {
	return ic.SetUserID(u.ID)
}

// Mutation returns the InfoMutation object of the builder.
func (ic *InfoCreate) Mutation() *InfoMutation {
	return ic.mutation
}

// Save creates the Info in the database.
func (ic *InfoCreate) Save(ctx context.Context) (*Info, error) {
	var (
		err  error
		node *Info
	)
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InfoCreate) SaveX(ctx context.Context) *Info {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InfoCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InfoCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InfoCreate) check() error {
	if _, ok := ic.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	return nil
}

func (ic *InfoCreate) sqlSave(ctx context.Context) (*Info, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (ic *InfoCreate) createSpec() (*Info, *sqlgraph.CreateSpec) {
	var (
		_node = &Info{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: info.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: info.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: info.FieldContent,
		})
		_node.Content = value
	}
	if nodes := ic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   info.UserTable,
			Columns: []string{info.UserColumn},
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
		_node.ID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InfoCreateBulk is the builder for creating many Info entities in bulk.
type InfoCreateBulk struct {
	config
	builders []*InfoCreate
}

// Save creates the Info entities in the database.
func (icb *InfoCreateBulk) Save(ctx context.Context) ([]*Info, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Info, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InfoMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InfoCreateBulk) SaveX(ctx context.Context) []*Info {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
