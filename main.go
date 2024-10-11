package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	// Apply modifications to the content
	modifiedContent := string(content)

	err = os.WriteFile(outputFilename, []byte(modifiedContent), 0o644)
	if err != nil {
		fmt.Println("Error writing to the output file:", err)
		return
	}

	if strings.Contains(inputFilename, "(hex)") {
		hexIndex := strings.Index(content, "(hex)")
		lastSpaceIndex := strings.LastIndex(content[:hexIndex], " ")
		hexWord := content[lastSpaceIndex+1 : hexIndex-1]

		decimalValue, err := strconv.ParseInt(hexWord, 16, 64)

		modifiedContent := strings.Replace(content, hexWord+" (hex)", fmt.Spintf("%d", decimalValue), 1)
	}

}
