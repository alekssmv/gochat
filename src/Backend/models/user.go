package models

import (
	"github.com/go-gorm/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) TableName() string {
	return "users"
}

