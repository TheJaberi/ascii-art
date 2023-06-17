package ascii
// containsSubstring checks if a string contains a specific substring.
func ContainsSubstring(str, substr string) bool {
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr { // Check if the current substring matches the target substring
			return true // Return true if a match is found
		}
	}
	return false // Return false if no match is found
}
