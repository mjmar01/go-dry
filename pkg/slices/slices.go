package slices

func ElementInSlice[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func Distinct[T comparable](slice []T) (distinct []T) {
	set := make(map[T]bool)
	for _, e := range slice {
		set[e] = true
	}
	for key := range set {
		distinct = append(distinct, key)
	}
	return
}
