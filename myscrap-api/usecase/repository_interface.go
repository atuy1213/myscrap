package usecase

import (
	"context"
	"database/sql"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
)

type UserRepositoryInterface interface {
	CreateUser(ctx *context.Context, db *sql.DB, user *dto.User) error
}

type ArticleRepositoryInterface interface {
	GetArticles(db *sql.DB, userID string) ([]*dto.Article, error)
}
