// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package derefed contains helper routines for simplifying the getting of
// optional fields of basic type. This allows you to get the value from the
// pointer even if it is nil, and if the pointer is nil it returns a value of
// second parameter.
package derefed

// Bool dereference a pointer bool from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Bool(v *bool, def bool) bool {
	if v == nil {
		return def
	}
	return *v
}

// Byte dereference a pointer byte from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Byte(v *byte, def byte) byte {
	if v == nil {
		return def
	}
	return *v
}

// Complex64 dereference a pointer complex64 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Complex64(v *complex64, def complex64) complex64 {
	if v == nil {
		return def
	}
	return *v
}

// Complex128 dereference a pointer complex128 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Complex128(v *complex128, def complex128) complex128 {
	if v == nil {
		return def
	}
	return *v
}

// Float32 dereference a pointer float32 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Float32(v *float32, def float32) float32 {
	if v == nil {
		return def
	}
	return *v
}

// Float64 dereference a pointer float64 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Float64(v *float64, def float64) float64 {
	if v == nil {
		return def
	}
	return *v
}

// Int dereference a pointer int from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Int(v *int, def int) int {
	if v == nil {
		return def
	}
	return *v
}

// Int8 dereference a pointer int8 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Int8(v *int8, def int8) int8 {
	if v == nil {
		return def
	}
	return *v
}

// Int16 dereference a pointer int16 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Int16(v *int16, def int16) int16 {
	if v == nil {
		return def
	}
	return *v
}

// Int32 dereference a pointer int32 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Int32(v *int32, def int32) int32 {
	if v == nil {
		return def
	}
	return *v
}

// Int64 dereference a pointer int64 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Int64(v *int64, def int64) int64 {
	if v == nil {
		return def
	}
	return *v
}

// Rune dereference a pointer rune from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Rune(v *rune, def rune) rune {
	if v == nil {
		return def
	}
	return *v
}

// String dereference a pointer string from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func String(v *string, def string) string {
	if v == nil {
		return def
	}
	return *v
}

// Uint dereference a pointer uint from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Uint(v *uint, def uint) uint {
	if v == nil {
		return def
	}
	return *v
}

// Uint8 dereference a pointer uint8 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Uint8(v *uint8, def uint8) uint8 {
	if v == nil {
		return def
	}
	return *v
}

// Uint16 dereference a pointer uint16 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Uint16(v *uint16, def uint16) uint16 {
	if v == nil {
		return def
	}
	return *v
}

// Uint32 dereference a pointer uint32 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Uint32(v *uint32, def uint32) uint32 {
	if v == nil {
		return def
	}
	return *v
}

// Uint64 dereference a pointer uint64 from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Uint64(v *uint64, def uint64) uint64 {
	if v == nil {
		return def
	}
	return *v
}

// Uintptr dereference a pointer uintptr from the structure of a literal or
// variable, and if the pointer is nil it returns a value of second parameter.
func Uintptr(v *uintptr, def uintptr) uintptr {
	if v == nil {
		return def
	}
	return *v
}
