package redis

import (
	// "os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func InitializeRedisClient() {
	sentinelMaster := "mymaster"
	sentinelAddress := "10.20.1.24:5000"
	redisPassword := "a-very-complex-password-here"
	// sentinelMaster := os.Getenv("SENTINEL_MASTER")
	// sentinelAddress := os.Getenv("SENTINEL_ADDRESS")
	// redisPassword := os.Getenv("REDIS_PASSWORD")

	logrus.WithFields(logrus.Fields{"sentinelMaster": sentinelMaster, "sentinelAddress": sentinelAddress}).Debug("Initializing Redis client.")

	RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    sentinelMaster,
		SentinelAddrs: []string{sentinelAddress},
		Password:      redisPassword,
		DB:            0,
	})
}
