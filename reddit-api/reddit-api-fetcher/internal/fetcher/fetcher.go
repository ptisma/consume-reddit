package fetcher

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"reddit-api-fetcher/internal/config"
	"strings"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type SubredditFetcher struct {
	encodedData   string
	basicAuth     string
	UserAgentName string
	URL           string
	Category      string
	NumOfPosts    int
}

func (s *SubredditFetcher) FetchToken(ctx context.Context, body string) {
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(s.encodedData))

	req.Header.Add("Authorization", "Basic "+s.basicAuth)
	req.Header.Add("User-Agent", s.UserAgentName)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

func (s *SubredditFetcher) FetchPosts(ctx context.Context, body string) {

}

func GetSubredditFetcher(config *config.SubredditFetcherConfig) *SubredditFetcher {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", config.RedditUsername)
	data.Add("password", config.RedditPassword)
	encodedData := data.Encode()
	basicAuth := basicAuth(config.ClientID, config.ClientSecret)

}
