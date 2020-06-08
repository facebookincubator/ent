// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// GalaxyUpdate is the builder for updating Galaxy entities.
type GalaxyUpdate struct {
	config
	hooks      []Hook
	mutation   *GalaxyMutation
	predicates []predicate.Galaxy
}

// Where adds a new predicate for the builder.
func (gu *GalaxyUpdate) Where(ps ...predicate.Galaxy) *GalaxyUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetName sets the name field.
func (gu *GalaxyUpdate) SetName(s string) *GalaxyUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetType sets the type field.
func (gu *GalaxyUpdate) SetType(ga galaxy.Type) *GalaxyUpdate {
	gu.mutation.SetType(ga)
	return gu
}

// AddPlanetIDs adds the planets edge to Planet by ids.
func (gu *GalaxyUpdate) AddPlanetIDs(ids ...int) *GalaxyUpdate {
	gu.mutation.AddPlanetIDs(ids...)
	return gu
}

// AddPlanets adds the planets edges to Planet.
func (gu *GalaxyUpdate) AddPlanets(p ...*Planet) *GalaxyUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.AddPlanetIDs(ids...)
}

// Mutation returns the GalaxyMutation object of the builder.
func (gu *GalaxyUpdate) Mutation() *GalaxyMutation {
	return gu.mutation
}

// RemovePlanetIDs removes the planets edge to Planet by ids.
func (gu *GalaxyUpdate) RemovePlanetIDs(ids ...int) *GalaxyUpdate {
	gu.mutation.RemovePlanetIDs(ids...)
	return gu
}

// RemovePlanets removes planets edges to Planet.
func (gu *GalaxyUpdate) RemovePlanets(p ...*Planet) *GalaxyUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gu.RemovePlanetIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GalaxyUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := gu.mutation.Name(); ok {
		if err := galaxy.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %w", err)
		}
	}
	if v, ok := gu.mutation.GetType(); ok {
		if err := galaxy.TypeValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"type\": %w", err)
		}
	}

	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GalaxyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GalaxyUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GalaxyUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GalaxyUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GalaxyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   galaxy.Table,
			Columns: galaxy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: galaxy.FieldID,
			},
		},
	}
	if ps := gu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: galaxy.FieldName,
		})
	}
	if value, ok := gu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: galaxy.FieldType,
		})
	}
	if nodes := gu.mutation.RemovedPlanetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   galaxy.PlanetsTable,
			Columns: []string{galaxy.PlanetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.PlanetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   galaxy.PlanetsTable,
			Columns: []string{galaxy.PlanetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{galaxy.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GalaxyUpdateOne is the builder for updating a single Galaxy entity.
type GalaxyUpdateOne struct {
	config
	hooks    []Hook
	mutation *GalaxyMutation
}

// SetName sets the name field.
func (guo *GalaxyUpdateOne) SetName(s string) *GalaxyUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetType sets the type field.
func (guo *GalaxyUpdateOne) SetType(ga galaxy.Type) *GalaxyUpdateOne {
	guo.mutation.SetType(ga)
	return guo
}

// AddPlanetIDs adds the planets edge to Planet by ids.
func (guo *GalaxyUpdateOne) AddPlanetIDs(ids ...int) *GalaxyUpdateOne {
	guo.mutation.AddPlanetIDs(ids...)
	return guo
}

// AddPlanets adds the planets edges to Planet.
func (guo *GalaxyUpdateOne) AddPlanets(p ...*Planet) *GalaxyUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.AddPlanetIDs(ids...)
}

// Mutation returns the GalaxyMutation object of the builder.
func (guo *GalaxyUpdateOne) Mutation() *GalaxyMutation {
	return guo.mutation
}

// RemovePlanetIDs removes the planets edge to Planet by ids.
func (guo *GalaxyUpdateOne) RemovePlanetIDs(ids ...int) *GalaxyUpdateOne {
	guo.mutation.RemovePlanetIDs(ids...)
	return guo
}

// RemovePlanets removes planets edges to Planet.
func (guo *GalaxyUpdateOne) RemovePlanets(p ...*Planet) *GalaxyUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return guo.RemovePlanetIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (guo *GalaxyUpdateOne) Save(ctx context.Context) (*Galaxy, error) {
	if v, ok := guo.mutation.Name(); ok {
		if err := galaxy.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %w", err)
		}
	}
	if v, ok := guo.mutation.GetType(); ok {
		if err := galaxy.TypeValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"type\": %w", err)
		}
	}

	var (
		err  error
		node *Galaxy
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GalaxyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GalaxyUpdateOne) SaveX(ctx context.Context) *Galaxy {
	ga, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ga
}

// Exec executes the query on the entity.
func (guo *GalaxyUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GalaxyUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GalaxyUpdateOne) sqlSave(ctx context.Context) (ga *Galaxy, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   galaxy.Table,
			Columns: galaxy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: galaxy.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Galaxy.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: galaxy.FieldName,
		})
	}
	if value, ok := guo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: galaxy.FieldType,
		})
	}
	if nodes := guo.mutation.RemovedPlanetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   galaxy.PlanetsTable,
			Columns: []string{galaxy.PlanetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.PlanetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   galaxy.PlanetsTable,
			Columns: []string{galaxy.PlanetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ga = &Galaxy{config: guo.config}
	_spec.Assign = ga.assignValues
	_spec.ScanValues = ga.scanValues()
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{galaxy.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ga, nil
}
