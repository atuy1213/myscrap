package dto

import "time"

type Article struct {
	UserID    string
	URL       string
	Title     string
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(userID, URL, title string, isRead bool, createdAt, updatedAt time.Time) *Article {
	return &Article{
		UserID:    userID,
		URL:       URL,
		Title:     title,
		IsRead:    isRead,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
