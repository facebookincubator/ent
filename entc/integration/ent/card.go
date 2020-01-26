// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges    CardEdges `json:"edges"`
	owner_id *string

	// StaticField defined by templates.
	StaticField string `json:"boring,omitempty"`
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Spec holds the value of the spec edge.
	Spec []*Spec
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // number
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Card) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(values ...interface{}) error {
	if m, n := len(values), len(card.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		c.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		c.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field number", values[2])
	} else if value.Valid {
		c.Number = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[3])
	} else if value.Valid {
		c.Name = value.String
	}
	values = values[4:]
	if len(values) == len(card.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			c.owner_id = new(string)
			*c.owner_id = strconv.FormatInt(value.Int64, 10)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Card.
func (c *Card) QueryOwner() *UserQuery {
	return (&CardClient{c.config}).QueryOwner(c)
}

// QuerySpec queries the spec edge of the Card.
func (c *Card) QuerySpec() *SpecQuery {
	return (&CardClient{c.config}).QuerySpec(c)
}

// Update returns a builder for updating this Card.
// Note that, you need to call Card.Unwrap() before calling this method, if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return (&CardClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", number=")
	builder.WriteString(c.Number)
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (c *Card) id() int {
	id, _ := strconv.Atoi(c.ID)
	return id
}

// Cards is a parsable slice of Card.
type Cards []*Card

func (c Cards) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
