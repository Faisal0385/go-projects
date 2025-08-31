package crud

import (
	"bufio"
	"fmt"

	"cli-user-management/config"
	"cli-user-management/models"
	"cli-user-management/utils"

)

// CreateUser handles the process of creating a new user in the system.
// It validates input for required fields, checks for duplicates, and saves to the database.
func CreateUser(reader *bufio.Reader) {
	// Display section title
	fmt.Println("\n== Create User ==")

	// Prompt for Name (cannot be empty, enforced by ReadRequired)
	name := utils.ReadRequired(reader, "Enter Name (required): ")

	// Validate Email input
	var email string
	for {
		// Ask for email (required)
		email = utils.ReadRequired(reader, "Enter Email (required): ")

		// Check email format (using custom validation helper)
		if !utils.IsValidEmail(email) {
			fmt.Println("❌ Invalid email format.")
			continue // keep asking until valid
		}

		// Check if email already exists in the database
		var count int64
		config.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
		if count > 0 {
			fmt.Println("❌ Email already exists.")
			continue // ask again if duplicate
		}
		break // valid and unique email -> exit loop
	}

	// Validate Phone input (optional)
	var phone string
	for {
		// Ask for phone (optional, can be left blank)
		phone = utils.ReadOptional(reader, "Enter Phone (optional): ")

		// Accept if empty or valid format
		if phone == "" || utils.IsValidPhone(phone) {
			break
		}
		// Otherwise, re-prompt user
		fmt.Println("❌ Invalid phone number.")
	}

	// Show user what they entered
	fmt.Printf("\nYou entered:\n  Name : %s\n  Email: %s\n  Phone: %s\n", name, email, phone)

	// Confirm before saving
	if !utils.AskYesNo(reader, "Save this user? (y/N): ") {
		fmt.Println("❌ Cancelled.") // user chose not to save
		return
	}

	// Create User object and insert into database
	user := models.User{Name: name, Email: email, Phone: phone}
	if err := config.DB.Create(&user).Error; err != nil {
		fmt.Println("\n❌ Failed to create user:", err)
		return
	}

	// Success message with generated ID
	fmt.Println("✅ User created with ID:", user.ID)
}
