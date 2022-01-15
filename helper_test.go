package gen

import (
	"bytes"
	"constraints"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var c1, c2, c3, c4, c5, c6 = make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int)

var tp1, tp2, tp3, tp4, tp5, tp6 = &testStruct{},
	&testStruct{},
	&testStruct{},
	&testStruct{},
	&testStruct{},
	&testStruct{}

var r1, r2, r3, r4, r5, r6 = &bytes.Buffer{},
	&bytes.Buffer{},
	&bytes.Buffer{},
	&bytes.Buffer{},
	&bytes.Buffer{},
	&bytes.Buffer{}

func GenTest[T any, U any](
	t *testing.T,
	testname string,
	testdata map[string]T,
	f func(t *testing.T, test T),
) {
	for name, test := range testdata {
		t.Run(fmt.Sprintf("%s-%s", testname, name), func(t *testing.T) {
			f(t, test)
		})
	}
}

// TODO: Move the below test harness to a separate generic test package.

// Initialize the random number generator.
func init() { rand.Seed(time.Now().Unix()) }

func Tst[U ~[]T, T any](
	t *testing.T,
	name string,
	data []U,
	f func(t *testing.T, test []T),
) {
	for _, test := range data {
		testname := fmt.Sprintf("%s-tests[%v]-values[%v]", name, len(data), len(test))
		t.Run(testname, func(t *testing.T) {
			f(t, test)
		})
	}
}

func Int[T constraints.Integer]() T {
	value := rand.Int()
	return T(value)
}

func Float[T constraints.Float]() T {
	return T(rand.Float64())
}

func Ints[T constraints.Integer](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Int[T]()
	}

	return out
}

func Floats[T constraints.Float](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Float[T]()
	}

	return out
}

func IntTests[T constraints.Integer](tests, cap int) [][]T {
	out := make([][]T, tests)

	for i := range out {
		out[i] = Ints[T](cap)
	}

	return out
}

func FloatTests[T constraints.Float](tests, cap int) [][]T {
	out := make([][]T, tests)

	for i := range out {
		out[i] = Floats[T](cap)
	}

	return out
}
