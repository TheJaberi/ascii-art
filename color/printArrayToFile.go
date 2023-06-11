package color

import(
	"os"
	"fmt"
)

func PrintArrayToFile(array []string, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, str := range array {
		fmt.Fprintln(file, str)
	}

	return nil
}