package database

import (
	"auth-service/internal/entity"
	"auth-service/internal/repository"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entity.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users(auth0_id, email, name) VALUES($1, $2, $3)",
		user.Auth0ID, user.Email, user.Name,
	)

	return err
}

func (r *UserRepo) GetByAuth0ID(auth0ID string) (*entity.User, error) {
	row := r.db.QueryRow(
		"SELECT id, auth0_id, email, name FROM users WHERE auth0_id = $1", auth0ID,
	)

	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Auth0ID, &user.Email, &user.Name)

	return user, err
}
