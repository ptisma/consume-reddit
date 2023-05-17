package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const EXCHANGE_NAME = "EXCHANGE_NAME"
const ROUTING_KEY = "ROUTING_KEY"
const BROKER_HOST = "BROKER_HOST"
const BROKER_USERNAME = "BROKER_USERNAME"
const BROKER_PASSWORD = "BROKER_PASSWORD"
const BROKER_PORT = "BROKER_PORT"
const AUTO_CREATE_RABBITMQ = "AUTO_CREATE_RABBITMQ"

type RabbitMQConfig struct {
	BrokerURI          string
	ExchangeName       string
	RoutingKey         string
	AutoCreateRabbitMQ bool
}

func GetRabbitMQConfig() (*RabbitMQConfig, error) {
	var err error
	exchangeName := os.Getenv(EXCHANGE_NAME)
	routingKey := os.Getenv(ROUTING_KEY)

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
	log.Println(brokerURI)

	return &RabbitMQConfig{BrokerURI: brokerURI, ExchangeName: exchangeName, RoutingKey: routingKey, AutoCreateRabbitMQ: autoCreateRabbitMQBool}, err

}

const REDDIT_USERNAME = "REDDIT_USERNAME"
const REDDIT_PASSWORD = "REDDIT_PASSWORD"
const CLIENT_ID = "CLIENT_ID"
const CLIENT_SECRET = "CLIENT_SECRET"
const USER_AGENT_NAME = "USER_AGENT_NAME"
const URL = "URL"
const CATEGORY = "CATEGORY"
const NUM_OF_POSTS = "NUM_OF_POSTS"

type SubredditFetcherConfig struct {
	RedditUsername string
	RedditPassword string
	ClientID       string
	ClientSecret   string
	UserAgentName  string
	URL            string
	Category       string
	NumOfPosts     int
}

func GetSubredditFetcherConfig() (*SubredditFetcherConfig, error) {
	var err error
	redditUsername := os.Getenv(REDDIT_USERNAME)
	redditPassword := os.Getenv(REDDIT_PASSWORD)
	clientID := os.Getenv(CLIENT_ID)
	clientSecret := os.Getenv(CLIENT_SECRET)
	userAgentName := os.Getenv(USER_AGENT_NAME)
	url := os.Getenv(URL)
	category := os.Getenv(CATEGORY)
	numOfPosts := os.Getenv(NUM_OF_POSTS)
	numOfPostsInt, err := strconv.Atoi(numOfPosts)
	if err != nil {
		return nil, err
	}

	return &SubredditFetcherConfig{RedditUsername: redditUsername, RedditPassword: redditPassword, ClientID: clientID, ClientSecret: clientSecret, UserAgentName: userAgentName, URL: url, Category: category, NumOfPosts: numOfPostsInt}, err

}
