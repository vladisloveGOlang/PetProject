package userService

import (
	"first/internal/web/users"
)

type UserService struct {
	urepo UserRepository
}

func CreateUserService(reponame UserRepository) *UserService {
	return &UserService{urepo: reponame}
}

func (s *UserService) GetAllMessages() ([]users.User, error) {
	return s.urepo.GetAllMessages()
}

func (s *UserService) CreateNewUser(face users.User) error {
	return s.urepo.CreateNewUser(face)
}

func (s *UserService) PatchUser(face users.User) (ans users.PatchUsers200JSONResponse, err error) {
	return s.urepo.PatchUser(face)
}

func (s *UserService) DeleteUserById(face users.User) (ans users.DeleteUsers200JSONResponse, err error) {
	return s.urepo.DeleteUserById(face)
}
