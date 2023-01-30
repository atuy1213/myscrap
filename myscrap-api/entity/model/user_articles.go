package model

import (
	"errors"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
)

type UserArticles struct {
	User     *User
	Articles []*Article
}

func NewUserArticles(user *dto.User, articles []*dto.Article) (*UserArticles, error) {

	var tmp []*Article
	for _, v := range articles {
		if user.UserID != v.UserID {
			return nil, errors.New("UserとArticleのUserIDが異なります。")
		}
		tmp = append(tmp, NewArticle(v))
	}

	return &UserArticles{
		User:     NewUser(user),
		Articles: tmp,
	}, nil
}
