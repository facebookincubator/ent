// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/schema"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescMixedString is the schema descriptor for mixed_string field.
	userDescMixedString := userMixinFields0[0].Descriptor()
	// user.DefaultMixedString holds the default value on creation for the mixed_string field.
	user.DefaultMixedString = userDescMixedString.Default.(string)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[4].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// userDescTitle is the schema descriptor for title field.
	userDescTitle := userFields[6].Descriptor()
	// user.DefaultTitle holds the default value on creation for the title field.
	user.DefaultTitle = userDescTitle.Default.(string)
}
