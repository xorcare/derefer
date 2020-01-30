package derefed_test

import (
	. "fmt"

	"github.com/xorcare/derefer/derefed"
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
var _string = "derefed"
var _uint = uint(19)
var _uint8 = uint8(8)
var _uint16 = uint16(16)
var _uint32 = uint32(32)
var _uint64 = uint64(94)
var _uintptr = uintptr(64)

func ExampleBool() {
	value := derefed.Bool(nil, true)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Bool(new(bool), true)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Bool(&_bool, false)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (bool)	"true"
	// dereference with default value new zero value: (bool)	"false"
	// dereference with default value variable value: (bool)	"true"
}

func ExampleByte() {
	value := derefed.Byte(nil, 13)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Byte(new(byte), 13)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Byte(&_byte, 13)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uint8)	"13"
	// dereference with default value new zero value: (uint8)	"0"
	// dereference with default value variable value: (uint8)	"32"
}

func ExampleComplex64() {
	value := derefed.Complex64(nil, 6.4)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Complex64(new(complex64), 6.4)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Complex64(&_complex64, 6.4)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (complex64)	"(6.4+0i)"
	// dereference with default value new zero value: (complex64)	"(0+0i)"
	// dereference with default value variable value: (complex64)	"(36.84787+0i)"
}

func ExampleComplex128() {
	value := derefed.Complex128(nil, 12.8)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Complex128(new(complex128), 12.8)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Complex128(&_complex128, 12.8)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (complex128)	"(12.8+0i)"
	// dereference with default value new zero value: (complex128)	"(0+0i)"
	// dereference with default value variable value: (complex128)	"(174.7652964+0i)"
}

func ExampleFloat32() {
	value := derefed.Float32(nil, 3.2)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Float32(new(float32), 3.2)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Float32(&_float32, 3.2)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (float32)	"3.2"
	// dereference with default value new zero value: (float32)	"0"
	// dereference with default value variable value: (float32)	"-43.529976"
}

func ExampleFloat64() {
	value := derefed.Float64(nil, 6.4)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Float64(new(float64), 6.4)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Float64(&_float64, 6.4)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (float64)	"6.4"
	// dereference with default value new zero value: (float64)	"0"
	// dereference with default value variable value: (float64)	"172.6233322"
}

func ExampleInt() {
	value := derefed.Int(nil, 42)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int(new(int), 42)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int(&_int, 42)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (int)	"42"
	// dereference with default value new zero value: (int)	"0"
	// dereference with default value variable value: (int)	"1"
}

func ExampleInt8() {
	value := derefed.Int8(nil, 127)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int8(new(int8), 127)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int8(&_int8, 127)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (int8)	"127"
	// dereference with default value new zero value: (int8)	"0"
	// dereference with default value variable value: (int8)	"8"
}

func ExampleInt16() {
	value := derefed.Int16(nil, 2033)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int16(new(int16), 2033)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int16(&_int16, 2033)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (int16)	"2033"
	// dereference with default value new zero value: (int16)	"0"
	// dereference with default value variable value: (int16)	"16"
}

func ExampleInt32() {
	value := derefed.Int32(nil, 2077)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int32(new(int32), 2077)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int32(&_int32, 2077)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (int32)	"2077"
	// dereference with default value new zero value: (int32)	"0"
	// dereference with default value variable value: (int32)	"32"
}

func ExampleInt64() {
	value := derefed.Int64(nil, 40000)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int64(new(int64), 40000)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Int64(&_int64, 40000)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (int64)	"40000"
	// dereference with default value new zero value: (int64)	"0"
	// dereference with default value variable value: (int64)	"12"
}

func ExampleRune() {
	value := derefed.Rune(nil, 13)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Rune(new(rune), 13)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Rune(&_rune, 13)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (int32)	"13"
	// dereference with default value new zero value: (int32)	"0"
	// dereference with default value variable value: (int32)	"127"
}

func ExampleString() {
	value := derefed.String(nil, "derefer")
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.String(new(string), "derefer")
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.String(&_string, "derefer")
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (string)	"derefer"
	// dereference with default value new zero value: (string)	""
	// dereference with default value variable value: (string)	"derefed"
}

func ExampleUint() {
	value := derefed.Uint(nil, 451)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint(new(uint), 451)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint(&_uint, 451)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uint)	"451"
	// dereference with default value new zero value: (uint)	"0"
	// dereference with default value variable value: (uint)	"19"
}

func ExampleUint8() {
	value := derefed.Uint8(nil, 255)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint8(new(uint8), 255)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint8(&_uint8, 255)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uint8)	"255"
	// dereference with default value new zero value: (uint8)	"0"
	// dereference with default value variable value: (uint8)	"8"
}

func ExampleUint16() {
	value := derefed.Uint16(nil, 2033)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint16(new(uint16), 2033)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint16(&_uint16, 2033)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uint16)	"2033"
	// dereference with default value new zero value: (uint16)	"0"
	// dereference with default value variable value: (uint16)	"16"
}

func ExampleUint32() {
	value := derefed.Uint32(nil, 2077)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint32(new(uint32), 2077)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint32(&_uint32, 2077)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uint32)	"2077"
	// dereference with default value new zero value: (uint32)	"0"
	// dereference with default value variable value: (uint32)	"32"
}

func ExampleUint64() {
	value := derefed.Uint64(nil, 20000)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint64(new(uint64), 20000)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uint64(&_uint64, 20000)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uint64)	"20000"
	// dereference with default value new zero value: (uint64)	"0"
	// dereference with default value variable value: (uint64)	"94"
}

func ExampleUintptr() {
	value := derefed.Uintptr(nil, 47)
	Printf("dereference with default value a nil pointer : (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uintptr(new(uintptr), 47)
	Printf("dereference with default value new zero value: (%T)\t%q\n", value, Sprint(value))

	value = derefed.Uintptr(&_uintptr, 47)
	Printf("dereference with default value variable value: (%T)\t%q\n", value, Sprint(value))

	// Output:
	// dereference with default value a nil pointer : (uintptr)	"47"
	// dereference with default value new zero value: (uintptr)	"0"
	// dereference with default value variable value: (uintptr)	"64"
}
