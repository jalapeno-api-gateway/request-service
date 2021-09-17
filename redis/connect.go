package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitializeRedisClient() {
	RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    os.Getenv("SENTINEL_MASTER"),
		SentinelAddrs: []string{os.Getenv("SENTINEL_ADDRESS")},
		Password:      os.Getenv("REDIS_PASSWORD"),
		DB:            0,
	})
}
