package service

import (
	"errors"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(reg request.RegisterRequest) (model.User, error)
	Login(email string, password string) (model.User, error)
}

type authService struct {
	// email      string
	// password   string
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *authService {
	return &authService{
		repository: repository,
	}
}

func (service *authService) Register(reg request.RegisterRequest) (model.User, error) {
	if reg.Password != reg.ConfirmPassword {
		return model.User{}, errors.New("Password and Confirmation Password are not match")
	}

	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(reg.Password), 14)
	reg_user := model.User{
		FullName: reg.FullName,
		Email:    reg.Email,
		Password: string(hashed_password),
		Role:     reg.Role,
	}

	user, err := service.repository.Create(reg_user)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (service *authService) Login(email string, password string) (model.User, error) {
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return user, err
}
