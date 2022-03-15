package gen

import (
	"testing"
)

type testStruct struct {
	A int
	B string
}

type DualSlice[T comparable] struct {
	a, b []T
}

func UniqueTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]DualSlice[T],
) {
	GenTest[DualSlice[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test DualSlice[T]) {
			result := Intersect(Unique(test.a), test.b)
			if len(result) != len(test.b) {
				t.Fatalf("expected full intersection; got %v", result)
			}
		})
}

func Test_Unique(t *testing.T) {
	UniqueTest(t, "string", map[string]DualSlice[string]{
		"single": {
			a: []string{"a"},
			b: []string{"a"},
		},
		"duplicate": {
			a: []string{"a", "a"},
			b: []string{"a"},
		},
		"multiple": {
			a: []string{"a", "b", "a"},
			b: []string{"a", "b"},
		},
		"multiple-no-duplicate": {
			a: []string{"a", "b", "c"},
			b: []string{"a", "b", "c"},
		},
		"empty": {
			a: []string{},
			b: []string{},
		},
	})

	UniqueTest(t, "int", map[string]DualSlice[int]{
		"single": {
			a: []int{1},
			b: []int{1},
		},
		"duplicate": {
			a: []int{1, 1},
			b: []int{1},
		},
		"multiple": {
			a: []int{1, 2, 1},
			b: []int{1, 2},
		},
		"multiple-no-duplicate": {
			a: []int{1, 2, 3},
			b: []int{1, 2, 3},
		},
		"empty": {
			a: []int{},
			b: []int{},
		},
	})

	UniqueTest(t, "float", map[string]DualSlice[float32]{
		"single": {
			a: []float32{1.0},
			b: []float32{1.0},
		},
		"duplicate": {
			a: []float32{1.0, 1.0},
			b: []float32{1.0},
		},
		"multiple": {
			a: []float32{1.0, 2.0, 1.0},
			b: []float32{1.0, 2.0},
		},
		"multiple-no-duplicate": {
			a: []float32{1.0, 2.0, 3.0},
			b: []float32{1.0, 2.0, 3.0},
		},
		"empty": {
			a: []float32{},
			b: []float32{},
		},
	})

	UniqueTest(t, "complex", map[string]DualSlice[complex64]{
		"single": {
			a: []complex64{1.0 + 2.0i},
			b: []complex64{1.0 + 2.0i},
		},
		"duplicate": {
			a: []complex64{1.0 + 2.0i, 1.0 + 2.0i},
			b: []complex64{1.0 + 2.0i},
		},
		"multiple": {
			a: []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			b: []complex64{1.0 + 2.0i, 2.0 + 3.0i},
		},
		"multiple-no-duplicate": {
			a: []complex64{1.0 + 2.0i, 2.0 + 3.0i, 4.0 + 5.0i},
			b: []complex64{1.0 + 2.0i, 2.0 + 3.0i, 4.0 + 5.0i},
		},
		"empty": {
			a: []complex64{},
			b: []complex64{},
		},
	})

	UniqueTest(t, "byte", map[string]DualSlice[byte]{
		"single": {
			a: []byte{'a'},
			b: []byte{'a'},
		},
		"duplicate": {
			a: []byte{'a', 'a'},
			b: []byte{'a'},
		},
		"multiple": {
			a: []byte{'a', 'b', 'a'},
			b: []byte{'a', 'b'},
		},
		"multiple-no-duplicate": {
			a: []byte{'a', 'b', 'c'},
			b: []byte{'a', 'b', 'c'},
		},
		"empty": {
			a: []byte{},
			b: []byte{},
		},
	})

	UniqueTest(t, "rune", map[string]DualSlice[rune]{
		"single": {
			a: []rune{'a'},
			b: []rune{'a'},
		},
		"duplicate": {
			a: []rune{'a', 'a'},
			b: []rune{'a'},
		},
		"multiple": {
			a: []rune{'a', 'b', 'a'},
			b: []rune{'a', 'b'},
		},
		"multiple-no-duplicate": {
			a: []rune{'a', 'b', 'c'},
			b: []rune{'a', 'b', 'c'},
		},
		"empty": {
			a: []rune{},
			b: []rune{},
		},
	})

	UniqueTest(t, "chan", map[string]DualSlice[chan int]{
		"single": {
			a: []chan int{c1},
			b: []chan int{c1},
		},
		"duplicate": {
			a: []chan int{c1, c1},
			b: []chan int{c1},
		},
		"multiple": {
			a: []chan int{c1, c2, c1},
			b: []chan int{c1, c2},
		},
		"multiple-no-duplicate": {
			a: []chan int{c1, c2, c3},
			b: []chan int{c1, c2, c3},
		},
		"empty": {
			a: []chan int{},
			b: []chan int{},
		},
	})

	UniqueTest(t, "struct", map[string]DualSlice[testStruct]{
		"single": {
			a: []testStruct{
				{A: 1, B: "a"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
			},
		},
		"duplicate": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 1, B: "a"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
			},
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
		},
		"empty": {
			a: []testStruct{},
			b: []testStruct{},
		},
	})

	UniqueTest(t, "struct-pointer", map[string]DualSlice[*testStruct]{
		"single": {
			a: []*testStruct{tp1},
			b: []*testStruct{tp1},
		},
		"duplicate": {
			a: []*testStruct{tp1, tp1},
			b: []*testStruct{tp1},
		},
		"multiple": {
			a: []*testStruct{tp1, tp2, tp1},
			b: []*testStruct{tp1, tp2},
		},
		"multiple-no-duplicate": {
			a: []*testStruct{tp1, tp2, tp3},
			b: []*testStruct{tp1, tp2, tp3},
		},
		"empty": {
			a: []*testStruct{},
			b: []*testStruct{},
		},
	})
}
