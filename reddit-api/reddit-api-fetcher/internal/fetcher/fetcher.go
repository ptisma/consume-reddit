package fetcher

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	var err error
	req, err := http.NewRequest("GET", s.URL, nil)
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	posts, _, _, _ := jsonparser.Get(body, "data", "children")

	jsonparser.ArrayEach(posts, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		//title, _, _, _ := jsonparser.Get(value, "data", "title")
		fmt.Println("Whole post data")
		fmt.Println(string(value))
		err = s.Producer.StorePost(ctx, value)
		fmt.Println(err)

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

	baseURL, _ := url.Parse(config.URL + "/" + config.Category)
	params := url.Values{}
	params.Add("limit", strconv.Itoa(config.NumOfPosts))
	baseURL.RawQuery = params.Encode()

	URL := baseURL.String()

	client := &http.Client{}

	return &SubredditFetcher{EncodedData: encodedData, BasicAuth: basicAuth, UserAgentName: config.UserAgentName, URL: URL, Client: client, Producer: producer}, err

}
