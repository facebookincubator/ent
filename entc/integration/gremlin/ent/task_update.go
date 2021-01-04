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
	"github.com/facebook/ent/entc/integration/ent/schema"
	"github.com/facebook/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebook/ent/entc/integration/gremlin/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where adds a new predicate for the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TaskUpdate) SetPriority(s schema.Priority) *TaskUpdate {
	tu.mutation.ResetPriority()
	tu.mutation.SetPriority(s)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePriority(s *schema.Priority) *TaskUpdate {
	if s != nil {
		tu.SetPriority(*s)
	}
	return tu
}

// AddPriority adds s to the "priority" field.
func (tu *TaskUpdate) AddPriority(s schema.Priority) *TaskUpdate {
	tu.mutation.AddPriority(s)
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Priority(); ok {
		if err := task.PriorityValidator(int(v)); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf("ent: validator failed for field \"priority\": %w", err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := tu.gremlin().Query()
	if err := tu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (tu *TaskUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(task.Label)
	for _, p := range tu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if value, ok := tu.mutation.Priority(); ok {
		v.Property(dsl.Single, task.FieldPriority, value)
	}
	if value, ok := tu.mutation.AddedPriority(); ok {
		v.Property(dsl.Single, task.FieldPriority, __.Union(__.Values(task.FieldPriority), __.Constant(value)).Sum())
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// SetPriority sets the "priority" field.
func (tuo *TaskUpdateOne) SetPriority(s schema.Priority) *TaskUpdateOne {
	tuo.mutation.ResetPriority()
	tuo.mutation.SetPriority(s)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePriority(s *schema.Priority) *TaskUpdateOne {
	if s != nil {
		tuo.SetPriority(*s)
	}
	return tuo
}

// AddPriority adds s to the "priority" field.
func (tuo *TaskUpdateOne) AddPriority(s schema.Priority) *TaskUpdateOne {
	tuo.mutation.AddPriority(s)
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Priority(); ok {
		if err := task.PriorityValidator(int(v)); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf("ent: validator failed for field \"priority\": %w", err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) gremlinSave(ctx context.Context) (*Task, error) {
	res := &gremlin.Response{}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	query, bindings := tuo.gremlin(id).Query()
	if err := tuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	t := &Task{config: tuo.config}
	if err := t.FromResponse(res); err != nil {
		return nil, err
	}
	return t, nil
}

func (tuo *TaskUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if value, ok := tuo.mutation.Priority(); ok {
		v.Property(dsl.Single, task.FieldPriority, value)
	}
	if value, ok := tuo.mutation.AddedPriority(); ok {
		v.Property(dsl.Single, task.FieldPriority, __.Union(__.Values(task.FieldPriority), __.Constant(value)).Sum())
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
