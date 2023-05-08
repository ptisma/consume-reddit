package main

import (
	"context"
	"fmt"
	"reddit-api-fetcher/internal/config"
	"reddit-api-fetcher/internal/fetcher"
	"reddit-api-fetcher/internal/producer"
)

func main() {
	// Test
	fmt.Println("Hello world")
	// Init a context
	ctx := context.Background()
	fmt.Println("ctx", ctx)

	// Load the producer's config
	producerConfig := config.GetRabbitMQConfig()

	fmt.Println("producerConfig", producerConfig)

	// Load the fetcher's config
	fetcherConfig, err := config.GetSubredditFetcherConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fetcherConfig", fetcherConfig)

	// Init a producer
	producer, err := producer.GetRabbitMQProducer(producerConfig)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("producer", producer)

	// Init a fetcher
	fetcher, err := fetcher.GetSubredditFetcher(fetcherConfig, producer)
	if err != nil {
		fmt.Println(err)
	}

	// for {
	// 	for i := 1; i <= 10; i++ {
	// 		if i%5 != 0 {
	// 			s := strconv.Itoa(i)
	// 			ctxTimeout, _ := context.WithTimeout(ctx, time.Second*5)
	// 			err := producer.StorePost(ctxTimeout, s)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 		} else {
	// 			fmt.Println("Sleeping...")
	// 			time.Sleep(60 * time.Second)
	// 		}

	// 	}
	// }

	token, err := fetcher.FetchToken()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token", token)

	err = fetcher.FetchPosts(ctx, token)
	if err != nil {
		fmt.Println(err)
	}

}
