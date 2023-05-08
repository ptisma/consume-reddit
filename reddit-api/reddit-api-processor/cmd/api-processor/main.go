package main

import (
	"fmt"
	"os"
	"reddit-api-processor/internal/config"
	"reddit-api-processor/internal/processor"
)

func main() {

	// Test
	fmt.Println("Hello world")

	consumerConfig := config.GetRabbitMQConfig()

	fmt.Println("consumerConfig", consumerConfig)

	dbConfig, err := config.GetPostgresConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dbConfig", dbConfig)

	processor, err := processor.GetProcessor(dbConfig, consumerConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = processor.AutoMigrate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	processor.WriteToDB()

}
