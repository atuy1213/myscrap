package domain

import (
	"context"
)

type AuthenticatorInterface interface {
	Signup(ctx *context.Context) (*model.User, error)
	Withdraw(userID string) error
}

type ClipperInterface interface {
	ClipArticle(ctx *context.Context, userID, URL string) (*model.Article, error)
	ClipArticles(userID string, URLs []string) error
}

type MyscrapListerInterface interface {
	ListArticles(userID string) ([]*model.Article, error)
	ListArticlesByGenre(userID, genre string) ([]*model.Article, error)
}
