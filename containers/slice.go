package containers

// Contain return true if given string is in the string slice.
func Contain[T comparable](slice []T, s T) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// Find return first index of given string. -1 for non existent.
func Find[T comparable](slice []T, s T) int {
	for i, item := range slice {
		if item == s {
			return i
		}
	}
	return -1
}

// FirstString returns the first element in a slice.
func First[T any](strSlice []T) (ret T) {
	if len(strSlice) == 0 {
		return
	}
	return strSlice[0]
}

// LastString returns the last element in a slice.
func Last[T any](strSlice []T) (ret T) {
	if len(strSlice) == 0 {
		return
	}
	return strSlice[len(strSlice)-1]
}

// Repeat returns a slice with `n` repeated same `element` string.
func Repeat[T any](n int, element T) []T {
	if n <= 0 {
		return nil
	}
	ret := make([]T, n)
	for i := 0; i < n; i++ {
		ret[i] = element
	}
	return ret
}

func InterfaceSlice[T any](arr []T) []interface{} {
	ret := make([]interface{}, len(arr))
	for i := range arr {
		ret[i] = arr[i]
	}
	return ret
}

// Clone copies the underlying array of given `strSlice` to a new slice.
func Clone[T any](strSlice []T) []T {
	if strSlice == nil {
		return strSlice
	}
	return append(strSlice[:0:0], strSlice...)
}

func Reverse[T any](slice []T) []T {
	ret := Clone(slice)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}
