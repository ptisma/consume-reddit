package initializers

import (
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func ConnectCache(config *Config) {
	hostURI := fmt.Sprintf("%s:%s", config.CacheHost, config.CachePort)
	Cache = redis.NewClient(&redis.Options{
		Addr:     hostURI,
		Username: config.RedisUser,
		Password: config.RedisPassword, // no password set
		DB:       config.RedisDBName,   // use default DB
	})

}
