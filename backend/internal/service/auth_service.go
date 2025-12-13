package service

import (
	"errors"

	"github.com/MananLedwani/taskmanager/backend/internal/models"
	"github.com/MananLedwani/taskmanager/backend/internal/repository"
	"github.com/MananLedwani/taskmanager/backend/internal/utils"
)

type AuthService interface {
	Signup(user *models.User) error
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Signup(user *models.User) error {

	existingUser, _ := s.userRepo.FindByEmail(user.Email)

	if existingUser != nil{
		return errors.New("user with the specified email already exist")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := s.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(email, password string) (string, error) {

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
