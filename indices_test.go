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

import (
	"testing"
)

type IndicesData[T comparable] struct {
	in      []T
	value   T
	indices []int
}

func IndicesTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]IndicesData[T],
) {
	GenTest[IndicesData[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test IndicesData[T]) {
			i := Indices(test.in, test.value)

			if !Equal(i, test.indices) {
				t.Fatalf("expected %v; got %v", test.indices, i)
			}
		})
}

func Test_Indices(t *testing.T) {
	IndicesTest(t, "string", map[string]IndicesData[string]{
		"single": {
			in:      []string{"a"},
			value:   "a",
			indices: []int{0},
		},
		"two": {
			in:      []string{"a", "b", "a"},
			value:   "a",
			indices: []int{0, 2},
		},
		"many": {
			in:      []string{"a", "b", "a", "b", "a", "a", "a", "b"},
			value:   "a",
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []string{"a", "b", "a"},
			value:   "c",
			indices: []int{},
		},
		"empty": {
			in:      []string{},
			value:   "a",
			indices: []int{},
		},
	})

	IndicesTest(t, "rune", map[string]IndicesData[rune]{
		"single": {
			in:      []rune{'a'},
			value:   'a',
			indices: []int{0},
		},
		"two": {
			in:      []rune{'a', 'b', 'a'},
			value:   'a',
			indices: []int{0, 2},
		},
		"many": {
			in:      []rune{'a', 'b', 'a', 'b', 'a', 'a', 'a', 'b'},
			value:   'a',
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []rune{'a', 'b', 'a'},
			value:   'c',
			indices: []int{},
		},
		"empty": {
			in:      []rune{},
			value:   'a',
			indices: []int{},
		},
	})

	IndicesTest(t, "byte", map[string]IndicesData[byte]{
		"single": {
			in:      []byte{'a'},
			value:   'a',
			indices: []int{0},
		},
		"second": {
			in:      []byte{'a', 'b', 'a'},
			value:   'a',
			indices: []int{0, 2},
		},
		"many": {
			in:      []byte{'a', 'b', 'a', 'b', 'a', 'a', 'a', 'b'},
			value:   'a',
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []byte{'a', 'b', 'a'},
			value:   'c',
			indices: []int{},
		},
		"empty": {
			in:      []byte{},
			value:   'a',
			indices: []int{},
		},
	})

	IndicesTest(t, "int", map[string]IndicesData[int]{
		"single": {
			in:      []int{1},
			value:   1,
			indices: []int{0},
		},
		"second": {
			in:      []int{1, 2, 1},
			value:   1,
			indices: []int{0, 2},
		},
		"many": {
			in:      []int{1, 2, 1, 2, 1, 1, 1, 2},
			value:   1,
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []int{1, 2, 1},
			value:   3,
			indices: []int{},
		},
		"empty": {
			in:      []int{},
			value:   1,
			indices: []int{},
		},
	})

	IndicesTest(t, "float64", map[string]IndicesData[float64]{
		"single": {
			in:      []float64{1.0},
			value:   1.0,
			indices: []int{0},
		},
		"second": {
			in:      []float64{1.0, 2.0, 1.0},
			value:   1.0,
			indices: []int{0, 2},
		},
		"many": {
			in:      []float64{1.0, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 2.0},
			value:   1.0,
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []float64{1.0, 2.0, 1.0},
			value:   3.0,
			indices: []int{},
		},
		"empty": {
			in:      []float64{},
			value:   1.0,
			indices: []int{},
		},
	})

	IndicesTest(t, "complex64", map[string]IndicesData[complex64]{
		"single": {
			in:      []complex64{1.0 + 2.0i},
			value:   1.0 + 2.0i,
			indices: []int{0},
		},
		"second": {
			in:      []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			value:   1.0 + 2.0i,
			indices: []int{0, 2},
		},
		"many": {
			in: []complex64{
				1.0 + 2.0i,
				2.0 + 3.0i,
				1.0 + 2.0i,
				2.0 + 3.0i,
				1.0 + 2.0i,
				1.0 + 2.0i,
				1.0 + 2.0i,
				2.0 + 3.0i,
			},
			value:   1.0 + 2.0i,
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			value:   3.0 + 4.0i,
			indices: []int{},
		},
		"empty": {
			in:      []complex64{},
			value:   1.0 + 2.0i,
			indices: []int{},
		},
	})

	IndicesTest(t, "chan", map[string]IndicesData[chan int]{
		"single": {
			in:      []chan int{c1},
			value:   c1,
			indices: []int{0},
		},
		"second": {
			in:      []chan int{c1, c2, c1},
			value:   c1,
			indices: []int{0, 2},
		},
		"many": {
			in: []chan int{
				c1,
				c2,
				c1,
				c2,
				c1,
				c1,
				c1,
				c2,
			},
			value:   c1,
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []chan int{c1, c2, c1},
			value:   c3,
			indices: []int{},
		},
		"empty": {
			in:      []chan int{},
			value:   c1,
			indices: []int{},
		},
	})

	IndicesTest(t, "struct", map[string]IndicesData[testStruct]{
		"single": {
			in:      []testStruct{{1, "a"}},
			value:   testStruct{1, "a"},
			indices: []int{0},
		},
		"second": {
			in:      []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
			value:   testStruct{1, "a"},
			indices: []int{0, 2},
		},
		"many": {
			in: []testStruct{
				{1, "a"},
				{2, "b"},
				{1, "a"},
				{2, "b"},
				{1, "a"},
				{1, "a"},
				{1, "a"},
				{2, "b"},
			},
			value:   testStruct{1, "a"},
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
			value:   testStruct{3, "c"},
			indices: []int{},
		},
		"empty": {
			in:      []testStruct{},
			value:   testStruct{1, "a"},
			indices: []int{},
		},
	})

	IndicesTest(t, "struct pointer", map[string]IndicesData[*testStruct]{
		"single": {
			in:      []*testStruct{tp1},
			value:   tp1,
			indices: []int{0},
		},
		"second": {
			in:      []*testStruct{tp1, tp2, tp1},
			value:   tp1,
			indices: []int{0, 2},
		},
		"many": {
			in: []*testStruct{
				tp1,
				tp2,
				tp1,
				tp2,
				tp1,
				tp1,
				tp1,
				tp2,
			},
			value:   tp1,
			indices: []int{0, 2, 4, 5, 6},
		},
		"does not have": {
			in:      []*testStruct{tp1, tp2, tp1},
			value:   tp3,
			indices: []int{},
		},
		"empty": {
			in:      []*testStruct{},
			value:   tp1,
			indices: []int{},
		},
	})
}
