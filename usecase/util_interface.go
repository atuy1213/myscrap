package usecase

import "time"

type UUIDUtilInterface interface {
	New() string
}

type TimeUtilInterface interface {
	NowUTC() *time.Time
}
