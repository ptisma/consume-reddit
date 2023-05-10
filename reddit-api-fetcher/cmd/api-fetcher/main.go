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

	// Acquire a token
	token, err := fetcher.FetchToken()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token", token)

	// Fetch and store posts
	err = fetcher.FetchAndStorePosts(ctx, token)
	if err != nil {
		fmt.Println(err)
	}

}
