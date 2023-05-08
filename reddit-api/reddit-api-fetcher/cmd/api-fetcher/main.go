package main

import (
	"context"
	"fmt"
	"log"
	"reddit-api-fetcher/internal/config"
	"reddit-api-fetcher/internal/fetcher"
	"reddit-api-fetcher/internal/producer"
)

func main() {
	// Init a context
	ctx := context.Background()
	log.Println("ctx", ctx)

	// Load the producer's config
	producerConfig := config.GetRabbitMQConfig()

	log.Println("producerConfig", producerConfig)

	// Load the fetcher's config
	fetcherConfig, err := config.GetSubredditFetcherConfig()
	if err != nil {
		log.Panicf(err)
	}
	log.Println("fetcherConfig", fetcherConfig)

	// Init a producer
	producer, err := producer.GetRabbitMQProducer(producerConfig)
	if err != nil {
		log.Panicf(err)
	}
	log.Println("producer", producer)

	// Init a fetcher
	fetcher, err := fetcher.GetSubredditFetcher(fetcherConfig, producer)
	if err != nil {
		log.Panicf(err)
	}

	// Acquire a token for Reddit API
	token, err := fetcher.FetchToken()
	if err != nil {
		log.Panicf(err)
	}
	log.Println("token", token)

	// Fetch and store posts
	err = fetcher.FetchAndStorePosts(ctx, token)
	if err != nil {
		fmt.Println(err)
	}

}
