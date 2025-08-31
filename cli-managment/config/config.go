package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

// Declare a global variable for the database connection
var DB *gorm.DB

// ConnectDB establishes a connection to the MySQL database
func ConnectDB() {
	// Try to load the .env file (environment variables)
	err := godotenv.Load()
	if err != nil {
		// If .env is not found, just show a warning but continue
		log.Println("⚠️ Warning: .env file not found")
	}

	// Fetch database credentials & config values from .env
	host := os.Getenv("DB_HOST") // Database host (e.g., localhost)
	port := os.Getenv("DB_PORT") // Database port (e.g., 3306)
	user := os.Getenv("DB_USER") // Database username
	pass := os.Getenv("DB_PASS") // Database password
	name := os.Getenv("DB_NAME") // Database name

	// Build DSN (Data Source Name) for MySQL connection
	// Example format: user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// Open connection to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// If connection fails → stop program with error message
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Assign the database connection to the global variable
	DB = db

	// Success message
	fmt.Println("✅ Database connected successfully!")
}
