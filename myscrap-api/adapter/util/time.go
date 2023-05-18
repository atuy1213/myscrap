package util

import "time"

type TimeUtil struct{}

func NewTimeUtil() *TimeUtil {
	return &TimeUtil{}
}

func (u *TimeUtil) NowUTC() *time.Time {
	nowUTC := time.Now().In(u.locationUTC())
	return &nowUTC
}

func (u *TimeUtil) locationUTC() *time.Location {
	UTC, _ := time.LoadLocation(time.UTC.String())
	return UTC
}
