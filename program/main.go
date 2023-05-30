package main

import (
	"fmt"
	"os"
	"picasso"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <text>")
		return
	}

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	str := os.Args[1]

	for _,letter := range str {
		if letter < ' ' || letter > '~' {
			fmt.Println("Use english chars only")
			os.Exit(0)
		}
	}
	asciiArray := make([]int, len(str))

	for i, char := range str {
		asciiArray[i] = int(char) - 32
	}

	multipleNewLines := false
	var lastcharacter rune
	for _, character := range str {
		if character == 'n' && lastcharacter == '\\' {
			multipleNewLines = true
		}
		lastcharacter = character
	}

	if multipleNewLines {
		words := strings.Split(str, "\\n")

		for _, currentWord := range words {
			for i := 0; i < 8; i++ {
		
				for _, currentLetter := range currentWord {
					_, err = file.Seek(0, 0) // Reset file cursor to the beginning
				if err != nil {
					fmt.Println("Error seeking file:", err)
					return
				}
		
					line := picasso.ReadLine(file, 2+(int(currentLetter)-32)*9+i)
					fmt.Print(line)
				}
				fmt.Println()
			}
			
		}
	} else {
		for i := 0; i < 8; i++ {
		
		for _, currentLetter := range str {
			_, err = file.Seek(0, 0) // Reset file cursor to the beginning
		if err != nil {
			fmt.Println("Error seeking file:", err)
			return
		}

			line := picasso.ReadLine(file, 2+(int(currentLetter)-32)*9+i)
			fmt.Print(line)
		}
		fmt.Println()
	}
	}
}

