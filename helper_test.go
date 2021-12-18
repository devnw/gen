package gen

import (
	"bytes"
	"fmt"
	"testing"
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
