package utils

import (
	"bufio"
	"fmt"
	"strings"
)

// ReadOptional is a helper function that reads input from the user.
// Unlike ReadRequired, this function allows the user to leave the input blank.
func ReadOptional(reader *bufio.Reader, prompt string) string {
	// Display the prompt message to the user
	fmt.Print(prompt)

	// Read the user input until a newline character is encountered
	// The function returns two values: the input string and an error
	// We are ignoring the error with `_` for simplicity
	text, _ := reader.ReadString('\n')

	// Trim any leading/trailing whitespace (including newline \n)
	// and return the cleaned string
	return strings.TrimSpace(text)
}
