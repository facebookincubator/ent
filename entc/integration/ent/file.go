// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// User holds the value of the "user" field.
	User *string `json:"user,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges         FileEdges `json:"edges"`
	type_id       *string
	group_file_id *string
	owner_id      *string
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Type holds the value of the type edge.
	Type *FileType
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerWithError returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) OwnerWithError() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TypeWithError returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) TypeWithError() (*FileType, error) {
	if e.loadedTypes[1] {
		if e.Type == nil {
			// The edge type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: filetype.Label}
		}
		return e.Type, nil
	}
	return nil, &NotLoadedError{edge: "type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // size
		&sql.NullString{}, // name
		&sql.NullString{}, // user
		&sql.NullString{}, // group
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*File) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // type_id
		&sql.NullInt64{}, // group_file_id
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(values ...interface{}) error {
	if m, n := len(values), len(file.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	f.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field size", values[0])
	} else if value.Valid {
		f.Size = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		f.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field user", values[2])
	} else if value.Valid {
		f.User = new(string)
		*f.User = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field group", values[3])
	} else if value.Valid {
		f.Group = value.String
	}
	values = values[4:]
	if len(values) == len(file.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field type_id", value)
		} else if value.Valid {
			f.type_id = new(string)
			*f.type_id = strconv.FormatInt(value.Int64, 10)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field group_file_id", value)
		} else if value.Valid {
			f.group_file_id = new(string)
			*f.group_file_id = strconv.FormatInt(value.Int64, 10)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			f.owner_id = new(string)
			*f.owner_id = strconv.FormatInt(value.Int64, 10)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the File.
func (f *File) QueryOwner() *UserQuery {
	return (&FileClient{f.config}).QueryOwner(f)
}

// QueryType queries the type edge of the File.
func (f *File) QueryType() *FileTypeQuery {
	return (&FileClient{f.config}).QueryType(f)
}

// Update returns a builder for updating this File.
// Note that, you need to call File.Unwrap() before calling this method, if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", name=")
	builder.WriteString(f.Name)
	if v := f.User; v != nil {
		builder.WriteString(", user=")
		builder.WriteString(*v)
	}
	builder.WriteString(", group=")
	builder.WriteString(f.Group)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (f *File) id() int {
	id, _ := strconv.Atoi(f.ID)
	return id
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
