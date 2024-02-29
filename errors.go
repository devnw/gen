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

import "fmt"

type IndexMismatchError[T any] struct {
	I        int
	Expected T
	Actual   T
}

func (e IndexMismatchError[T]) Error() string {
	return fmt.Sprintf("index mismatch: expected %v, actual %v", e.Expected, e.Actual)
}

type LengthMismatchError struct {
	Expected int
	Actual   int
}

func (e LengthMismatchError) Error() string {
	return fmt.Sprintf("length mismatch: expected %d, actual %d", e.Expected, e.Actual)
}
