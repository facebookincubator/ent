// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
)

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CityQuery when eager-loading is set.
	Edges CityEdges `json:"edges"`
}

// CityEdges holds the relations/edges for other nodes in the graph.
type CityEdges struct {
	// Streets holds the value of the streets edge.
	Streets []*Street
}

// scanValues returns the types for scanning values from sql.Rows.
func (*City) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the City fields.
func (c *City) assignValues(values ...interface{}) error {
	if m, n := len(values), len(city.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		c.Name = value.String
	}
	return nil
}

// QueryStreets queries the streets edge of the City.
func (c *City) QueryStreets() *StreetQuery {
	return (&CityClient{c.config}).QueryStreets(c)
}

// Update returns a builder for updating this City.
// Note that, you need to call City.Unwrap() before calling this method, if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *City) Update() *CityUpdateOne {
	return (&CityClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *City) Unwrap() *City {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *City) String() string {
	var builder strings.Builder
	builder.WriteString("City(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Cities is a parsable slice of City.
type Cities []*City

func (c Cities) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
