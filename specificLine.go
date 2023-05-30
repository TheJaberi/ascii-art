package picasso

import (
	"bufio"
	"io"
)

func ReadLine(theFile io.Reader, lineNum int) string {
	sc := bufio.NewScanner(theFile)
	currentLine := 0
	for sc.Scan() {
		currentLine++
		if currentLine == lineNum {
			return sc.Text()
		}

	}
	return ""
}
