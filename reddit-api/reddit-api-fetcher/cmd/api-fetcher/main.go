package main

import (
	"context"
	"fmt"
	"log"
	"reddit-api-fetcher/internal/config"
	"reddit-api-fetcher/internal/producer"
	"strconv"
	"time"
)

func main() {
	// Test
	fmt.Println("Hello world")
	// Init a context
	ctx := context.Background()

	// Load the producer's config
	config := config.GetRabbitMQConfig()

	// Init a producer
	producer, err := producer.GetRabbitMQProducer(config)
	if err != nil {
		log.Panic(err)
	}
	for {
		for i := 1; i <= 10; i++ {
			if i%5 != 0 {
				s := strconv.Itoa(i)
				ctxTimeout, _ := context.WithTimeout(ctx, time.Second*5)
				err := producer.StorePost(ctxTimeout, s)
				if err != nil {
					log.Panic(err)
				}
			} else {
				fmt.Println("Sleeping...")
				time.Sleep(60 * time.Second)
			}

		}
	}

}
