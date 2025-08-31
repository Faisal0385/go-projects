package crud

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"cli-user-management/config"
	"cli-user-management/models"

)

// ReadUser allows the user to look up a single user by their ID and display their details.
func ReadUser(reader *bufio.Reader) {
	// Display section title
	fmt.Println("\n== Read User By ID ==")

	// Prompt the user to enter the ID of the user they want to read
	fmt.Print("Enter User ID: ")

	// Read input from the terminal
	idInput, _ := reader.ReadString('\n')

	// Remove leading/trailing spaces and newline characters
	idInput = strings.TrimSpace(idInput)

	// Convert the input string to an integer
	id, err := strconv.Atoi(idInput)
	if err != nil {
		// If conversion fails (input is not a number), show an error and exit
		fmt.Println("❌ Invalid ID.")
		return
	}

	// Declare a User variable to hold the result
	var user models.User

	// Fetch the user from the database by ID using GORM's First method
	result := config.DB.First(&user, id)

	// If user not found, show an error and exit
	if result.Error != nil {
		fmt.Println("❌ User not found.")
		return
	}

	// Display the user's details in a readable format
	fmt.Printf(
		"\nUser Details:\n  ID: %d\n  Name: %s\n  Email: %s\n  Phone: %s\n  Created At: %s\n",
		user.ID,
		user.Name,
		user.Email,
		user.Phone,
		user.CreatedAt.Format("2006-01-02 15:04:05"), // Format timestamp nicely
	)
}
