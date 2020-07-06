// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent/card"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/user"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCard = "Card"
	TypeUser = "User"
)

// CardMutation represents an operation that mutate the Cards
// nodes in the graph.
type CardMutation struct {
	config
	op            Op
	typ           string
	id            *int
	number        *string
	name          *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Card, error)
}

var _ ent.Mutation = (*CardMutation)(nil)

// cardOption allows to manage the mutation configuration using functional options.
type cardOption func(*CardMutation)

// newCardMutation creates new mutation for $n.Name.
func newCardMutation(c config, op Op, opts ...cardOption) *CardMutation {
	m := &CardMutation{
		config:        c,
		op:            op,
		typ:           TypeCard,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCardID sets the id field of the mutation.
func withCardID(id int) cardOption {
	return func(m *CardMutation) {
		var (
			err   error
			once  sync.Once
			value *Card
		)
		m.oldValue = func(ctx context.Context) (*Card, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Card.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCard sets the old Card of the mutation.
func withCard(node *Card) cardOption {
	return func(m *CardMutation) {
		m.oldValue = func(context.Context) (*Card, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CardMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CardMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CardMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetNumber sets the number field.
func (m *CardMutation) SetNumber(s string) {
	m.number = &s
}

// Number returns the number value in the mutation.
func (m *CardMutation) Number() (r string, exists bool) {
	v := m.number
	if v == nil {
		return
	}
	return *v, true
}

// OldNumber returns the old number value of the Card.
// If the Card object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *CardMutation) OldNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNumber is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNumber: %w", err)
	}
	return oldValue.Number, nil
}

// ResetNumber reset all changes of the "number" field.
func (m *CardMutation) ResetNumber() {
	m.number = nil
}

// SetName sets the name field.
func (m *CardMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *CardMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Card.
// If the Card object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *CardMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of name.
func (m *CardMutation) ClearName() {
	m.name = nil
	m.clearedFields[card.FieldName] = struct{}{}
}

// NameCleared returns if the field name was cleared in this mutation.
func (m *CardMutation) NameCleared() bool {
	_, ok := m.clearedFields[card.FieldName]
	return ok
}

// ResetName reset all changes of the "name" field.
func (m *CardMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, card.FieldName)
}

// SetCreatedAt sets the created_at field.
func (m *CardMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *CardMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Card.
// If the Card object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *CardMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *CardMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetOwnerID sets the owner edge to User by id.
func (m *CardMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the owner edge to User.
func (m *CardMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *CardMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the owner id in the mutation.
func (m *CardMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the owner ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *CardMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner reset all changes of the "owner" edge.
func (m *CardMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Op returns the operation name.
func (m *CardMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Card).
func (m *CardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CardMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.number != nil {
		fields = append(fields, card.FieldNumber)
	}
	if m.name != nil {
		fields = append(fields, card.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, card.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CardMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case card.FieldNumber:
		return m.Number()
	case card.FieldName:
		return m.Name()
	case card.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *CardMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case card.FieldNumber:
		return m.OldNumber(ctx)
	case card.FieldName:
		return m.OldName(ctx)
	case card.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Card field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CardMutation) SetField(name string, value ent.Value) error {
	switch name {
	case card.FieldNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNumber(v)
		return nil
	case card.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case card.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CardMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CardMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CardMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Card numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CardMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(card.FieldName) {
		fields = append(fields, card.FieldName)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CardMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CardMutation) ClearField(name string) error {
	switch name {
	case card.FieldName:
		m.ClearName()
		return nil
	}
	return fmt.Errorf("unknown Card nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CardMutation) ResetField(name string) error {
	switch name {
	case card.FieldNumber:
		m.ResetNumber()
		return nil
	case card.FieldName:
		m.ResetName()
		return nil
	case card.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CardMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, card.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CardMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case card.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CardMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, card.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CardMutation) EdgeCleared(name string) bool {
	switch name {
	case card.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CardMutation) ClearEdge(name string) error {
	switch name {
	case card.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Card unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CardMutation) ResetEdge(name string) error {
	switch name {
	case card.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Card edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	version            *int
	addversion         *int
	name               *string
	worth              *uint
	addworth           *uint
	clearedFields      map[string]struct{}
	cards              map[int]struct{}
	removedcards       map[int]struct{}
	friends            map[int]struct{}
	removedfriends     map[int]struct{}
	best_friend        *int
	clearedbest_friend bool
	done               bool
	oldValue           func(context.Context) (*User, error)
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the id field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetVersion sets the version field.
func (m *UserMutation) SetVersion(i int) {
	m.version = &i
	m.addversion = nil
}

// Version returns the version value in the mutation.
func (m *UserMutation) Version() (r int, exists bool) {
	v := m.version
	if v == nil {
		return
	}
	return *v, true
}

// OldVersion returns the old version value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldVersion(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldVersion is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVersion: %w", err)
	}
	return oldValue.Version, nil
}

// AddVersion adds i to version.
func (m *UserMutation) AddVersion(i int) {
	if m.addversion != nil {
		*m.addversion += i
	} else {
		m.addversion = &i
	}
}

// AddedVersion returns the value that was added to the version field in this mutation.
func (m *UserMutation) AddedVersion() (r int, exists bool) {
	v := m.addversion
	if v == nil {
		return
	}
	return *v, true
}

// ResetVersion reset all changes of the "version" field.
func (m *UserMutation) ResetVersion() {
	m.version = nil
	m.addversion = nil
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetWorth sets the worth field.
func (m *UserMutation) SetWorth(u uint) {
	m.worth = &u
	m.addworth = nil
}

// Worth returns the worth value in the mutation.
func (m *UserMutation) Worth() (r uint, exists bool) {
	v := m.worth
	if v == nil {
		return
	}
	return *v, true
}

// OldWorth returns the old worth value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldWorth(ctx context.Context) (v uint, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldWorth is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldWorth requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWorth: %w", err)
	}
	return oldValue.Worth, nil
}

// AddWorth adds u to worth.
func (m *UserMutation) AddWorth(u uint) {
	if m.addworth != nil {
		*m.addworth += u
	} else {
		m.addworth = &u
	}
}

// AddedWorth returns the value that was added to the worth field in this mutation.
func (m *UserMutation) AddedWorth() (r uint, exists bool) {
	v := m.addworth
	if v == nil {
		return
	}
	return *v, true
}

// ClearWorth clears the value of worth.
func (m *UserMutation) ClearWorth() {
	m.worth = nil
	m.addworth = nil
	m.clearedFields[user.FieldWorth] = struct{}{}
}

// WorthCleared returns if the field worth was cleared in this mutation.
func (m *UserMutation) WorthCleared() bool {
	_, ok := m.clearedFields[user.FieldWorth]
	return ok
}

// ResetWorth reset all changes of the "worth" field.
func (m *UserMutation) ResetWorth() {
	m.worth = nil
	m.addworth = nil
	delete(m.clearedFields, user.FieldWorth)
}

// AddCardIDs adds the cards edge to Card by ids.
func (m *UserMutation) AddCardIDs(ids ...int) {
	if m.cards == nil {
		m.cards = make(map[int]struct{})
	}
	for i := range ids {
		m.cards[ids[i]] = struct{}{}
	}
}

// RemoveCardIDs removes the cards edge to Card by ids.
func (m *UserMutation) RemoveCardIDs(ids ...int) {
	if m.removedcards == nil {
		m.removedcards = make(map[int]struct{})
	}
	for i := range ids {
		m.removedcards[ids[i]] = struct{}{}
	}
}

// RemovedCards returns the removed ids of cards.
func (m *UserMutation) RemovedCardsIDs() (ids []int) {
	for id := range m.removedcards {
		ids = append(ids, id)
	}
	return
}

// CardsIDs returns the cards ids in the mutation.
func (m *UserMutation) CardsIDs() (ids []int) {
	for id := range m.cards {
		ids = append(ids, id)
	}
	return
}

// ResetCards reset all changes of the "cards" edge.
func (m *UserMutation) ResetCards() {
	m.cards = nil
	m.removedcards = nil
}

// AddFriendIDs adds the friends edge to User by ids.
func (m *UserMutation) AddFriendIDs(ids ...int) {
	if m.friends == nil {
		m.friends = make(map[int]struct{})
	}
	for i := range ids {
		m.friends[ids[i]] = struct{}{}
	}
}

// RemoveFriendIDs removes the friends edge to User by ids.
func (m *UserMutation) RemoveFriendIDs(ids ...int) {
	if m.removedfriends == nil {
		m.removedfriends = make(map[int]struct{})
	}
	for i := range ids {
		m.removedfriends[ids[i]] = struct{}{}
	}
}

// RemovedFriends returns the removed ids of friends.
func (m *UserMutation) RemovedFriendsIDs() (ids []int) {
	for id := range m.removedfriends {
		ids = append(ids, id)
	}
	return
}

// FriendsIDs returns the friends ids in the mutation.
func (m *UserMutation) FriendsIDs() (ids []int) {
	for id := range m.friends {
		ids = append(ids, id)
	}
	return
}

// ResetFriends reset all changes of the "friends" edge.
func (m *UserMutation) ResetFriends() {
	m.friends = nil
	m.removedfriends = nil
}

// SetBestFriendID sets the best_friend edge to User by id.
func (m *UserMutation) SetBestFriendID(id int) {
	m.best_friend = &id
}

// ClearBestFriend clears the best_friend edge to User.
func (m *UserMutation) ClearBestFriend() {
	m.clearedbest_friend = true
}

// BestFriendCleared returns if the edge best_friend was cleared.
func (m *UserMutation) BestFriendCleared() bool {
	return m.clearedbest_friend
}

// BestFriendID returns the best_friend id in the mutation.
func (m *UserMutation) BestFriendID() (id int, exists bool) {
	if m.best_friend != nil {
		return *m.best_friend, true
	}
	return
}

// BestFriendIDs returns the best_friend ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// BestFriendID instead. It exists only for internal usage by the builders.
func (m *UserMutation) BestFriendIDs() (ids []int) {
	if id := m.best_friend; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetBestFriend reset all changes of the "best_friend" edge.
func (m *UserMutation) ResetBestFriend() {
	m.best_friend = nil
	m.clearedbest_friend = false
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.version != nil {
		fields = append(fields, user.FieldVersion)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.worth != nil {
		fields = append(fields, user.FieldWorth)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldVersion:
		return m.Version()
	case user.FieldName:
		return m.Name()
	case user.FieldWorth:
		return m.Worth()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldVersion:
		return m.OldVersion(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldWorth:
		return m.OldWorth(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldVersion:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVersion(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldWorth:
		v, ok := value.(uint)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWorth(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addversion != nil {
		fields = append(fields, user.FieldVersion)
	}
	if m.addworth != nil {
		fields = append(fields, user.FieldWorth)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldVersion:
		return m.AddedVersion()
	case user.FieldWorth:
		return m.AddedWorth()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldVersion:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddVersion(v)
		return nil
	case user.FieldWorth:
		v, ok := value.(uint)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddWorth(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldWorth) {
		fields = append(fields, user.FieldWorth)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldWorth:
		m.ClearWorth()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldVersion:
		m.ResetVersion()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldWorth:
		m.ResetWorth()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.cards != nil {
		edges = append(edges, user.EdgeCards)
	}
	if m.friends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	if m.best_friend != nil {
		edges = append(edges, user.EdgeBestFriend)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCards:
		ids := make([]ent.Value, 0, len(m.cards))
		for id := range m.cards {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.friends))
		for id := range m.friends {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeBestFriend:
		if id := m.best_friend; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	if m.removedcards != nil {
		edges = append(edges, user.EdgeCards)
	}
	if m.removedfriends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCards:
		ids := make([]ent.Value, 0, len(m.removedcards))
		for id := range m.removedcards {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.removedfriends))
		for id := range m.removedfriends {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.clearedbest_friend {
		edges = append(edges, user.EdgeBestFriend)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeBestFriend:
		return m.clearedbest_friend
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeBestFriend:
		m.ClearBestFriend()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCards:
		m.ResetCards()
		return nil
	case user.EdgeFriends:
		m.ResetFriends()
		return nil
	case user.EdgeBestFriend:
		m.ResetBestFriend()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
