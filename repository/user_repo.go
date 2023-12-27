package repository

import (
	model "rmzstartup/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(ID string) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(ID string) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Save(user model.User) (model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
