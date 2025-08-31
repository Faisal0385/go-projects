package crud

import (
	"fmt"
	"strings"

	"cli-user-management/config"
	"cli-user-management/models"
)

// ListUsers retrieves all users from the database and displays them in a table format.
func ListUsers() {
	// Display section title
	fmt.Println("\n== All Users ==")

	// Create a slice to hold multiple User records
	var users []models.User

	// Fetch all users from the database using GORM
	result := config.DB.Find(&users)

	// Check for errors during the database query
	if result.Error != nil {
		fmt.Println("❌ Failed to fetch users:", result.Error)
		return
	}

	// If no users are found, show a message and return
	if len(users) == 0 {
		fmt.Println("❌ No users found.")
		return
	}

	// Print table headers with formatting
	fmt.Printf("%-5s %-20s %-30s %-15s\n", "ID", "Name", "Email", "Phone")
	fmt.Println(strings.Repeat("-", 75)) // Print a separator line

	// Loop through each user and print their details in formatted columns
	for _, u := range users {
		fmt.Printf("%-5d %-20s %-30s %-15s\n", u.ID, u.Name, u.Email, u.Phone)
	}
}
