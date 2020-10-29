// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/examples/traversal/ent/group"
	"github.com/facebook/ent/examples/traversal/ent/user"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges       GroupEdges `json:"edges"`
	group_admin *int
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Users holds the value of the users edge.
	Users []*User
	// Admin holds the value of the admin edge.
	Admin *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// AdminOrErr returns the Admin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) AdminOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Admin == nil {
			// The edge admin was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Admin, nil
	}
	return nil, &NotLoadedError{edge: "admin"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Group) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // group_admin
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(values ...interface{}) error {
	if m, n := len(values), len(group.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		gr.Name = value.String
	}
	values = values[1:]
	if len(values) == len(group.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field group_admin", value)
		} else if value.Valid {
			gr.group_admin = new(int)
			*gr.group_admin = int(value.Int64)
		}
	}
	return nil
}

// QueryUsers queries the users edge of the Group.
func (gr *Group) QueryUsers() *UserQuery {
	return (&GroupClient{config: gr.config}).QueryUsers(gr)
}

// QueryAdmin queries the admin edge of the Group.
func (gr *Group) QueryAdmin() *UserQuery {
	return (&GroupClient{config: gr.config}).QueryAdmin(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
