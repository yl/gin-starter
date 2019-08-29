package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/yangliulnn/gin-starter/configs"
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
		log.Fatalln(err)
	}
}
