package userService

type UserService struct {
	urepo UserRepository
}

func CreateUserService(reponame UserRepository) *UserService {
	return &UserService{urepo: reponame}
}

func (s *UserService) GetAllMessages(mess usermess) ([]usermess, error) {
	return s.urepo.GetAllMessages(mess)
}
