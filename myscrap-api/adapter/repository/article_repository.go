package repository

import (
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
