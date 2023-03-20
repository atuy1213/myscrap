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

	var input *input.ClipArticleInput
	bt, _ := io.ReadAll(r.Body)
	fmt.Println(string(bt))
	fmt.Println()
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(input); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(input)

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
