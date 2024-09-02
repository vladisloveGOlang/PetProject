package userService

import (
	"gorm.io/gorm"
)

// описание данных пользователя
type userData struct {
	gorm.Model
	email    string
	password string
}
