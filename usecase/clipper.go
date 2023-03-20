package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
	"github.com/atuy1213/myscrap/myscrap-api/entity/domain"
	"github.com/atuy1213/myscrap/myscrap-api/entity/model"
)

type Clipper struct {
	db                *sql.DB
	articleRepository ArticleRepositoryInterface
	timeUtil          TimeUtilInterface
}

func NewClipper(
	db *sql.DB,
	articleRepository ArticleRepositoryInterface,
	timeUtil TimeUtilInterface,
) domain.ClipperInterface {
	return &Clipper{
		db:                db,
		articleRepository: articleRepository,
		timeUtil:          timeUtil,
	}
}

func (u *Clipper) ClipArticle(ctx *context.Context, userID, URL string) (*model.Article, error) {
	now := u.timeUtil.NowUTC()
	article := dto.NewArticle(userID, URL, "Title", false, *now, *now)
	if err := u.articleRepository.CreateArticle(ctx, u.db, article); err != nil {
		return nil, err
	}
	return model.NewArticle(article), nil
}

func (u *Clipper) ClipArticles(userID string, URLs []string) error {
	fmt.Println("ClipArticles")
	return nil
}
