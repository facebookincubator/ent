// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/migrate/entv1/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Renamed applies equality check predicate on the "renamed" field. It's identical to RenamedEQ.
func Renamed(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRenamed), v))
	})
}

// Blob applies equality check predicate on the "blob" field. It's identical to BlobEQ.
func Blob(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlob), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Workplace applies equality check predicate on the "workplace" field. It's identical to WorkplaceEQ.
func Workplace(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkplace), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int32) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int32) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int32) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickname), v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNickname), v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNickname), v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickname), v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickname), v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickname), v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickname), v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickname), v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickname), v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickname), v))
	})
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickname), v))
	})
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickname), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAddress)))
	})
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAddress)))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// RenamedEQ applies the EQ predicate on the "renamed" field.
func RenamedEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRenamed), v))
	})
}

// RenamedNEQ applies the NEQ predicate on the "renamed" field.
func RenamedNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRenamed), v))
	})
}

// RenamedIn applies the In predicate on the "renamed" field.
func RenamedIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRenamed), v...))
	})
}

// RenamedNotIn applies the NotIn predicate on the "renamed" field.
func RenamedNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRenamed), v...))
	})
}

// RenamedGT applies the GT predicate on the "renamed" field.
func RenamedGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRenamed), v))
	})
}

// RenamedGTE applies the GTE predicate on the "renamed" field.
func RenamedGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRenamed), v))
	})
}

// RenamedLT applies the LT predicate on the "renamed" field.
func RenamedLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRenamed), v))
	})
}

// RenamedLTE applies the LTE predicate on the "renamed" field.
func RenamedLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRenamed), v))
	})
}

// RenamedContains applies the Contains predicate on the "renamed" field.
func RenamedContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRenamed), v))
	})
}

// RenamedHasPrefix applies the HasPrefix predicate on the "renamed" field.
func RenamedHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRenamed), v))
	})
}

// RenamedHasSuffix applies the HasSuffix predicate on the "renamed" field.
func RenamedHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRenamed), v))
	})
}

// RenamedIsNil applies the IsNil predicate on the "renamed" field.
func RenamedIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRenamed)))
	})
}

// RenamedNotNil applies the NotNil predicate on the "renamed" field.
func RenamedNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRenamed)))
	})
}

// RenamedEqualFold applies the EqualFold predicate on the "renamed" field.
func RenamedEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRenamed), v))
	})
}

// RenamedContainsFold applies the ContainsFold predicate on the "renamed" field.
func RenamedContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRenamed), v))
	})
}

// BlobEQ applies the EQ predicate on the "blob" field.
func BlobEQ(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlob), v))
	})
}

// BlobNEQ applies the NEQ predicate on the "blob" field.
func BlobNEQ(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlob), v))
	})
}

// BlobIn applies the In predicate on the "blob" field.
func BlobIn(vs ...[]byte) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBlob), v...))
	})
}

// BlobNotIn applies the NotIn predicate on the "blob" field.
func BlobNotIn(vs ...[]byte) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBlob), v...))
	})
}

// BlobGT applies the GT predicate on the "blob" field.
func BlobGT(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlob), v))
	})
}

// BlobGTE applies the GTE predicate on the "blob" field.
func BlobGTE(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlob), v))
	})
}

// BlobLT applies the LT predicate on the "blob" field.
func BlobLT(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlob), v))
	})
}

// BlobLTE applies the LTE predicate on the "blob" field.
func BlobLTE(v []byte) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlob), v))
	})
}

// BlobIsNil applies the IsNil predicate on the "blob" field.
func BlobIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBlob)))
	})
}

// BlobNotNil applies the NotNil predicate on the "blob" field.
func BlobNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBlob)))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	})
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// WorkplaceEQ applies the EQ predicate on the "workplace" field.
func WorkplaceEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkplace), v))
	})
}

// WorkplaceNEQ applies the NEQ predicate on the "workplace" field.
func WorkplaceNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkplace), v))
	})
}

// WorkplaceIn applies the In predicate on the "workplace" field.
func WorkplaceIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkplace), v...))
	})
}

// WorkplaceNotIn applies the NotIn predicate on the "workplace" field.
func WorkplaceNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkplace), v...))
	})
}

// WorkplaceGT applies the GT predicate on the "workplace" field.
func WorkplaceGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkplace), v))
	})
}

// WorkplaceGTE applies the GTE predicate on the "workplace" field.
func WorkplaceGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkplace), v))
	})
}

// WorkplaceLT applies the LT predicate on the "workplace" field.
func WorkplaceLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkplace), v))
	})
}

// WorkplaceLTE applies the LTE predicate on the "workplace" field.
func WorkplaceLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkplace), v))
	})
}

// WorkplaceContains applies the Contains predicate on the "workplace" field.
func WorkplaceContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWorkplace), v))
	})
}

// WorkplaceHasPrefix applies the HasPrefix predicate on the "workplace" field.
func WorkplaceHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWorkplace), v))
	})
}

// WorkplaceHasSuffix applies the HasSuffix predicate on the "workplace" field.
func WorkplaceHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWorkplace), v))
	})
}

// WorkplaceIsNil applies the IsNil predicate on the "workplace" field.
func WorkplaceIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWorkplace)))
	})
}

// WorkplaceNotNil applies the NotNil predicate on the "workplace" field.
func WorkplaceNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWorkplace)))
	})
}

// WorkplaceEqualFold applies the EqualFold predicate on the "workplace" field.
func WorkplaceEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWorkplace), v))
	})
}

// WorkplaceContainsFold applies the ContainsFold predicate on the "workplace" field.
func WorkplaceContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWorkplace), v))
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSpouse applies the HasEdge predicate on the "spouse" edge.
func HasSpouse() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpouseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SpouseTable, SpouseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpouseWith applies the HasEdge predicate on the "spouse" edge with a given conditions (other predicates).
func HasSpouseWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SpouseTable, SpouseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCar applies the HasEdge predicate on the "car" edge.
func HasCar() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CarTable, CarFieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CarTable, CarColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCarWith applies the HasEdge predicate on the "car" edge with a given conditions (other predicates).
func HasCarWith(preds ...predicate.Car) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CarInverseTable, CarFieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CarTable, CarColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
