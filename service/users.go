package service

import (
	"goLang-userManage79/model"
	"goLang-userManage79/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userRequest model.UsersRequest) (model.Users, error)
	GetUsers(offset int, limit int) ([]model.Users, error, int64)
	GetUser(id int) (model.Users, error)
	UpdateUser(ID int, userRequest model.UsersRequest) (model.Users, error)
	DeleteUser(ID int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUser(userRequest model.UsersRequest) (model.Users, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	user := model.Users{
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Password: password,
	}
	newUser, err := s.userRepository.CreateUser(user)
	return newUser, err
}

func (s *userService) GetUsers(offset int, limit int) ([]model.Users, error, int64) {
	return s.userRepository.GetUsers(offset, limit)
}

func (s *userService) GetUser(ID int) (model.Users, error) {
	return s.userRepository.GetUser(ID)
}

func (s *userService) UpdateUser(ID int, userRequest model.UsersRequest) (model.Users, error) {
	user, err := s.userRepository.GetUser(ID)
	if err != nil {
		return user, err
	}

	user.Name = userRequest.Name
	user.Username = userRequest.Username
	password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	user.Password = password

	newUser, err := s.userRepository.UpdateUser(user)
	return newUser, err
}

func (s *userService) DeleteUser(ID int) error {
	return s.userRepository.DeleteUser(ID)
}
