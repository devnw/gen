package gen

import (
	"io"
	"testing"
)

type SliceData[T comparable] struct {
	data     []T
	expected Map[T, struct{}]
}

func SliceTest[T comparable](
	t *testing.T,
	testname string,
	testdata map[string]SliceData[T],
) {
	GenTest[SliceData[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test SliceData[T]) {
			out := Slice[T](test.data)

			for k1, v1 := range out.Map() {
				v2, exists := test.expected[k1]

				if !exists {
					t.Errorf("unexpected key: %v", k1)
				}

				if v1 != v2 {
					t.Errorf("expected %v, got %v", v2, v1)
				}
			}
		})
}

func Test_Slice(t *testing.T) {
	SliceTest(t, "string", map[string]SliceData[string]{
		"single": {
			data: []string{"a"},
			expected: Map[string, struct{}]{
				"a": struct{}{},
			},
		},
		"duplicate": {
			data: []string{"a", "b", "a"},
			expected: Map[string, struct{}]{
				"a": struct{}{},
				"b": struct{}{},
			},
		},
		"empty": {
			data:     []string{},
			expected: Map[string, struct{}]{},
		},
	})

	SliceTest(t, "rune", map[string]SliceData[rune]{
		"single": {
			data: []rune{'a'},
			expected: Map[rune, struct{}]{
				'a': struct{}{},
			},
		},
		"duplicate": {
			data: []rune{'a', 'b', 'a'},
			expected: Map[rune, struct{}]{
				'a': struct{}{},
				'b': struct{}{},
			},
		},
		"empty": {
			data:     []rune{},
			expected: Map[rune, struct{}]{},
		},
	})

	SliceTest(t, "byte", map[string]SliceData[byte]{
		"single": {
			data: []byte{'a'},
			expected: Map[byte, struct{}]{
				'a': struct{}{},
			},
		},
		"duplicate": {
			data: []byte{'a', 'b', 'a'},
			expected: Map[byte, struct{}]{
				'a': struct{}{},
				'b': struct{}{},
			},
		},
		"empty": {
			data:     []byte{},
			expected: Map[byte, struct{}]{},
		},
	})

	SliceTest(t, "int", map[string]SliceData[int]{
		"single": {
			data: []int{1},
			expected: Map[int, struct{}]{
				1: struct{}{},
			},
		},
		"duplicate": {
			data: []int{1, 2, 1},
			expected: Map[int, struct{}]{
				1: struct{}{},
				2: struct{}{},
			},
		},
		"empty": {
			data:     []int{},
			expected: Map[int, struct{}]{},
		},
	})

	SliceTest(t, "float64", map[string]SliceData[float64]{
		"single": {
			data: []float64{1.0},
			expected: Map[float64, struct{}]{
				1.0: struct{}{},
			},
		},
		"duplicate": {
			data: []float64{1.0, 2.0, 1.0},
			expected: Map[float64, struct{}]{
				1.0: struct{}{},
				2.0: struct{}{},
			},
		},
		"empty": {
			data:     []float64{},
			expected: Map[float64, struct{}]{},
		},
	})

	SliceTest(t, "complex128", map[string]SliceData[complex128]{
		"single": {
			data: []complex128{1.0 + 2.0i},
			expected: Map[complex128, struct{}]{
				1.0 + 2.0i: struct{}{},
			},
		},
		"duplicate": {
			data: []complex128{1.0 + 2.0i, 3.0 + 4.0i, 1.0 + 2.0i},
			expected: Map[complex128, struct{}]{
				1.0 + 2.0i: struct{}{},
				3.0 + 4.0i: struct{}{},
			},
		},
		"empty": {
			data:     []complex128{},
			expected: Map[complex128, struct{}]{},
		},
	})

	SliceTest(t, "chan int", map[string]SliceData[chan int]{
		"single": {
			data: []chan int{c1},
			expected: Map[chan int, struct{}]{
				c1: struct{}{},
			},
		},
		"duplicate": {
			data: []chan int{c1, c2, c1},
			expected: Map[chan int, struct{}]{
				c1: struct{}{},
				c2: struct{}{},
			},
		},
		"empty": {
			data:     []chan int{},
			expected: Map[chan int, struct{}]{},
		},
	})

	SliceTest(t, "struct", map[string]SliceData[testStruct]{
		"single": {
			data: []testStruct{
				{
					A: 1,
					B: "a",
				},
			},
			expected: Map[testStruct, struct{}]{
				{
					A: 1,
					B: "a",
				}: struct{}{},
			},
		},
		"duplicate": {
			data: []testStruct{
				{
					A: 1,
					B: "a",
				},
				{
					A: 2,
					B: "b",
				},
				{
					A: 1,
					B: "a",
				},
			},
			expected: Map[testStruct, struct{}]{
				{
					A: 1,
					B: "a",
				}: struct{}{},
				{
					A: 2,
					B: "b",
				}: struct{}{},
			},
		},
		"empty": {
			data:     []testStruct{},
			expected: Map[testStruct, struct{}]{},
		},
	})

	SliceTest(t, "struct pointer", map[string]SliceData[*testStruct]{
		"single": {
			data: []*testStruct{tp1},
			expected: Map[*testStruct, struct{}]{
				tp1: struct{}{},
			},
		},
		"duplicate": {
			data: []*testStruct{tp1, tp2, tp1},
			expected: Map[*testStruct, struct{}]{
				tp1: struct{}{},
				tp2: struct{}{},
			},
		},
		"empty": {
			data:     []*testStruct{},
			expected: Map[*testStruct, struct{}]{},
		},
	})

	SliceTest(t, "interface", map[string]SliceData[io.Reader]{
		"single": {
			data: []io.Reader{r1},
			expected: Map[io.Reader, struct{}]{
				r1: struct{}{},
			},
		},
		"duplicate": {
			data: []io.Reader{r1, r2, r1},
			expected: Map[io.Reader, struct{}]{
				r1: struct{}{},
				r2: struct{}{},
			},
		},
		"empty": {
			data:     []io.Reader{},
			expected: Map[io.Reader, struct{}]{},
		},
	})
}
