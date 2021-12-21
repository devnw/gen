# `gen` The Generic General Use Go Functions Library

> ***NOTE:*** This project is currently under active development
> and REQUIRES Go 1.18beta1 or higher.

[![Build & Test Action Status](https://github.com/structsdev/gen/actions/workflows/build.yml/badge.svg)](https://github.com/structsdev/gen/actions)
[![Go Report Card](https://goreportcard.com/badge/go.structs.dev/gen)](https://goreportcard.com/report/go.structs.dev/gen)
[![codecov](https://codecov.io/gh/structsdev/gen/branch/main/graph/badge.svg)](https://codecov.io/gh/structsdev/gen)
[![Go Reference](https://pkg.go.dev/badge/go.structs.dev/gen.svg)](#documentation)
[![License: Apache 2.0](https://img.shields.io/badge/license-Apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

`gen` is a generic general use Go functions library with the intention of
replacing duplicated code where the same functionality is needed across multiple
types, and provides a common interface for the functionality.

The library is designed to be used in a wide variety of projects and is
designed to be easy to use.

## Installation

```bash
go get -u go.structs.dev/gen@latest
```

## Import

It is recommended that you use a `.` import for this library so that the
library functions are available in the current namespace without the `gen`
prefix.

Example:

```go

    import . "go.structs.dev/gen"

    func main() {
        // ...

        m := Map[string, string]{
            "foo": "bar",
            "bar": "baz",
        }

        // Get a slice of map keys
        keys := m.Keys()

        // ...

        slice1 := []string{"foo", "bar"}
        slice2 := []string{"floob", "bar"}

        // Call the Dedup method without the `gen` prefix.
        unique := Dedup(slice1, slice2)
    }

```

## Documentation

Until pkg.go.dev support the 1.18 release, here is the documentation for the
current version of the library.

The types defined in this library are:

Map and Slice which provide helper functions for maps and slices. They are alias
types for the standard Go map and slice types so you can cast from a map or
slice to a Map or Slice directly without a type assertion to access the methods.

### FUNCTIONS

**`func Compare[T comparable](a, b []T) error`**
> Compare returns nil if the two input slices are equal otherwise it returns
    an error indicating the issue

**`func Diff[T comparable](a, b []T) []T`**
>Diff returns the symmetric difference between the two input slices

**`func Equal[T comparable](a, b []T) bool`**
>Equal returns true if the two input slices are exactly equal

**`func Exclude[T comparable](a []T, b ...T) []T`**
>Exclude returns a slice of values from the input slice minus the values
    supplied in the second argument

**`func Has[T comparable](in []T, v T) bool`**
>Has determines if the input slice contains the input value

**`func Index[T comparable](in []T, v T) int`**
>Index returns the index of the first occurrence of the input value

**`func Indices[T comparable](in []T, v T) []int`**
>Indices returns a slice of indices for all occurrences of value `v`

**`func Intersect[T comparable](a, b []T) []T`**
>Intersect returns the intersection between the two input slices

**`func Match[T comparable](a, b []T) bool`**
>Match returns true if the two input slices contain equivalent values NOTE:
    Match ignores ordering

**`func Unique[U ~[]T, T comparable](in U) []T`**
>Unique returns a slice of unique values from the input slice

## Benchmarks

To execute the benchmarks, run the following command:

```bash
go test -bench=. ./...
```

To view benchmarks over time for the `main` branch of the repository they can
be seen on our [Benchmark Report Card](https://structsdev.github.io/gen/dev/bench/).
