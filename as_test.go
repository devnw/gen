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
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

// These tests are going to be much shorter since this is going to be
// compiler enforced for the most part
func Test_As_Reader(t *testing.T) {
	r1 := &bytes.Buffer{}
	r2 := &bufio.Reader{}
	r3 := &bufio.ReadWriter{}
	r4 := strings.NewReader("test")
	r5 := io.Reader(nil)

	out := As[io.Reader](r1, r2, r3, r4, r5)
	for _, v := range out {
		switch v.(type) {
		case nil:
		case *bufio.Reader:
		case *bufio.ReadWriter:
		case *bytes.Buffer:
		case *strings.Reader:
		default:
			t.Fatalf("expected *bufio.Reader, *bytes.Buffer, or *strings.Reader; got %T", v)
		}
	}
}

func Test_As_Reader_Empty(t *testing.T) {
	out := As[io.Reader]()
	if len(out) != 0 {
		t.Fatalf("expected empty slice; got %v", out)
	}
}

func Test_As_Writer(t *testing.T) {
	r1 := &bytes.Buffer{}
	r2 := &bufio.Writer{}
	r3 := &bufio.ReadWriter{}
	r4 := io.Writer(nil)

	out := As[io.Writer](r1, r2, r3, r4)
	for _, v := range out {
		switch v.(type) {
		case nil:
		case *bufio.Writer:
		case *bufio.ReadWriter:
		case *bytes.Buffer:
		default:
			t.Fatalf("expected *bufio.Reader, *bytes.Buffer, or *strings.Reader; got %T", v)
		}
	}
}

func Test_As_Writer_Empty(t *testing.T) {
	out := As[io.Writer]()
	if len(out) != 0 {
		t.Fatalf("expected empty slice; got %v", out)
	}
}
