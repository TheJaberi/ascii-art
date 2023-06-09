package main

import (
	"color"
	"fmt"
	"os"
	"strings"
)

func main() {
	var colorFlagCount = 0
	var colorFlag string
	var pos int
	var str string
	var specified string
	var specific = false
	// Check the number of command-line arguments
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage:\ngo run . <text>\ngo run . --color=<color> <text>\ngo run . --color=<color> <specific letters> <text>")
		return
	} else if len(os.Args) == 2  {
		if strings.HasPrefix(os.Args[1], "--color="){
			fmt.Println("Usage:\ngo run . <text>\ngo run . --color=<color> <text>\ngo run . --color=<color> <specific letters> <text>")
			return
		}
		str = os.Args[1]
	}else if len(os.Args) == 3 {
		for i, arg := range os.Args[1:] {
			if strings.HasPrefix(arg, "--color=") {
				colorFlag = strings.TrimPrefix(arg, "--color=")
				pos = i
				colorFlagCount++
			}
		}
		if colorFlagCount < 1 || colorFlagCount > 1 {
			fmt.Println("Usage:\ngo run . <text>\ngo run . --color=<color> <text>\ngo run . --color=<color> <specific letters> <text>")
			return
		}
		if pos == 1 {
			str = os.Args[2]
			
		}else {
			str = os.Args[1]
		}
	} else if len(os.Args) == 4 {
		for i, arg := range os.Args {
			if strings.HasPrefix(arg, "--color=") {
				colorFlag = strings.TrimPrefix(arg, "--color=")
				pos = i
				colorFlagCount++
			}
		}
		if colorFlagCount < 1 || colorFlagCount > 1 {
			fmt.Println("Usage:\ngo run . <text>\ngo run . --color=<color> <text>\ngo run . --color=<color> <specific letters> <text>")
			return
		}
		if pos == 1 {
			if len(os.Args[2]) > len(os.Args[3]) {
				if color.ContainsSubstring(os.Args[2], os.Args[3]){
				str = os.Args[2]
				specified = os.Args[3]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			} else {
				if color.ContainsSubstring(os.Args[3], os.Args[2]){
					str = os.Args[3]
					specified = os.Args[2]
					} else {
						fmt.Println("you specified something unavailable")
						return
					}
			}
		} else if pos == 2 {
			if len(os.Args[1]) > len(os.Args[3]) {
				if color.ContainsSubstring(os.Args[2], os.Args[3]){
				str = os.Args[1]
				specified = os.Args[3]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			} else {
				if color.ContainsSubstring(os.Args[3], os.Args[2]){
					str = os.Args[3]
					specified = os.Args[1]
					} else {
						fmt.Println("you specified something unavailable")
						return
					}
			}
		}else {
			if len(os.Args[1]) > len(os.Args[2]) {
				if color.ContainsSubstring(os.Args[1], os.Args[2]){
				str = os.Args[1]
				specified = os.Args[2]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			} else {
				if color.ContainsSubstring(os.Args[2], os.Args[1]){
					str = os.Args[2]
					specified = os.Args[1]
					} else {
						fmt.Println("you specified something unavailable")
						return
					}
			}
		}
		specific = true
	}

	if specific {
	fmt.Println(specified)
	}
	fmt.Println(colorFlag)
	
	
	// Open the "standard.txt" file
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	

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