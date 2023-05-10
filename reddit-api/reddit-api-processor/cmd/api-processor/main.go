package main

import (
	"log"
	"reddit-api-processor/internal/config"
	"reddit-api-processor/internal/processor"
)

func main() {

	consumerConfig := config.GetRabbitMQConfig()

	log.Println("consumerConfig", consumerConfig)

	dbConfig, err := config.GetPostgresConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("dbConfig", dbConfig)

	processor, err := processor.GetProcessor(dbConfig, consumerConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = processor.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	ch, err := processor.ReadFromBroker()
	if err != nil {
		log.Fatal(err)
	}

	for d := range ch {
		log.Printf("Message:%s\n", string(d.Body))
		post, err := processor.Process(d.Body, d.Type)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Post:", post)
		err = processor.WriteToDB(&post)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Over with receive")

}
