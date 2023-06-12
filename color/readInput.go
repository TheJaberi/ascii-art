package color

import (
	"io"
	"os"
	"fmt"
	"bufio"
)


func ReadInput(file io.Reader) [][]string {
	var arr [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the line
		line := scanner.Text()

		// Initialize a slice to hold characters
		chars := make([]string, len(line))

		// Split line into characters
		for i, char := range line {
			chars[i] = string(char)
		}

		// Add characters to the 2D array
		arr = append(arr, chars)
	}

	// Check for errors during scanning
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		os.Exit(0)
	}

	return arr
}