package main

import (
	"fmt"
	"os"
	"reverse"
	"strings"
)

func main() {
	// Check the number of command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <text>")
		return
	}

	flag := "--reverse="
	str := os.Args[1] // Get the input text from command-line argument

	if strings.HasPrefix(str, flag) {
		// Extract the value after the equal sign
		value := str[len(flag):]

		// Process the value
		fmt.Println("Value:", value)
	

	// Open the "standard.txt" file
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	allAscii := reverse.ReadLines(file)
	for i := 0; i <= 854;i++ {
	fmt.Println(allAscii[i])
	}
   

	// Validate characters in the input text
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			fmt.Println("Use English chars only")
			os.Exit(0)
		}
	}
	}
}
