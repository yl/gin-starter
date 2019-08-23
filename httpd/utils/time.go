package utils

import (
	"fmt"
	"go-trading/configs"
	"time"
)

type Time struct{}

func NewTime() *Time {
	return &Time{}
}

func (T *Time) Format(t time.Time) string {
	config := configs.App
	fmt.Println(config)
	l, _ := time.LoadLocation(config.Timezone)
	return t.In(l).Format(config.TimeFormat)
}
