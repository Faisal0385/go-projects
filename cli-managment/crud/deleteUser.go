package crud

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"cli-user-management/config"
	"cli-user-management/models"
)

// DeleteUser allows the user to permanently delete a user by ID from the database.
func DeleteUser(reader *bufio.Reader) {
	// Display section title
	fmt.Println("\n== Delete User ==")

	// Prompt the user to enter the ID of the user to delete
	fmt.Print("Enter User ID: ")

	// Read input from terminal
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput) // Remove spaces and newline

	// Convert the input string to an integer
	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("❌ Invalid ID.") // Input is not a valid number
		return
	}

	// Declare a User variable to hold the database record
	var user models.User

	// Fetch the user from the database using the given ID
	result := config.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println("❌ User not found.") // No user exists with this ID
		return
	}

	// Ask for confirmation before deleting
	fmt.Printf("Are you sure you want to delete user '%s' (ID: %d)? (y/N): ", user.Name, user.ID)
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))

	// Cancel deletion if the user does not confirm
	if confirm != "y" && confirm != "yes" {
		fmt.Println("❌ Cancelled.")
		return
	}

	// Permanently delete the user from the database (Unscoped bypasses soft delete)
	if err := config.DB.Unscoped().Delete(&user).Error; err != nil {
		fmt.Println("❌ Failed to delete user:", err)
		return
	}

	// Success message
	fmt.Println("✅ User deleted successfully.")
}
