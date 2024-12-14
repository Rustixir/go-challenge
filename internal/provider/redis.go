package provider

import (
	"sync"

	"github.com/Rustixir/go-challenge/pkg/config"
	"github.com/redis/go-redis/v9"
)

var redisOnce sync.Once
var redisClient *redis.Client

func GetRedis() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.Config.DB.Redis.Addr,
		})
	})
	return redisClient
}
