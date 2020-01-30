// Copyright © 2019 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package derefer contains helper routines for simplifying the getting of
// optional fields of basic type. This allows you to get the value from the
// pointer even if it is nil, because in this case the zero value of the
// specified type will be received.
package derefer

import (
	"reflect"
	"time"
)

// Any dereference a pointer any type from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Any(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Ptr {
		if !r.IsNil() {
			return r.Elem().Interface()
		}
		return reflect.New(r.Type().Elem()).Elem().Interface()
	}
	return r.Interface()
}

// Bool dereference a pointer bool from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Bool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// Byte dereference a pointer byte from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Byte(v *byte) byte {
	if v == nil {
		return 0
	}
	return *v
}

// Complex64 dereference a pointer complex64 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Complex64(v *complex64) complex64 {
	if v == nil {
		return 0
	}
	return *v
}

// Complex128 dereference a pointer complex128 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Complex128(v *complex128) complex128 {
	if v == nil {
		return 0
	}
	return *v
}

// Float32 dereference a pointer float32 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Float32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

// Float64 dereference a pointer float64 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Float64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

// Int dereference a pointer int from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Int(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// Int8 dereference a pointer int8 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Int8(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

// Int16 dereference a pointer int16 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Int16(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

// Int32 dereference a pointer int32 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Int32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// Int64 dereference a pointer int64 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Int64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

// Rune dereference a pointer rune from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Rune(v *rune) rune {
	if v == nil {
		return 0
	}
	return *v
}

// String dereference a pointer string from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func String(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// Time dereference a pointer time.Time from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Time(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}

// Uint dereference a pointer uint from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Uint(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

// Uint8 dereference a pointer uint8 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Uint8(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint16 dereference a pointer uint16 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Uint16(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint32 dereference a pointer uint32 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Uint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint64 dereference a pointer uint64 from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Uint64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

// Uintptr dereference a pointer uintptr from the structure of a literal or
// variable, and if the pointer is nil it returns a zero value of this type.
func Uintptr(v *uintptr) uintptr {
	if v == nil {
		return 0
	}
	return *v
}
