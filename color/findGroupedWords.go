package color

import "strings"

// FindGroupedWord searches for the grouped word in str and returns a slice of starting and ending indices.
func FindGroupedWord(groupedWord string, str string) [][]int {
	var indices [][]int
	start := -1
	offset := 0

	for {
		start = strings.Index(str[offset:], groupedWord)
		if start == -1 {
			break
		}
		start += offset
		end := start + len(groupedWord) - 1
		indices = append(indices, []int{start, end})
		offset = end + 1
	}

	return indices
}