package gen

// Slice wraps a slice of values with helper functions
type Slice[T comparable] []T

// Map returns a map of the given slice where the values of the slice
// are the Map keys
func (s Slice[T]) Map() Map[T, struct{}] {
	out := map[T]struct{}{}

	// Ensure that the slice is not nil or empty
	if len(s) == 0 {
		return out
	}

	// Add each element to the map
	// NOTE: This will by default overwrite any existing value
	for _, v := range s {
		out[v] = struct{}{}
	}

	return out
}
