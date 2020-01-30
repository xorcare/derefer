package derefer_test

import (
	. "fmt"
	"time"

	"github.com/xorcare/derefer"
)

var _bool = true
var _byte = byte(32)
var _complex64 = complex64(36.8478697)
var _complex128 = complex128(174.7652964)
var _float32 = float32(-43.5299773)
var _float64 = float64(172.6233322)
var _int = int(1)
var _int8 = int8(8)
var _int16 = int16(16)
var _int32 = int32(32)
var _int64 = int64(12)
var _rune = rune(127)
var _string = "derefer"
var _uint = uint(19)
var _uint8 = uint8(8)
var _uint16 = uint16(16)
var _uint32 = uint32(32)
var _uint64 = uint64(94)
var _uintptr = uintptr(64)
var _time = time.Time{}.AddDate(2033, 1, 1)

func ExampleAny() {
	{
		var tm *time.Time
		value := derefer.Any(tm).(time.Time)
		Printf("dereference typed nil: (%T)\t%q\n", value, Sprint(value))
	}
	{
		value := derefer.Any(time.Time{}).(time.Time)
		Printf("no dereference a value: (%T)\t%q\n", value, Sprint(value))
	}
	{
		value := derefer.Any(new(time.Time)).(time.Time)
		Printf("dereference new zero time: (%T)\t%q\n", value, Sprint(value))
	}
	{
		value := derefer.Any(nil)
		Printf("dereference non typed nil: (%T)\t%q\n", value, Sprint(value))
	}

	// Output:
	// dereference typed nil: (time.Time)	"0001-01-01 00:00:00 +0000 UTC"
	// no dereference a value: (time.Time)	"0001-01-01 00:00:00 +0000 UTC"
	// dereference new zero time: (time.Time)	"0001-01-01 00:00:00 +0000 UTC"
	// dereference non typed nil: (<nil>)	"<nil>"
}

func ExampleBool() {
	value := derefer.Bool(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Bool(new(bool))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Bool(&_bool)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (bool)	"false"
	// dereference new zero value: (bool)	"false"
	// dereference variable value: (bool)	"true"
}

func ExampleByte() {
	value := derefer.Byte(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Byte(new(byte))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Byte(&_byte)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uint8)	"0"
	// dereference new zero value: (uint8)	"0"
	// dereference variable value: (uint8)	"32"
}

func ExampleComplex64() {
	value := derefer.Complex64(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Complex64(new(complex64))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Complex64(&_complex64)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (complex64)	"(0+0i)"
	// dereference new zero value: (complex64)	"(0+0i)"
	// dereference variable value: (complex64)	"(36.84787+0i)"
}

func ExampleComplex128() {
	value := derefer.Complex128(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Complex128(new(complex128))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Complex128(&_complex128)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (complex128)	"(0+0i)"
	// dereference new zero value: (complex128)	"(0+0i)"
	// dereference variable value: (complex128)	"(174.7652964+0i)"
}

func ExampleFloat32() {
	value := derefer.Float32(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Float32(new(float32))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Float32(&_float32)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (float32)	"0"
	// dereference new zero value: (float32)	"0"
	// dereference variable value: (float32)	"-43.529976"
}

func ExampleFloat64() {
	value := derefer.Float64(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Float64(new(float64))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Float64(&_float64)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (float64)	"0"
	// dereference new zero value: (float64)	"0"
	// dereference variable value: (float64)	"172.6233322"
}

func ExampleInt() {
	value := derefer.Int(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int(new(int))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int(&_int)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (int)	"0"
	// dereference new zero value: (int)	"0"
	// dereference variable value: (int)	"1"
}

func ExampleInt8() {
	value := derefer.Int8(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int8(new(int8))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int8(&_int8)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (int8)	"0"
	// dereference new zero value: (int8)	"0"
	// dereference variable value: (int8)	"8"
}

func ExampleInt16() {
	value := derefer.Int16(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int16(new(int16))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int16(&_int16)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (int16)	"0"
	// dereference new zero value: (int16)	"0"
	// dereference variable value: (int16)	"16"
}

func ExampleInt32() {
	value := derefer.Int32(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int32(new(int32))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int32(&_int32)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (int32)	"0"
	// dereference new zero value: (int32)	"0"
	// dereference variable value: (int32)	"32"
}

func ExampleInt64() {
	value := derefer.Int64(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int64(new(int64))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Int64(&_int64)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (int64)	"0"
	// dereference new zero value: (int64)	"0"
	// dereference variable value: (int64)	"12"
}

func ExampleRune() {
	value := derefer.Rune(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Rune(new(rune))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Rune(&_rune)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (int32)	"0"
	// dereference new zero value: (int32)	"0"
	// dereference variable value: (int32)	"127"
}

func ExampleString() {
	value := derefer.String(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.String(new(string))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.String(&_string)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (string)	""
	// dereference new zero value: (string)	""
	// dereference variable value: (string)	"derefer"
}

func ExampleUint() {
	value := derefer.Uint(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint(new(uint))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint(&_uint)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uint)	"0"
	// dereference new zero value: (uint)	"0"
	// dereference variable value: (uint)	"19"
}

func ExampleUint8() {
	value := derefer.Uint8(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint8(new(uint8))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint8(&_uint8)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uint8)	"0"
	// dereference new zero value: (uint8)	"0"
	// dereference variable value: (uint8)	"8"
}

func ExampleUint16() {
	value := derefer.Uint16(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint16(new(uint16))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint16(&_uint16)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uint16)	"0"
	// dereference new zero value: (uint16)	"0"
	// dereference variable value: (uint16)	"16"
}

func ExampleUint32() {
	value := derefer.Uint32(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint32(new(uint32))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint32(&_uint32)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uint32)	"0"
	// dereference new zero value: (uint32)	"0"
	// dereference variable value: (uint32)	"32"
}

func ExampleUint64() {
	value := derefer.Uint64(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint64(new(uint64))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uint64(&_uint64)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uint64)	"0"
	// dereference new zero value: (uint64)	"0"
	// dereference variable value: (uint64)	"94"
}

func ExampleUintptr() {
	value := derefer.Uintptr(nil)
	Printf("dereference a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uintptr(new(uintptr))
	Printf("dereference new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Uintptr(&_uintptr)
	Printf("dereference variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference a nil pointer : (uintptr)	"0"
	// dereference new zero value: (uintptr)	"0"
	// dereference variable value: (uintptr)	"64"
}

func ExampleTime() {
	value := derefer.Time(nil)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefer.Time(new(time.Time))
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefer.Time(&_time)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (time.Time)	"0001-01-01 00:00:00 +0000 UTC"
	// dereference with default value new zero value: (time.Time)	"0001-01-01 00:00:00 +0000 UTC"
	// dereference with default value variable value: (time.Time)	"2034-02-02 00:00:00 +0000 UTC"
}
