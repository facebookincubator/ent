// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package file

import (
	"strconv"

	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(id)
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.EQ(id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.NEQ(id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i], _ = strconv.Atoi(ids[i])
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
		func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Within(v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i], _ = strconv.Atoi(ids[i])
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
		func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Without(v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GT(id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GTE(id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LT(id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LTE(id))
		},
	)
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.EQ(v))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.EQ(v))
		},
	)
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.EQ(v))
		},
	)
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.EQ(v))
		},
	)
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.NEQ(v))
		},
	)
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldSize), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.Within(v...))
		},
	)
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldSize), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.Without(v...))
		},
	)
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.GT(v))
		},
	)
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.GTE(v))
		},
	)
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.LT(v))
		},
	)
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.LTE(v))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.NEQ(v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldName), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Within(v...))
		},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Without(v...))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GT(v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GTE(v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LT(v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LTE(v))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Containing(v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.StartingWith(v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EndingWith(v))
		},
	)
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.EQ(v))
		},
	)
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.NEQ(v))
		},
	)
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldUser), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.Within(v...))
		},
	)
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldUser), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.Without(v...))
		},
	)
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.GT(v))
		},
	)
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.GTE(v))
		},
	)
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.LT(v))
		},
	)
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.LTE(v))
		},
	)
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.Containing(v))
		},
	)
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.StartingWith(v))
		},
	)
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldUser), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUser, p.EndingWith(v))
		},
	)
}

// UserIsNil applies the IsNil predicate on the "user" field.
func UserIsNil() predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldUser)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldUser)
		},
	)
}

// UserNotNil applies the NotNil predicate on the "user" field.
func UserNotNil() predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldUser)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldUser)
		},
	)
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.EQ(v))
		},
	)
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.NEQ(v))
		},
	)
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldGroup), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.Within(v...))
		},
	)
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldGroup), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.Without(v...))
		},
	)
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.GT(v))
		},
	)
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.GTE(v))
		},
	)
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.LT(v))
		},
	)
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.LTE(v))
		},
	)
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.Containing(v))
		},
	)
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.StartingWith(v))
		},
	)
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldGroup), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldGroup, p.EndingWith(v))
		},
	)
}

// GroupIsNil applies the IsNil predicate on the "group" field.
func GroupIsNil() predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldGroup)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldGroup)
		},
	)
}

// GroupNotNil applies the NotNil predicate on the "group" field.
func GroupNotNil() predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldGroup)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldGroup)
		},
	)
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(OwnerTable, FieldID),
				sql.Edge(sql.M2O, true, OwnerTable, OwnerColumn),
			)
			sql.HasNeighbors(s, step)
		},
		func(t *dsl.Traversal) {
			t.InE(OwnerInverseLabel).InV()
		},
	)
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(FieldID).From(builder.Table(OwnerInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(OwnerColumn), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p(tr)
			}
			t.InE(OwnerInverseLabel).Where(tr).InV()
		},
	)
}

// HasType applies the HasEdge predicate on the "type" edge.
func HasType() predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(TypeTable, FieldID),
				sql.Edge(sql.M2O, true, TypeTable, TypeColumn),
			)
			sql.HasNeighbors(s, step)
		},
		func(t *dsl.Traversal) {
			t.InE(TypeInverseLabel).InV()
		},
	)
}

// HasTypeWith applies the HasEdge predicate on the "type" edge with a given conditions (other predicates).
func HasTypeWith(preds ...predicate.FileType) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(FieldID).From(builder.Table(TypeInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(TypeColumn), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p(tr)
			}
			t.InE(TypeInverseLabel).Where(tr).InV()
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.And(trs...))
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.Or(trs...))
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			p(s.Not())
		},
		func(tr *dsl.Traversal) {
			t := __.New()
			p(t)
			tr.Where(__.Not(t))
		},
	)
}
