package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/atuy1213/myscrap/myscrap-api/adapter/dto"
	"github.com/atuy1213/myscrap/myscrap-api/entity/domain"
	"github.com/atuy1213/myscrap/myscrap-api/entity/model"
)

type MyscrapAuthenticator struct {
	db             *sql.DB
	userRepository UserRepositoryInterface
	UUIDUtil       UUIDUtilInterface
	timeUtil       TimeUtilInterface
}

func NewMyscrapAuthenticator(
	db *sql.DB,
	userRepository UserRepositoryInterface,
	UUIDUtil UUIDUtilInterface,
	tiemUtil TimeUtilInterface,
) domain.AuthenticatorInterface {
	return &MyscrapAuthenticator{
		db:             db,
		userRepository: userRepository,
		UUIDUtil:       UUIDUtil,
		timeUtil:       tiemUtil,
	}
}

func (u *MyscrapAuthenticator) Signup(ctx *context.Context) (*model.User, error) {

	userID := u.UUIDUtil.New()
	accessKey := u.UUIDUtil.New()
	secretAccessKey := u.UUIDUtil.New()
	now := u.timeUtil.NowUTC()
	user := dto.NewUser(userID, accessKey, secretAccessKey, *now, *now)

	if err := u.userRepository.CreateUser(ctx, u.db, user); err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (u *MyscrapAuthenticator) Withdraw(userID string) error {
	fmt.Println("UnregisterForMyscrap")
	return nil
}
