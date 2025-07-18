package services

import (
	"errors"
	"to-do-list/internal/repository"
	"to-do-list/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}

type userSvc struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userSvc{repo: r}
}

func (s *userSvc) Register(user *model.User) error {
	existingUser, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("email sudah digunakan")
	}

	user.ID = uuid.New()

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	return s.repo.Create(user)
}

func (s *userSvc) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("email tidak ditemukan")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("password salah")
	}

	return user, nil
}
