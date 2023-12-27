package repository

import (
	model "rmzstartup/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) Save(user model.User) (model.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
