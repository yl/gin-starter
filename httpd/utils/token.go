package utils

import (
	"github.com/google/uuid"
	"go-trading/services/redis"
	"time"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) Generate(id int) (string, error) {
	token := uuid.New().String()
	err := redis.Redis.Set("token:"+token, id, 24*time.Hour).Err()
	return token, err
}
