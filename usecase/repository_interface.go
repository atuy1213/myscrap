package usecase

import (
	"context"
	"database/sql"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
)

type UserRepositoryInterface interface {
	// create
	CreateUser(ctx *context.Context, db *sql.DB, user *dto.User) error
}

type ArticleRepositoryInterface interface {
	// read
	GetArticles(db *sql.DB, userID string) ([]*dto.Article, error)

	// create
	CreateArticle(ctx *context.Context, db *sql.DB, article *dto.Article) error
}
