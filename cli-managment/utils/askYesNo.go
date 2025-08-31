package utils

import (
	"bufio"
	"fmt"
	"strings"

)

// AskYesNo displays a prompt to the user and waits for a Yes/No input.
// It returns true if the user enters "y" or "yes", otherwise false.
func AskYesNo(reader *bufio.Reader, prompt string) bool {
	// Show the prompt message (e.g., "Do you want to continue? (y/n): ")
	fmt.Print(prompt)

	// Read user input until the Enter key is pressed
	text, _ := reader.ReadString('\n')

	// Trim whitespace and convert input to lowercase for easier comparison
	text = strings.TrimSpace(strings.ToLower(text))

	// Return true if input is "y" or "yes", false otherwise
	return text == "y" || text == "yes"
}