package config

import "os"

const BROKER_URI = "BROKER_URI"
const TOPIC_NAME = "TOPIC_NAME"
const ROUTING_KEY = "ROUTING_KEY"

type RabbitMQConfig struct {
	BrokerURI  string
	TopicName  string
	RoutingKey string
}

func GetRabbitMQConfig() *RabbitMQConfig {

	brokerURI := os.Getenv(BROKER_URI)
	topicName := os.Getenv(TOPIC_NAME)
	routingKey := os.Getenv(ROUTING_KEY)

	return &RabbitMQConfig{BrokerURI: brokerURI, TopicName: topicName, RoutingKey: routingKey}

}

type PostgresConfig struct {
	BrokerURI  string
	TopicName  string
	RoutingKey string
}
