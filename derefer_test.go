// Copyright © 2019 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package derefer

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func equal(t testing.TB, expected, actual interface{}) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}

	t.Logf("expected: %q (%T), actual: %q (%T)", fmt.Sprint(expected), expected, fmt.Sprint(actual), actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("values obtained are not equal, expected: %v, actual: %v", expected, actual)
	}
}

var _bool = true
var _byte = byte(1)
var _complex64 = complex64(1.1)
var _complex128 = complex128(1.1)
var _float32 = float32(1.1)
var _float64 = float64(1.1)
var _int = int(1)
var _int8 = int8(8)
var _int16 = int16(16)
var _int32 = int32(32)
var _int64 = int64(64)
var _rune = rune(1)
var _string = "derefer"
var _uint = uint(1)
var _uint8 = uint8(8)
var _uint16 = uint16(16)
var _uint32 = uint32(32)
var _uint64 = uint64(64)
var _uintptr = uintptr(64)

func TestAny(t *testing.T) {
	equal(t, nil, Any(nil))
	equal(t, false, Any(new(bool)).(bool))
	equal(t, _bool, Any(&_bool).(bool))
	equal(t, _bool, Any(_bool).(bool))
	{
		var prt *string
		equal(t, "", Any(prt))
	}
}

func TestBool(t *testing.T) {
	equal(t, false, Bool(nil))
	equal(t, false, Bool(new(bool)))
	equal(t, true, Bool(&_bool))
}

func TestByte(t *testing.T) {
	equal(t, byte(0), Byte(nil))
	equal(t, byte(0), Byte(new(byte)))
	equal(t, byte(1), Byte(&_byte))
}

func TestComplex64(t *testing.T) {
	equal(t, complex64(0), Complex64(nil))
	equal(t, complex64(0), Complex64(new(complex64)))
	equal(t, complex64(1.1), Complex64(&_complex64))
}

func TestComplex128(t *testing.T) {
	equal(t, complex128(0), Complex128(nil))
	equal(t, complex128(0), Complex128(new(complex128)))
	equal(t, complex128(1.1), Complex128(&_complex128))
}

func TestFloat32(t *testing.T) {
	equal(t, float32(0), Float32(nil))
	equal(t, float32(0), Float32(new(float32)))
	equal(t, float32(1.1), Float32(&_float32))
}

func TestFloat64(t *testing.T) {
	equal(t, float64(0), Float64(nil))
	equal(t, float64(0), Float64(new(float64)))
	equal(t, float64(1.1), Float64(&_float64))
}

func TestInt(t *testing.T) {
	equal(t, 0, Int(nil))
	equal(t, 0, Int(new(int)))
	equal(t, 1, Int(&_int))
}

func TestInt8(t *testing.T) {
	equal(t, int8(0), Int8(nil))
	equal(t, int8(0), Int8(new(int8)))
	equal(t, int8(8), Int8(&_int8))
}

func TestInt16(t *testing.T) {
	equal(t, int16(0), Int16(nil))
	equal(t, int16(0), Int16(new(int16)))
	equal(t, int16(16), Int16(&_int16))
}

func TestInt32(t *testing.T) {
	equal(t, int32(0), Int32(nil))
	equal(t, int32(0), Int32(new(int32)))
	equal(t, int32(32), Int32(&_int32))
}

func TestInt64(t *testing.T) {
	equal(t, int64(0), Int64(nil))
	equal(t, int64(0), Int64(new(int64)))
	equal(t, int64(64), Int64(&_int64))
}

func TestRune(t *testing.T) {
	equal(t, rune(0), Rune(nil))
	equal(t, rune(0), Rune(new(rune)))
	equal(t, rune(1), Rune(&_rune))
}

func TestString(t *testing.T) {
	equal(t, "", String(nil))
	equal(t, "", String(new(string)))
	equal(t, "derefer", String(&_string))
}

func TestUint(t *testing.T) {
	equal(t, uint(0), Uint(nil))
	equal(t, uint(0), Uint(new(uint)))
	equal(t, uint(1), Uint(&_uint))
}

func TestUint8(t *testing.T) {
	equal(t, uint8(0), Uint8(nil))
	equal(t, uint8(0), Uint8(new(uint8)))
	equal(t, uint8(8), Uint8(&_uint8))
}

func TestUint16(t *testing.T) {
	equal(t, uint16(0), Uint16(nil))
	equal(t, uint16(0), Uint16(new(uint16)))
	equal(t, uint16(16), Uint16(&_uint16))
}

func TestUint32(t *testing.T) {
	equal(t, uint32(0), Uint32(nil))
	equal(t, uint32(0), Uint32(new(uint32)))
	equal(t, uint32(32), Uint32(&_uint32))
}

func TestUint64(t *testing.T) {
	equal(t, uint64(0), Uint64(nil))
	equal(t, uint64(0), Uint64(new(uint64)))
	equal(t, uint64(64), Uint64(&_uint64))
}

func TestUintptr(t *testing.T) {
	equal(t, uintptr(0), Uintptr(nil))
	equal(t, uintptr(0), Uintptr(new(uintptr)))
	equal(t, uintptr(64), Uintptr(&_uintptr))
}

func TestTime(t *testing.T) {
	equal(t, time.Time{}, Time(nil))
	equal(t, time.Time{}, Time(new(time.Time)))
	{
		now := time.Now()
		equal(t, now, Time(&now))
	}
}
