// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/car"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/predicate"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age             *int32
	addage          *int32
	name            *string
	nickname        *string
	address         *string
	clearaddress    bool
	renamed         *string
	clearrenamed    bool
	blob            *[]byte
	clearblob       bool
	state           *user.State
	clearstate      bool
	parent          map[int]struct{}
	children        map[int]struct{}
	spouse          map[int]struct{}
	car             map[int]struct{}
	clearedParent   bool
	removedChildren map[int]struct{}
	clearedSpouse   bool
	clearedCar      bool
	predicates      []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int32) *UserUpdate {
	uu.age = &i
	uu.addage = nil
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int32) *UserUpdate {
	if uu.addage == nil {
		uu.addage = &i
	} else {
		*uu.addage += i
	}
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetNickname sets the nickname field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.nickname = &s
	return uu
}

// SetAddress sets the address field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.address = &s
	return uu
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of address.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.address = nil
	uu.clearaddress = true
	return uu
}

// SetRenamed sets the renamed field.
func (uu *UserUpdate) SetRenamed(s string) *UserUpdate {
	uu.renamed = &s
	return uu
}

// SetNillableRenamed sets the renamed field if the given value is not nil.
func (uu *UserUpdate) SetNillableRenamed(s *string) *UserUpdate {
	if s != nil {
		uu.SetRenamed(*s)
	}
	return uu
}

// ClearRenamed clears the value of renamed.
func (uu *UserUpdate) ClearRenamed() *UserUpdate {
	uu.renamed = nil
	uu.clearrenamed = true
	return uu
}

// SetBlob sets the blob field.
func (uu *UserUpdate) SetBlob(b []byte) *UserUpdate {
	uu.blob = &b
	return uu
}

// ClearBlob clears the value of blob.
func (uu *UserUpdate) ClearBlob() *UserUpdate {
	uu.blob = nil
	uu.clearblob = true
	return uu
}

// SetState sets the state field.
func (uu *UserUpdate) SetState(u user.State) *UserUpdate {
	uu.state = &u
	return uu
}

// SetNillableState sets the state field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(u *user.State) *UserUpdate {
	if u != nil {
		uu.SetState(*u)
	}
	return uu
}

// ClearState clears the value of state.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.state = nil
	uu.clearstate = true
	return uu
}

// SetParentID sets the parent edge to User by id.
func (uu *UserUpdate) SetParentID(id int) *UserUpdate {
	if uu.parent == nil {
		uu.parent = make(map[int]struct{})
	}
	uu.parent[id] = struct{}{}
	return uu
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uu *UserUpdate) SetNillableParentID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetParentID(*id)
	}
	return uu
}

// SetParent sets the parent edge to User.
func (uu *UserUpdate) SetParent(u *User) *UserUpdate {
	return uu.SetParentID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uu *UserUpdate) AddChildIDs(ids ...int) *UserUpdate {
	if uu.children == nil {
		uu.children = make(map[int]struct{})
	}
	for i := range ids {
		uu.children[ids[i]] = struct{}{}
	}
	return uu
}

// AddChildren adds the children edges to User.
func (uu *UserUpdate) AddChildren(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddChildIDs(ids...)
}

// SetSpouseID sets the spouse edge to User by id.
func (uu *UserUpdate) SetSpouseID(id int) *UserUpdate {
	if uu.spouse == nil {
		uu.spouse = make(map[int]struct{})
	}
	uu.spouse[id] = struct{}{}
	return uu
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uu *UserUpdate) SetNillableSpouseID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetSpouseID(*id)
	}
	return uu
}

// SetSpouse sets the spouse edge to User.
func (uu *UserUpdate) SetSpouse(u *User) *UserUpdate {
	return uu.SetSpouseID(u.ID)
}

// SetCarID sets the car edge to Car by id.
func (uu *UserUpdate) SetCarID(id int) *UserUpdate {
	if uu.car == nil {
		uu.car = make(map[int]struct{})
	}
	uu.car[id] = struct{}{}
	return uu
}

// SetNillableCarID sets the car edge to Car by id if the given value is not nil.
func (uu *UserUpdate) SetNillableCarID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetCarID(*id)
	}
	return uu
}

// SetCar sets the car edge to Car.
func (uu *UserUpdate) SetCar(c *Car) *UserUpdate {
	return uu.SetCarID(c.ID)
}

// ClearParent clears the parent edge to User.
func (uu *UserUpdate) ClearParent() *UserUpdate {
	uu.clearedParent = true
	return uu
}

// RemoveChildIDs removes the children edge to User by ids.
func (uu *UserUpdate) RemoveChildIDs(ids ...int) *UserUpdate {
	if uu.removedChildren == nil {
		uu.removedChildren = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedChildren[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveChildren removes children edges to User.
func (uu *UserUpdate) RemoveChildren(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveChildIDs(ids...)
}

// ClearSpouse clears the spouse edge to User.
func (uu *UserUpdate) ClearSpouse() *UserUpdate {
	uu.clearedSpouse = true
	return uu
}

// ClearCar clears the car edge to Car.
func (uu *UserUpdate) ClearCar() *UserUpdate {
	uu.clearedCar = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.name != nil {
		if err := user.NameValidator(*uu.name); err != nil {
			return 0, fmt.Errorf("entv1: validator failed for field \"name\": %v", err)
		}
	}
	if uu.state != nil {
		if err := user.StateValidator(*uu.state); err != nil {
			return 0, fmt.Errorf("entv1: validator failed for field \"state\": %v", err)
		}
	}
	if len(uu.parent) > 1 {
		return 0, errors.New("entv1: multiple assignments on a unique edge \"parent\"")
	}
	if len(uu.spouse) > 1 {
		return 0, errors.New("entv1: multiple assignments on a unique edge \"spouse\"")
	}
	if len(uu.car) > 1 {
		return 0, errors.New("entv1: multiple assignments on a unique edge \"car\"")
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.age; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uu.addage; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uu.nickname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
	}
	if value := uu.address; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldAddress,
		})
	}
	if uu.clearaddress {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAddress,
		})
	}
	if value := uu.renamed; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldRenamed,
		})
	}
	if uu.clearrenamed {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldRenamed,
		})
	}
	if value := uu.blob; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBlob,
		})
	}
	if uu.clearblob {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBlob,
		})
	}
	if value := uu.state; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldState,
		})
	}
	if uu.clearstate {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldState,
		})
	}
	if uu.clearedParent {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.parent; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.removedChildren; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.children; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.clearedSpouse {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.spouse; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.clearedCar {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.car; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id              int
	age             *int32
	addage          *int32
	name            *string
	nickname        *string
	address         *string
	clearaddress    bool
	renamed         *string
	clearrenamed    bool
	blob            *[]byte
	clearblob       bool
	state           *user.State
	clearstate      bool
	parent          map[int]struct{}
	children        map[int]struct{}
	spouse          map[int]struct{}
	car             map[int]struct{}
	clearedParent   bool
	removedChildren map[int]struct{}
	clearedSpouse   bool
	clearedCar      bool
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int32) *UserUpdateOne {
	uuo.age = &i
	uuo.addage = nil
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int32) *UserUpdateOne {
	if uuo.addage == nil {
		uuo.addage = &i
	} else {
		*uuo.addage += i
	}
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetNickname sets the nickname field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.nickname = &s
	return uuo
}

// SetAddress sets the address field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.address = &s
	return uuo
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of address.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.address = nil
	uuo.clearaddress = true
	return uuo
}

// SetRenamed sets the renamed field.
func (uuo *UserUpdateOne) SetRenamed(s string) *UserUpdateOne {
	uuo.renamed = &s
	return uuo
}

// SetNillableRenamed sets the renamed field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRenamed(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRenamed(*s)
	}
	return uuo
}

// ClearRenamed clears the value of renamed.
func (uuo *UserUpdateOne) ClearRenamed() *UserUpdateOne {
	uuo.renamed = nil
	uuo.clearrenamed = true
	return uuo
}

// SetBlob sets the blob field.
func (uuo *UserUpdateOne) SetBlob(b []byte) *UserUpdateOne {
	uuo.blob = &b
	return uuo
}

// ClearBlob clears the value of blob.
func (uuo *UserUpdateOne) ClearBlob() *UserUpdateOne {
	uuo.blob = nil
	uuo.clearblob = true
	return uuo
}

// SetState sets the state field.
func (uuo *UserUpdateOne) SetState(u user.State) *UserUpdateOne {
	uuo.state = &u
	return uuo
}

// SetNillableState sets the state field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(u *user.State) *UserUpdateOne {
	if u != nil {
		uuo.SetState(*u)
	}
	return uuo
}

// ClearState clears the value of state.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.state = nil
	uuo.clearstate = true
	return uuo
}

// SetParentID sets the parent edge to User by id.
func (uuo *UserUpdateOne) SetParentID(id int) *UserUpdateOne {
	if uuo.parent == nil {
		uuo.parent = make(map[int]struct{})
	}
	uuo.parent[id] = struct{}{}
	return uuo
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableParentID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetParentID(*id)
	}
	return uuo
}

// SetParent sets the parent edge to User.
func (uuo *UserUpdateOne) SetParent(u *User) *UserUpdateOne {
	return uuo.SetParentID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uuo *UserUpdateOne) AddChildIDs(ids ...int) *UserUpdateOne {
	if uuo.children == nil {
		uuo.children = make(map[int]struct{})
	}
	for i := range ids {
		uuo.children[ids[i]] = struct{}{}
	}
	return uuo
}

// AddChildren adds the children edges to User.
func (uuo *UserUpdateOne) AddChildren(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddChildIDs(ids...)
}

// SetSpouseID sets the spouse edge to User by id.
func (uuo *UserUpdateOne) SetSpouseID(id int) *UserUpdateOne {
	if uuo.spouse == nil {
		uuo.spouse = make(map[int]struct{})
	}
	uuo.spouse[id] = struct{}{}
	return uuo
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSpouseID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSpouseID(*id)
	}
	return uuo
}

// SetSpouse sets the spouse edge to User.
func (uuo *UserUpdateOne) SetSpouse(u *User) *UserUpdateOne {
	return uuo.SetSpouseID(u.ID)
}

// SetCarID sets the car edge to Car by id.
func (uuo *UserUpdateOne) SetCarID(id int) *UserUpdateOne {
	if uuo.car == nil {
		uuo.car = make(map[int]struct{})
	}
	uuo.car[id] = struct{}{}
	return uuo
}

// SetNillableCarID sets the car edge to Car by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCarID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCarID(*id)
	}
	return uuo
}

// SetCar sets the car edge to Car.
func (uuo *UserUpdateOne) SetCar(c *Car) *UserUpdateOne {
	return uuo.SetCarID(c.ID)
}

// ClearParent clears the parent edge to User.
func (uuo *UserUpdateOne) ClearParent() *UserUpdateOne {
	uuo.clearedParent = true
	return uuo
}

// RemoveChildIDs removes the children edge to User by ids.
func (uuo *UserUpdateOne) RemoveChildIDs(ids ...int) *UserUpdateOne {
	if uuo.removedChildren == nil {
		uuo.removedChildren = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedChildren[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveChildren removes children edges to User.
func (uuo *UserUpdateOne) RemoveChildren(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveChildIDs(ids...)
}

// ClearSpouse clears the spouse edge to User.
func (uuo *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	uuo.clearedSpouse = true
	return uuo
}

// ClearCar clears the car edge to Car.
func (uuo *UserUpdateOne) ClearCar() *UserUpdateOne {
	uuo.clearedCar = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.name != nil {
		if err := user.NameValidator(*uuo.name); err != nil {
			return nil, fmt.Errorf("entv1: validator failed for field \"name\": %v", err)
		}
	}
	if uuo.state != nil {
		if err := user.StateValidator(*uuo.state); err != nil {
			return nil, fmt.Errorf("entv1: validator failed for field \"state\": %v", err)
		}
	}
	if len(uuo.parent) > 1 {
		return nil, errors.New("entv1: multiple assignments on a unique edge \"parent\"")
	}
	if len(uuo.spouse) > 1 {
		return nil, errors.New("entv1: multiple assignments on a unique edge \"spouse\"")
	}
	if len(uuo.car) > 1 {
		return nil, errors.New("entv1: multiple assignments on a unique edge \"car\"")
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.age; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uuo.addage; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: user.FieldAge,
		})
	}
	if value := uuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uuo.nickname; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
	}
	if value := uuo.address; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldAddress,
		})
	}
	if uuo.clearaddress {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldAddress,
		})
	}
	if value := uuo.renamed; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldRenamed,
		})
	}
	if uuo.clearrenamed {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldRenamed,
		})
	}
	if value := uuo.blob; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBlob,
		})
	}
	if uuo.clearblob {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBlob,
		})
	}
	if value := uuo.state; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldState,
		})
	}
	if uuo.clearstate {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldState,
		})
	}
	if uuo.clearedParent {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.parent; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.removedChildren; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.children; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.clearedSpouse {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.spouse; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.clearedCar {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.car; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
