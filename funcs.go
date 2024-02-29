// Copyright 2022 Developer Network
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

// Unique returns a slice of unique values from the input slice
func Unique[U ~[]T, T comparable](in U) []T {
	return Slice[T](in).Map().Keys()
}

// Has determines if the input slice contains the input value
func Has[T comparable](in []T, v T) bool {
	return Index(in, v) != -1
}

// Match returns true if the two input slices contain
// equivalent values
// NOTE: Match ignores ordering
func Match[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := map[T]int{}
	for _, v := range a {
		counts[v]++
	}

	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}

	return true
}

// Indices returns a slice of indices for all occurrences of value `v`
func Indices[T comparable](in []T, v T) []int {
	out := []int{}

	for i, vv := range in {
		if vv == v {
			out = append(out, i)
		}
	}

	return out
}

// Index returns the index of the first occurrence of the input value
func Index[T comparable](in []T, v T) int {
	for i, vv := range in {
		if vv == v {
			return i
		}
	}
	return -1
}

func comp[T comparable](a, b []T) int {
	for i, v := range a {
		if v != b[i] {
			return i
		}
	}

	return -1
}

// Equal returns true if the two input slices are exactly equal
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	if comp(a, b) == -1 {
		return true
	}

	return false
}

// Compare returns nil if the two input slices are equal
// otherwise it returns an error indicating the issue
func Compare[T comparable](a, b []T) error {
	if len(a) != len(b) {
		return LengthMismatchError{
			Expected: len(a),
			Actual:   len(b),
		}
	}

	if len(a) == 0 {
		return nil
	}

	i := comp(a, b)
	if i == -1 {
		return nil
	}

	return IndexMismatchError[T]{i, a[i], b[i]}
}

// Intersect returns the intersection between the two input slices
func Intersect[T comparable](a, b []T) []T {
	out := []T{}

	for _, v := range a {
		if Has(b, v) {
			out = append(out, v)
		}
	}

	return out
}

// Diff returns the symmetric difference between the two input slices
func Diff[T comparable](a, b []T) []T {
	// Determine intersection of `a` and `b`
	intersect := Intersect(a, b)

	// Append `a` and `b`
	//nolint:gocritic // this is on purpose
	all := append(a, b...)

	// Exclude intersection from `all`
	return Exclude(all, intersect...)
}

// Exclude returns a slice of values from the input slice minus
// the values supplied in the second argument
func Exclude[T comparable](a []T, b ...T) []T {
	out := []T{}

	// Convert exclusions slice to map
	excludes := Slice[T](b).Map()

	// Iterate over data and create a new slice with
	// exclusions removed
	for _, v := range a {
		_, exists := excludes[v]
		if !exists {
			out = append(out, v)
		}
	}

	return out
}

// As allows you to cast N values passed in through a variadic
// argument to a slice of N values. This is useful for casting
// disparate struct types to slices of implemented interfaces
//
// Example: As[io.Reader](&bytes.Buffer{}, &bufio.Reader{})
func As[T any](in ...T) []T {
	return in
}

// ReadOnly accepts a slice of channels and returns a slice
// of channels that are read-only.
func ReadOnly[U chan T, T any](in ...U) []<-chan T {
	out := make([]<-chan T, 0, len(in))

	for _, v := range in {
		out = append(out, v)
	}

	return out
}

// WriteOnly accepts a slice of channels and returns a slice
// of channels that are write-only.
func WriteOnly[U chan T, T any](in ...U) []chan<- T {
	out := make([]chan<- T, 0, len(in))

	for _, v := range in {
		out = append(out, v)
	}

	return out
}

// channel defines a constraint that a value must be a channel
// or a read-only channel of a specific type
type channel[T any] interface {
	chan T | chan<- T
}

// Close closes all channels in the input slice
//
// NOTE: If a channel is already closed any panic will
// be ignored and the channel will be skipped
func Close[U channel[T], T any](in ...U) {
	for _, v := range in {
		func() {
			defer func() {
				_ = recover()
			}()

			close(v)
		}()
	}
}
