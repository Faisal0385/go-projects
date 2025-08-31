package crud

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"cli-user-management/config"
	"cli-user-management/models"
	"cli-user-management/utils"
)

// UpdateUser allows the user to modify existing user information by ID.
func UpdateUser(reader *bufio.Reader) {
	// Display section title
	fmt.Println("\n== Update User ==")

	// Prompt for the user ID
	fmt.Print("Enter User ID: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput) // Trim spaces/newline

	// Convert ID string to integer
	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("❌ Invalid ID.") // Input is not a number
		return
	}

	// Fetch the user from the database
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		fmt.Println("❌ User not found.") // No user exists with this ID
		return
	}

	// Show current information
	fmt.Printf("\nCurrent Info:\n  Name: %s\n  Email: %s\n  Phone: %s\n", user.Name, user.Email, user.Phone)

	// Prompt for new Name (optional, leave blank to keep current)
	fmt.Print("Enter New Name (leave blank to keep current): ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	if newName != "" {
		user.Name = newName
	}

	// Prompt for new Email (optional)
	fmt.Print("Enter New Email (leave blank to keep current): ")
	newEmail, _ := reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)
	if newEmail != "" {
		// Validate new email format
		if !utils.IsValidEmail(newEmail) {
			fmt.Println("❌ Invalid email format.")
			return
		}
		user.Email = newEmail
	}

	// Prompt for new Phone (optional)
	fmt.Print("Enter New Phone (leave blank to keep current): ")
	newPhone, _ := reader.ReadString('\n')
	newPhone = strings.TrimSpace(newPhone)
	if newPhone != "" {
		// Validate new phone format
		if !utils.IsValidPhone(newPhone) {
			fmt.Println("❌ Invalid phone number.")
			return
		}
		user.Phone = newPhone
	}

	// Show updated information before saving
	fmt.Printf("\nUpdated Info:\n  Name: %s\n  Email: %s\n  Phone: %s\n", user.Name, user.Email, user.Phone)

	// Ask for confirmation before saving changes
	if !utils.AskYesNo(reader, "Save changes? (y/N): ") {
		fmt.Println("❌ Cancelled.")
		return
	}

	// Save the updated user information to the database
	if err := config.DB.Save(&user).Error; err != nil {
		fmt.Println("❌ Failed to update user:", err)
		return
	}

	// Success message
	fmt.Println("✅ User updated successfully.")
}
