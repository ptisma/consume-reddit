package main

import (
	"fmt"
	"log"
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

	ch, err := processor.ReadFromBroker()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for d := range ch {
		log.Printf("Message:%s", string(d.Body))
		fmt.Println()
		post := processor.Process(d.Body, d.Type)
		log.Println("Writing to DB")
		processor.WriteToDB(&post)
	}

	log.Println("Over with receive")

}
