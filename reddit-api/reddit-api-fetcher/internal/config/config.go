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
