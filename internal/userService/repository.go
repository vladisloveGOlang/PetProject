package userService

import (
	//"errors"
	//"fmt"

	"first/internal/web/users"
	"log"

	//"os/user"

	"gorm.io/gorm"
	//"first/internal/web/messages"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAllMessages() ([]users.User, error)
	CreateNewUser(face users.User) error
	PatchUser(face users.User) (ans users.PatchUsers200JSONResponse, err error)
	DeleteUserById(face users.User) (ans users.DeleteUsers200JSONResponse, err error)
}

func (r *userRepository) PatchUser(face users.User) (ans users.PatchUsers200JSONResponse, err error) {

	result := r.db.Model(&face).Omit("id").Updates(map[string]interface{}{"id": face.Id, "email": face.Email, "password": face.Password})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	status := true

	ansf := users.PatchUsers200JSONResponse{
		Id:      face.Id,
		Changed: &status,
	}

	return ansf, nil
}

func (r *userRepository) DeleteUserById(face users.User) (ans users.DeleteUsers200JSONResponse, err error) {
	result := r.db.Delete(&face)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	status := true

	ansf := users.DeleteUsers200JSONResponse{
		Id:      face.Id,
		Changed: &status,
	}
	return ansf, nil

}

func (r *userRepository) CreateNewUser(face users.User) error {
	result := r.db.Create(&face)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) GetAllMessages() ([]users.User, error) {
	var userList []users.User
	result := r.db.Table("users").Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func CreateUsersRepsitory(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
