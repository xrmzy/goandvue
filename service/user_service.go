package service

import (
	model "rmzstartup/model/entity"
	"rmzstartup/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type UserService interface {
	RegisterUser(input RegisterUserInput) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func (s *userService) RegisterUser(input RegisterUserInput) (model.User, error) {
	user := model.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}
