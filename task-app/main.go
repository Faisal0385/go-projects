package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Task model
type Task struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

var db *gorm.DB

func main() {
	// Connect to MySQL
	dsn := "root:your_mysql_password@tcp(127.0.0.1:3306)/taskdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate table
	db.AutoMigrate(&Task{})

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Serve the form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Handle form submission
	http.HandleFunc("/add-task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		taskName := r.FormValue("task")
		fmt.Println("New Task:", taskName)

		// Save to database
		task := Task{Name: taskName}
		result := db.Create(&task)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Task saved successfully: %s", taskName)
	})

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
