package initializers

import (
	"fmt"
	"log"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func ConnectCache(config *Config) {
	hostURI := fmt.Sprintf("%s:%s", config.CacheHost, config.CachePort)
	db, err := strconv.Atoi(config.CacheDBName)
	if err != nil {
		log.Fatal("Failed to connect to the Cache")
	}
	Cache = redis.NewClient(&redis.Options{
		Addr:     hostURI,
		Username: config.CacheUserName,
		Password: config.CacheUserPassword, // no password set
		DB:       db,                       // use default DB
	})
	fmt.Println("🚀 Connected Successfully to the Cache")

}
