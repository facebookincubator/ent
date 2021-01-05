// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/migrate/entv2/car"
	"github.com/facebook/ent/entc/integration/migrate/entv2/pet"
	"github.com/facebook/ent/entc/integration/migrate/entv2/predicate"
	"github.com/facebook/ent/entc/integration/migrate/entv2/user"
	"github.com/facebook/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetMixedString sets the "mixed_string" field.
func (uu *UserUpdate) SetMixedString(s string) *UserUpdate {
	uu.mutation.SetMixedString(s)
	return uu
}

// SetNillableMixedString sets the "mixed_string" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMixedString(s *string) *UserUpdate {
	if s != nil {
		uu.SetMixedString(*s)
	}
	return uu
}

// SetMixedEnum sets the "mixed_enum" field.
func (uu *UserUpdate) SetMixedEnum(ue user.MixedEnum) *UserUpdate {
	uu.mutation.SetMixedEnum(ue)
	return uu
}

// SetNillableMixedEnum sets the "mixed_enum" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMixedEnum(ue *user.MixedEnum) *UserUpdate {
	if ue != nil {
		uu.SetMixedEnum(*ue)
	}
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// SetBuffer sets the "buffer" field.
func (uu *UserUpdate) SetBuffer(b []byte) *UserUpdate {
	uu.mutation.SetBuffer(b)
	return uu
}

// ClearBuffer clears the value of the "buffer" field.
func (uu *UserUpdate) ClearBuffer() *UserUpdate {
	uu.mutation.ClearBuffer()
	return uu
}

// SetTitle sets the "title" field.
func (uu *UserUpdate) SetTitle(s string) *UserUpdate {
	uu.mutation.SetTitle(s)
	return uu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTitle(s *string) *UserUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// SetNewName sets the "new_name" field.
func (uu *UserUpdate) SetNewName(s string) *UserUpdate {
	uu.mutation.SetNewName(s)
	return uu
}

// SetNillableNewName sets the "new_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNewName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNewName(*s)
	}
	return uu
}

// ClearNewName clears the value of the "new_name" field.
func (uu *UserUpdate) ClearNewName() *UserUpdate {
	uu.mutation.ClearNewName()
	return uu
}

// SetBlob sets the "blob" field.
func (uu *UserUpdate) SetBlob(b []byte) *UserUpdate {
	uu.mutation.SetBlob(b)
	return uu
}

// ClearBlob clears the value of the "blob" field.
func (uu *UserUpdate) ClearBlob() *UserUpdate {
	uu.mutation.ClearBlob()
	return uu
}

// SetState sets the "state" field.
func (uu *UserUpdate) SetState(u user.State) *UserUpdate {
	uu.mutation.SetState(u)
	return uu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(u *user.State) *UserUpdate {
	if u != nil {
		uu.SetState(*u)
	}
	return uu
}

// ClearState clears the value of the "state" field.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.mutation.ClearState()
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(u user.Status) *UserUpdate {
	uu.mutation.SetStatus(u)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(u *user.Status) *UserUpdate {
	if u != nil {
		uu.SetStatus(*u)
	}
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// SetWorkplace sets the "workplace" field.
func (uu *UserUpdate) SetWorkplace(s string) *UserUpdate {
	uu.mutation.SetWorkplace(s)
	return uu
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWorkplace(s *string) *UserUpdate {
	if s != nil {
		uu.SetWorkplace(*s)
	}
	return uu
}

// ClearWorkplace clears the value of the "workplace" field.
func (uu *UserUpdate) ClearWorkplace() *UserUpdate {
	uu.mutation.ClearWorkplace()
	return uu
}

// AddCarIDs adds the "car" edge to the Car entity by IDs.
func (uu *UserUpdate) AddCarIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCarIDs(ids...)
	return uu
}

// AddCar adds the "car" edges to the Car entity.
func (uu *UserUpdate) AddCar(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarIDs(ids...)
}

// SetPetsID sets the "pets" edge to the Pet entity by ID.
func (uu *UserUpdate) SetPetsID(id int) *UserUpdate {
	uu.mutation.SetPetsID(id)
	return uu
}

// SetNillablePetsID sets the "pets" edge to the Pet entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePetsID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetPetsID(*id)
	}
	return uu
}

// SetPets sets the "pets" edge to the Pet entity.
func (uu *UserUpdate) SetPets(p *Pet) *UserUpdate {
	return uu.SetPetsID(p.ID)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (uu *UserUpdate) AddFriendIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFriendIDs(ids...)
	return uu
}

// AddFriends adds the "friends" edges to the User entity.
func (uu *UserUpdate) AddFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFriendIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCar clears all "car" edges to the Car entity.
func (uu *UserUpdate) ClearCar() *UserUpdate {
	uu.mutation.ClearCar()
	return uu
}

// RemoveCarIDs removes the "car" edge to Car entities by IDs.
func (uu *UserUpdate) RemoveCarIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCarIDs(ids...)
	return uu
}

// RemoveCar removes "car" edges to Car entities.
func (uu *UserUpdate) RemoveCar(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarIDs(ids...)
}

// ClearPets clears the "pets" edge to the Pet entity.
func (uu *UserUpdate) ClearPets() *UserUpdate {
	uu.mutation.ClearPets()
	return uu
}

// ClearFriends clears all "friends" edges to the User entity.
func (uu *UserUpdate) ClearFriends() *UserUpdate {
	uu.mutation.ClearFriends()
	return uu
}

// RemoveFriendIDs removes the "friends" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFriendIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFriendIDs(ids...)
	return uu
}

// RemoveFriends removes "friends" edges to User entities.
func (uu *UserUpdate) RemoveFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFriendIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.MixedEnum(); ok {
		if err := user.MixedEnumValidator(v); err != nil {
			return &ValidationError{Name: "mixed_enum", err: fmt.Errorf("entv2: validator failed for field \"mixed_enum\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("entv2: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("entv2: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("entv2: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Schema:  uu.UserSchema,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.MixedString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldMixedString,
		})
	}
	if value, ok := uu.mutation.MixedEnum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldMixedEnum,
		})
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uu.mutation.Buffer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: user.FieldBuffer,
		})
	}
	if uu.mutation.BufferCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBuffer,
		})
	}
	if value, ok := uu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTitle,
		})
	}
	if value, ok := uu.mutation.NewName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNewName,
		})
	}
	if uu.mutation.NewNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldNewName,
		})
	}
	if value, ok := uu.mutation.Blob(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: user.FieldBlob,
		})
	}
	if uu.mutation.BlobCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBlob,
		})
	}
	if value, ok := uu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldState,
		})
	}
	if uu.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldState,
		})
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if uu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uu.mutation.Workplace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldWorkplace,
		})
	}
	if uu.mutation.WorkplaceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldWorkplace,
		})
	}
	if uu.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
	if nodes := uu.mutation.RemovedCarIDs(); len(nodes) > 0 && !uu.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
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
	if nodes := uu.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !uu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetMixedString sets the "mixed_string" field.
func (uuo *UserUpdateOne) SetMixedString(s string) *UserUpdateOne {
	uuo.mutation.SetMixedString(s)
	return uuo
}

// SetNillableMixedString sets the "mixed_string" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMixedString(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMixedString(*s)
	}
	return uuo
}

// SetMixedEnum sets the "mixed_enum" field.
func (uuo *UserUpdateOne) SetMixedEnum(ue user.MixedEnum) *UserUpdateOne {
	uuo.mutation.SetMixedEnum(ue)
	return uuo
}

// SetNillableMixedEnum sets the "mixed_enum" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMixedEnum(ue *user.MixedEnum) *UserUpdateOne {
	if ue != nil {
		uuo.SetMixedEnum(*ue)
	}
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// SetBuffer sets the "buffer" field.
func (uuo *UserUpdateOne) SetBuffer(b []byte) *UserUpdateOne {
	uuo.mutation.SetBuffer(b)
	return uuo
}

// ClearBuffer clears the value of the "buffer" field.
func (uuo *UserUpdateOne) ClearBuffer() *UserUpdateOne {
	uuo.mutation.ClearBuffer()
	return uuo
}

// SetTitle sets the "title" field.
func (uuo *UserUpdateOne) SetTitle(s string) *UserUpdateOne {
	uuo.mutation.SetTitle(s)
	return uuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTitle(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// SetNewName sets the "new_name" field.
func (uuo *UserUpdateOne) SetNewName(s string) *UserUpdateOne {
	uuo.mutation.SetNewName(s)
	return uuo
}

// SetNillableNewName sets the "new_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNewName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNewName(*s)
	}
	return uuo
}

// ClearNewName clears the value of the "new_name" field.
func (uuo *UserUpdateOne) ClearNewName() *UserUpdateOne {
	uuo.mutation.ClearNewName()
	return uuo
}

// SetBlob sets the "blob" field.
func (uuo *UserUpdateOne) SetBlob(b []byte) *UserUpdateOne {
	uuo.mutation.SetBlob(b)
	return uuo
}

// ClearBlob clears the value of the "blob" field.
func (uuo *UserUpdateOne) ClearBlob() *UserUpdateOne {
	uuo.mutation.ClearBlob()
	return uuo
}

// SetState sets the "state" field.
func (uuo *UserUpdateOne) SetState(u user.State) *UserUpdateOne {
	uuo.mutation.SetState(u)
	return uuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(u *user.State) *UserUpdateOne {
	if u != nil {
		uuo.SetState(*u)
	}
	return uuo
}

// ClearState clears the value of the "state" field.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.mutation.ClearState()
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(u user.Status) *UserUpdateOne {
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(u *user.Status) *UserUpdateOne {
	if u != nil {
		uuo.SetStatus(*u)
	}
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// SetWorkplace sets the "workplace" field.
func (uuo *UserUpdateOne) SetWorkplace(s string) *UserUpdateOne {
	uuo.mutation.SetWorkplace(s)
	return uuo
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWorkplace(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetWorkplace(*s)
	}
	return uuo
}

// ClearWorkplace clears the value of the "workplace" field.
func (uuo *UserUpdateOne) ClearWorkplace() *UserUpdateOne {
	uuo.mutation.ClearWorkplace()
	return uuo
}

// AddCarIDs adds the "car" edge to the Car entity by IDs.
func (uuo *UserUpdateOne) AddCarIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCarIDs(ids...)
	return uuo
}

// AddCar adds the "car" edges to the Car entity.
func (uuo *UserUpdateOne) AddCar(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarIDs(ids...)
}

// SetPetsID sets the "pets" edge to the Pet entity by ID.
func (uuo *UserUpdateOne) SetPetsID(id int) *UserUpdateOne {
	uuo.mutation.SetPetsID(id)
	return uuo
}

// SetNillablePetsID sets the "pets" edge to the Pet entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePetsID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPetsID(*id)
	}
	return uuo
}

// SetPets sets the "pets" edge to the Pet entity.
func (uuo *UserUpdateOne) SetPets(p *Pet) *UserUpdateOne {
	return uuo.SetPetsID(p.ID)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFriendIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFriendIDs(ids...)
	return uuo
}

// AddFriends adds the "friends" edges to the User entity.
func (uuo *UserUpdateOne) AddFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFriendIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCar clears all "car" edges to the Car entity.
func (uuo *UserUpdateOne) ClearCar() *UserUpdateOne {
	uuo.mutation.ClearCar()
	return uuo
}

// RemoveCarIDs removes the "car" edge to Car entities by IDs.
func (uuo *UserUpdateOne) RemoveCarIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCarIDs(ids...)
	return uuo
}

// RemoveCar removes "car" edges to Car entities.
func (uuo *UserUpdateOne) RemoveCar(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarIDs(ids...)
}

// ClearPets clears the "pets" edge to the Pet entity.
func (uuo *UserUpdateOne) ClearPets() *UserUpdateOne {
	uuo.mutation.ClearPets()
	return uuo
}

// ClearFriends clears all "friends" edges to the User entity.
func (uuo *UserUpdateOne) ClearFriends() *UserUpdateOne {
	uuo.mutation.ClearFriends()
	return uuo
}

// RemoveFriendIDs removes the "friends" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFriendIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFriendIDs(ids...)
	return uuo
}

// RemoveFriends removes "friends" edges to User entities.
func (uuo *UserUpdateOne) RemoveFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFriendIDs(ids...)
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
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

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.MixedEnum(); ok {
		if err := user.MixedEnumValidator(v); err != nil {
			return &ValidationError{Name: "mixed_enum", err: fmt.Errorf("entv2: validator failed for field \"mixed_enum\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf("entv2: validator failed for field \"nickname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("entv2: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("entv2: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Schema:  uuo.UserSchema,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.MixedString(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldMixedString,
		})
	}
	if value, ok := uuo.mutation.MixedEnum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldMixedEnum,
		})
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNickname,
		})
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uuo.mutation.Buffer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: user.FieldBuffer,
		})
	}
	if uuo.mutation.BufferCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBuffer,
		})
	}
	if value, ok := uuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTitle,
		})
	}
	if value, ok := uuo.mutation.NewName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldNewName,
		})
	}
	if uuo.mutation.NewNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldNewName,
		})
	}
	if value, ok := uuo.mutation.Blob(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: user.FieldBlob,
		})
	}
	if uuo.mutation.BlobCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: user.FieldBlob,
		})
	}
	if value, ok := uuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldState,
		})
	}
	if uuo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldState,
		})
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldStatus,
		})
	}
	if uuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldStatus,
		})
	}
	if value, ok := uuo.mutation.Workplace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldWorkplace,
		})
	}
	if uuo.mutation.WorkplaceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldWorkplace,
		})
	}
	if uuo.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
	if nodes := uuo.mutation.RemovedCarIDs(); len(nodes) > 0 && !uuo.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
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
	if nodes := uuo.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !uuo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
