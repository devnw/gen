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

func Test_Close(t *testing.T) {
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

func Test_Close_WriteOnly(t *testing.T) {
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
