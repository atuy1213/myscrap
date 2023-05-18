package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atuy1213/myscrap/myscrap-api/entity/domain"
)

type AuthenticationController struct {
	authenticator domain.AuthenticatorInterface
}

func NewAuthenticationController(
	authenticator domain.AuthenticatorInterface,
) *AuthenticationController {
	return &AuthenticationController{
		authenticator: authenticator,
	}
}

func (u *AuthenticationController) Signup(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	user, err := u.authenticator.Signup(&ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusCreated)
}
