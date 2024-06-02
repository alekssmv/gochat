package validators

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `validate:"required,min=3,max=20"`
	Password string `validate:"required,min=3,max=20"`
}

// Валидация пользователя по структуре User
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}