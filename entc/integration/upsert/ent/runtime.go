// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/ent/entc/integration/upsert/ent/schema"
	"entgo.io/ent/entc/integration/upsert/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescUpdateCount is the schema descriptor for updateCount field.
	userDescUpdateCount := userFields[1].Descriptor()
	// user.DefaultUpdateCount holds the default value on creation for the updateCount field.
	user.DefaultUpdateCount = userDescUpdateCount.Default.(int)
}
