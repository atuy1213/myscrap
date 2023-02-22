package repository

import (
	"context"
	"database/sql"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
	"github.com/atuy1213/myscrap/myscrap-api/usecase"
)

type ArticleRepository struct{}

func NewArticleRepository() usecase.ArticleRepositoryInterface {
	return &ArticleRepository{}
}

func (u *ArticleRepository) GetArticles(db *sql.DB, userID string) ([]*dto.Article, error) {
	return nil, nil
}

func (u *ArticleRepository) CreateArticle(ctx *context.Context, db *sql.DB, article *dto.Article) error {
	sql := `INSERT INTO "article" (user_id, url, title, is_read, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sql, article.UserID, article.URL, article.Title, article.IsRead, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
