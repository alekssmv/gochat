package models

import (
	"gorm.io/gorm"
)

// User represents a user in the database
type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null" json:"username"`
	Password string `gorm:"size:100;not null" json:"password"`
}
