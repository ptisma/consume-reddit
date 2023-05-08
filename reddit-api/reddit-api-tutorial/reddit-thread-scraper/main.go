package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/buger/jsonparser"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

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

func main() {

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", "beskucnik_na_feru")
	data.Add("password", "Wizard21298")

	encodedData := data.Encode()

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(encodedData))
	failOnError(err, "REQUEST")
	req.Header.Add("Authorization", "Basic "+basicAuth("95wusYL4TfUPURtBnswGQg", "qlMMTd2bN3ijKT7wrRlJtGpxQVyZ2w"))
	req.Header.Add("User-Agent", "reddit-api-scraper-test:0.01 by Pepe")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	res, err := client.Do(req)
	failOnError(err, "RESPONSE")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	failOnError(err, "BODY")
	fmt.Println(string(body))

	var token Token
	json.Unmarshal([]byte(body), &token)
	fmt.Println(token)

	// req, err = http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	// failOnError(err, "REQUEST 2")
	// req.Header.Add("Authorization", "bearer "+token.AccessToken)
	// req.Header.Add("User-Agent", "reddit-api-scraper-test:0.01 by Pepe")

	// res, err = client.Do(req)
	// failOnError(err, "RESPONSE 2")
	// defer res.Body.Close()

	// body, err = ioutil.ReadAll(res.Body)
	// failOnError(err, "BODY 2")
	// fmt.Println(string(body))

	baseURL := "https://oauth.reddit.com"
	resource := "/r/croatia/hot"
	params := url.Values{}
	params.Add("limit", "2")

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u) // "http://example.com/path?param1=value1&param2=value2"
	fmt.Println("URL:", urlStr)

	req, err = http.NewRequest("GET", urlStr, nil)
	failOnError(err, "REQUEST 3")
	req.Header.Add("Authorization", "bearer "+token.AccessToken)
	req.Header.Add("User-Agent", "reddit-api-scraper-test:0.01 by Pepe")

	res, err = client.Do(req)
	failOnError(err, "RESPONSE 3")
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	failOnError(err, "BODY 3")
	fmt.Println("Whole response")
	fmt.Println(string(body))
	fmt.Println("Posts")
	value1, _, _, _ := jsonparser.Get(body, "data")

	value2, _, _, _ := jsonparser.Get(value1, "children")
	//fmt.Println(string(value2))

	// Prints only titles of the posts
	jsonparser.ArrayEach(value2, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		//title, _, _, _ := jsonparser.Get(value, "data", "title")
		fmt.Println("Whole post data")
		fmt.Println(string(value))
	})

}
