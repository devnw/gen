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

import "context"

// Slice wraps a slice of values with helper functions
type Slice[T comparable] []T

// Map returns a map of the given slice where the values of the slice
// are the Map keys
func (s Slice[T]) Map() Map[T, any] {
	out := map[T]any{}

	// Ensure that the slice is not nil or empty
	if len(s) == 0 {
		return out
	}

	// Add each element to the map
	// NOTE: This will by default overwrite any existing value
	for _, v := range s {
		out[v] = nil
	}

	return out
}

// Chan converts the slice to a channel of type T
//
// NOTE: This function does NOT use a buffered channel.
func (s Slice[T]) Chan(ctx context.Context) <-chan T {
	ctx, _ = _ctx(ctx)
	out := make(chan T)

	go func(out chan<- T) {
		defer close(out)

		for _, v := range s {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}(out)

	return out
}
