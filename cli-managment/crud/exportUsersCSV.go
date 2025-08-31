package crud

import (
	"encoding/csv"
	"fmt"
	"os"

	"cli-user-management/config"
	"cli-user-management/models"
)

// ExportUsersCSV retrieves all users from the database and exports them to a CSV file.
func ExportUsersCSV() {
	// Display section title
	fmt.Println("\n== Export Users to CSV ==")

	// Fetch all users from the database
	var users []models.User
	result := config.DB.Find(&users)

	// Check for errors during the database query
	if result.Error != nil {
		fmt.Println("❌ Failed to fetch users:", result.Error)
		return
	}

	// If no users are found, show a warning and exit
	if len(users) == 0 {
		fmt.Println("⚠️ No users to export.")
		return
	}

	// Create a CSV file named "users_export.csv"
	file, err := os.Create("users_export.csv")
	if err != nil {
		fmt.Println("❌ Could not create CSV file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a CSV writer for the file
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all buffered data is written to the file

	// Write CSV header row
	writer.Write([]string{"ID", "Name", "Email", "Phone", "Created At"})

	// Write each user's data as a row in the CSV
	for _, u := range users {
		writer.Write([]string{
			fmt.Sprintf("%d", u.ID), // Convert ID to string
			u.Name,                  // User name
			u.Email,                 // User email
			u.Phone,                 // User phone
			u.CreatedAt.Format("2006-01-02 15:04:05"), // Format timestamp as string
		})
	}

	// Success message
	fmt.Println("✅ Users exported successfully to users_export.csv")
}
