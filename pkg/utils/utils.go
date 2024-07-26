package utils

// Compare returns an integer comparing two values.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func Compare[T int | float64 | string](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Swap swaps the values of two variables of the same type.
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// Max returns the larger of x or y.
func Max[T int | float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y.
func Min[T int | float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Contains checks if a slice contains a specific element.
func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// RemoveDuplicates removes duplicate elements from a slice.
func RemoveDuplicates[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	result := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
