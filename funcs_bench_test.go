package gen

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func Benchmark_Unique(b *testing.B) {

	data := []int{
		1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Unique(data)
		if len(out) != 10 {
			b.Errorf("expected %d, got %d", len(data), len(out))
		}
	}
}

func Benchmark_Unique_NoUnique(b *testing.B) {

	data := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Unique(data)
		if len(out) != 10 {
			b.Errorf("expected %d, got %d", len(data), len(out))
		}
	}
}

func Benchmark_Has(b *testing.B) {

	data := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for n := 0; n < b.N; n++ {
		if Has(data, 10) != true {
			b.Errorf("expected true, got false")
		}
	}
}

func Benchmark_Has_DoesntHave(b *testing.B) {

	data := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for n := 0; n < b.N; n++ {
		if Has(data, 11) != false {
			b.Errorf("expected false, got true")
		}
	}
}

func Benchmark_Match(b *testing.B) {

	dataa := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	datab := []int{
		10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	for n := 0; n < b.N; n++ {
		if Match(dataa, datab) != true {
			b.Errorf("expected true, got false")
		}
	}
}

func Benchmark_Match_Doesnt(b *testing.B) {

	dataa := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	datab := []int{
		10, 9, 8, 7, 6, 5, 4, 3, 2, 99,
	}

	for n := 0; n < b.N; n++ {
		if Match(dataa, datab) != false {
			b.Errorf("expected false, got true")
		}
	}
}

func Benchmark_Indices(b *testing.B) {

	data := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Indices(data, 2)
		if len(out) != 4 {
			b.Errorf("expected %d, got %d", 4, len(out))
		}
	}
}

func Benchmark_Indices_None(b *testing.B) {

	data := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Indices(data, 11)
		if len(out) != 0 {
			b.Errorf("expected %d, got %d", 0, len(out))
		}
	}
}

func Benchmark_Index(b *testing.B) {

	data := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Index(data, 7)
		if out != 6 {
			b.Errorf("expected %d, got %d", 6, out)
		}
	}
}

func Benchmark_Index_None(b *testing.B) {

	data := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Index(data, 11)
		if out != -1 {
			b.Errorf("expected %d, got %d", 1, out)
		}
	}
}

func Benchmark_comp(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		if comp(dataa, datab) != -1 {
			b.Errorf("expected match, got mismatch")
		}
	}
}

func Benchmark_comp_nomatch(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 8,
	}

	for n := 0; n < b.N; n++ {
		if comp(dataa, datab) != 9 {
			b.Errorf("expected mismatch, got match")
		}
	}
}

func Benchmark_Equal(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		if Equal(dataa, datab) != true {
			b.Errorf("expected true, got false")
		}
	}
}

func Benchmark_Equal_not(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 11,
	}

	for n := 0; n < b.N; n++ {
		if Equal(dataa, datab) != false {
			b.Errorf("expected false, got true")
		}
	}
}

func Benchmark_Compare(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		err := Compare(dataa, datab)
		if err != nil {
			b.Errorf("expected nil, got %v", err)
		}
	}
}

func Benchmark_Compare_err(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 11,
	}

	for n := 0; n < b.N; n++ {
		err := Compare(dataa, datab)
		if err == nil {
			b.Errorf("expected error, got nil")
		}
	}
}

func Benchmark_Intersect_full(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Intersect(dataa, datab)
		if len(out) != 10 {
			b.Errorf("expected %d, got %d", 10, len(out))
		}
	}
}

func Benchmark_Intersect_partial(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Intersect(dataa, datab)
		if len(out) != 6 {
			b.Errorf("expected %d, got %d", 6, len(out))
		}
	}
}

func Benchmark_Intersect_none(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	}

	for n := 0; n < b.N; n++ {
		out := Intersect(dataa, datab)
		if len(out) != 0 {
			b.Errorf("expected %d, got %d", 0, len(out))
		}
	}
}

func Benchmark_Diff_all(b *testing.B) {

	dataa := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	datab := []int{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}

	for n := 0; n < b.N; n++ {
		out := Diff(dataa, datab)
		if len(out) != 20 {
			b.Errorf("expected %d, got %d", 20, len(out))
		}
	}
}

func Benchmark_Diff_partial(b *testing.B) {

	dataa := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	datab := []int{
		1, 2, 3, 4, 5, 16, 17, 18, 19, 20,
	}

	for n := 0; n < b.N; n++ {
		out := Diff(dataa, datab)
		if len(out) != 10 {
			b.Errorf("expected %d, got %d", 10, len(out))
		}
	}
}

func Benchmark_Diff_none(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	datab := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	for n := 0; n < b.N; n++ {
		out := Diff(dataa, datab)
		if len(out) != 0 {
			b.Errorf("expected %d, got %d", 0, len(out))
		}
	}
}

func Benchmark_Exclude(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	value := 2

	for n := 0; n < b.N; n++ {
		out := Exclude(dataa, value)
		if len(out) != 6 {
			b.Errorf("expected %d, got %d", 6, len(out))
		}
	}
}

func Benchmark_Exclude_none(b *testing.B) {

	dataa := []int{
		1, 2, 2, 4, 5, 2, 7, 8, 2, 10,
	}

	value := 12

	for n := 0; n < b.N; n++ {
		out := Exclude(dataa, value)
		if len(out) != 10 {
			b.Errorf("expected %d, got %d", 6, len(out))
		}
	}
}

func Benchmark_As(b *testing.B) {
	r1 := &bytes.Buffer{}
	r2 := &bufio.Reader{}
	r3 := &bufio.ReadWriter{}
	r4 := strings.NewReader("test")
	r5 := io.Reader(nil)

	for n := 0; n < b.N; n++ {
		out := As[io.Reader](r1, r2, r3, r4, r5)
		if len(out) != 5 {
			b.Errorf("expected %d, got %d", 5, len(out))
		}
	}
}
