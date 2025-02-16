package db

import (
	"fmt"
	"log"

	"merch-shop/internal/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(cfg config.RedisConfig) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})

	log.Println("Connected to Redis")
}
