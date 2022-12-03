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

func Repeat[T any](count int, elements ...T) (out []T) {
	out = make([]T, 0, len(elements)*count)
	for i := 0; i < count; i++ {
		out = append(out, elements...)
	}
	return
}

func IndexOf[T comparable](slice []T, value T) int {
	for i, element := range slice {
		if element == value {
			return i
		}
	}
	return -1
}
