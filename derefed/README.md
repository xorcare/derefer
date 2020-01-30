# derefed

[![goreportcard.com](https://goreportcard.com/badge/github.com/xorcare/derefer/derefed)][GRC]
[![godoc.org](https://godoc.org/github.com/xorcare/derefer/derefed?status.svg)][DOC]

Package derefed contains helper routines for simplifying the getting of optional
fields of basic type. This allows you to get the value from the pointer even,
but if the pointer is nil it returns a value of second parameter.

## Examples

Examples of using the library are presented on [godoc.org][EXD]
and in the [source library code][EXS].

## FAQ

| Question | Source |
| -------- | ------ |
| How to dereference **bool** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ bool = derefed.Bool(new(bool), true)` |
| How to dereference **byte** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ byte = derefed.Byte(new(byte), 13)` |
| How to dereference **complex64** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ complex64 = derefed.Complex64(new(complex64), 6.4)` |
| How to dereference **complex128** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ complex128 = derefed.Complex128(new(complex128), 12.8)` |
| How to dereference **float32** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ float32 = derefed.Float32(new(float32), 3.2)` |
| How to dereference **float64** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ float64 = derefed.Float64(new(float64), 6.4)` |
| How to dereference **int** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ int = derefed.Int(new(int), 42)` |
| How to dereference **int8** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ int8 = derefed.Int8(new(int8), 127)` |
| How to dereference **int16** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ int16 = derefed.Int16(new(int16), 16)` |
| How to dereference **int32** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ int32 = derefed.Int32(new(int32), 32)` |
| How to dereference **int64** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ int64 = derefed.Int64(new(int64), 64)` |
| How to dereference **rune** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ rune = derefed.Rune(new(rune), 67)` |
| How to dereference **string** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ string = derefed.String(new(string), "derefed")` |
| How to dereference **uint** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ uint = derefed.Uint(new(uint), 451)` |
| How to dereference **uint8** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ uint8 = derefed.Uint8(new(uint8), 255)` |
| How to dereference **uint16** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ uint16 = derefed.Uint16(new(uint16), 2033)` |
| How to dereference **uint32** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ uint32 = derefed.Uint32(new(uint32), 2077)` |
| How to dereference **uint64** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ uint64 = derefed.Uint64(new(uint64), 20000)` |
| How to dereference **uintptr** pointer from a struct literal or variable and received custom value if pointer is nil? | `var _ uintptr = derefed.Uintptr(new(uintptr), 2112)` |

## License

© Vasiliy Vasilyuk, 2019-2020

Released under the [BSD 3-Clause License][LIC].

[LIC]: https://github.com/xorcare/derefer/blob/master/LICENSE 'BSD 3-Clause "New" or "Revised" License'
[EXD]: https://godoc.org/github.com/xorcare/derefer/derefed#pkg-examples 'Examples of using package derefer'
[EXS]: https://github.com/xorcare/derefer/blob/master/derefed/example_test.go 'Examples source file'
[GRC]: https://goreportcard.com/report/github.com/xorcare/derefer/derefed 'A web application that generates a report on the quality of an open source go project'
[DOC]: https://godoc.org/github.com/xorcare/derefer 'GoDoc hosts documentation for Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad'
