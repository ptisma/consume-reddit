package initializers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func ConnectCache(config *Config) (err error) {
	hostURI := fmt.Sprintf("%s:%s", config.CacheHost, config.CachePort)
	db, err := strconv.Atoi(config.CacheDBName)
	if err != nil {
		return err
	}
	Cache = redis.NewClient(&redis.Options{
		Addr:     hostURI,
		Username: config.CacheUserName,
		Password: config.CacheUserPassword, // no password set
		DB:       db,                       // use default DB
	})
	ctx := context.Background()
	res := Cache.Ping(ctx)
	log.Println("res:", res)
	return err
}
