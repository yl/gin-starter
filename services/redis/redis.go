package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/yangliulnn/gin-starter/configs"
	"github.com/yangliulnn/gin-starter/httpd/utils/log"
)

var Redis *redis.Client

func Setup() {
	config := configs.Redis

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: configs.Redis.Password,
		DB:       configs.Redis.DB,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		log.Log.Error(err)
	}
}
