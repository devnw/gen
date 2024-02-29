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
	"context"
	"testing"
	"time"
)

type SliceData[T comparable] struct {
	data     []T
	expected Map[T, any]
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
			expected: Map[string, any]{
				"a": nil,
			},
		},
		"duplicate": {
			data: []string{"a", "b", "a"},
			expected: Map[string, any]{
				"a": nil,
				"b": nil,
			},
		},
		"empty": {
			data:     []string{},
			expected: Map[string, any]{},
		},
	})

	SliceTest(t, "rune", map[string]SliceData[rune]{
		"single": {
			data: []rune{'a'},
			expected: Map[rune, any]{
				'a': nil,
			},
		},
		"duplicate": {
			data: []rune{'a', 'b', 'a'},
			expected: Map[rune, any]{
				'a': nil,
				'b': nil,
			},
		},
		"empty": {
			data:     []rune{},
			expected: Map[rune, any]{},
		},
	})

	SliceTest(t, "byte", map[string]SliceData[byte]{
		"single": {
			data: []byte{'a'},
			expected: Map[byte, any]{
				'a': nil,
			},
		},
		"duplicate": {
			data: []byte{'a', 'b', 'a'},
			expected: Map[byte, any]{
				'a': nil,
				'b': nil,
			},
		},
		"empty": {
			data:     []byte{},
			expected: Map[byte, any]{},
		},
	})

	SliceTest(t, "int", map[string]SliceData[int]{
		"single": {
			data: []int{1},
			expected: Map[int, any]{
				1: nil,
			},
		},
		"duplicate": {
			data: []int{1, 2, 1},
			expected: Map[int, any]{
				1: nil,
				2: nil,
			},
		},
		"empty": {
			data:     []int{},
			expected: Map[int, any]{},
		},
	})

	SliceTest(t, "float64", map[string]SliceData[float64]{
		"single": {
			data: []float64{1.0},
			expected: Map[float64, any]{
				1.0: nil,
			},
		},
		"duplicate": {
			data: []float64{1.0, 2.0, 1.0},
			expected: Map[float64, any]{
				1.0: nil,
				2.0: nil,
			},
		},
		"empty": {
			data:     []float64{},
			expected: Map[float64, any]{},
		},
	})

	SliceTest(t, "complex128", map[string]SliceData[complex128]{
		"single": {
			data: []complex128{1.0 + 2.0i},
			expected: Map[complex128, any]{
				1.0 + 2.0i: nil,
			},
		},
		"duplicate": {
			data: []complex128{1.0 + 2.0i, 3.0 + 4.0i, 1.0 + 2.0i},
			expected: Map[complex128, any]{
				1.0 + 2.0i: nil,
				3.0 + 4.0i: nil,
			},
		},
		"empty": {
			data:     []complex128{},
			expected: Map[complex128, any]{},
		},
	})

	SliceTest(t, "chan int", map[string]SliceData[chan int]{
		"single": {
			data: []chan int{c1},
			expected: Map[chan int, any]{
				c1: nil,
			},
		},
		"duplicate": {
			data: []chan int{c1, c2, c1},
			expected: Map[chan int, any]{
				c1: nil,
				c2: nil,
			},
		},
		"empty": {
			data:     []chan int{},
			expected: Map[chan int, any]{},
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
			expected: Map[testStruct, any]{
				{
					A: 1,
					B: "a",
				}: nil,
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
			expected: Map[testStruct, any]{
				{
					A: 1,
					B: "a",
				}: nil,
				{
					A: 2,
					B: "b",
				}: nil,
			},
		},
		"empty": {
			data:     []testStruct{},
			expected: Map[testStruct, any]{},
		},
	})

	SliceTest(t, "struct pointer", map[string]SliceData[*testStruct]{
		"single": {
			data: []*testStruct{tp1},
			expected: Map[*testStruct, any]{
				tp1: nil,
			},
		},
		"duplicate": {
			data: []*testStruct{tp1, tp2, tp1},
			expected: Map[*testStruct, any]{
				tp1: nil,
				tp2: nil,
			},
		},
		"empty": {
			data:     []*testStruct{},
			expected: Map[*testStruct, any]{},
		},
	})
}

func ChanTest[U ~[]T, T comparable](
	t *testing.T,
	name string,
	data []U,
) {
	Tst(
		t,
		name,
		data,
		func(t *testing.T, data []T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			out := Slice[T](data).Chan(ctx)

			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					t.Error("context cancelled")
					return
				case out, ok := <-out:
					if !ok {
						if i != len(data) {
							t.Fatalf("closed prematurely, expected %v, got %v", len(data)-1, i)
						}

						return
					}

					if out != data[i] {
						t.Errorf("expected %v, got %v", data[i], out)
					}
				}
			}
		})
}

func Test_Slice_Chan(t *testing.T) {
	ChanTest(t, "int8", IntTests[int8](100, 1000))
	ChanTest(t, "uint8", IntTests[uint8](100, 1000))
	ChanTest(t, "uint8", IntTests[uint8](100, 1000))
	ChanTest(t, "uint16", IntTests[uint16](100, 1000))
	ChanTest(t, "int32", IntTests[int32](100, 1000))
	ChanTest(t, "uint32", IntTests[uint32](100, 1000))
	ChanTest(t, "int64", IntTests[int64](100, 1000))
	ChanTest(t, "uint64", IntTests[uint64](100, 1000))
	ChanTest(t, "float32", FloatTests[float32](100, 1000))
	ChanTest(t, "float64", FloatTests[float64](100, 1000))
}

func Test_Slice_Chan_CtxCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	out := Slice[int]([]int{1, 2, 3}).Chan(ctx)

	select {
	case <-time.After(time.Second):
		t.Error("context cancelled")
	case <-out:
	}
}

func Test_Slice_Chan_NilCtx(t *testing.T) {
	//nolint:staticcheck // this is on purpose
	out := Slice[int]([]int{1, 2, 3}).Chan(nil)

	select {
	case <-time.After(time.Second):
		t.Error("context cancelled")
	case <-out:
	}
}
