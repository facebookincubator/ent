// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

type Conversion struct {
	ent.Schema
}

// Fields of the Conversion.
func (Conversion) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		// convert integer fields to string
		// Postgres uses the same type for int8 and int16
		// Postgres loses unsigned so we have assume value is signed
		field.String("int8_to_string").MaxLen(6),
		field.String("uint8_to_string").MaxLen(6),
		field.String("int16_to_string").MaxLen(6),
		field.String("uint16_to_string").MaxLen(6),
		field.String("int32_to_string").MaxLen(12),
		field.String("uint32_to_string").MaxLen(12),
		field.String("int64_to_string").MaxLen(21),
		field.String("uint64_to_string").MaxLen(21),
	}
}
