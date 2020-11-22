// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/facebook/ent/entc/integration/ent/schema"
	"github.com/facebook/ent/entc/integration/gremlin/ent/card"
	"github.com/facebook/ent/entc/integration/gremlin/ent/fieldtype"
	"github.com/facebook/ent/entc/integration/gremlin/ent/file"
	"github.com/facebook/ent/entc/integration/gremlin/ent/group"
	"github.com/facebook/ent/entc/integration/gremlin/ent/groupinfo"
	"github.com/facebook/ent/entc/integration/gremlin/ent/task"
	"github.com/facebook/ent/entc/integration/gremlin/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinFields0 := cardMixin[0].Fields()
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreateTime is the schema descriptor for create_time field.
	cardDescCreateTime := cardMixinFields0[0].Descriptor()
	// card.DefaultCreateTime holds the default value on creation for the create_time field.
	card.DefaultCreateTime = cardDescCreateTime.Default.(func() time.Time)
	// cardDescUpdateTime is the schema descriptor for update_time field.
	cardDescUpdateTime := cardMixinFields0[1].Descriptor()
	// card.DefaultUpdateTime holds the default value on creation for the update_time field.
	card.DefaultUpdateTime = cardDescUpdateTime.Default.(func() time.Time)
	// card.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	card.UpdateDefaultUpdateTime = cardDescUpdateTime.UpdateDefault.(func() time.Time)
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[0].Descriptor()
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescName is the schema descriptor for name field.
	cardDescName := cardFields[1].Descriptor()
	// card.NameValidator is a validator for the "name" field. It is called by the builders before save.
	card.NameValidator = cardDescName.Validators[0].(func(string) error)
	fieldtypeFields := schema.FieldType{}.Fields()
	_ = fieldtypeFields
	// fieldtypeDescValidateOptionalInt32 is the schema descriptor for validate_optional_int32 field.
	fieldtypeDescValidateOptionalInt32 := fieldtypeFields[15].Descriptor()
	// fieldtype.ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	fieldtype.ValidateOptionalInt32Validator = fieldtypeDescValidateOptionalInt32.Validators[0].(func(int32) error)
	// fieldtypeDescNdir is the schema descriptor for ndir field.
	fieldtypeDescNdir := fieldtypeFields[27].Descriptor()
	// fieldtype.NdirValidator is a validator for the "ndir" field. It is called by the builders before save.
	fieldtype.NdirValidator = fieldtypeDescNdir.Validators[0].(func(string) error)
	// fieldtypeDescLink is the schema descriptor for link field.
	fieldtypeDescLink := fieldtypeFields[30].Descriptor()
	// fieldtype.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	fieldtype.LinkValidator = fieldtypeDescLink.Validators[0].(func(string) error)
	// fieldtypeDescMAC is the schema descriptor for mac field.
	fieldtypeDescMAC := fieldtypeFields[45].Descriptor()
	// fieldtype.MACValidator is a validator for the "mac" field. It is called by the builders before save.
	fieldtype.MACValidator = fieldtypeDescMAC.Validators[0].(func(string) error)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[0].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	filetypeFields := schema.FileType{}.Fields()
	_ = filetypeFields
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescActive is the schema descriptor for active field.
	groupDescActive := groupFields[0].Descriptor()
	// group.DefaultActive holds the default value on creation for the active field.
	group.DefaultActive = groupDescActive.Default.(bool)
	// groupDescType is the schema descriptor for type field.
	groupDescType := groupFields[2].Descriptor()
	// group.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	group.TypeValidator = func() func(string) error {
		validators := groupDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// groupDescMaxUsers is the schema descriptor for max_users field.
	groupDescMaxUsers := groupFields[3].Descriptor()
	// group.DefaultMaxUsers holds the default value on creation for the max_users field.
	group.DefaultMaxUsers = groupDescMaxUsers.Default.(int)
	// group.MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	group.MaxUsersValidator = groupDescMaxUsers.Validators[0].(func(int) error)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[4].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = func() func(string) error {
		validators := groupDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	groupinfoFields := schema.GroupInfo{}.Fields()
	_ = groupinfoFields
	// groupinfoDescMaxUsers is the schema descriptor for max_users field.
	groupinfoDescMaxUsers := groupinfoFields[1].Descriptor()
	// groupinfo.DefaultMaxUsers holds the default value on creation for the max_users field.
	groupinfo.DefaultMaxUsers = groupinfoDescMaxUsers.Default.(int)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescPriority is the schema descriptor for priority field.
	taskDescPriority := taskFields[0].Descriptor()
	// task.DefaultPriority holds the default value on creation for the priority field.
	task.DefaultPriority = schema.Priority(taskDescPriority.Default.(int))
	// task.PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	task.PriorityValidator = taskDescPriority.Validators[0].(func(int) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescOptionalInt is the schema descriptor for optional_int field.
	userDescOptionalInt := userMixinFields0[0].Descriptor()
	// user.OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	user.OptionalIntValidator = userDescOptionalInt.Validators[0].(func(int) error)
	// userDescLast is the schema descriptor for last field.
	userDescLast := userFields[2].Descriptor()
	// user.DefaultLast holds the default value on creation for the last field.
	user.DefaultLast = userDescLast.Default.(string)
}
