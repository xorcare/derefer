# derefer

[![travis-ci.org](https://travis-ci.org/xorcare/derefer.svg?branch=master)][TCI]
[![codecov.io](https://codecov.io/gh/xorcare/derefer/badge.svg)][COV]
[![goreportcard.com](https://goreportcard.com/badge/github.com/xorcare/derefer?)][GRC]
[![godoc.org](https://godoc.org/github.com/xorcare/derefer?status.svg)][DOC]

Package derefer contains helper routines for simplifying the getting of
optional fields of basic type. This allows you to get the value from the
pointer even if it is nil, because in this case the zero value of the
specified type will be received.

## Installation

```bash
go get github.com/xorcare/derefer
```

## Examples

Examples of using the library are presented on [godoc.org][EXD]
and in the [source library code][EXS].

## FAQ

| Question | Source |
| -------- | ------ |
| How to dereference **bool** pointer from a struct literal or variable? | `var _ bool = derefer.Bool(new(bool))` |
| How to dereference **byte** pointer from a struct literal or variable? | `var _ byte = derefer.Byte(new(byte))` |
| How to dereference **complex64** pointer from a struct literal or variable? | `var _ complex64 = derefer.Complex64(new(complex64))` |
| How to dereference **complex128** pointer from a struct literal or variable? | `var _ complex128 = derefer.Complex128(new(complex128))` |
| How to dereference **float32** pointer from a struct literal or variable? | `var _ float32 = derefer.Float32(new(float32))` |
| How to dereference **float64** pointer from a struct literal or variable? | `var _ float64 = derefer.Float64(new(float64))` |
| How to dereference **int** pointer from a struct literal or variable? | `var _ int = derefer.Int(new(int))` |
| How to dereference **int8** pointer from a struct literal or variable? | `var _ int8 = derefer.Int8(new(int8))` |
| How to dereference **int16** pointer from a struct literal or variable? | `var _ int16 = derefer.Int16(new(int16))` |
| How to dereference **int32** pointer from a struct literal or variable? | `var _ int32 = derefer.Int32(new(int32))` |
| How to dereference **int64** pointer from a struct literal or variable? | `var _ int64 = derefer.Int64(new(int64))` |
| How to dereference **rune** pointer from a struct literal or variable? | `var _ rune = derefer.Rune(new(rune))` |
| How to dereference **string** pointer from a struct literal or variable? | `var _ string = derefer.String(new(string))` |
| How to dereference **uint** pointer from a struct literal or variable? | `var _ uint = derefer.Uint(new(uint))` |
| How to dereference **uint8** pointer from a struct literal or variable? | `var _ uint8 = derefer.Uint8(new(uint8))` |
| How to dereference **uint16** pointer from a struct literal or variable? | `var _ uint16 = derefer.Uint16(new(uint16))` |
| How to dereference **uint32** pointer from a struct literal or variable? | `var _ uint32 = derefer.Uint32(new(uint32))` |
| How to dereference **uint64** pointer from a struct literal or variable? | `var _ uint64 = derefer.Uint64(new(uint64))` |
| How to dereference **uintptr** pointer from a struct literal or variable? | `var _ uintptr = derefer.Uintptr(new(uintptr))` |

## License

© Vasiliy Vasilyuk, 2019

Released under the [BSD 3-Clause License][LIC].

[LIC]: https://github.com/xorcare/derefer/blob/master/LICENSE 'BSD 3-Clause "New" or "Revised" License'
[EXD]: https://godoc.org/github.com/xorcare/derefer#pkg-examples 'Examples of using package derefer'
[EXS]: https://github.com/xorcare/derefer/blob/master/example_test.go 'Examples source file'
[TCI]: https://travis-ci.org/xorcare/derefer 'Travis CI is a hosted continuous integration service used to build and test software projects hosted at GitHub'
[COV]: https://codecov.io/gh/xorcare/derefer/branch/master/graph/badge.svg 'Codecov is a code coverage tool, which is available for GitHub, Bitbucket and GitLab'
[GRC]: https://goreportcard.com/report/github.com/xorcare/derefer 'A web application that generates a report on the quality of an open source go project'
[DOC]: https://godoc.org/github.com/xorcare/derefer 'GoDoc hosts documentation for Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad'
