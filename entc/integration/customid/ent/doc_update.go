// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/schema/field"
)

// DocUpdate is the builder for updating Doc entities.
type DocUpdate struct {
	config
	hooks    []Hook
	mutation *DocMutation
}

// Where appends a list predicates to the DocUpdate builder.
func (du *DocUpdate) Where(ps ...predicate.Doc) *DocUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetText sets the "text" field.
func (du *DocUpdate) SetText(s string) *DocUpdate {
	du.mutation.SetText(s)
	return du
}

// SetNillableText sets the "text" field if the given value is not nil.
func (du *DocUpdate) SetNillableText(s *string) *DocUpdate {
	if s != nil {
		du.SetText(*s)
	}
	return du
}

// ClearText clears the value of the "text" field.
func (du *DocUpdate) ClearText() *DocUpdate {
	du.mutation.ClearText()
	return du
}

// SetParentID sets the "parent" edge to the Doc entity by ID.
func (du *DocUpdate) SetParentID(id schema.DocID) *DocUpdate {
	du.mutation.SetParentID(id)
	return du
}

// SetNillableParentID sets the "parent" edge to the Doc entity by ID if the given value is not nil.
func (du *DocUpdate) SetNillableParentID(id *schema.DocID) *DocUpdate {
	if id != nil {
		du = du.SetParentID(*id)
	}
	return du
}

// SetParent sets the "parent" edge to the Doc entity.
func (du *DocUpdate) SetParent(d *Doc) *DocUpdate {
	return du.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Doc entity by IDs.
func (du *DocUpdate) AddChildIDs(ids ...schema.DocID) *DocUpdate {
	du.mutation.AddChildIDs(ids...)
	return du
}

// AddChildren adds the "children" edges to the Doc entity.
func (du *DocUpdate) AddChildren(d ...*Doc) *DocUpdate {
	ids := make([]schema.DocID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddChildIDs(ids...)
}

// Mutation returns the DocMutation object of the builder.
func (du *DocUpdate) Mutation() *DocMutation {
	return du.mutation
}

// ClearParent clears the "parent" edge to the Doc entity.
func (du *DocUpdate) ClearParent() *DocUpdate {
	du.mutation.ClearParent()
	return du
}

// ClearChildren clears all "children" edges to the Doc entity.
func (du *DocUpdate) ClearChildren() *DocUpdate {
	du.mutation.ClearChildren()
	return du
}

// RemoveChildIDs removes the "children" edge to Doc entities by IDs.
func (du *DocUpdate) RemoveChildIDs(ids ...schema.DocID) *DocUpdate {
	du.mutation.RemoveChildIDs(ids...)
	return du
}

// RemoveChildren removes "children" edges to Doc entities.
func (du *DocUpdate) RemoveChildren(d ...*Doc) *DocUpdate {
	ids := make([]schema.DocID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DocUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DocUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DocUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DocUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DocUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doc.Table,
			Columns: doc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: doc.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doc.FieldText,
		})
	}
	if du.mutation.TextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: doc.FieldText,
		})
	}
	if du.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doc.ParentTable,
			Columns: []string{doc.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doc.ParentTable,
			Columns: []string{doc.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !du.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DocUpdateOne is the builder for updating a single Doc entity.
type DocUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocMutation
}

// SetText sets the "text" field.
func (duo *DocUpdateOne) SetText(s string) *DocUpdateOne {
	duo.mutation.SetText(s)
	return duo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (duo *DocUpdateOne) SetNillableText(s *string) *DocUpdateOne {
	if s != nil {
		duo.SetText(*s)
	}
	return duo
}

// ClearText clears the value of the "text" field.
func (duo *DocUpdateOne) ClearText() *DocUpdateOne {
	duo.mutation.ClearText()
	return duo
}

// SetParentID sets the "parent" edge to the Doc entity by ID.
func (duo *DocUpdateOne) SetParentID(id schema.DocID) *DocUpdateOne {
	duo.mutation.SetParentID(id)
	return duo
}

// SetNillableParentID sets the "parent" edge to the Doc entity by ID if the given value is not nil.
func (duo *DocUpdateOne) SetNillableParentID(id *schema.DocID) *DocUpdateOne {
	if id != nil {
		duo = duo.SetParentID(*id)
	}
	return duo
}

// SetParent sets the "parent" edge to the Doc entity.
func (duo *DocUpdateOne) SetParent(d *Doc) *DocUpdateOne {
	return duo.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Doc entity by IDs.
func (duo *DocUpdateOne) AddChildIDs(ids ...schema.DocID) *DocUpdateOne {
	duo.mutation.AddChildIDs(ids...)
	return duo
}

// AddChildren adds the "children" edges to the Doc entity.
func (duo *DocUpdateOne) AddChildren(d ...*Doc) *DocUpdateOne {
	ids := make([]schema.DocID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddChildIDs(ids...)
}

// Mutation returns the DocMutation object of the builder.
func (duo *DocUpdateOne) Mutation() *DocMutation {
	return duo.mutation
}

// ClearParent clears the "parent" edge to the Doc entity.
func (duo *DocUpdateOne) ClearParent() *DocUpdateOne {
	duo.mutation.ClearParent()
	return duo
}

// ClearChildren clears all "children" edges to the Doc entity.
func (duo *DocUpdateOne) ClearChildren() *DocUpdateOne {
	duo.mutation.ClearChildren()
	return duo
}

// RemoveChildIDs removes the "children" edge to Doc entities by IDs.
func (duo *DocUpdateOne) RemoveChildIDs(ids ...schema.DocID) *DocUpdateOne {
	duo.mutation.RemoveChildIDs(ids...)
	return duo
}

// RemoveChildren removes "children" edges to Doc entities.
func (duo *DocUpdateOne) RemoveChildren(d ...*Doc) *DocUpdateOne {
	ids := make([]schema.DocID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DocUpdateOne) Select(field string, fields ...string) *DocUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Doc entity.
func (duo *DocUpdateOne) Save(ctx context.Context) (*Doc, error) {
	var (
		err  error
		node *Doc
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DocUpdateOne) SaveX(ctx context.Context) *Doc {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DocUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DocUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DocUpdateOne) sqlSave(ctx context.Context) (_node *Doc, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doc.Table,
			Columns: doc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: doc.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Doc.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, doc.FieldID)
		for _, f := range fields {
			if !doc.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != doc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doc.FieldText,
		})
	}
	if duo.mutation.TextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: doc.FieldText,
		})
	}
	if duo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doc.ParentTable,
			Columns: []string{doc.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doc.ParentTable,
			Columns: []string{doc.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !duo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Doc{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
