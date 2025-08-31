package models

import "gorm.io/gorm"

// User represents the users table in the database
type User struct {
	// gorm.Model adds default fields:
	// ID (primary key), CreatedAt, UpdatedAt, DeletedAt (for soft delete)
	gorm.Model

	// Name field (VARCHAR(100), cannot be NULL)
	Name string `gorm:"size:100;not null"`

	// Email field (VARCHAR(100), must be unique, cannot be NULL)
	Email string `gorm:"size:100;uniqueIndex;not null"`

	// Phone field (VARCHAR(20), optional → can be empty)
	Phone string `gorm:"size:20"`
}
