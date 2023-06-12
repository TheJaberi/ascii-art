package color

import (
	"os"
	"strings"
)

func Compare(arr []string, allAscii []string) string {
	strResult := ""
	if len(arr) == 0 || len(allAscii) == 0 {
		os.Exit(0)
	}

	for _, wholeAsciiLine := range arr {
		if wholeAsciiLine == "" {
			continue
		}

		splitAscii := strings.Split(wholeAsciiLine, "\n")

		startIndex := -1
		idx := 0
		for i := 0; i < len(allAscii); i++ {
			if allAscii[i] == splitAscii[idx] {
				if idx == 0 {
					startIndex = i
				}
				idx++
					
				if idx == len(splitAscii) {
					numResult := ((startIndex - 2) / 9) + 33
					if numResult < 32 {
						numResult = 32
					}
					strResult += string(numResult)
					idx = 0
				}
			} else {
				startIndex = -1
				idx = 0
			}
		}
	}

	return strResult
}
