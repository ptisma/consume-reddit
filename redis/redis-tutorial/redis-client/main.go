package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("hello world")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "kek",
		Password: "kek", // no password set
		DB:       0,     // use default DB
	})

	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("foo", val)

}
