package util

import (
	"log"

	"github.com/atuy1213/myscrap/myscrap-api/usecase"
	"github.com/google/uuid"
)

type GoogleUUID struct{}

func NewGoogleUUID() usecase.UUIDUtilInterface {
	return &GoogleUUID{}
}

func (u *GoogleUUID) New() string {
	ID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	return ID.String()
}
