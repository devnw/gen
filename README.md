# `gen` The Generic General Use Go Functions Library

> ***NOTE:*** This project is currently under active development
> and REQUIRES Go 1.18beta1 or higher.

[![Build & Test Action Status](https://github.com/structsdev/gen/actions/workflows/build.yml/badge.svg)](https://github.com/structsdev/gen/actions)
[![Go Report Card](https://goreportcard.com/badge/go.structs.dev/gen)](https://goreportcard.com/report/go.structs.dev/gen)
[![codecov](https://codecov.io/gh/structsdev/gen/branch/main/graph/badge.svg)](https://codecov.io/gh/structsdev/gen)
[![Go Reference](https://pkg.go.dev/badge/go.structs.dev/gen.svg)](https://pkg.go.dev/go.structs.dev/gen)
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

## Benchmarks

To execute the benchmarks, run the following command:

```bash
go test -bench=. ./...
```

To view benchmarks over time for the `main` branch of the repository they can
be seen on our [Benchmark Report Card](https://structsdev.github.io/gen/dev/bench/).
