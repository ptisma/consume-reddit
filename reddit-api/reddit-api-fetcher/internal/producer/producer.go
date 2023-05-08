package producer

import (
	"context"
	"reddit-api-fetcher/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	Exchange   *amqp.Channel
	TopicName  string
	RoutingKey string
}

func (r *RabbitMQProducer) StorePost(ctx context.Context, body []byte, category string) error {
	err := r.Exchange.PublishWithContext(ctx,
		r.TopicName,  // exchange
		r.RoutingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Type:        category,
		})

	return err

}

func GetRabbitMQProducer(config *config.RabbitMQConfig) (*RabbitMQProducer, error) {
	conn, err := amqp.Dial(config.BrokerURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(
		config.TopicName, // name
		"topic",          // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQProducer{Exchange: ch, TopicName: config.TopicName, RoutingKey: config.RoutingKey}, nil

}
