package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Passive, assume it already exists, its on k8 cluster CRDS
	err = ch.ExchangeDeclarePassive(
		"reddit-exchange", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// err = ch.ExchangeDeclare(
	// 	"logs_topic", // name
	// 	"topic",      // type
	// 	true,         // durable
	// 	false,        // auto-deleted
	// 	false,        // internal
	// 	false,        // no-wait
	// 	nil,          // arguments
	// )
	failOnError(err, "Failed to declare an exchange")

	// q, err := ch.QueueDeclarePassive(
	// 	"logs_queue", // name
	// 	true,         // durable
	// 	false,        // delete when unused
	// 	false,        // exclusive
	// 	false,        // no-wait
	// 	nil,          // arguments
	// )

	// Create queue on producer side just to see if they are going to be saved for consumer
	// It works
	// q, err := ch.QueueDeclare(
	// 	"logs_queue", // name
	// 	true,         // durable
	// 	false,        // delete when unused
	// 	false,        // exclusive
	// 	false,        // no-wait
	// 	nil,          // arguments
	// )
	failOnError(err, "Failed to declare an queue")
	// err = ch.QueueBind(
	// 	q.Name,                // queue name
	// 	severityFrom(os.Args), // routing key
	// 	"logs_topic",          // exchange
	// 	false,                 //noWait
	// 	nil)
	// failOnError(err, "Failed to bind a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"reddit-exchange",          // exchange
		severityFrom(os.Args), // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
