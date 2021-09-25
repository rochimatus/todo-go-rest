package service

import "todo-go-rest/repository"

type Service struct {
	AuthService AuthService
	JWTService  JWTService
}

func CreateService(repo *repository.Repository) *Service {

	return &Service{
		AuthService: NewAuthService(repo.UserRepository),
		JWTService:  NewJWTService(),
	}
}
