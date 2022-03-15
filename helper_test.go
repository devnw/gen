package gen

import (
	"bytes"
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

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type float interface {
	~float32 | ~float64
}

func Int[T integer]() T {
	value := rand.Int()
	return T(value)
}

func Float[T float]() T {
	return T(rand.Float64())
}

func Ints[T integer](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Int[T]()
	}

	return out
}

func Floats[T float](size int) []T {
	out := make([]T, size)

	for i := range out {
		out[i] = Float[T]()
	}

	return out
}

func IntTests[T integer](tests, cap int) [][]T {
	out := make([][]T, tests)

	for i := range out {
		out[i] = Ints[T](cap)
	}

	return out
}

func FloatTests[T float](tests, cap int) [][]T {
	out := make([][]T, tests)

	for i := range out {
		out[i] = Floats[T](cap)
	}

	return out
}
