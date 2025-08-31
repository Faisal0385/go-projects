package crud

import (
	"bufio"
	"fmt"
	"strings"

	"cli-user-management/config"
	"cli-user-management/models"
)

// SearchUser allows the user to search for users by name or email keyword.
func SearchUser(reader *bufio.Reader) {
	// Display section title
	fmt.Println("\n== Search User ==")

	// Prompt the user to enter a search keyword
	fmt.Print("Enter name or email keyword: ")

	// Read input from terminal
	keyword, _ := reader.ReadString('\n')

	// Remove leading/trailing spaces and newline characters
	keyword = strings.TrimSpace(keyword)

	// If keyword is empty, show an error and exit
	if keyword == "" {
		fmt.Println("❌ Keyword cannot be empty.")
		return
	}

	// Slice to hold matched users
	var users []models.User

	// Query database: search for users where name OR email contains the keyword
	// The % signs are wildcards for SQL LIKE operator
	result := config.DB.Where("name LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&users)

	// Check for errors during the query
	if result.Error != nil {
		fmt.Println("❌ Error searching users:", result.Error)
		return
	}

	// If no users matched, show a message
	if len(users) == 0 {
		fmt.Println("⚠️ No users found.")
		return
	}

	// Display the search results
	fmt.Println("\n🔎 Search Results:")
	for _, u := range users {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s\n", u.ID, u.Name, u.Email, u.Phone)
	}
}
