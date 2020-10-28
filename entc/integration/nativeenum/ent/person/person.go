// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package person

import (
	"fmt"

	"github.com/facebook/ent/entc/integration/nativeenum/ent/mood"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMood holds the string denoting the mood field in the database.
	FieldMood = "mood"

	// Table holds the table name of the person in the database.
	Table = "persons"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldMood,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// MoodValidator is a validator for the "mood" field enum values. It is called by the builders before save.
func MoodValidator(m mood.Mood) error {
	switch m {
	case "happy", "ok", "sad":
		return nil
	default:
		return fmt.Errorf("person: invalid enum value for mood field: %q", m)
	}
}