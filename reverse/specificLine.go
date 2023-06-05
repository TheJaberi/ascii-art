package reverse

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadLines reads all lines from the given file (io.Reader) and returns them as a slice of strings.
// Each line of the file is stored in an index of the slice.
// If an error occurs during file reading, it prints an error message and exits the program.
func ReadLines(theFile io.Reader) []string {
	sc := bufio.NewScanner(theFile)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(0)
	}

	return lines
}
