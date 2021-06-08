// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/metadata"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// MetadataCreate is the builder for creating a Metadata entity.
type MetadataCreate struct {
	config
	mutation *MetadataMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (mc *MetadataCreate) SetAge(i int) *MetadataCreate {
	mc.mutation.SetAge(i)
	return mc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (mc *MetadataCreate) SetNillableAge(i *int) *MetadataCreate {
	if i != nil {
		mc.SetAge(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MetadataCreate) SetID(i int) *MetadataCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mc *MetadataCreate) SetUserID(id int) *MetadataCreate {
	mc.mutation.SetUserID(id)
	return mc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mc *MetadataCreate) SetNillableUserID(id *int) *MetadataCreate {
	if id != nil {
		mc = mc.SetUserID(*id)
	}
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MetadataCreate) SetUser(u *User) *MetadataCreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mc *MetadataCreate) Mutation() *MetadataMutation {
	return mc.mutation
}

// Save creates the Metadata in the database.
func (mc *MetadataCreate) Save(ctx context.Context) (*Metadata, error) {
	var (
		err  error
		node *Metadata
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetadataCreate) SaveX(ctx context.Context) *Metadata {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mc *MetadataCreate) defaults() {
	if _, ok := mc.mutation.Age(); !ok {
		v := metadata.DefaultAge
		mc.mutation.SetAge(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetadataCreate) check() error {
	if _, ok := mc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	return nil
}

func (mc *MetadataCreate) sqlSave(ctx context.Context) (*Metadata, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (mc *MetadataCreate) createSpec() (*Metadata, *sqlgraph.CreateSpec) {
	var (
		_node = &Metadata{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metadata.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: metadata.FieldAge,
		})
		_node.Age = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.UserTable,
			Columns: []string{metadata.UserColumn},
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

// MetadataCreateBulk is the builder for creating many Metadata entities in bulk.
type MetadataCreateBulk struct {
	config
	builders []*MetadataCreate
}

// Save creates the Metadata entities in the database.
func (mcb *MetadataCreateBulk) Save(ctx context.Context) ([]*Metadata, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metadata, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetadataMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetadataCreateBulk) SaveX(ctx context.Context) []*Metadata {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
