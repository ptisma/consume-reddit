package config

import (
	"os"
	"strconv"
)

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

const DB_HOST_URI = "DB_HOST_URI"
const DB_USER_NAME = "DB_USER_NAME"
const DB_USER_PASSWORD = "DB_USER_PASSWORD"
const DB_NAME = "DB_NAME"
const DB_PORT = "DB_PORT"

type PostgresConfig struct {
	DBHostURI      string
	DBUsername     string
	DBUserPassword string
	DBName         string
	DBPort         int
}

func GetPostgresConfig() (*PostgresConfig, error) {
	var err error
	dbHostUri := os.Getenv(DB_HOST_URI)
	dbUserName := os.Getenv(DB_USER_NAME)
	dbUserPassword := os.Getenv(DB_USER_PASSWORD)
	dbName := os.Getenv(DB_NAME)
	dbPortStr := os.Getenv(DB_PORT)
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, err
	}

	return &PostgresConfig{DBHostURI: dbHostUri, DBUsername: dbUserName, DBUserPassword: dbUserPassword, DBName: dbName, DBPort: dbPort}, err

}
