// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table the holds the users relation/edge. The primary key declared below.
	UsersTable = "group_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// AdminTable is the table the holds the admin relation/edge.
	AdminTable = "groups"
	// AdminInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminInverseTable = "users"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "group_admin"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// Columns holds all SQL unique columns for group fields that can be used in conflict constraints.
var UniqueColumns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_admin",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"group_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
