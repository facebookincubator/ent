// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package card

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldOwnerID,
}

// Columns holds all SQL unique columns for card fields that can be used in conflict constraints.
var UniqueColumns = []string{
	FieldID,
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

// ValidConstraintColumn reports if the column name is valid for use as a conflict column.
func ValidConstraintColumn(column string) bool {
	for i := range UniqueColumns {
		if column == UniqueColumns[i] {
			return true
		}
	}
	return false
}
