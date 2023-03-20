package usecase

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
)

type MyscrapLister struct {
	db                *sql.DB
	articleRepository ArticleRepositoryInterface
}

func (u *MyscrapLister) ListArticles(userID string) ([]*dto.Article, error) {
	articles, err := u.articleRepository.GetArticles(u.db, userID)
	if err != nil {
		log.Println(err)
	}
	return articles, nil
}

func (u *MyscrapLister) ListArticlesByGenre(userID, genre string) ([]*dto.Article, error) {
	fmt.Println("UnregisterForMyscrap")
	return nil, nil
}
