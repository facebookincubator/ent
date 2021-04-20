// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/entc/integration/customid/ent/car"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/mixinid"
	"entgo.io/ent/entc/integration/customid/ent/note"
	"entgo.io/ent/entc/integration/customid/ent/pet"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blobFields := schema.Blob{}.Fields()
	_ = blobFields
	// blobDescUUID is the schema descriptor for uuid field.
	blobDescUUID := blobFields[1].Descriptor()
	// blob.DefaultUUID holds the default value on creation for the uuid field.
	blob.DefaultUUID = blobDescUUID.Default.(func() uuid.UUID)
	// blobDescID is the schema descriptor for id field.
	blobDescID := blobFields[0].Descriptor()
	// blob.DefaultID holds the default value on creation for the id field.
	blob.DefaultID = blobDescID.Default.(func() uuid.UUID)
	carMixin := schema.Car{}.Mixin()
	carMixinFields0 := carMixin[0].Fields()
	_ = carMixinFields0
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescBeforeID is the schema descriptor for before_id field.
	carDescBeforeID := carMixinFields0[0].Descriptor()
	// car.BeforeIDValidator is a validator for the "before_id" field. It is called by the builders before save.
	car.BeforeIDValidator = carDescBeforeID.Validators[0].(func(float64) error)
	// carDescAfterID is the schema descriptor for after_id field.
	carDescAfterID := carMixinFields0[2].Descriptor()
	// car.AfterIDValidator is a validator for the "after_id" field. It is called by the builders before save.
	car.AfterIDValidator = carDescAfterID.Validators[0].(func(float64) error)
	// carDescID is the schema descriptor for id field.
	carDescID := carMixinFields0[1].Descriptor()
	// car.IDValidator is a validator for the "id" field. It is called by the builders before save.
	car.IDValidator = carDescID.Validators[0].(func(int) error)
	docFields := schema.Doc{}.Fields()
	_ = docFields
	// docDescID is the schema descriptor for id field.
	docDescID := docFields[0].Descriptor()
	// doc.DefaultID holds the default value on creation for the id field.
	doc.DefaultID = docDescID.Default.(func() schema.DocID)
	// doc.IDValidator is a validator for the "id" field. It is called by the builders before save.
	doc.IDValidator = func() func(string) error {
		validators := docDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	mixinidMixin := schema.MixinID{}.Mixin()
	mixinidMixinFields0 := mixinidMixin[0].Fields()
	_ = mixinidMixinFields0
	mixinidFields := schema.MixinID{}.Fields()
	_ = mixinidFields
	// mixinidDescID is the schema descriptor for id field.
	mixinidDescID := mixinidMixinFields0[0].Descriptor()
	// mixinid.DefaultID holds the default value on creation for the id field.
	mixinid.DefaultID = mixinidDescID.Default.(func() uuid.UUID)
	noteFields := schema.Note{}.Fields()
	_ = noteFields
	// noteDescID is the schema descriptor for id field.
	noteDescID := noteFields[0].Descriptor()
	// note.DefaultID holds the default value on creation for the id field.
	note.DefaultID = noteDescID.Default.(func() schema.NoteID)
	// note.IDValidator is a validator for the "id" field. It is called by the builders before save.
	note.IDValidator = func() func(string) error {
		validators := noteDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.DefaultID holds the default value on creation for the id field.
	pet.DefaultID = petDescID.Default.(func() string)
	// pet.IDValidator is a validator for the "id" field. It is called by the builders before save.
	pet.IDValidator = func() func(string) error {
		validators := petDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
