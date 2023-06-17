package ascii

import (
	"bufio"
	"io"
	"fmt"
	"os"
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
	fmt.Println("Error in file holding ascii, make sure it's the same as in the repo")
	os.Exit(0)
	return ""
}
