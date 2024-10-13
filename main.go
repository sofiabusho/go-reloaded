package main

import (
	"fmt"
	"os"
)

func main() {
	// Checks if exactly two arguments are passed (input and output filenames)
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	// Assign input and output filenames from command line arguments
	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	// Read the content from the filename input
	content, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Println("Error reading the input file:", err)
		return
	}

	if len(content) == 0 {
		fmt.Println("The input file is empty, please add some content to it!")
	}

	// Convert byte content to string for manipulation
	modifiedContent := handleAllModifications(string(content))

	// Write the modified content to the output file
	err = os.WriteFile(outputFilename, []byte(modifiedContent), 0o644)
	if err != nil {
		fmt.Println("Error writing to the output file:", err)
		return
	}
}
