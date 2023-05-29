package main

import (
	"fmt"
	"os"
	"picasso"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <text>")
		return
	}

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	str := strings.TrimSpace(os.Args[1])
	asciiArray := make([]int, len(str))

	for i, char := range str {
		asciiArray[i] = int(char) - 32
	}

	res := make([]string, 8)
	for i := 0; i < 8; i++ {
		_, err = file.Seek(0, 0) // Reset file cursor to the beginning
		if err != nil {
			fmt.Println("Error seeking file:", err)
			return
		}

		for _, ascii := range asciiArray {
			line := picasso.ReadLine(file, 2+ascii*9+i)
			res[i] += line
		}
		fmt.Println(res[i])
	}
}
