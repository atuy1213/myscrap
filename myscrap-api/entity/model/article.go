package model

import "github.com/atuy1213/myscrap/myscrap-api/adapter/dto"

type Article struct {
	UserID string
	URL    string
	Title  string
	IsRead bool
}

func NewArticle(article *dto.Article) *Article {
	return &Article{
		UserID: article.UserID,
		URL:    article.URL,
		Title:  article.Title,
		IsRead: article.IsRead,
	}
}
