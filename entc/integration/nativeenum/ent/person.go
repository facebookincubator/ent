// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/nativeenum/ent/mood"
	"github.com/facebook/ent/entc/integration/nativeenum/ent/person"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Mood holds the value of the "mood" field.
	Mood mood.Mood `json:"mood,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // mood
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(values ...interface{}) error {
	if m, n := len(values), len(person.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pe.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mood", values[0])
	} else if value.Valid {
		pe.Mood = mood.Mood(value.String)
	}
	return nil
}

// Update returns a builder for updating this Person.
// Note that, you need to call Person.Unwrap() before calling this method, if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return (&PersonClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", mood=")
	builder.WriteString(fmt.Sprintf("%v", pe.Mood))
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person

func (pe Persons) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}