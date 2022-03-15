package gen

import (
	"testing"
)

func ExcludeTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]DSExpected[T],
) {
	GenTest[DSExpected[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test DSExpected[T]) {
			out := Exclude(test.a, test.b...)

			err := Compare(out, test.expected)
			if err != nil {
				t.Errorf("Expected %v, got %v", test.expected, out)
			}
		})
}

func Test_Exclude(t *testing.T) {
	ExcludeTest(t, "string", map[string]DSExpected[string]{
		"single": {
			a:        []string{"a"},
			b:        []string{"a"},
			expected: []string{},
		},
		"duplicate": {
			a:        []string{"a", "a"},
			b:        []string{"a"},
			expected: []string{},
		},
		"multiple": {
			a:        []string{"a", "b", "a"},
			b:        []string{"a", "a", "a"},
			expected: []string{"b"},
		},
		"multiple-no-duplicate": {
			a:        []string{"a", "b", "c"},
			b:        []string{"d", "e", "f"},
			expected: []string{"a", "b", "c"},
		},
		"length-mismatch": {
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b"},
			expected: []string{"c"},
		},
		"empty": {
			a:        []string{},
			b:        []string{},
			expected: []string{},
		},
	})

	ExcludeTest(t, "rune", map[string]DSExpected[rune]{
		"single": {
			a:        []rune{'a'},
			b:        []rune{'a'},
			expected: []rune{},
		},
		"duplicate": {
			a:        []rune{'a', 'a'},
			b:        []rune{'a'},
			expected: []rune{},
		},
		"multiple": {
			a:        []rune{'a', 'b', 'a'},
			b:        []rune{'a', 'a', 'a'},
			expected: []rune{'b'},
		},
		"multiple-no-duplicate": {
			a:        []rune{'a', 'b', 'c'},
			b:        []rune{'d', 'e', 'f'},
			expected: []rune{'a', 'b', 'c'},
		},
		"length-mismatch": {
			a:        []rune{'a', 'b', 'c'},
			b:        []rune{'a', 'b'},
			expected: []rune{'c'},
		},
		"empty": {
			a:        []rune{},
			b:        []rune{},
			expected: []rune{},
		},
	})

	ExcludeTest(t, "byte", map[string]DSExpected[byte]{
		"single": {
			a:        []byte{'a'},
			b:        []byte{'a'},
			expected: []byte{},
		},
		"duplicate": {
			a:        []byte{'a', 'a'},
			b:        []byte{'a'},
			expected: []byte{},
		},
		"multiple": {
			a:        []byte{'a', 'b', 'a'},
			b:        []byte{'a', 'a', 'a'},
			expected: []byte{'b'},
		},
		"multiple-no-duplicate": {
			a:        []byte{'a', 'b', 'c'},
			b:        []byte{'d', 'e', 'f'},
			expected: []byte{'a', 'b', 'c'},
		},
		"length-mismatch": {
			a:        []byte{'a', 'b', 'c'},
			b:        []byte{'a', 'b'},
			expected: []byte{'c'},
		},
		"empty": {
			a:        []byte{},
			b:        []byte{},
			expected: []byte{},
		},
	})

	ExcludeTest(t, "int", map[string]DSExpected[int]{
		"single": {
			a:        []int{1},
			b:        []int{1},
			expected: []int{},
		},
		"duplicate": {
			a:        []int{1, 1},
			b:        []int{1},
			expected: []int{},
		},
		"multiple": {
			a:        []int{1, 2, 1},
			b:        []int{1, 1, 1},
			expected: []int{2},
		},
		"multiple-no-duplicate": {
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
		"length-mismatch": {
			a:        []int{1, 2, 3},
			b:        []int{1, 2},
			expected: []int{3},
		},
		"empty": {
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
	})

	ExcludeTest(t, "float64", map[string]DSExpected[float64]{
		"single": {
			a:        []float64{1.0},
			b:        []float64{1.0},
			expected: []float64{},
		},
		"duplicate": {
			a:        []float64{1.0, 1.0},
			b:        []float64{1.0},
			expected: []float64{},
		},
		"multiple": {
			a:        []float64{1.0, 2.0, 1.0},
			b:        []float64{1.0, 1.0, 1.0},
			expected: []float64{2.0},
		},
		"multiple-no-duplicate": {
			a:        []float64{1.0, 2.0, 3.0},
			b:        []float64{4.0, 5.0, 6.0},
			expected: []float64{1.0, 2.0, 3.0},
		},
		"length-mismatch": {
			a:        []float64{1.0, 2.0, 3.0},
			b:        []float64{1.0, 2.0},
			expected: []float64{3.0},
		},
		"empty": {
			a:        []float64{},
			b:        []float64{},
			expected: []float64{},
		},
	})

	ExcludeTest(t, "complex", map[string]DSExpected[complex128]{
		"single": {
			a:        []complex128{1.0 + 2.0i},
			b:        []complex128{1.0 + 2.0i},
			expected: []complex128{},
		},
		"duplicate": {
			a:        []complex128{1.0 + 2.0i, 1.0 + 2.0i},
			b:        []complex128{1.0 + 2.0i},
			expected: []complex128{},
		},
		"multiple": {
			a:        []complex128{1.0 + 2.0i, 2.0 + 3.0i, 1.0 + 2.0i},
			b:        []complex128{1.0 + 2.0i, 1.0 + 2.0i, 1.0 + 2.0i},
			expected: []complex128{2.0 + 3.0i},
		},
		"multiple-no-duplicate": {
			a:        []complex128{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
			b:        []complex128{4.0 + 5.0i, 5.0 + 6.0i, 6.0 + 7.0i},
			expected: []complex128{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
		},
		"length-mismatch": {
			a:        []complex128{1.0 + 2.0i, 2.0 + 3.0i, 3.0 + 4.0i},
			b:        []complex128{1.0 + 2.0i, 2.0 + 3.0i},
			expected: []complex128{3.0 + 4.0i},
		},
		"empty": {
			a:        []complex128{},
			b:        []complex128{},
			expected: []complex128{},
		},
	})

	ExcludeTest(t, "chan", map[string]DSExpected[chan int]{
		"single": {
			a:        []chan int{c1},
			b:        []chan int{c1},
			expected: []chan int{},
		},
		"duplicate": {
			a:        []chan int{c1, c1},
			b:        []chan int{c1},
			expected: []chan int{},
		},
		"multiple": {
			a:        []chan int{c1, c2, c1},
			b:        []chan int{c1, c1, c1},
			expected: []chan int{c2},
		},
		"multiple-no-duplicate": {
			a:        []chan int{c1, c2, c3},
			b:        []chan int{c4, c5, c6},
			expected: []chan int{c1, c2, c3},
		},
		"length-mismatch": {
			a:        []chan int{c1, c2, c3},
			b:        []chan int{c1, c2},
			expected: []chan int{c3},
		},
		"empty": {
			a:        []chan int{},
			b:        []chan int{},
			expected: []chan int{},
		},
	})

	ExcludeTest(t, "struct", map[string]DSExpected[testStruct]{
		"single": {
			a:        []testStruct{{A: 1, B: "a"}},
			b:        []testStruct{{A: 1, B: "a"}},
			expected: []testStruct{},
		},
		"duplicate": {
			a:        []testStruct{{A: 1, B: "a"}, {A: 1, B: "a"}},
			b:        []testStruct{{A: 1, B: "a"}},
			expected: []testStruct{},
		},
		"multiple": {
			a:        []testStruct{{A: 1, B: "a"}, {A: 2, B: "b"}, {A: 1, B: "a"}},
			b:        []testStruct{{A: 1, B: "a"}, {A: 1, B: "a"}, {A: 1, B: "a"}},
			expected: []testStruct{{A: 2, B: "b"}},
		},
		"multiple-no-duplicate": {
			a:        []testStruct{{A: 1, B: "a"}, {A: 2, B: "b"}, {A: 3, B: "c"}},
			b:        []testStruct{{A: 4, B: "d"}, {A: 5, B: "e"}, {A: 6, B: "f"}},
			expected: []testStruct{{A: 1, B: "a"}, {A: 2, B: "b"}, {A: 3, B: "c"}},
		},
		"length-mismatch": {
			a:        []testStruct{{A: 1, B: "a"}, {A: 2, B: "b"}, {A: 3, B: "c"}},
			b:        []testStruct{{A: 1, B: "a"}, {A: 2, B: "b"}},
			expected: []testStruct{{A: 3, B: "c"}},
		},
		"empty": {
			a:        []testStruct{},
			b:        []testStruct{},
			expected: []testStruct{},
		},
	})

	ExcludeTest(t, "pointer", map[string]DSExpected[*testStruct]{
		"single": {
			a:        []*testStruct{tp1},
			b:        []*testStruct{tp1},
			expected: []*testStruct{},
		},
		"duplicate": {
			a:        []*testStruct{tp1, tp1},
			b:        []*testStruct{tp1},
			expected: []*testStruct{},
		},
		"multiple": {
			a:        []*testStruct{tp1, tp2, tp1},
			b:        []*testStruct{tp1, tp1, tp1},
			expected: []*testStruct{tp2},
		},
		"multiple-no-duplicate": {
			a:        []*testStruct{tp1, tp2, tp3},
			b:        []*testStruct{tp4, tp5, tp6},
			expected: []*testStruct{tp1, tp2, tp3},
		},
		"length-mismatch": {
			a:        []*testStruct{tp1, tp2, tp3},
			b:        []*testStruct{tp1, tp2},
			expected: []*testStruct{tp3},
		},
		"empty": {
			a:        []*testStruct{},
			b:        []*testStruct{},
			expected: []*testStruct{},
		},
	})
}
