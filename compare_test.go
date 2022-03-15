package gen

import (
	"testing"
)

type DSErr[T comparable] struct {
	a, b []T
	err  error
}

func CompareTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]DSErr[T],
) {
	GenTest[DSErr[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test DSErr[T]) {
			err := Compare(test.a, test.b)

			if err != test.err {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}

			if !Equal(test.a, test.b) && test.err == nil {
				t.Errorf("Expected %v, got %v", test.a, test.b)
			}
		})
}

func Test_Compare(t *testing.T) {
	CompareTest(t, "string", map[string]DSErr[string]{
		"single": {
			a:   []string{"a"},
			b:   []string{"a"},
			err: nil,
		},
		"duplicate": {
			a:   []string{"a", "a"},
			b:   []string{"a", "a"},
			err: nil,
		},
		"multiple": {
			a:   []string{"a", "b", "a"},
			b:   []string{"a", "b", "a"},
			err: nil,
		},
		"multiple-no-duplicate": {
			a:   []string{"a", "b", "c"},
			b:   []string{"a", "b", "c"},
			err: nil,
		},
		"length-mismatch": {
			a:   []string{"a", "b", "c"},
			b:   []string{"a", "b"},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []string{"a", "b", "c"},
			b:   []string{"a", "b", "d"},
			err: ErrIndexMismatch[string]{2, "c", "d"},
		},
		"empty": {
			a:   []string{},
			b:   []string{},
			err: nil,
		},
	})

	CompareTest(t, "rune", map[string]DSErr[rune]{
		"single": {
			a:   []rune{'a'},
			b:   []rune{'a'},
			err: nil,
		},
		"duplicate": {
			a:   []rune{'a', 'a'},
			b:   []rune{'a', 'a'},
			err: nil,
		},
		"multiple": {
			a:   []rune{'a', 'b', 'a'},
			b:   []rune{'a', 'b', 'a'},
			err: nil,
		},
		"multiple-no-duplicate": {
			a:   []rune{'a', 'b', 'c'},
			b:   []rune{'a', 'b', 'c'},
			err: nil,
		},
		"length-mismatch": {
			a:   []rune{'a', 'b', 'c'},
			b:   []rune{'a', 'b'},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []rune{'a', 'b', 'c'},
			b:   []rune{'a', 'b', 'd'},
			err: ErrIndexMismatch[rune]{2, 'c', 'd'},
		},
		"empty": {
			a:   []rune{},
			b:   []rune{},
			err: nil,
		},
	})

	CompareTest(t, "byte", map[string]DSErr[byte]{
		"single": {
			a:   []byte{'a'},
			b:   []byte{'a'},
			err: nil,
		},
		"duplicate": {
			a:   []byte{'a', 'a'},
			b:   []byte{'a', 'a'},
			err: nil,
		},
		"multiple": {
			a:   []byte{'a', 'b', 'a'},
			b:   []byte{'a', 'b', 'a'},
			err: nil,
		},
		"multiple-no-duplicate": {
			a:   []byte{'a', 'b', 'c'},
			b:   []byte{'a', 'b', 'c'},
			err: nil,
		},
		"length-mismatch": {
			a:   []byte{'a', 'b', 'c'},
			b:   []byte{'a', 'b'},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []byte{'a', 'b', 'c'},
			b:   []byte{'a', 'b', 'd'},
			err: ErrIndexMismatch[byte]{2, 'c', 'd'},
		},
		"empty": {
			a:   []byte{},
			b:   []byte{},
			err: nil,
		},
	})

	CompareTest(t, "int", map[string]DSErr[int]{
		"single": {
			a:   []int{1},
			b:   []int{1},
			err: nil,
		},
		"duplicate": {
			a:   []int{1, 1},
			b:   []int{1, 1},
			err: nil,
		},
		"multiple": {
			a:   []int{1, 2, 1},
			b:   []int{1, 2, 1},
			err: nil,
		},
		"length-mismatch": {
			a:   []int{1, 2, 1},
			b:   []int{1, 2},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []int{1, 2, 1},
			b:   []int{1, 2, 3},
			err: ErrIndexMismatch[int]{2, 1, 3},
		},
		"empty": {
			a:   []int{},
			b:   []int{},
			err: nil,
		},
	})

	CompareTest(t, "float", map[string]DSErr[float32]{
		"single": {
			a:   []float32{1.0},
			b:   []float32{1.0},
			err: nil,
		},
		"duplicate": {
			a:   []float32{1.0, 1.0},
			b:   []float32{1.0, 1.0},
			err: nil,
		},
		"multiple": {
			a:   []float32{1.0, 2.0, 1.0},
			b:   []float32{1.0, 2.0, 1.0},
			err: nil,
		},
		"length-mismatch": {
			a:   []float32{1.0, 2.0, 1.0},
			b:   []float32{1.0, 2.0},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []float32{1.0, 2.0, 1.0},
			b:   []float32{1.0, 2.0, 3.0},
			err: ErrIndexMismatch[float32]{2, 1.0, 3.0},
		},
		"empty": {
			a:   []float32{},
			b:   []float32{},
			err: nil,
		},
	})

	CompareTest(t, "complex", map[string]DSErr[complex64]{
		"single": {
			a:   []complex64{1.0 + 2.0i},
			b:   []complex64{1.0 + 2.0i},
			err: nil,
		},
		"duplicate": {
			a:   []complex64{1.0 + 2.0i, 1.0 + 2.0i},
			b:   []complex64{1.0 + 2.0i, 1.0 + 2.0i},
			err: nil,
		},
		"multiple": {
			a:   []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			b:   []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			err: nil,
		},
		"length-mismatch": {
			a:   []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			b:   []complex64{1.0 + 2.0i, 2.0 + 3.0i},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			b:   []complex64{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
			err: ErrIndexMismatch[complex64]{2, 1.0 + 2.0i, 3.0 + 4.0i},
		},
		"empty": {
			a:   []complex64{},
			b:   []complex64{},
			err: nil,
		},
	})

	CompareTest(t, "chan", map[string]DSErr[chan int]{
		"single": {
			a:   []chan int{c1},
			b:   []chan int{c1},
			err: nil,
		},
		"duplicate": {
			a:   []chan int{c1, c2},
			b:   []chan int{c1, c2},
			err: nil,
		},
		"multiple": {
			a:   []chan int{c1, c2, c3},
			b:   []chan int{c1, c2, c3},
			err: nil,
		},
		"length-mismatch": {
			a:   []chan int{c1, c2, c3},
			b:   []chan int{c1, c2},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []chan int{c1, c2, c3},
			b:   []chan int{c1, c2, c2},
			err: ErrIndexMismatch[chan int]{2, c3, c2},
		},
		"empty": {
			a:   []chan int{},
			b:   []chan int{},
			err: nil,
		},
	})

	CompareTest(t, "struct", map[string]DSErr[testStruct]{
		"single": {
			a: []testStruct{
				{A: 1, B: "a"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
			},
			err: nil,
		},
		"duplicate": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
			},
			err: nil,
		},
		"multiple": {
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
			err: nil,
		},
		"length-mismatch": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
			},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 3, B: "c"},
			},
			b: []testStruct{
				{A: 1, B: "a"},
				{A: 2, B: "b"},
				{A: 4, B: "d"},
			},
			err: ErrIndexMismatch[testStruct]{
				2,
				testStruct{A: 3, B: "c"},
				testStruct{A: 4, B: "d"},
			},
		},
		"empty": {
			a:   []testStruct{},
			b:   []testStruct{},
			err: nil,
		},
	})

	CompareTest(t, "pointer", map[string]DSErr[*testStruct]{
		"single": {
			a:   []*testStruct{tp1},
			b:   []*testStruct{tp1},
			err: nil,
		},
		"duplicate": {
			a:   []*testStruct{tp1, tp2},
			b:   []*testStruct{tp1, tp2},
			err: nil,
		},
		"multiple": {
			a:   []*testStruct{tp1, tp2, tp3},
			b:   []*testStruct{tp1, tp2, tp3},
			err: nil,
		},
		"length-mismatch": {
			a:   []*testStruct{tp1, tp2, tp3},
			b:   []*testStruct{tp1, tp2},
			err: ErrLengthMismatch{3, 2},
		},
		"index-mismatch": {
			a:   []*testStruct{tp1, tp2, tp3},
			b:   []*testStruct{tp1, tp2, tp2},
			err: ErrIndexMismatch[*testStruct]{2, tp3, tp2},
		},
		"empty": {
			a:   []*testStruct{},
			b:   []*testStruct{},
			err: nil,
		},
	})
}
