package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cli-user-management/config"
	"cli-user-management/crud"
	"cli-user-management/models"
)

func main() {

	// Step 1: Connect to the database
	config.ConnectDB()

	// Step 2: AutoMigrate ensures that the User table is created if not exists
	config.DB.AutoMigrate(&models.User{})

	// Step 3: Reader for CLI input
	reader := bufio.NewReader(os.Stdin)

	// Step 4: Start menu loop
	for {

		// Step 5a: Printing Menu
		fmt.Println("\n== CLI User Management Menu ==")
		fmt.Println("1. Create User")
		fmt.Println("2. Read User by ID")
		fmt.Println("3. List All Users")
		fmt.Println("4. Update User")
		fmt.Println("5. Delete User")
		fmt.Println("6. Search User by Name/Email")
		fmt.Println("7. Export Users to CSV")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		// Step 5b: Read user choice
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// Step 6: Perform action based on choice
		switch choice {
		case "1":
			crud.CreateUser(reader) // Call Create user function
		case "2":
			crud.ReadUser(reader) // Call Read user function
		case "3":
			crud.ListUsers() // Call List users function
		case "4":
			crud.UpdateUser(reader) // Call Update user function
		case "5":
			crud.DeleteUser(reader) // Call Delete user function
		case "6":
			crud.SearchUser(reader) // Call Search user function
		case "7":
			crud.ExportUsersCSV() // Call Export to CSV function
		case "0":
			fmt.Println("Goodbye!") // Exit message
			return
		default:
			fmt.Println("❌ Invalid option. Try again.") // Invalid input
		}
	}
}
