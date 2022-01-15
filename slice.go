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
