package usecase

import (
	"context"
	"net/http"
)

type AuthenticationControllerInterface interface {
	Signup(w http.ResponseWriter, r *http.Request)
}

type ListControllerInterface interface {
	ListArticles(ctx *context.Context)
}

type ClipControllerInteface interface {
	ClipArticle(w http.ResponseWriter, r *http.Request)
}
