package color

import (
	"strings"
)

func ColumnSlicer(twoDArray [][]string) []string {
	result := make([]string, 0)
	for col := 0; col < len(twoDArray[0]); col++ {
		isEmptyColumn := true
		for row := 0; row < len(twoDArray); row++ {
			if twoDArray[row][col] != " " {
				isEmptyColumn = false
				break
			}
		}
		if isEmptyColumn {
			slice := make([]string, 0)
			for row := 0; row < len(twoDArray); row++ {
				slice = append(slice, strings.Join(twoDArray[row][:col+1], ""))
			}
			result = append(result, strings.Join(slice, "\n"))
			for row := 0; row < len(twoDArray); row++ {
				twoDArray[row] = twoDArray[row][col+1:]
			}
			col = -1 // reset column index
		}
	}
	// Add remaining data if any
	if len(twoDArray[0]) > 0 {
		slice := make([]string, 0)
		for row := 0; row < len(twoDArray); row++ {
			slice = append(slice, strings.Join(twoDArray[row], ""))
		}
		result = append(result, strings.Join(slice, "\n"))
	}
	return result
}