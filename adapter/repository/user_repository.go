package repository

import (
	"context"
	"database/sql"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
	"github.com/atuy1213/myscrap/myscrap-api/usecase"
)

type UserRepository struct{}

func NewUserRepository() usecase.UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(ctx *context.Context, db *sql.DB, user *dto.User) error {
	sql := `INSERT INTO "user" (user_id, access_key_id, secret_access_key, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sql, user.UserID, user.AccessKeyID, user.SecretAccessKey, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
