package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time
	Category  string
}

func main() {
	fmt.Println("hello world")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "default",
		Password: "redis", // no password set
		DB:       0,       // use default DB
	})

	ctx := context.Background()
	res := client.Ping(ctx)
	fmt.Println("res", res)
	post := Post{1, "Test", "hello world!", time.Now(), "Hot"}
	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Set(ctx, "Hot:2023-01-01:2023-01-02", b, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "Hot:2023-01-01:2023-01-02").Result()
	if err != nil {
		panic(err)
	}

	posts := []Post{Post{2, "Test2", "hello world2!", time.Now(), "Hot"}}
	b, err = json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Set(ctx, "Hot:2023-02-02:2023-02-03", b, 0).Err()
	if err != nil {
		panic(err)
	}
	val, err = client.Get(ctx, "Hot:2023-02-02:2023-02-03").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	fmt.Println(reflect.TypeOf(val))
	var postsNew []Post
	err = json.Unmarshal([]byte(val), &postsNew)
	fmt.Println(postsNew)

}
