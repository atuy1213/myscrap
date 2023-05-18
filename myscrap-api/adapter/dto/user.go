package dto

import "time"

type User struct {
	UserID          string
	AccessKeyID     string
	SecretAccessKey string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewUser(userID, accessKeyID, secretAccessKey string, createdAt, updatedAt time.Time) *User {
	return &User{
		UserID:          userID,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
