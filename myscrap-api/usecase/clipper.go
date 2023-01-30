package usecase

import "fmt"

type MyscrapClipper struct{}

func (u *MyscrapClipper) ClipArticle(userID, URL string) error {
	fmt.Println("ClipArticle")
	return nil
}

func (u *MyscrapClipper) ClipArticles(userID string, URLs []string) error {
	fmt.Println("ClipArticles")
	return nil
}
