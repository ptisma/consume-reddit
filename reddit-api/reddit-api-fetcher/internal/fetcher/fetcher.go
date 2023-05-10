package fetcher

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"reddit-api-fetcher/internal/config"
	"reddit-api-fetcher/internal/producer"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiration  uint64 `json:"expires_in"`
	Scope       string `json:"scope"`
}

type SubredditFetcher struct {
	EncodedData   string
	BasicAuth     string
	UserAgentName string
	URL           string
	Category      string
	NumOfPosts    int
	Client        *http.Client
	Producer      *producer.RabbitMQProducer
}

func (s *SubredditFetcher) FetchToken() (Token, error) {
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(s.EncodedData))
	if err != nil {
		return Token{}, err
	}

	req.Header.Add("Authorization", "Basic "+s.BasicAuth)
	req.Header.Add("User-Agent", s.UserAgentName)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := s.Client.Do(req)
	if err != nil {
		return Token{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Token{}, err
	}
	var token Token
	err = json.Unmarshal([]byte(body), &token)
	if err != nil {
		return Token{}, err
	}

	return token, err

}

func (s *SubredditFetcher) FetchAndStorePosts(ctx context.Context, token Token) error {

	baseURL, _ := url.Parse(s.URL + "/" + s.Category)
	params := url.Values{}
	params.Add("limit", strconv.Itoa(s.NumOfPosts))
	baseURL.RawQuery = params.Encode()

	URL := baseURL.String()

	var err error
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "bearer "+token.AccessToken)
	req.Header.Add("User-Agent", s.UserAgentName)

	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	posts, _, _, _ := jsonparser.Get(body, "data", "children")

	jsonparser.ArrayEach(posts, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		//title, _, _, _ := jsonparser.Get(value, "data", "title")
		log.Println("Whole post data")
		log.Println(string(value))
		err = s.Producer.StorePost(ctx, value, s.Category)
		log.Println(err)

	})

	return err

}

func GetSubredditFetcher(config *config.SubredditFetcherConfig, producer *producer.RabbitMQProducer) (*SubredditFetcher, error) {
	var err error
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", config.RedditUsername)
	data.Add("password", config.RedditPassword)
	encodedData := data.Encode()
	basicAuth := basicAuth(config.ClientID, config.ClientSecret)

	client := &http.Client{}

	return &SubredditFetcher{EncodedData: encodedData, BasicAuth: basicAuth, UserAgentName: config.UserAgentName, URL: config.URL, Category: config.Category, NumOfPosts: config.NumOfPosts, Client: client, Producer: producer}, err

}
