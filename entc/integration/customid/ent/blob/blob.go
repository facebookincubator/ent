// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package blob

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the blob type in the database.
	Label = "blob"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeLinks holds the string denoting the links edge name in mutations.
	EdgeLinks = "links"
	// Table holds the table name of the blob in the database.
	Table = "blobs"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "blobs"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "blob_parent"
	// LinksTable is the table the holds the links relation/edge. The primary key declared below.
	LinksTable = "blob_links"
)

// Columns holds all SQL columns for blob fields.
var Columns = []string{
	FieldID,
	FieldUUID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "blobs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"blob_parent",
}

var (
	// LinksPrimaryKey and LinksColumn2 are the table columns denoting the
	// primary key for the links relation (M2M).
	LinksPrimaryKey = []string{"blob_id", "link_id"}
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

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
