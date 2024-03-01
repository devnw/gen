# `gen` - Generic Go Utility Library

[![Build & Test Action Status](https://github.com/devnw/gen/actions/workflows/build.yml/badge.svg)](https://github.com/devnw/gen/actions)
[![Go Report Card](https://goreportcard.com/badge/go.devnw.com/gen)](https://goreportcard.com/report/go.devnw.com/gen)
[![Go Reference](https://pkg.go.dev/badge/go.devnw.com/gen.svg)](https://pkg.go.dev/go.devnw.com/gen)
[![License: Apache 2.0](https://img.shields.io/badge/license-Apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

`gen` is a generic general use Go functions library with the intention of
replacing duplicated code where the same functionality is needed across multiple
types, and provides a common interface for the functionality.

The library is designed to be used in a wide variety of projects and is
designed to be easy to use.

## Installation

```bash
go get -u go.devnw.com/gen@latest
```

## Import

Example:

```go

import "go.devnw.com/gen"

func main() {
    // ...

    m := gen.Map[string, string]{
        "foo": "bar",
        "bar": "baz",
    }

    // Get a slice of map keys
    keys := m.Keys()

    // ...

    slice1 := []string{"foo", "bar"}
    slice2 := []string{"floob", "bar"}

    // Call the Unique method without the `gen` prefix.
    unique := gen.Unique(slice1, slice2)

    existingmap := map[string]int{
        "foo": 1223,
        "bar": 111,
    }

    // Cast to the generic Map
    m := gen.Map[string, int](existingmap)

    keys = m.Keys()
    values := m.Values()

    // ...

    existingSlice := []string{"foo", "bar", "baz"}

    // Cast to the generic Slice
    s := gen.Slice[string](existingSlice)

    // Convert slice to map

    m := s.Map()
}

```

## Benchmarks

To execute the benchmarks, run the following command:

```bash
go test -bench=. ./...
```

To view benchmarks over time for the `main` branch of the repository they can
be seen on our [Benchmark Report Card].

[Benchmark Report Card]: https://devnw.github.io/gen/dev/bench/
