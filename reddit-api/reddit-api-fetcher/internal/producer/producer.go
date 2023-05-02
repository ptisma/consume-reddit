package producer

import (
	"context"
	"fmt"
	"reddit-api-fetcher/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

// type IProducer interface {
// 	storeMessage()
// }

// type IBasicProducer interface {
// 	SetBrokeriURI(brokerURI string)
// 	GetBrokerURI() string
// 	SetTopicName(topicName string)
// 	GetTopicName() string
// }
// type BasicProducer struct {
// 	BrokerURI string
// 	TopicName string
// }

// func (b *BasicProducer) SetBrokeriURI(brokerURI string) {
// 	b.BrokerURI = brokerURI
// }

// func (b *BasicProducer) GetBrokerURI() string {
// 	return b.BrokerURI
// }

// func (b *BasicProducer) SetTopicName(topicName string) {
// 	b.TopicName = topicName
// }

// func (b *BasicProducer) GetTopicName() string {
// 	return b.TopicName
// }

type RabbitMQProducer struct {
	// IBasicProducer
	Exchange   *amqp.Channel
	TopicName  string
	RoutingKey string
}

func (r *RabbitMQProducer) StorePost(ctx context.Context, body string) error {
	err := r.Exchange.PublishWithContext(ctx,
		r.TopicName,  // exchange
		r.RoutingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err

}

func GetRabbitMQProducer(config *config.RabbitMQConfig) (*RabbitMQProducer, error) {

	fmt.Println(config.BrokerURI)
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
