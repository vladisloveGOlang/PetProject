package userService

import (
	"gorm.io/gorm"
)

// описание данных пользователя
type usermess struct {
	gorm.Model
	email    string
	password string
}
