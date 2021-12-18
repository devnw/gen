package gen

import "fmt"

type ErrIndexMismatch[T any] struct {
	I        int
	Expected T
	Actual   T
}

func (e ErrIndexMismatch[T]) Error() string {
	return fmt.Sprintf("index mismatch: expected %v, actual %v", e.Expected, e.Actual)
}

type ErrLengthMismatch struct {
	Expected int
	Actual   int
}

func (e ErrLengthMismatch) Error() string {
	return fmt.Sprintf("length mismatch: expected %d, actual %d", e.Expected, e.Actual)
}
