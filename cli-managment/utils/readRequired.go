package utils

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadRequired repeatedly prompts the user for input until they enter a non-empty value.
// It ensures that the returned string is never empty.
func ReadRequired(reader *bufio.Reader, prompt string) string {
	for {
		// Show the prompt message (e.g., "Enter your name: ")
		fmt.Print(prompt)

		// Read user input until the Enter key is pressed
		text, _ := reader.ReadString('\n')

		// Remove leading/trailing spaces and newline characters
		text = strings.TrimSpace(text)

		// If the user entered something (not empty), return it
		if text != "" {
			return text
		}

		// Otherwise, show an error message and ask again
		fmt.Println("❌ This field is required.")
	}
}
