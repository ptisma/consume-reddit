package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var rdb *redis.client
var REDIS_URL = os.Getenv("REDIS_URL")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")

func connectToRedis(ctx context.Context, REDIS_URL, REDIS_PASSWORD string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: REDIS_PASSWORD, // no password set
		DB:       0,              // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err == nil {
		log.Println(pong, err)
	} else {
		log.Panic(err)
	}

}

func writeToRedis(ctx context.Context, key, val string) error {
	err := rdb.Set(ctx, key, val, 0).Err()
	return err

}

func readFromRedis(ctx context.Context) (string, error) {
	val, err := rdb.Get(ctx, "foo").Result()

	return fmt.Sprintf("%v", val), err

}

func main() {

	ctx := context.Background()
	connectToRedis(ctx, REDIS_URL, REDIS_PASSWORD)

	writeToRedis(ctx, "foo", "bar")

	val, err := readFromRedis(ctx)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Value fetched from Redis is:", val)

	for {
		time.Sleep(5 * time.Second)
	}
}
