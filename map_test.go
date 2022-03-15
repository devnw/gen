package gen

import (
	"testing"
)

type MapData[T comparable] struct {
	in     Map[T, T]
	keys   []T
	values []T
}

func MapTest[T comparable](t *testing.T, testname string, testdata map[string]MapData[T]) {
	GenTest[MapData[T], T](
		t,
		testname,
		testdata,
		func(t *testing.T, test MapData[T]) {
			keys := test.in.Keys()
			if len(keys) != len(test.keys) {
				t.Fatalf("expected %d keys, got %d", len(test.keys), len(keys))
			}

			if !Match(keys, test.keys) {
				t.Fatalf("expected %v, got %v", test.keys, keys)
			}

			values := test.in.Values()
			if len(values) != len(test.values) {
				t.Fatalf("expected %d values, got %d", len(test.values), len(values))
			}

			if !Match(values, test.values) {
				t.Fatalf("expected %v, got %v", test.values, values)
			}
		})
}

func Test_nil(t *testing.T) {
	var m Map[string, bool]
	ks := m.Keys()
	if len(ks) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(ks))
	}

	vs := m.Values()
	if len(vs) != 0 {
		t.Errorf("Expected 0 values, got %d", len(vs))
	}
}

func Test_Map(t *testing.T) {
	MapTest(t, "string", map[string]MapData[string]{
		"one": {
			in: Map[string, string]{"a": "b"},
			keys: []string{
				"a",
			},
			values: []string{
				"b",
			},
		},
		"two": {
			in: Map[string, string]{"a": "b", "c": "d"},
			keys: []string{
				"a",
				"c",
			},
			values: []string{
				"b",
				"d",
			},
		},
		"three": {
			in: Map[string, string]{"a": "b", "c": "d", "e": "f"},
			keys: []string{
				"a",
				"c",
				"e",
			},
			values: []string{
				"b",
				"d",
				"f",
			},
		},
	})

	MapTest(t, "rune", map[string]MapData[rune]{
		"one": {
			in: Map[rune, rune]{'a': 'b'},
			keys: []rune{
				'a',
			},
			values: []rune{
				'b',
			},
		},
		"two": {
			in: Map[rune, rune]{'a': 'b', 'c': 'd'},
			keys: []rune{
				'a',
				'c',
			},
			values: []rune{
				'b',
				'd',
			},
		},
		"three": {
			in: Map[rune, rune]{'a': 'b', 'c': 'd', 'e': 'f'},
			keys: []rune{
				'a',
				'c',
				'e',
			},
			values: []rune{
				'b',
				'd',
				'f',
			},
		},
	})

	MapTest(t, "byte", map[string]MapData[byte]{
		"one": {
			in: Map[byte, byte]{'a': 'b'},
			keys: []byte{
				'a',
			},
			values: []byte{
				'b',
			},
		},
		"two": {
			in: Map[byte, byte]{'a': 'b', 'c': 'd'},
			keys: []byte{
				'a',
				'c',
			},
			values: []byte{
				'b',
				'd',
			},
		},
		"three": {
			in: Map[byte, byte]{'a': 'b', 'c': 'd', 'e': 'f'},
			keys: []byte{
				'a',
				'c',
				'e',
			},
			values: []byte{
				'b',
				'd',
				'f',
			},
		},
	})

	MapTest(t, "int", map[string]MapData[int]{
		"one": {
			in: Map[int, int]{1: 2},
			keys: []int{
				1,
			},
			values: []int{
				2,
			},
		},
		"two": {
			in: Map[int, int]{1: 2, 3: 4},
			keys: []int{
				1,
				3,
			},
			values: []int{
				2,
				4,
			},
		},
		"three": {
			in: Map[int, int]{1: 2, 3: 4, 5: 6},
			keys: []int{
				1,
				3,
				5,
			},
			values: []int{
				2,
				4,
				6,
			},
		},
	})

	MapTest(t, "float", map[string]MapData[float32]{
		"one": {
			in: Map[float32, float32]{1.0: 2.0},
			keys: []float32{
				1.0,
			},
			values: []float32{
				2.0,
			},
		},
		"two": {
			in: Map[float32, float32]{1.0: 2.0, 3.0: 4.0},
			keys: []float32{
				1.0,
				3.0,
			},
			values: []float32{
				2.0,
				4.0,
			},
		},
		"three": {
			in: Map[float32, float32]{1.0: 2.0, 3.0: 4.0, 5.0: 6.0},
			keys: []float32{
				1.0,
				3.0,
				5.0,
			},
			values: []float32{
				2.0,
				4.0,
				6.0,
			},
		},
	})

	MapTest(t, "complex", map[string]MapData[complex64]{
		"one": {
			in: Map[complex64, complex64]{1 + 2i: 3 + 4i},
			keys: []complex64{
				1 + 2i,
			},
			values: []complex64{
				3 + 4i,
			},
		},
		"two": {
			in: Map[complex64, complex64]{1 + 2i: 3 + 4i, 5 + 6i: 7 + 8i},
			keys: []complex64{
				1 + 2i,
				5 + 6i,
			},
			values: []complex64{
				3 + 4i,
				7 + 8i,
			},
		},
		"three": {
			in: Map[complex64, complex64]{1 + 2i: 3 + 4i, 5 + 6i: 7 + 8i, 9 + 10i: 11 + 12i},
			keys: []complex64{
				1 + 2i,
				5 + 6i,
				9 + 10i,
			},
			values: []complex64{
				3 + 4i,
				7 + 8i,
				11 + 12i,
			},
		},
	})

	MapTest(t, "struct", map[string]MapData[testStruct]{
		"one": {
			in: Map[testStruct, testStruct]{
				{1, "a"}: {2, "b"},
			},
			keys: []testStruct{
				{1, "a"},
			},
			values: []testStruct{
				{2, "b"},
			},
		},
		"two": {
			in: Map[testStruct, testStruct]{
				{1, "a"}: {2, "b"},
				{3, "c"}: {4, "d"},
			},
			keys: []testStruct{
				{1, "a"},
				{3, "c"},
			},
			values: []testStruct{
				{2, "b"},
				{4, "d"},
			},
		},
		"three": {
			in: Map[testStruct, testStruct]{
				{1, "a"}: {2, "b"},
				{3, "c"}: {4, "d"},
				{5, "e"}: {6, "f"},
			},
			keys: []testStruct{
				{1, "a"},
				{3, "c"},
				{5, "e"},
			},
			values: []testStruct{
				{2, "b"},
				{4, "d"},
				{6, "f"},
			},
		},
	})

	MapTest(t, "struct pointer", map[string]MapData[*testStruct]{
		"one": {
			in: Map[*testStruct, *testStruct]{
				tp1: tp2,
			},
			keys: []*testStruct{
				tp1,
			},
			values: []*testStruct{
				tp2,
			},
		},
		"two": {
			in: Map[*testStruct, *testStruct]{
				tp1: tp2,
				tp3: tp4,
			},
			keys: []*testStruct{
				tp1,
				tp3,
			},
			values: []*testStruct{
				tp2,
				tp4,
			},
		},
		"three": {
			in: Map[*testStruct, *testStruct]{
				tp1: tp2,
				tp3: tp4,
				tp5: tp6,
			},
			keys: []*testStruct{
				tp1,
				tp3,
				tp5,
			},
			values: []*testStruct{
				tp2,
				tp4,
				tp6,
			},
		},
	})

	MapTest(t, "chan", map[string]MapData[chan int]{
		"one": {
			in: Map[chan int, chan int]{
				c1: c2,
			},
			keys: []chan int{
				c1,
			},
			values: []chan int{
				c2,
			},
		},
		"two": {
			in: Map[chan int, chan int]{
				c1: c2,
				c3: c4,
			},
			keys: []chan int{
				c1,
				c3,
			},
			values: []chan int{
				c2,
				c4,
			},
		},
		"three": {
			in: Map[chan int, chan int]{
				c1: c2,
				c3: c4,
				c5: c6,
			},
			keys: []chan int{
				c1,
				c3,
				c5,
			},
			values: []chan int{
				c2,
				c4,
				c6,
			},
		},
	})
}
