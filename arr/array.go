package arr

// Identify whether the given array contians the given element.
func Contains[T comparable](arr []T, e T) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}
