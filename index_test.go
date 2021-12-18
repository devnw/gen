package gen

import (
	"io"
	"testing"
)

type IndexData[T comparable] struct {
	in    []T
	value T
	index int
}

func IndexTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]IndexData[T],
) {
	GenTest[IndexData[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test IndexData[T]) {
			i := Index(test.in, test.value)

			if i != test.index {
				t.Fatalf("expected %v; got %v", test.index, i)
			}
		})
}

func Test_Index(t *testing.T) {
	IndexTest(t, "string", map[string]IndexData[string]{
		"single": {
			in:    []string{"a"},
			value: "a",
			index: 0,
		},
		"second": {
			in:    []string{"a", "b", "a"},
			value: "b",
			index: 1,
		},
		"does not have": {
			in:    []string{"a", "b", "a"},
			value: "c",
			index: -1,
		},
		"empty": {
			in:    []string{},
			value: "a",
			index: -1,
		},
	})

	IndexTest(t, "rune", map[string]IndexData[rune]{
		"single": {
			in:    []rune{'a'},
			value: 'a',
			index: 0,
		},
		"second": {
			in:    []rune{'a', 'b', 'a'},
			value: 'b',
			index: 1,
		},
		"does not have": {
			in:    []rune{'a', 'b', 'a'},
			value: 'c',
			index: -1,
		},
		"empty": {
			in:    []rune{},
			value: 'a',
			index: -1,
		},
	})

	IndexTest(t, "byte", map[string]IndexData[byte]{
		"single": {
			in:    []byte{'a'},
			value: 'a',
			index: 0,
		},
		"second": {
			in:    []byte{'a', 'b', 'a'},
			value: 'b',
			index: 1,
		},
		"does not have": {
			in:    []byte{'a', 'b', 'a'},
			value: 'c',
			index: -1,
		},
		"empty": {
			in:    []byte{},
			value: 'a',
			index: -1,
		},
	})

	IndexTest(t, "int", map[string]IndexData[int]{
		"single": {
			in:    []int{1},
			value: 1,
			index: 0,
		},
		"second": {
			in:    []int{1, 2, 1},
			value: 2,
			index: 1,
		},
		"does not have": {
			in:    []int{1, 2, 1},
			value: 3,
			index: -1,
		},
		"empty": {
			in:    []int{},
			value: 1,
			index: -1,
		},
	})

	IndexTest(t, "float", map[string]IndexData[float32]{
		"single": {
			in:    []float32{1.0},
			value: 1.0,
			index: 0,
		},
		"second": {
			in:    []float32{1.0, 2.0, 1.0},
			value: 2.0,
			index: 1,
		},
		"does not have": {
			in:    []float32{1.0, 2.0, 1.0},
			value: 3.0,
			index: -1,
		},
		"empty": {
			in:    []float32{},
			value: 1.0,
			index: -1,
		},
	})

	IndexTest(t, "complex", map[string]IndexData[complex64]{
		"single": {
			in:    []complex64{1 + 2i},
			value: 1 + 2i,
			index: 0,
		},
		"second": {
			in:    []complex64{1 + 2i, 2 + 3i, 1 + 2i},
			value: 2 + 3i,
			index: 1,
		},
		"does not have": {
			in:    []complex64{1 + 2i, 2 + 3i, 1 + 2i},
			value: 3 + 4i,
			index: -1,
		},
		"empty": {
			in:    []complex64{},
			value: 1 + 2i,
			index: -1,
		},
	})

	IndexTest(t, "chan", map[string]IndexData[chan int]{
		"single": {
			in:    []chan int{c1},
			value: c1,
			index: 0,
		},
		"second": {
			in:    []chan int{c1, c2, c1},
			value: c2,
			index: 1,
		},
		"does not have": {
			in:    []chan int{c1, c2, c1},
			value: c3,
			index: -1,
		},
		"empty": {
			in:    []chan int{},
			value: c1,
			index: -1,
		},
	})

	IndexTest(t, "struct", map[string]IndexData[testStruct]{
		"single": {
			in:    []testStruct{{1, "a"}},
			value: testStruct{1, "a"},
			index: 0,
		},
		"second": {
			in:    []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
			value: testStruct{2, "b"},
			index: 1,
		},
		"does not have": {
			in:    []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
			value: testStruct{3, "c"},
			index: -1,
		},
		"empty": {
			in:    []testStruct{},
			value: testStruct{1, "a"},
			index: -1,
		},
	})

	IndexTest(t, "pointer", map[string]IndexData[*testStruct]{
		"single": {
			in:    []*testStruct{tp1},
			value: tp1,
			index: 0,
		},
		"second": {
			in:    []*testStruct{tp1, tp2, tp1},
			value: tp2,
			index: 1,
		},
		"does not have": {
			in:    []*testStruct{tp1, tp2, tp1},
			value: tp3,
			index: -1,
		},
		"empty": {
			in:    []*testStruct{},
			value: tp1,
			index: -1,
		},
	})

	IndexTest(t, "interface", map[string]IndexData[io.Reader]{
		"single": {
			in:    []io.Reader{r1},
			value: r1,
			index: 0,
		},
		"second": {
			in:    []io.Reader{r1, r2, r1},
			value: r2,
			index: 1,
		},
		"does not have": {
			in:    []io.Reader{r1, r2, r1},
			value: r3,
			index: -1,
		},
		"empty": {
			in:    []io.Reader{},
			value: r1,
			index: -1,
		},
	})
}
