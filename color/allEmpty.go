package color

// AllEmpty checks if all elements in the given string slice are empty strings.
// It returns true if all elements are empty, and false otherwise.
func AllEmpty(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			return false
		}
	}
	return true
}
