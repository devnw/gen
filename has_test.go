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

type HasData[T comparable] struct {
	in    []T
	value T
	has   bool
}

func HasTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]HasData[T],
) {
	GenTest[HasData[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test HasData[T]) {
			out := Has(test.in, test.value)

			if out != test.has {
				t.Fatalf("expected %v; got %v", test.has, out)
			}
		})
}

func Test_Has(t *testing.T) {
	HasTest(t, "string", map[string]HasData[string]{
		"single": {
			in:    []string{"a"},
			value: "a",
			has:   true,
		},
		"duplicate": {
			in:    []string{"a", "b", "a"},
			value: "a",
			has:   true,
		},
		"does not have": {
			in:    []string{"a", "b", "a"},
			value: "c",
			has:   false,
		},
		"empty": {
			in:    []string{},
			value: "a",
			has:   false,
		},
	})

	HasTest(t, "rune", map[string]HasData[rune]{
		"single": {
			in:    []rune{'a'},
			value: 'a',
			has:   true,
		},
		"duplicate": {
			in:    []rune{'a', 'b', 'a'},
			value: 'a',
			has:   true,
		},
		"does not have": {
			in:    []rune{'a', 'b', 'a'},
			value: 'c',
			has:   false,
		},
		"empty": {
			in:    []rune{},
			value: 'a',
			has:   false,
		},
	})

	HasTest(t, "byte", map[string]HasData[byte]{
		"single": {
			in:    []byte{'a'},
			value: 'a',
			has:   true,
		},
		"duplicate": {
			in:    []byte{'a', 'b', 'a'},
			value: 'a',
			has:   true,
		},
		"does not have": {
			in:    []byte{'a', 'b', 'a'},
			value: 'c',
			has:   false,
		},
		"empty": {
			in:    []byte{},
			value: 'a',
			has:   false,
		},
	})

	HasTest(t, "int", map[string]HasData[int]{
		"single": {
			in:    []int{1},
			value: 1,
			has:   true,
		},
		"duplicate": {
			in:    []int{1, 2, 1},
			value: 1,
			has:   true,
		},
		"multiple": {
			in:    []int{1, 2, 3},
			value: 1,
			has:   true,
		},
		"multiple-no-duplicate": {
			in:    []int{1, 2, 3},
			value: 4,
			has:   false,
		},
		"empty": {
			in:    []int{},
			value: 1,
			has:   false,
		},
	})

	HasTest(t, "float", map[string]HasData[float32]{
		"single": {
			in:    []float32{1.0},
			value: 1.0,
			has:   true,
		},
		"duplicate": {
			in:    []float32{1.0, 2.0, 1.0},
			value: 1.0,
			has:   true,
		},
		"multiple": {
			in:    []float32{1.0, 2.0, 3.0},
			value: 1.0,
			has:   true,
		},
		"multiple-no-duplicate": {
			in:    []float32{1.0, 2.0, 3.0},
			value: 4.0,
			has:   false,
		},
		"empty": {
			in:    []float32{},
			value: 1.0,
			has:   false,
		},
	})

	HasTest(t, "chan", map[string]HasData[chan int]{
		"single": {
			in:    []chan int{c1},
			value: c1,
			has:   true,
		},
		"duplicate": {
			in:    []chan int{c1, c2, c1},
			value: c1,
			has:   true,
		},
		"multiple": {
			in:    []chan int{c1, c2, c3},
			value: c1,
			has:   true,
		},
		"multiple-no-duplicate": {
			in:    []chan int{c1, c2},
			value: c3,
			has:   false,
		},
		"empty": {
			in:    []chan int{},
			value: c1,
			has:   false,
		},
	})

	HasTest(t, "struct", map[string]HasData[testStruct]{
		"single": {
			in:    []testStruct{{1, "a"}},
			value: testStruct{1, "a"},
			has:   true,
		},
		"duplicate": {
			in: []testStruct{
				{1, "a"},
				{2, "b"},
				{1, "a"},
			},
			value: testStruct{1, "a"},
			has:   true,
		},
		"multiple": {
			in: []testStruct{
				{1, "a"},
				{2, "b"},
				{3, "c"},
			},
			value: testStruct{1, "a"},
			has:   true,
		},
		"multiple-no-duplicate": {
			in: []testStruct{
				{1, "a"},
				{2, "b"},
				{3, "c"},
			},
			value: testStruct{4, "d"},
			has:   false,
		},
		"empty": {
			in:    []testStruct{},
			value: testStruct{1, "a"},
			has:   false,
		}})

	HasTest(t, "pointer", map[string]HasData[*testStruct]{
		"single": {
			in:    []*testStruct{tp1},
			value: tp1,
			has:   true,
		},
		"duplicate": {
			in: []*testStruct{
				tp1,
				tp2,
				tp1,
			},
			value: tp1,
			has:   true,
		},
		"multiple": {
			in: []*testStruct{
				tp1,
				tp2,
				tp3,
			},
			value: tp1,
			has:   true,
		},
		"multiple-no-duplicate": {
			in: []*testStruct{
				tp1,
				tp2,
			},
			value: tp3,
			has:   false,
		},
		"empty": {
			in:    []*testStruct{},
			value: tp1,
			has:   false,
		},
	})
}
