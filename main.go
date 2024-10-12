package main

import (
	"fmt"
	"os"
	"regexp"
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

	// Convert byte content to string for manipulation
	modifiedContent := string(content)

	// Check if the file contains "(hex)" or "(bin)" and replace the hex value with decimal and the bin value with binary
	modifiedContent = handleHexAndBin(modifiedContent)

	modifiedContent = handleAllModifications(modifiedContent)

	// Write the modified content to the output file
	err = os.WriteFile(outputFilename, []byte(modifiedContent), 0o644)
	if err != nil {
		fmt.Println("Error writing to the output file:", err)
		return
	}
}

func handleHexAndBin(content string) string {
	// Regex to find hex and bin patterns
	hexPattern := regexp.MustCompile(`(\b[0-9a-fA-F]+)\s*\(hex\)`)
	binPattern := regexp.MustCompile(`(\b[01]+)\s*\(bin\)`)

	// Replace hex values with decimal equivalents
	content = hexPattern.ReplaceAllStringFunc(content, func(hexMatch string) string {
		hexValue := hexPattern.FindStringSubmatch(hexMatch)[1]
		decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
		if err != nil {
			return hexMatch // Return original if convencion fails
		}
		return fmt.Sprintf("%d", decimalValue)
	})

	// Replace bin values with decimal equivalents
	content = binPattern.ReplaceAllStringFunc(content, func(binMatch string) string {
		binValue := binPattern.FindStringSubmatch(binMatch)[1]
		decimalValue, err := strconv.ParseInt(binValue, 2, 64)
		if err != nil {
			return binMatch // Return original if conversion fails
		}
		return fmt.Sprintf("%d", decimalValue)
	})

	return content
}

// Function to handle text modifications (to be implemented)
func handleAllModifications(content string) string {
	// Handle (up), (low) and (cap)
	content = handleTextModifications(content)

	// Handle punctuation formatting
	content = adjustPunctuation(content)

	// Handle 'a' to 'an'
	content = handleAtoAn(content)

	return content
}

func handleTextModifications(text string) string {
	// Handle (up, n), (low, n) and (cap, n)
	re := regexp.MustCompile(`\s*\((up|low|cap)(?:,\s*(\d+))?\)`)
	matches := re.FindAllStringSubmatchIndex(text, -1)

	// Check if there are any matches
	if matches == nil {
		return text // Return the original text if no matches are found
	}

	// Apply case modifications for each match
	for i := len(matches) - 1; i >= 0; i-- { // Iterate backward to prevent index shifting
		match := matches[i]
		modType := text[match[2]:match[3]] // Get the modification type (up/low/cap)

		count := 1
		if match[4] != -1 { // Check if a number is provided for (up, n), (low, n), (cap, n)
			count, _ = strconv.Atoi(text[match[4]:match[5]])
		}

		// Determine which function to use based on (up), (low), or (cap)
		var caseFunc func(string) string
		switch modType {
		case "up":
			caseFunc = strings.ToUpper
		case "low":
			caseFunc = strings.ToLower
		case "cap":
			caseFunc = capitalizeWord
		}

		// Extract the part of the text before the marker and apply case modification
		words := strings.Fields(text[:match[0]]) // Extract the words before the marker

		startWord := len(words) - count // Identify the starting word
		if startWord < 0 {
			startWord = 0 // Ensure we do not go out of range
		}

		for j := startWord; j < len(words); j++ {
			words[j] = caseFunc(words[j])
		}

		// Rebuild the text without the marker and with the modifications applied
		text = strings.Join(words, " ") + text[match[1]:]
	}

	return text
}

// Capitilize the first letter of a word
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// Function to adjust punctuation
func adjustPunctuation(text string) string {

	// Remove spaces around punctuation
	text = regexp.MustCompile(`\s*([.,!?;:]+)\s*`).ReplaceAllString(text, "$1 ")

	// Format single quotes
	text = regexp.MustCompile(`'\s*([^']+?)\s*'`).ReplaceAllString(text, "$1 ")

	// Handle multiple words within single quotes
	text = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(text, "$1 ")

	return text
}

// Handle 'a' to 'an'
func handleAtoAn(text string) string {
	// Regex to find "a" followed by a word starting with a vowel
	return regexp.MustCompile(`\ba\s+([aeiouhAEIOUH]\w*)`).ReplaceAllString(text, "an $1")
}
