package model

import "github.com/atuy1213/myscrap/myscrap-api/adapter/dto"

type User struct {
	UserID          string
	AccessKeyID     string
	SecretAccessKey string
}

func NewUser(user *dto.User) *User {
	return &User{
		UserID:          user.UserID,
		AccessKeyID:     user.AccessKeyID,
		SecretAccessKey: user.SecretAccessKey,
	}
}
