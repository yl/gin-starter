package utils

import (
	"time"

	"github.com/yangliulnn/gin-starter/configs"
)

type Time struct{}

func NewTime() *Time {
	return &Time{}
}

func (T *Time) Format(t time.Time) string {
	config := configs.App
	l, _ := time.LoadLocation(config.Timezone)
	return t.In(l).Format(config.TimeFormat)
}
