package gen

import (
	"io"
	"testing"
)

type DSExpected[T comparable] struct {
	a, b     []T
	expected []T
}

func IntersectTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]DSExpected[T],
) {
	GenTest[DSExpected[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test DSExpected[T]) {
			out := Intersect(test.a, test.b)

			err := Compare(out, test.expected)
			if err != nil {
				t.Errorf("Expected %v, got %v", test.expected, out)
			}
		})
}

func Test_Intersect(t *testing.T) {
	IntersectTest(t, "string", map[string]DSExpected[string]{
		"single": {
			a:        []string{"a"},
			b:        []string{"a"},
			expected: []string{"a"},
		},
		"offset-index": {
			a:        []string{"a", "b"},
			b:        []string{"b", "a"},
			expected: []string{"a", "b"},
		},
		"duplicate": {
			a:        []string{"a", "b", "a"},
			b:        []string{"a", "b", "a"},
			expected: []string{"a", "b", "a"},
		},
		"multiple-no-duplicate": {
			a:        []string{"a", "b", "c"},
			b:        []string{"d", "e", "f"},
			expected: []string{},
		},
		"length-mismatch": {
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b"},
			expected: []string{"a", "b"},
		},
		"index-mismatch": {
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "d"},
			expected: []string{"a", "b"},
		},
		"empty": {
			a:        []string{},
			b:        []string{},
			expected: []string{},
		},
	})

	IntersectTest(t, "rune", map[string]DSExpected[rune]{
		"single": {
			a:        []rune{'a'},
			b:        []rune{'a'},
			expected: []rune{'a'},
		},
		"offset-index": {
			a:        []rune{'a', 'b'},
			b:        []rune{'b', 'a'},
			expected: []rune{'a', 'b'},
		},
		"duplicate": {
			a:        []rune{'a', 'b', 'a'},
			b:        []rune{'a', 'b', 'a'},
			expected: []rune{'a', 'b', 'a'},
		},
		"multiple-no-duplicate": {
			a:        []rune{'a', 'b', 'c'},
			b:        []rune{'d', 'e', 'f'},
			expected: []rune{},
		},
		"length-mismatch": {
			a:        []rune{'a', 'b', 'c'},
			b:        []rune{'a', 'b'},
			expected: []rune{'a', 'b'},
		},
		"index-mismatch": {
			a:        []rune{'a', 'b', 'c'},
			b:        []rune{'a', 'b', 'd'},
			expected: []rune{'a', 'b'},
		},
		"empty": {
			a:        []rune{},
			b:        []rune{},
			expected: []rune{},
		},
	})

	IntersectTest(t, "byte", map[string]DSExpected[byte]{
		"single": {
			a:        []byte{'a'},
			b:        []byte{'a'},
			expected: []byte{'a'},
		},
		"offset-index": {
			a:        []byte{'a', 'b'},
			b:        []byte{'b', 'a'},
			expected: []byte{'a', 'b'},
		},
		"duplicate": {
			a:        []byte{'a', 'b', 'a'},
			b:        []byte{'a', 'b', 'a'},
			expected: []byte{'a', 'b', 'a'},
		},
		"multiple-no-duplicate": {
			a:        []byte{'a', 'b', 'c'},
			b:        []byte{'d', 'e', 'f'},
			expected: []byte{},
		},
		"length-mismatch": {
			a:        []byte{'a', 'b', 'c'},
			b:        []byte{'a', 'b'},
			expected: []byte{'a', 'b'},
		},
		"index-mismatch": {
			a:        []byte{'a', 'b', 'c'},
			b:        []byte{'a', 'b', 'd'},
			expected: []byte{'a', 'b'},
		},
		"empty": {
			a:        []byte{},
			b:        []byte{},
			expected: []byte{},
		},
	})

	IntersectTest(t, "int", map[string]DSExpected[int]{
		"single": {
			a:        []int{1},
			b:        []int{1},
			expected: []int{1},
		},
		"offset-index": {
			a:        []int{1, 2},
			b:        []int{2, 1},
			expected: []int{1, 2},
		},
		"duplicate": {
			a:        []int{1, 2, 1},
			b:        []int{1, 2, 1},
			expected: []int{1, 2, 1},
		},
		"multiple-no-duplicate": {
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: []int{},
		},
		"length-mismatch": {
			a:        []int{1, 2, 3},
			b:        []int{1, 2},
			expected: []int{1, 2},
		},
		"index-mismatch": {
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 4},
			expected: []int{1, 2},
		},
		"empty": {
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
	})

	IntersectTest(t, "float32", map[string]DSExpected[float32]{
		"single": {
			a:        []float32{1.0},
			b:        []float32{1.0},
			expected: []float32{1.0},
		},
		"offset-index": {
			a:        []float32{1.0, 2.0},
			b:        []float32{2.0, 1.0},
			expected: []float32{1.0, 2.0},
		},
		"duplicate": {
			a:        []float32{1.0, 2.0, 1.0},
			b:        []float32{1.0, 2.0, 1.0},
			expected: []float32{1.0, 2.0, 1.0},
		},
		"multiple-no-duplicate": {
			a:        []float32{1.0, 2.0, 3.0},
			b:        []float32{4.0, 5.0, 6.0},
			expected: []float32{},
		},
		"length-mismatch": {
			a:        []float32{1.0, 2.0, 3.0},
			b:        []float32{1.0, 2.0},
			expected: []float32{1.0, 2.0},
		},
		"index-mismatch": {
			a:        []float32{1.0, 2.0, 3.0},
			b:        []float32{1.0, 2.0, 4.0},
			expected: []float32{1.0, 2.0},
		},
		"empty": {
			a:        []float32{},
			b:        []float32{},
			expected: []float32{},
		},
	})

	IntersectTest(t, "complex64", map[string]DSExpected[complex64]{
		"single": {
			a:        []complex64{1.0 + 2.0i},
			b:        []complex64{1.0 + 2.0i},
			expected: []complex64{1.0 + 2.0i},
		},
		"offset-index": {
			a:        []complex64{1.0 + 2.0i, 2.0 + 3.0i},
			b:        []complex64{2.0 + 3.0i, 1.0 + 2.0i},
			expected: []complex64{1.0 + 2.0i, 2.0 + 3.0i},
		},
		"duplicate": {
			a:        []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			b:        []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			expected: []complex64{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
		},
		"multiple-no-duplicate": {
			a:        []complex64{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
			b:        []complex64{4.0 + 5.0i, 5.0 + 6.0i, 6.0 + 7.0i},
			expected: []complex64{},
		},
		"length-mismatch": {
			a:        []complex64{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
			b:        []complex64{1.0 + 2.0i, 2.0 + 3.0i},
			expected: []complex64{1.0 + 2.0i, 2.0 + 3.0i},
		},
		"index-mismatch": {
			a:        []complex64{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
			b:        []complex64{1.0 + 2.0i, 2.0 + 3.0i, 4.0 + 5.0i},
			expected: []complex64{1.0 + 2.0i, 2.0 + 3.0i},
		},
		"empty": {
			a:        []complex64{},
			b:        []complex64{},
			expected: []complex64{},
		},
	})

	IntersectTest(t, "channel", map[string]DSExpected[chan int]{
		"single": {
			a:        []chan int{c1},
			b:        []chan int{c1},
			expected: []chan int{c1},
		},
		"offset-index": {
			a:        []chan int{c1, c2},
			b:        []chan int{c2, c1},
			expected: []chan int{c1, c2},
		},
		"duplicate": {
			a:        []chan int{c1, c2, c1},
			b:        []chan int{c1, c2, c1},
			expected: []chan int{c1, c2, c1},
		},
		"multiple-no-duplicate": {
			a:        []chan int{c1, c2, c3},
			b:        []chan int{c4, c5, c6},
			expected: []chan int{},
		},
		"length-mismatch": {
			a:        []chan int{c1, c2, c3},
			b:        []chan int{c1, c2},
			expected: []chan int{c1, c2},
		},
		"index-mismatch": {
			a:        []chan int{c1, c2, c3},
			b:        []chan int{c1, c2, c4},
			expected: []chan int{c1, c2},
		},
		"empty": {
			a:        []chan int{},
			b:        []chan int{},
			expected: []chan int{},
		},
	})

	IntersectTest(t, "struct", map[string]DSExpected[testStruct]{
		"single": {
			a:        []testStruct{{1, "a"}},
			b:        []testStruct{{1, "a"}},
			expected: []testStruct{{1, "a"}},
		},
		"offset-index": {
			a:        []testStruct{{1, "a"}, {2, "b"}},
			b:        []testStruct{{2, "b"}, {1, "a"}},
			expected: []testStruct{{1, "a"}, {2, "b"}},
		},
		"duplicate": {
			a:        []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
			b:        []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
			expected: []testStruct{{1, "a"}, {2, "b"}, {1, "a"}},
		},
		"multiple-no-duplicate": {
			a:        []testStruct{{1, "a"}, {2, "b"}, {3, "c"}},
			b:        []testStruct{{4, "d"}, {5, "e"}, {6, "f"}},
			expected: []testStruct{},
		},
		"length-mismatch": {
			a:        []testStruct{{1, "a"}, {2, "b"}, {3, "c"}},
			b:        []testStruct{{1, "a"}, {2, "b"}},
			expected: []testStruct{{1, "a"}, {2, "b"}},
		},
		"index-mismatch": {
			a:        []testStruct{{1, "a"}, {2, "b"}, {3, "c"}},
			b:        []testStruct{{1, "a"}, {2, "b"}, {4, "d"}},
			expected: []testStruct{{1, "a"}, {2, "b"}},
		},
		"empty": {
			a:        []testStruct{},
			b:        []testStruct{},
			expected: []testStruct{},
		},
	})

	IntersectTest(t, "pointer", map[string]DSExpected[*testStruct]{
		"single": {
			a:        []*testStruct{tp1},
			b:        []*testStruct{tp1},
			expected: []*testStruct{tp1},
		},
		"offset-index": {
			a:        []*testStruct{tp1, tp2},
			b:        []*testStruct{tp2, tp1},
			expected: []*testStruct{tp1, tp2},
		},
		"duplicate": {
			a:        []*testStruct{tp1, tp2, tp1},
			b:        []*testStruct{tp1, tp2, tp1},
			expected: []*testStruct{tp1, tp2, tp1},
		},
		"multiple-no-duplicate": {
			a:        []*testStruct{tp1, tp2, tp3},
			b:        []*testStruct{tp4, tp5, tp6},
			expected: []*testStruct{},
		},
		"length-mismatch": {
			a:        []*testStruct{tp1, tp2, tp3},
			b:        []*testStruct{tp1, tp2},
			expected: []*testStruct{tp1, tp2},
		},
		"index-mismatch": {
			a:        []*testStruct{tp1, tp2, tp3},
			b:        []*testStruct{tp1, tp2, tp4},
			expected: []*testStruct{tp1, tp2},
		},
		"empty": {
			a:        []*testStruct{},
			b:        []*testStruct{},
			expected: []*testStruct{},
		},
	})

	IntersectTest(t, "interface", map[string]DSExpected[io.Reader]{
		"single": {
			a:        []io.Reader{r1},
			b:        []io.Reader{r1},
			expected: []io.Reader{r1},
		},
		"offset-index": {
			a:        []io.Reader{r1, r2},
			b:        []io.Reader{r2, r1},
			expected: []io.Reader{r1, r2},
		},
		"duplicate": {
			a:        []io.Reader{r1, r2, r1},
			b:        []io.Reader{r1, r2, r1},
			expected: []io.Reader{r1, r2, r1},
		},
		"multiple-no-duplicate": {
			a:        []io.Reader{r1, r2, r3},
			b:        []io.Reader{r4, r5, r6},
			expected: []io.Reader{},
		},
		"length-mismatch": {
			a:        []io.Reader{r1, r2, r3},
			b:        []io.Reader{r1, r2},
			expected: []io.Reader{r1, r2},
		},
		"index-mismatch": {
			a:        []io.Reader{r1, r2, r3},
			b:        []io.Reader{r1, r2, r4},
			expected: []io.Reader{r1, r2},
		},
		"empty": {
			a:        []io.Reader{},
			b:        []io.Reader{},
			expected: []io.Reader{},
		},
	})
}
