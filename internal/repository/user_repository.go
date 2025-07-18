package repository

import (
	"errors"
	"to-do-list/database"
	"to-do-list/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}
