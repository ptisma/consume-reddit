package config

import (
	"fmt"
	"os"
	"strconv"
)

const BROKER_HOST = "BROKER_HOST"
const EXCHANGE_NAME = "EXCHANGE_NAME"
const QUEUE_NAME = "QUEUE_NAME"
const ROUTING_KEY = "ROUTING_KEY"
const BROKER_USERNAME = "BROKER_USERNAME"
const BROKER_PASSWORD = "BROKER_PASSWORD"
const BROKER_PORT = "BROKER_PORT"
const AUTO_CREATE_RABBITMQ = "AUTO_CREATE_RABBITMQ"

type RabbitMQConfig struct {
	BrokerURI          string
	ExchangeName       string
	RoutingKey         string
	QueueName          string
	AutoCreateRabbitMQ bool
}

func GetRabbitMQConfig() (*RabbitMQConfig, error) {
	var err error
	exchangeName := os.Getenv(EXCHANGE_NAME)
	routingKey := os.Getenv(ROUTING_KEY)
	queueName := os.Getenv(QUEUE_NAME)

	brokerHost := os.Getenv(BROKER_HOST)
	brokerUsername := os.Getenv(BROKER_USERNAME)
	brokerPassword := os.Getenv(BROKER_PASSWORD)
	brokerPort := os.Getenv(BROKER_PORT)
	brokerPortInt, err := strconv.Atoi(brokerPort)
	if err != nil {
		return nil, err
	}
	autoCreateRabbitMQ := os.Getenv(AUTO_CREATE_RABBITMQ)
	autoCreateRabbitMQBool, err := strconv.ParseBool(autoCreateRabbitMQ)
	if err != nil {
		return nil, err
	}

	brokerURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", brokerUsername, brokerPassword, brokerHost, brokerPortInt)

	return &RabbitMQConfig{BrokerURI: brokerURI, ExchangeName: exchangeName, RoutingKey: routingKey, QueueName: queueName, AutoCreateRabbitMQ: autoCreateRabbitMQBool}, err

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
