package model

import (
	"time"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
)

type Article struct {
	UserID    string
	URL       string
	Title     string
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(article *dto.Article) *Article {
	return &Article{
		UserID:    article.UserID,
		URL:       article.URL,
		Title:     article.Title,
		IsRead:    article.IsRead,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}
