// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package pet

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldLicensedAt holds the string denoting the licensed_at field in the database.
	FieldLicensedAt = "licensed_at"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the pet in the database.
	Table = "pets"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "pets"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_pets"
)

// Columns holds all SQL columns for pet fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldLicensedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Pet type.
var ForeignKeys = []string{
	"user_pets",
}
