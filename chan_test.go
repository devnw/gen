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

import "testing"

var chandata = []chan int{c1, c2, c3, c4, c5, c6}

func Test_ReadOnly(t *testing.T) {
	out := ReadOnly(chandata...)

	if len(out) != len(chandata) {
		t.Errorf("expected 6, got %d", len(out))
	}

	for i, v := range out {
		if v != chandata[i] {
			t.Errorf("expected %v, got %v", chandata[i], v)
		}
	}
}

func Test_WriteOnly(t *testing.T) {
	out := WriteOnly(chandata...)

	if len(out) != len(chandata) {
		t.Errorf("expected 6, got %d", len(out))
	}

	for i, v := range out {
		if v != chandata[i] {
			t.Errorf("expected %v, got %v", chandata[i], v)
		}
	}
}

func Test_Close(_ *testing.T) {
	c1, c2, c3, c4, c5, c6 := make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int)

	go Close(c1, c2, c3, c4, c5, c6)

	<-c1
	<-c2
	<-c3
	<-c4
	<-c5
	<-c6
}

func Test_Close_WriteOnly(_ *testing.T) {
	c1, c2, c3, c4, c5, c6 := make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int)

	go Close(WriteOnly(c1, c2, c3, c4, c5, c6)...)

	<-c1
	<-c2
	<-c3
	<-c4
	<-c5
	<-c6
}
