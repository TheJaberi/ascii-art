package main

import (
	"color"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check the number of command-line arguments
	if len(os.Args) <= 2 {
		fmt.Println("Usage: go run . <text>")
		return
	}

	colorFlag := flag.String("color", "white", "Specify the color") // Define a string flag named "color" with a default value of "red"

	flag.Parse() // Parse the command-line flags

	chosencolor := *colorFlag // Access the value of the "color" flag

	textColor := color.ColorSelector(chosencolor)

	fmt.Print(textColor)

	// Open the "standard.txt" file
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Get the input text from command-line argument
	str := os.Args[1]

	// Validate characters in the input text
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			fmt.Println("Use English chars only")
			os.Exit(0)
		}
	}

	// Split the input text into words and check if it contains only new lines
	words := strings.Split(str, "\\n")
	if color.AllEmpty(words) {
		words = words[1:]
	}

	// Iterate over each word in the input text
	for current, currentWord := range words {
		// Skip empty words and print a newline
		if words[current] == "" {
			fmt.Println()
			continue
		}

		// Display each line of the stylized text for the current word
		for i := 0; i < 8; i++ {

			// Iterate over each letter in the current word
			for _, currentLetter := range currentWord {
				_, err = file.Seek(0, 0) // Reset file cursor to the beginning
				if err != nil {
					fmt.Println("Error seeking file:", err)
					return
				}

				// Read the appropriate line from the "standard.txt" file
				line := color.ReadLine(file, 2+(int(currentLetter)-32)*9+i)
				fmt.Print(line)
			}
			fmt.Println()
		}
	}
}
