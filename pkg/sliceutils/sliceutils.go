package sliceutils

func Transform[T, V any](slice []T, transformation func(T) V) []V {
	newSlice := []V{}
	for i := 0; i < len(slice); i++ {
		newSlice = append(newSlice, transformation(slice[i]))
	}
	return newSlice
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	newSlice := []T{}
	for i := 0; i < len(slice); i++ {
		if predicate(slice[i]) {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}
