package userService

import (
	//"errors"
	//"fmt"

	"gorm.io/gorm"
	//"first/internal/web/messages"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAllMessages(mess usermess) ([]usermess, error)
}

func CreateUsersRepsitory(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllMessages(mess usermess) ([]usermess, error) {
	var userList []usermess
	result := r.db.Table("users").Find(&userList, &mess)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}
