package repository

import (
	"auth-service/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByAuth0ID(auth0Id string) (*entity.User, error)
}
