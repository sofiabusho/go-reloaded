package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Function to handle all the text modifications (to be implemented)
func handleAllModifications(content string) string {

	// Handle hex and bin
	content = handleHexAndBin(content)

	// Handle (up), (low) and (cap)
	content = handleTextModifications(content)

	// Handle punctuation formatting
	content = adjustPunctuation(content)

	// Handle 'a' to 'an'
	content = handleAtoAn(content)

	return content
}

// Function to handle hex and bin modifications
func handleHexAndBin(content string) string {
	// Regex looks for sequences of hexadecimal digits followed by (hex) and  binary numbers followed by (bin)
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

// Function to handle the (up, n), (low, n) and (cap, n)
func handleTextModifications(text string) string {

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

		startWord := len(words) - count // Identify the starting word based on the 'count'
		if startWord < 0 {
			startWord = 0 // Ensure we do not go out of range, by making sure 'startWord' is not negative
		}

		// Loop through the words starting from 'startWord'
		for j := startWord; j < len(words); j++ {
			words[j] = caseFunc(words[j]) // Apply the 'caseFunc' to modify the case of each word
		}

		// Rebuild the text without the marker and with the modifications applied
		text = strings.Join(words, " ") + text[match[1]:]
	}

	return text
}

// Functon to capitilize the first letter of a word
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// Function to adjust punctuation
func adjustPunctuation(text string) string {

	// Remove spaces around individual punctuation marks
	text = regexp.MustCompile(`\s*([.,!?;:])`).ReplaceAllString(text, "$1")

	// Add a space after punctuation followed directly by a word character
	text = regexp.MustCompile(`([.,!?;:])([A-Za-z])`).ReplaceAllString(text, "$1 $2")

	// Handle multiple words within single quotes
	text = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(text, "'$1'")

	// Handle ellipses (...) or multiple punctuation marks like !? by ensuring no space after them
	text = regexp.MustCompile(`([.]{3}|[!?]{2,})([^\s])`).ReplaceAllString(text, "$1 $2")

	return text
}

// Function to handle 'a' to 'an'
func handleAtoAn(text string) string {
	re := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH]\w*)`)
	// Regex to find "a" followed by a word starting with a vowel
	return re.ReplaceAllStringFunc(text, func(match string) string {
		if match[0] == 'A' {
			return "An " + match[2:] // Capitalized replacement for "A"
		}
		return "an " + match[2:] // Lowercase replacement for "a"
	})
}
