package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/input"
	"github.com/atuy1213/myscrap/myscrap-api/entity/domain"
	"github.com/atuy1213/myscrap/myscrap-api/usecase"
)

type ClipController struct {
	clipper domain.ClipperInterface
}

func NewClipController(
	clipper domain.ClipperInterface,
) usecase.ClipControllerInteface {
	return &ClipController{
		clipper: clipper,
	}
}

func (u *ClipController) ClipArticle(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var input input.ClipArticleInput
	dec := json.NewDecoder(r.Body)
	bytes, _ := io.ReadAll(r.Body)
	fmt.Println(string(bytes))
	if err := dec.Decode(&input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article, err := u.clipper.ClipArticle(&ctx, "", "")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(article)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusCreated)
}
