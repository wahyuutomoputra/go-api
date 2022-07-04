package users

type UserService interface {
	FindAll() ([]Users, error)
}

type userService struct {
	userRepository UserRepository
}

func NewServiceUser(repository UserRepository) *userService {
	return &userService{userRepository: repository}
}

func (s userService) FindAll() ([]Users, error) {
	return s.userRepository.FindAll()
}
