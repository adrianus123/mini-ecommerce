package usecase

import (
	"auth-service/internal/entity"
	"auth-service/internal/infrastructure/auth0"
	"auth-service/internal/repository"
)

type AuthUsecase struct {
	auth0 *auth0.Auth0Client
	repo  repository.UserRepository
}

func NewAuthUsecase(auth0Client *auth0.Auth0Client, userRepo repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		auth0: auth0Client,
		repo:  userRepo,
	}
}

func (u *AuthUsecase) Register(email, password, name string) error {
	auth0ID, err := u.auth0.Register(email, password)
	if err != nil {
		return err
	}

	return u.repo.Create(&entity.User{
		Auth0ID: auth0ID,
		Email:   email,
		Name:    name,
	})
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	return u.auth0.Login(email, password)
}
