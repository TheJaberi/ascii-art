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
	var isWithinRange bool
	var outflag = false
	var revflag = false
	var outputFlag = ""
	var reverseFlag = ""
	var indices [][]int
	var outputArray []string
	var currentLine string
	var fileName = "standard.txt"

	// Check the number of command-line arguments
	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println("Usage:\ngo run . [OPTION] [STYLE] [STRING]\nExample: go run . --color=red standard Hello")
		return
	}

	//check for reverse
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--reverse=") {
			reverseFlag = strings.TrimPrefix(arg, "--reverse=")
			if reverseFlag == "" {
				return
			} else if !strings.HasSuffix(reverseFlag, ".txt"){
				reverseFlag += ".txt"
			}
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			revflag = true
			break
		}
	}
	if revflag == true {
		if len(os.Args) != 1 {
			fmt.Println("when using the reverse flag you should only have one argument which is the flag")
			return
		}
		revfile, err := os.ReadFile(reverseFlag)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		//make the input into a 2d array
		//inputArray := color.ReadInput(revfile)

		//make the 2d array as a 1d array joined by "\n"
		//inputAsLine := color.ColumnSlicer(inputArray)

		// Open the file
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		color.ReverseAsciiArt(revfile, file)
		//asciiArray := color.LinesToArray(file)

		//fmt.Println(color.Compare(inputAsLine, asciiArray))

		return
	}


	// Identify and remove style argument
	for i, arg := range os.Args {
		if arg == "standard" || arg == "shadow" || arg == "thinkertoy" {
			fileName = arg + ".txt"
			// Remove the style argument from os.Args
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			break
		}
	}

	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--output=") {
			outputFlag = strings.TrimPrefix(arg, "--output=")
			if outputFlag == "" {
				return
			} else if !strings.HasSuffix(outputFlag, ".txt"){
				outputFlag += ".txt"
			}
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			outflag = true
			break
		}
	}

	if fileName == "" {
		fmt.Println("Please specify a valid style: standard, shadow, or thinkertoy.")
		return
	} else if len(os.Args) == 2 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			fmt.Println("you cannot have a flag as something to print")
			return
		} else {
			str = os.Args[1]
		}
	} else if len(os.Args) == 3 {
		for i, arg := range os.Args {
			if strings.HasPrefix(arg, "--color=") {
				colorFlag = strings.TrimPrefix(arg, "--color=")
				pos = i
				colorFlagCount++
			}
		}
		if colorFlagCount < 1 || colorFlagCount > 1 {
			fmt.Println("Usage:\ngo run . [OPTION] [STRING]\nExample: go run . --color=red Hello")
			return
		}
		if pos == 1 {
			str = os.Args[2]
		} else {
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
			fmt.Println("Usage:\ngo run . [OPTION] [STRING]\nExample: go run . --color=red eeetee")
			return
		}
		if pos == 1 {
			if len(os.Args[2]) > len(os.Args[3]) {
				if color.ContainsSubstring(os.Args[2], os.Args[3]) {
					str = os.Args[2]
					specified = os.Args[3]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			} else {
				if color.ContainsSubstring(os.Args[3], os.Args[2]) {
					str = os.Args[3]
					specified = os.Args[2]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			}
		} else if pos == 2 {
			if len(os.Args[1]) > len(os.Args[3]) {
				if color.ContainsSubstring(os.Args[1], os.Args[3]) {
					str = os.Args[1]
					specified = os.Args[3]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			} else {
				if color.ContainsSubstring(os.Args[3], os.Args[1]) {
					str = os.Args[3]
					specified = os.Args[1]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			}
		} else {
			if len(os.Args[1]) > len(os.Args[2]) {
				if color.ContainsSubstring(os.Args[1], os.Args[2]) {
					str = os.Args[1]
					specified = os.Args[2]
				} else {
					fmt.Println("you specified something unavailable")
					return
				}
			} else {
				if color.ContainsSubstring(os.Args[2], os.Args[1]) {
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

	// Open the file
	file, err := os.Open(fileName)
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

		// Check if the specified string is a grouped word in the current word
		if specific {
			indices = color.FindGroupedWord(specified, currentWord)
		}

		// Display each line of the stylized text for the current word
		for i := 0; i < 8; i++ {
			// Iterate over each letter in the current word
			for j, currentLetter := range currentWord {
				_, err = file.Seek(0, 0) // Reset file cursor to the beginning
				if err != nil {
					fmt.Println("Error seeking file:", err)
					return
				}

				// Read the appropriate line from the "standard.txt" file
				line := color.ReadLine(file, 2+(int(currentLetter)-32)*9+i)

				// Check if the current index is within any of the grouped word ranges
				for _, indexPair := range indices {
					if j >= indexPair[0] && j <= indexPair[1] {
						isWithinRange = true
						break
					}
				}

				// Apply the specified color if the current index is within the range, otherwise apply the default color
				if outflag {
					currentLine += line
				} else if specific && isWithinRange {
					fmt.Print(color.ColorSelector(colorFlag) + line + color.ColorSelector("reset"))
					isWithinRange = false
				} else if !specific && colorFlagCount == 1 {
					fmt.Print(color.ColorSelector(colorFlag) + line + color.ColorSelector("reset"))
				} else {
					fmt.Print(line)
				}
			}
			if outflag {
				outputArray = append(outputArray, currentLine)
				currentLine = ""
			} else if !outflag {
				fmt.Println()
			}
		}
	}
	if outflag {
	err := color.PrintArrayToFile(outputArray, outputFlag)
		if err != nil {
			fmt.Println("Error:", err)
		return
		}
	}
}
