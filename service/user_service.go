package service

import (
	"errors"
	"rmzstartup/helper"
	model "rmzstartup/model/entity"
	"rmzstartup/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input helper.RegisterUserInput) (model.User, error)
	Login(input helper.LoginInputUser) (model.User, error)
	CheckEmailAvalaible(input helper.CheckEmailInput) (bool, error)
	SaveAvatar(ID, fileLocation string) (model.User, error)
	GetUserByID(ID string) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func (s *userService) RegisterUser(input helper.RegisterUserInput) (model.User, error) {
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

func (s *userService) Login(input helper.LoginInputUser) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, errors.New("no user found on that email")
	}
	if user.Id == uuid.Nil {
		return user, errors.New("user not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) CheckEmailAvalaible(input helper.CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.Id == uuid.Nil {
		return true, nil
	}
	return false, nil
}

func (s *userService) SaveAvatar(ID, fileLocation string) (model.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation
	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *userService) GetUserByID(ID string) (model.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.Id == uuid.Nil {
		return user, errors.New("no user found on that ID")
	}
	return user, nil
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}
