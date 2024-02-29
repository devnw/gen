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

type DSMatch[T comparable] struct {
	a, b  []T
	match bool
}

func MatchTest[T comparable](t *testing.T, testname string, testdata map[string]DSMatch[T]) {
	GenTest[DSMatch[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test DSMatch[T]) {
			match := Match(test.a, test.b)

			if match != test.match {
				t.Fatalf("expected %v; got %v", test.match, match)
			}
		})
}

func Test_Match(t *testing.T) {
	MatchTest(t, "string", map[string]DSMatch[string]{
		"single": {
			a:     []string{"a"},
			b:     []string{"a"},
			match: true,
		},
		"mismatch len, same value": {

			a:     []string{"a", "a"},
			b:     []string{"a"},
			match: false,
		},
		"same, first two, mismatch end": {
			a:     []string{"a", "b", "a"},
			b:     []string{"a", "b"},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []string{"a", "b", "c"},
			b:     []string{"a", "b", "c"},
			match: true,
		},
		"multiple-mismatch": {
			a:     []string{"a", "b", "c"},
			b:     []string{"a", "c", "c"},
			match: false,
		},
		"duplicate in b": {
			a:     []string{"a", "b", "c", "c"},
			b:     []string{"a", "b", "c", "c"},
			match: true,
		},
		"empty": {
			a:     []string{},
			b:     []string{},
			match: true,
		},
	})

	MatchTest(t, "rune", map[string]DSMatch[rune]{
		"single": {
			a:     []rune{'a'},
			b:     []rune{'a'},
			match: true,
		},
		"mismatch len, same value": {
			a:     []rune{'a', 'a'},
			b:     []rune{'a'},
			match: false,
		},
		"same, first two, mismatch end": {
			a:     []rune{'a', 'b', 'a'},
			b:     []rune{'a', 'b'},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []rune{'a', 'b', 'c'},
			b:     []rune{'a', 'b', 'c'},
			match: true,
		},
		"multiple-mismatch": {
			a:     []rune{'a', 'b', 'c'},
			b:     []rune{'a', 'c', 'c'},
			match: false,
		},
		"duplicate in b": {
			a:     []rune{'a', 'b', 'c', 'c'},
			b:     []rune{'a', 'b', 'c', 'c'},
			match: true,
		},
		"empty": {
			a:     []rune{},
			b:     []rune{},
			match: true,
		},
	})

	MatchTest(t, "byte", map[string]DSMatch[byte]{
		"single": {
			a:     []byte{'a'},
			b:     []byte{'a'},
			match: true,
		},
		"mismatch len, same value": {
			a:     []byte{'a', 'a'},
			b:     []byte{'a'},
			match: false,
		},
		"same, first two, mismatch end": {
			a:     []byte{'a', 'b', 'a'},
			b:     []byte{'a', 'b'},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []byte{'a', 'b', 'c'},
			b:     []byte{'a', 'b', 'c'},
			match: true,
		},
		"multiple-mismatch": {
			a:     []byte{'a', 'b', 'c'},
			b:     []byte{'a', 'c', 'c'},
			match: false,
		},
		"duplicate in b": {
			a:     []byte{'a', 'b', 'c', 'c'},
			b:     []byte{'a', 'b', 'c', 'c'},
			match: true,
		},
		"empty": {
			a:     []byte{},
			b:     []byte{},
			match: true,
		},
	})

	MatchTest(t, "int", map[string]DSMatch[int]{
		"single": {
			a:     []int{1},
			b:     []int{1},
			match: true,
		},
		"duplicate": {
			a:     []int{1, 1},
			b:     []int{1},
			match: false,
		},
		"multiple": {
			a:     []int{1, 2, 1},
			b:     []int{1, 2},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []int{1, 2, 3},
			b:     []int{1, 2, 3},
			match: true,
		},
		"multiple-mismatch": {
			a:     []int{1, 2, 3},
			b:     []int{1, 3, 3},
			match: false,
		},
		"duplicate in b": {
			a:     []int{1, 2, 3, 3},
			b:     []int{1, 2, 3, 3},
			match: true,
		},
		"empty": {
			a:     []int{},
			b:     []int{},
			match: true,
		},
	})

	MatchTest(t, "float32", map[string]DSMatch[float32]{
		"single": {
			a:     []float32{1.0},
			b:     []float32{1.0},
			match: true,
		},
		"duplicate": {
			a:     []float32{1.0, 1.0},
			b:     []float32{1.0},
			match: false,
		},
		"multiple": {
			a:     []float32{1.0, 2.0, 1.0},
			b:     []float32{1.0, 2.0},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []float32{1.0, 2.0, 3.0},
			b:     []float32{1.0, 2.0, 3.0},
			match: true,
		},
		"multiple-mismatch": {
			a:     []float32{1.0, 2.0, 3.0},
			b:     []float32{1.0, 3.0, 3.0},
			match: false,
		},
		"duplicate in b": {
			a:     []float32{1.0, 2.0, 3.0, 3.0},
			b:     []float32{1.0, 2.0, 3.0, 3.0},
			match: true,
		},
		"empty": {
			a:     []float32{},
			b:     []float32{},
			match: true,
		},
	})

	MatchTest(t, "complex", map[string]DSMatch[complex128]{
		"single": {
			a:     []complex128{1 + 2i},
			b:     []complex128{1 + 2i},
			match: true,
		},
		"mismatch len, same value": {
			a:     []complex128{1 + 2i, 1 + 2i},
			b:     []complex128{1 + 2i},
			match: false,
		},
		"same, first two, mismatch end": {
			a:     []complex128{1 + 2i, 2 + 3i, 1 + 2i},
			b:     []complex128{1 + 2i, 2 + 3i},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []complex128{1 + 2i, 2 + 3i, 3 + 4i},
			b:     []complex128{1 + 2i, 2 + 3i, 3 + 4i},
			match: true,
		},
		"multiple-mismatch": {
			a:     []complex128{1 + 2i, 2 + 3i, 3 + 4i},
			b:     []complex128{1 + 2i, 3 + 4i, 3 + 4i},
			match: false,
		},
		"duplicate in b": {
			a:     []complex128{1 + 2i, 2 + 3i, 3 + 4i, 3 + 4i},
			b:     []complex128{1 + 2i, 2 + 3i, 3 + 4i, 3 + 4i},
			match: true,
		},
		"empty": {
			a:     []complex128{},
			b:     []complex128{},
			match: true,
		},
	})

	MatchTest(t, "chan", map[string]DSMatch[chan int]{
		"single": {
			a:     []chan int{c1},
			b:     []chan int{c1},
			match: true,
		},
		"duplicate": {
			a:     []chan int{c1, c2},
			b:     []chan int{c1},
			match: false,
		},
		"multiple": {
			a:     []chan int{c1, c2, c3},
			b:     []chan int{c1, c2},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []chan int{c1, c2, c3},
			b:     []chan int{c1, c2, c3},
			match: true,
		},
		"multiple-mismatch": {
			a:     []chan int{c1, c2, c3},
			b:     []chan int{c1, c3, c3},
			match: false,
		},
		"duplicate in b": {
			a:     []chan int{c1, c2, c3, c3},
			b:     []chan int{c1, c2, c3, c3},
			match: true,
		},
		"empty": {
			a:     []chan int{},
			b:     []chan int{},
			match: true,
		},
	})

	MatchTest(t, "struct", map[string]DSMatch[testStruct]{
		"single": {
			a: []testStruct{
				{A: 1, B: "a"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
			},
			match: true,
		},
		"duplicate": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 1, B: "a"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
			},
			match: false,
		},
		"multiple": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 1, B: "a"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
			},
			match: false,
		},
		"multiple-no-duplicate": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
			},
			match: true,
		},
		"multiple-mismatch": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 3, B: "c"},
				{A: 3, B: "c"},
			},
			match: false,
		},
		"duplicate in b": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
				{A: 4, B: "c"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
				{A: 4, B: "c"},
			},
			match: true,
		},
		"empty": {
			a:     []testStruct{},
			b:     []testStruct{},
			match: true,
		},
	})

	MatchTest(t, "pointer_struct", map[string]DSMatch[*testStruct]{
		"single": {
			a:     []*testStruct{tp1},
			b:     []*testStruct{tp1},
			match: true,
		},
		"duplicate": {
			a:     []*testStruct{tp1, tp2},
			b:     []*testStruct{tp1},
			match: false,
		},
		"multiple": {
			a:     []*testStruct{tp1, tp2, tp3},
			b:     []*testStruct{tp1, tp2},
			match: false,
		},
		"multiple-no-duplicate": {
			a:     []*testStruct{tp1, tp2, tp3},
			b:     []*testStruct{tp1, tp2, tp3},
			match: true,
		},
		"multiple-mismatch": {
			a:     []*testStruct{tp1, tp2, tp3},
			b:     []*testStruct{tp1, tp3, tp3},
			match: false,
		},
		"duplicate in b": {
			a:     []*testStruct{tp1, tp2, tp3, tp3},
			b:     []*testStruct{tp1, tp2, tp3, tp3},
			match: true,
		},
		"empty": {
			a:     []*testStruct{},
			b:     []*testStruct{},
			match: true,
		},
	})
}
